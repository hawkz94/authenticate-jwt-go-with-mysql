package server

import (
	"log"
	"os"

	"tecmentor-api/server/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	port := os.Getenv("APP_PORT")
	return Server{
		port:   port,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Print("Server is running on port", s.port)
	log.Fatal(router.Run(":" + s.port))
}
