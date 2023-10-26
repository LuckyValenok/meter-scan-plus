package api

import (
	"backend/internal/api/service"
	"backend/internal/config"
	"backend/internal/store"
)

type Service service.Service

func NewService(config config.ApiConfig) (*Service, error) {
	dbConfig := config.Database
	connection, err := store.OpenConnection(dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Schema)
	if err != nil {
		return nil, err
	}
	s := &Service{
		DB:     connection,
		Router: configureRouter(),
	}
	return s, nil
}

func (s *Service) Start() error {
	go func() {
		// TODO: сюда порт
		if err := s.Router.Run(); err != nil {
			panic(err)
		}
	}()
	return nil
}

func (s *Service) Shutdown() error {
	return store.ShutdownConnection(s.DB)
}
