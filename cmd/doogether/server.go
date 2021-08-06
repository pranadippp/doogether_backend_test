package main

import (
	"go/internal/pkg/config"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	conf   *config.AppConfig
	engine *gin.Engine
}

func NewServer(conf *config.AppConfig, engine *gin.Engine) *Server {
	return &Server{
		conf:   conf,
		engine: engine,
	}
}

func (s *Server) Run() {
	if err := s.engine.Run(s.conf.GetString("server.address")); err != nil {
		log.Fatalf("failed to run server : %v", err)
	}
}

func (s *Server) RunTls() {
	if err := s.engine.RunTLS(s.conf.GetString("server.secureaddress"), s.conf.GetString("server.pem"), s.conf.GetString("server.key")); err != nil {
		log.Fatalf("failed to run server : %v", err)
	}
}
