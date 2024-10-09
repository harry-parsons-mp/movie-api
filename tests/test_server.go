package tests

import (
	"fmt"
	"log"
	"movie-api/server"
	"movie-api/server/routes"
)

type TestServer struct {
	S *server.Server
}

func NewTestServer() *TestServer {
	ts := &TestServer{S: server.NewServer("./test-database.db")}
	routes.InitialiseRoutes(ts.S)
	return ts
}

func (ts *TestServer) ClearTable(table string) {
	err := ts.S.Db.Exec(fmt.Sprintf("DELETE FROM %v", table)).Error
	if err != nil {
		log.Fatalf("Failed to clear table: %v", err)
	}
	err = ts.S.Db.Exec(fmt.Sprintf("DELETE FROM sqlite_sequence WHERE name='%v'", table)).Error
	if err != nil {
		log.Fatalf("Failed to reset AUTO_INCREMENT on table: %v", err)
	}
}
