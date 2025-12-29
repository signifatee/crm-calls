package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	config     *Config
	logger     *logrus.Logger
	httpServer *http.Server
}

func New(config *Config, handler http.Handler) *Server {
	return &Server{
		config: config,
		httpServer: &http.Server{
			Addr:           config.BindAddr,
			Handler:        handler,
			MaxHeaderBytes: config.MaxHeaderBytes,
			ReadTimeout:    config.ReadTimeout,
			WriteTimeout:   config.WriteTimeout,
		},
		logger: logrus.New(),
	}
}

func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Infof("Starting API server on %s", s.config.BindAddr)

	return s.httpServer.ListenAndServe()
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
