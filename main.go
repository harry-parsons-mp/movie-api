package main

import (
	"movie-api/server"
	"movie-api/server/routes"
)

func main() {
	s := server.NewServer("./database.db")
	routes.InitialiseRoutes(s)
	s.Start()
}
