package pkg

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Server struct {
	logger *logrus.Logger
}

func NewServerHTTP(logger *logrus.Logger) *Server {
	return &Server{
		logger: logger,
	}
}

func (s *Server) Launch(config *Config, router http.Handler) error {
	server := http.Server{
		Addr:         config.ServerHTTPMain.BindAddr,
		Handler:      router,
		ReadTimeout:  time.Duration(config.ServerHTTPMain.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.ServerHTTPMain.WriteTimeout) * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
