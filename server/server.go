package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"movie-api/models"
	"movie-api/repos"
	"os"
)

type Server struct {
	Db    *gorm.DB
	Echo  *echo.Echo
	Repos *repos.Repos
}

func NewServer(path string) *Server {
	s := Server{}
	s.InitialiseDB(path)
	s.Repos = repos.NewRepos(s.Db)
	return &s
}
func (server *Server) InitialiseDB(path string) {
	server.Echo = echo.New()

	Db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect to Db")
	}
	Db.AutoMigrate(&models.Movie{}, &models.Review{}, &models.User{})
	server.Db = Db

}
func (server *Server) CloseDB() {
	db, err := server.Db.DB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.Close()

}
func (server *Server) Start() {
	server.Echo.Logger.Fatal(server.Echo.Start(":1234"))
}
func (server *Server) DeleteDB(path string) {
	log.Println("Deleting db...")
	os.Remove(path)

}
