package gameserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"text/template"
	"time"

	"github.com/JohnNON/gamewithnums/internal/app/model"
	"github.com/JohnNON/gamewithnums/internal/app/store"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gorilla/csrf"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
)

const (
	win  = "Win"
	cont = "Continue"
)

var funcMap = template.FuncMap{
	"add": func(a, b int) int {
		return a + b
	},
	"formated": func(t int) string {
		var h int = int((time.Duration(t) * time.Second).Hours())
		var m int = int((time.Duration(t) * time.Second).Minutes()) % 60
		var s int = (t % 3600) % 60
		var str string
		if h < 10 {
			str = str + "0"
		}
		str = str + strconv.Itoa(h) + ":"

		if m < 10 {
			str = str + "0"
		}
		str = str + strconv.Itoa(m) + ":"

		if s < 10 {
			str = str + "0"
		}
		str = str + strconv.Itoa(s)

		return str
	},
}

func renderTemplate(s string) (*template.Template, error) {
	return template.ParseFiles(
		s,
		"./internal/templates/head.html",
		"./internal/templates/header.html",
		"./internal/templates/message.html",
		"./internal/templates/scripts.html",
		"./internal/templates/base.html")
}

func renderFuncTemplate(s string) (*template.Template, error) {
	tmpl := template.New(s)
	tmpl.Funcs(funcMap)
	return tmpl.ParseFiles(
		s,
		"./internal/templates/head.html",
		"./internal/templates/header.html",
		"./internal/templates/message.html",
		"./internal/templates/scripts.html",
		"./internal/templates/base.html")
}

func (s *server) handleIndexPage() http.HandlerFunc {
	var templateIndexPage *template.Template
	templateIndexPage = template.Must(renderFuncTemplate("./internal/templates/index.html"))

	var rcEasy, rcMedium, rcHard *[]model.Record
	var err error
	var wg sync.WaitGroup

	return func(w http.ResponseWriter, r *http.Request) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rcEasy, err = s.store.Record().GetAllRecords("4")
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			rcMedium, err = s.store.Record().GetAllRecords("8")
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			rcHard, err = s.store.Record().GetAllRecords("12")
		}()
		wg.Wait()

		if err != nil {
			if err != store.ErrRecordNotFound {
				s.error(w, r, http.StatusInternalServerError, err)
				return
			}
		}

		pageData := map[string]interface{}{
			"Title":   "Games' Records!",
			"Message": "Games' Records!",
			"Records": struct{}{},
		}

		if s.checkForMenu(r) {
			pageData["Login"] = struct{}{}
		}

		if rcEasy != nil {
			pageData["RecordsEasy"] = rcEasy
		}

		if rcMedium != nil {
			pageData["RecordsMedium"] = rcMedium
		}

		if rcHard != nil {
			pageData["RecordsHard"] = rcHard
		}

		err = templateIndexPage.ExecuteTemplate(w, "base", pageData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (s *server) handleLogin() http.HandlerFunc {
	type login struct {
		Email    string
		Password string
		_        string `schema:"Csrf"`
	}

	validateLogin := func(user *login) error {
		return validation.ValidateStruct(
			user,
			validation.Field(&user.Email, validation.Required, is.Email),
			validation.Field(&user.Password, validation.Required, is.Alphanumeric),
		)

	}

	var templateLoginPage *template.Template
	templateLoginPage = template.Must(renderTemplate("./internal/templates/login.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if r.Method == "POST" {

			l := &login{}
			s.readForm(r, l)

			err := validateLogin(l)
			if err == nil {

				u, err := s.store.User().FindByEmail(l.Email)
				if err != nil || !u.ComparePassword(l.Password) {
					session.Values["user_message"] = errIncorrectEmailOrPassword.Error()
					if err := s.sessionStore.Save(r, w, session); err != nil {
						s.error(w, r, http.StatusInternalServerError, err)
						return
					}
					http.Redirect(w, r, "/login", http.StatusFound)
					return
				}

				session.Values["user_id"] = u.ID
				s.clearGameSession(session)
				if err := s.sessionStore.Save(r, w, session); err != nil {
					s.error(w, r, http.StatusInternalServerError, err)
					return
				}
				http.Redirect(w, r, "/index", http.StatusFound)
				return

			}

			session.Values["user_message"] = errIncorrectEmailOrPassword.Error()
			if err := s.sessionStore.Save(r, w, session); err != nil {
				s.error(w, r, http.StatusInternalServerError, err)
				return
			}

			http.Redirect(w, r, "/login", http.StatusFound)

		} else {

			if s.checkForMenu(r) {
				http.Redirect(w, r, "/index", http.StatusFound)
				return
			}

			var message string
			mes, ok := session.Values["user_message"]
			fmt.Println(ok, mes)
			if !ok || mes == nil {
				message = "Login for game!"
			} else {
				message = mes.(string)
				session.Values["user_message"] = nil
				if err := s.sessionStore.Save(r, w, session); err != nil {
					s.error(w, r, http.StatusInternalServerError, err)
					return
				}
			}

			pageData := map[string]interface{}{
				"Title":          "Login!",
				"Message":        message,
				csrf.TemplateTag: csrf.TemplateField(r),
			}

			err := templateLoginPage.ExecuteTemplate(w, "base", pageData)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

		}
	}
}

func (s *server) clearGameSession(session *sessions.Session) {
	session.Values["user_message"] = nil
	session.Values["user_game"] = nil
	session.Values["user_game_start"] = nil
	session.Values["user_game_diff"] = nil
	session.Values["user_game_rounds"] = nil
	session.Values["user_game_duration"] = nil
}

func (s *server) handleLogout() http.HandlerFunc {
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

		if err := s.store.Round().DeleteByUserID(strconv.Itoa(id.(int))); err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = nil
		s.clearGameSession(session)
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		http.Redirect(w, r, "/index", http.StatusFound)
	}
}

