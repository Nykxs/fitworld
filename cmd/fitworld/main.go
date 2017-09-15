package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/nykxs/fitworld/http"
	"github.com/nykxs/fitworld/pg"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "fitworld_demo"
)

func SetupStore() (*pg.Store, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return pg.NewStore(psqlInfo)
}

func main() {
	store, err := SetupStore()
	if err != nil {
		log.Fatal(err)
	}
	UserService := pg.NewUserService(store)

	HTTPServer := http.NewServer(UserService)
	if err := HTTPServer.Setup(); err != nil {
		log.Fatal(err)
	}
	if err := HTTPServer.Start(); err != nil {
		log.Fatal(err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	if err := HTTPServer.Stop(); err != nil {
		log.Fatal(err)
	}
}
