package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pe-Gomes/short-url/util"
)

type Server struct {
	config util.Config
	router *gin.Engine
}

func NewServer(config util.Config) (*Server, error) {
	server := &Server{
		config: config,
	}
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
