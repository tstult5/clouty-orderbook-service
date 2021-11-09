package controllers

import (
	"fmt"
	"log"
	"net/http"
  //"time"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver - unused

	"github.com/tstult5/clouty-orderbook-service/api/models"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
  //DB environment vars unused for sqlite, placeholder for postgres, etc.

	//var err error

  server.DB, _  = gorm.Open("sqlite3", "./clouty.db")

  //statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS orderbook (id INTEGER PRIMARY KEY, orderbookname TEXT, created_at DATE)")
  //statement.Exec()
  //statement, _ := db.Prepare("INSERT INTO orderbook (orderbookname, created_at) VALUES (?, ?)")
  //statement.Exec("orderBookOne", time.Now())

	server.DB.Debug().AutoMigrate(&models.OrderBook{}) //database migration

	server.Router = mux.NewRouter()
  //defer db.Close()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
