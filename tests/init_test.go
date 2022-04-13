package tests

import (
	"log"
	"os"
	"testing"

	"jjrepos/golang/api/database"
	"jjrepos/golang/api/server"

	unitTest "github.com/Valiben/gin_unit_test"
)

func init() {
	setupServer()
	newLog := log.New(os.Stdout, "", log.Llongfile|log.Ldate|log.Ltime)
	unitTest.SetLog(newLog)
	setupDB()
}

func setupDB() {
	database.Connect()
	database.Db.Exec("DELETE FROM books")
}

func setupServer() {
	router := server.SetupRoutes()
	unitTest.SetRouter(router)
}
func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
