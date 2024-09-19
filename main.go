package main

import "movie-api/server"

func main() {
	s := server.Server{}
	s.InitialiseDB()
	s.InitialiseRoutes()

}
