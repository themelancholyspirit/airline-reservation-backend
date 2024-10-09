package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/themelancholyspirit/airline-reservation-system/storage"
	"github.com/themelancholyspirit/airline-reservation-system/util"
)

func (s *Server) setupCORS() {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // Replace with your frontend URL
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = true

	s.Router.Use(cors.New(config))
}

type Server struct {
	ListenAddr string
	Router     *gin.Engine
	Storage    storage.Storage
}

func NewServer(listenAddr string, storage storage.Storage) *Server {
	server := &Server{
		Router:     gin.Default(),
		ListenAddr: listenAddr,
		Storage:    storage,
	}
	server.setupCORS()
	return server
}

func (s *Server) SetupRoutes() {
	s.Router.GET("/", s.handleMain)
	s.Router.GET("/flights", s.handleListFlights)
	s.Router.POST("/signup", s.handleSignup)
	s.Router.POST("/login", s.handleLogin)
	s.Router.GET("/profile", util.AuthMiddleware(s.handleGetProfile))
	s.Router.PUT("/profile", util.AuthMiddleware(s.handleUpdateProfile))
}

func (s *Server) handleMain(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to the Airline Reservation System"})
}
