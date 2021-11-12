package api

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/tstult5/clouty-orderbook-service/api/utils/initdb"
	"github.com/tstult5/clouty-orderbook-service/api/controllers"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
  initdb.Load(server.DB)
	server.Run(":8080")

}
