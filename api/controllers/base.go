package controllers

import (
	"fmt"
	"log"
	"net/http"
  //"time"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
  _ "github.com/tstult5/clouty-orderbook-service/docs"
  httpSwagger "github.com/swaggo/http-swagger"
	//_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver - unused

	"github.com/tstult5/clouty-orderbook-service/api/models"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
  //DB environment vars unused for sqlite, placeholder for postgres, etc.

	var err error

  server.DB, err  = gorm.Open("sqlite3","file::memory:?cache=shared")

  if err != nil {
    fmt.Println(err)
    fmt.Println("Unable to open DB for initialization")
  }

	server.DB.Debug().AutoMigrate(&models.Order{}) //database migration
	server.Router = mux.NewRouter()
  //defer db.Close()
  // Swagger
  server.Router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
