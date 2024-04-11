package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/pe-Gomes/short-url/infra/db/repository"
	"github.com/pe-Gomes/short-url/util"
)

type Server struct {
	config util.Config
	store  db.Store
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
	}))

	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.POST("/users", server.createUser)
	router.GET("/users/:email", server.getUserByEmail)
	router.GET("/users", server.listUsers)
	router.DELETE("/users/:id", server.deleteUserById)

	router.POST("/url", server.createShortURL)
	router.GET("/:slug", server.getShortURLBySlug)
	router.GET("/url", server.listURLByUser)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
