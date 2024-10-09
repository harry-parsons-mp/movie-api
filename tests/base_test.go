package tests

import (
	"os"
	"testing"
)

var ts *TestServer

func TestMain(m *testing.M) {
	ts = NewTestServer()
	defer ts.S.CloseDB()

	code := m.Run()
	os.Exit(code)
}
