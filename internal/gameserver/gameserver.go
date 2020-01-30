package gameserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

// GameServer - структура, хранящая сервер и его настройки
type GameServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New - функция для инициализации GameServer
func New(config *Config) *GameServer {
	return &GameServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start - осуществляет запуск сервера
func (s *GameServer) Start() error {
	if err := s.configLogger(); err != nil {
		return err
	}

	s.configRouter()

	s.logger.Info("starting game server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// configLogger - конфигурирует логгер
func (s *GameServer) configLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *GameServer) configRouter() {
	s.router.HandleFunc("/", s.handleHello())
}

func (s *GameServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello!")
	}
}
