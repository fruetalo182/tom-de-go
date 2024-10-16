package server

import (
	"tom-de-go/internal/config"
	"tom-de-go/internal/handlers"
	"tom-de-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	config *config.Config
}

func New(cfg *config.Config) *Server {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.SecurityHeaders())

	r.GET("/", handlers.ServeImage(cfg.ImagePath))

	return &Server{
		router: r,
		config: cfg,
	}
}

func (s *Server) Run() error {
	return s.router.Run(":" + s.config.Port)
}