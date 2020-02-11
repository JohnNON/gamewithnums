package gameserver

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/JohnNON/gamewithnums/internal/app/model"
	"github.com/gorilla/csrf"
)

func (s *server) handleGameEndCheck() http.HandlerFunc {
	type respond struct {
		Val bool `json:"val"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		res := &respond{}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		game, ok := session.Values["user_game"]
		if !ok || game == nil {
			s.response(w, r, http.StatusOK, res)
			return
		}
		res.Val = true

		w.Header().Set("X-CSRF-Token", csrf.Token(r))
		s.response(w, r, http.StatusOK, res)
	}
}

func (s *server) handleNewGame() http.HandlerFunc {
	type request struct {
		Val string `json:"val"`
	}

	randInt := func(min int, max int) int {
		rand.Seed(time.Now().UTC().UnixNano())
		return rand.Intn(max - min)
	}

	randomString := func(l int) string {
		var str string
		for i := 0; i < l; i++ {
			str = str + strconv.Itoa(randInt(0, 10))
			time.Sleep(time.Nanosecond)
		}
		return str
	}

	return func(w http.ResponseWriter, r *http.Request) {
		rnd := &model.Round{}

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		l, err := strconv.Atoi(req.Val)

		if err != nil {
			l = 4
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok || id == nil {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		game := randomString(l)
		t := time.Now().UTC().Format("2006-01-02 15:04:05")

		if err := s.store.Round().DeleteByUserID(strconv.Itoa(id.(int))); err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		rnd.UserID = id.(int)
		rnd.Difficulty = l
		rnd.GameNumber = game
		rnd.GameTime = t

		if err := s.store.Round().Create(rnd); err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.clearGameSession(session)

		session.Values["user_game"] = game
		session.Values["user_game_start"] = t
		session.Values["user_game_diff"] = l
		session.Values["user_game_rounds"] = 0

		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		req.Val = "OK"
		w.Header().Set("X-CSRF-Token", csrf.Token(r))
		s.response(w, r, http.StatusOK, req)
	}
}

func (s *server) handleGameRoundCheck() http.HandlerFunc {
	type request struct {
		Val string `json:"val"`
	}

	type respond struct {
		Val    string `json:"val"`
		Status string `json:"status"`
		Code   string `json:"code"`
	}

	check := func(game, reqs string) string {
		res := ""
		if reqs != "" {
			for i := range reqs {
				if strings.Index(game[i:], string(reqs[i])) == 0 {
					res = res + "B"
				} else if strings.Index(game, string(reqs[i])) >= 0 {
					res = res + "K"
				} else {
					res = res + "X"
				}
			}
		} else {
			for range game {
				res = res + "X"
			}
		}

		return res
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		res := &respond{}
		rnd := &model.Round{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		game, ok := session.Values["user_game"]
		if !ok || game == nil {
			s.error(w, r, http.StatusInternalServerError, errInternalServerError)
			return
		}

		diff, ok := session.Values["user_game_diff"]
		if !ok || diff == nil {
			s.error(w, r, http.StatusInternalServerError, errInternalServerError)
			return
		}

		round, ok := session.Values["user_game_rounds"]
		if !ok || round == nil {
			s.error(w, r, http.StatusNotImplemented, errInternalServerError)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok || id == nil {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		if game.(string) == req.Val {
			res.Val = req.Val[:diff.(int)]
			res.Status = win
			res.Code = ""

			gameStart, ok := session.Values["user_game_start"]
			if !ok || gameStart == nil {
				s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
				return
			}

			rec := &model.Record{}
			rec.UserID = id.(int)
			rec.Difficulty = diff.(int)
			rec.RoundCount = round.(int) + 1
			t, err := time.Parse("2006-01-02 15:04:05", gameStart.(string))
			if err != nil {
				s.error(w, r, http.StatusInternalServerError, errNotAuthenticated)
				return
			}

			rec.GameTime = int(time.Now().UTC().Sub(t) / time.Second)

			duration, ok := session.Values["user_game_duration"]
			if ok && duration != nil {
				rec.GameTime = rec.GameTime + duration.(int)
			}

			err = s.store.Record().Create(rec)
			if err != nil {
				s.error(w, r, http.StatusInternalServerError, errNotAuthenticated)
				return
			}

			if err := s.store.Round().DeleteByUserID(strconv.Itoa(id.(int))); err != nil {
				s.error(w, r, http.StatusInternalServerError, err)
				return
			}

			s.clearGameSession(session)
		} else {
			res.Val = req.Val
			res.Status = cont
			res.Code = check(game.(string), req.Val)

			t := time.Now().UTC().Format("2006-01-02 15:04:05")

			rnd.UserID = id.(int)
			rnd.GameTime = t
			rnd.Inpt = req.Val
			rnd.Outpt = res.Code

			if err := s.store.Round().Create(rnd); err != nil {
				s.error(w, r, http.StatusInternalServerError, err)
				return
			}
		}

		session.Values["user_game_rounds"] = round.(int) + 1

		if err := s.sessionStore.Save(r, w, session); err != nil {
			fmt.Println(123456789)
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		w.Header().Set("X-CSRF-Token", csrf.Token(r))
		s.response(w, r, http.StatusOK, res)
	}
}

func (s *server) handleLoadGame() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok || id == nil {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		rounds, err := s.store.Round().FindByUserID(strconv.Itoa(id.(int)))
		rnd := *rounds

		if len(rnd) < 1 {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		t1, err := time.Parse("2006-01-02 15:04:05", rnd[0].GameTime)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, errNotAuthenticated)
			return
		}

		t2, err := time.Parse("2006-01-02 15:04:05", rnd[len(rnd)-1].GameTime)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, errNotAuthenticated)
			return
		}

		session.Values["user_game"] = rnd[0].GameNumber
		session.Values["user_game_start"] = time.Now().UTC().Format("2006-01-02 15:04:05")
		session.Values["user_game_diff"] = rnd[0].Difficulty
		session.Values["user_game_rounds"] = len(*rounds) - 1
		session.Values["user_game_duration"] = int(t2.Sub(t1) / time.Second)

		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		w.Header().Set("X-CSRF-Token", csrf.Token(r))
		s.response(w, r, http.StatusOK, rounds)
	}
}
