package main

import "movie-api/server"

func main() {
	s := server.Server{}
	s.InitialiseDB("./database.db")
	s.ConfigCors()
	s.InitialiseRoutes()

}