func (s *server) handleRegistration() http.HandlerFunc {
	type registration struct {
		Email          string
		Password       string
		PasswordRepeat string
		_              string `schema:"Csrf"`
	}

	requiredIf := func(cond bool) validation.RuleFunc {
		return func(value interface{}) error {
			if cond {
				return validation.Validate(value, validation.Required)
			}

			return nil
		}
	}

	validateLogin := func(user *registration) error {
		return validation.ValidateStruct(
			user,
			validation.Field(&user.Email,
				validation.Required,
				is.Email),
			validation.Field(&user.Password,
				validation.Required,
				validation.By(requiredIf(user.Password == user.PasswordRepeat)),
				validation.Length(8, 128),
				is.Alphanumeric),
			validation.Field(&user.PasswordRepeat,
				validation.Required,
				validation.By(requiredIf(user.Password == user.PasswordRepeat)),
				validation.Length(8, 128),
				is.Alphanumeric),
		)

	}

	var templateRegistrationPage *template.Template
	templateRegistrationPage = template.Must(renderTemplate("./internal/templates/registration.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if r.Method == "POST" {

			reg := &registration{}
			s.readForm(r, reg)

			err := validateLogin(reg)
			if err == nil {

				_, err := s.store.User().FindByEmail(reg.Email)
				if err == nil {
					session.Values["user_message"] = "email is registered already"
					if err := s.sessionStore.Save(r, w, session); err != nil {
						s.error(w, r, http.StatusInternalServerError, err)
						return
					}
					http.Redirect(w, r, "/login", http.StatusFound)
					return
				}
				if reg.Password == reg.PasswordRepeat {
					u := &model.User{
						Email:    reg.Email,
						Password: reg.Password,
					}

					if err := s.store.User().Create(u); err != nil {
						s.error(w, r, http.StatusUnprocessableEntity, err)
						return
					}

					u.Sanitize()
				}
				session.Values["user_message"] = "you are registered successfully"
				if err := s.sessionStore.Save(r, w, session); err != nil {
					s.error(w, r, http.StatusInternalServerError, err)
					return
				}
				http.Redirect(w, r, "/login", http.StatusFound)
				return

			}

			session.Values["user_message"] = "wrong entered values"
			if err := s.sessionStore.Save(r, w, session); err != nil {
				s.error(w, r, http.StatusInternalServerError, err)
				return
			}
			http.Redirect(w, r, "/registration", http.StatusFound)

		} else {

			if s.checkForMenu(r) {
				http.Redirect(w, r, "/index", http.StatusFound)
				return
			}

			pageData := map[string]interface{}{
				"Title":          "Registration!",
				csrf.TemplateTag: csrf.TemplateField(r),
			}

			err := templateRegistrationPage.ExecuteTemplate(w, "base", pageData)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

		}
	}
}

func (s *server) handleGamePage() http.HandlerFunc {
	var templateGamePage *template.Template
	templateGamePage = template.Must(renderTemplate("./internal/templates/game.html"))

	return func(w http.ResponseWriter, r *http.Request) {

		pageData := map[string]interface{}{
			"Title":          "Game!",
			csrf.TemplateTag: csrf.TemplateField(r),
		}

		if s.checkForMenu(r) {
			pageData["Login"] = struct{}{}
		}

		err := templateGamePage.ExecuteTemplate(w, "base", pageData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	}
}

func (s *server) readForm(r *http.Request, data interface{}) {
	r.ParseForm()
	decoder := schema.NewDecoder()
	decodeErr := decoder.Decode(data, r.PostForm)
	if decodeErr != nil {
		log.Printf("error mapping parsed form data to struct: %e\n", decodeErr)
	}
}

func (s *server) handleGameStatisticPage() http.HandlerFunc {
	var templateIndexPage *template.Template
	templateIndexPage = template.Must(renderFuncTemplate("./internal/templates/index.html"))

	var rcEasy, rcMedium, rcHard *[]model.Record
	var err error
	var wg sync.WaitGroup

	return func(w http.ResponseWriter, r *http.Request) {
		u := r.Context().Value(ctxKeyUser).(*model.User)
		id := strconv.Itoa(u.ID)
		wg.Add(1)
		go func() {
			defer wg.Done()
			rcEasy, err = s.store.Record().FindByUserID(id, "4")
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			rcMedium, err = s.store.Record().FindByUserID(id, "8")
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			rcHard, err = s.store.Record().FindByUserID(id, "12")
		}()
		wg.Wait()

		if err != nil {
			if err != store.ErrRecordNotFound {
				s.error(w, r, http.StatusInternalServerError, err)
				return
			}
		}

		pageData := map[string]interface{}{
			"Title":   "Users' Records!",
			"Message": fmt.Sprintf("%s Records!", u.Email),
		}

		if s.checkForMenu(r) {
			pageData["Login"] = struct{}{}
		}

		if rcEasy != nil {
			pageData["RecordsEasy"] = *rcEasy
		}

		if rcMedium != nil {
			pageData["RecordsMedium"] = *rcMedium
		}

		if rcHard != nil {
			pageData["RecordsHard"] = *rcHard
		}

		err = templateIndexPage.ExecuteTemplate(w, "base", pageData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (s *server) checkForMenu(r *http.Request) bool {
	session, err := s.sessionStore.Get(r, sessionName)
	if err != nil {
		return false
	}

	id, ok := session.Values["user_id"]
	if !ok || id == nil {
		return false
	}

	return true
}
