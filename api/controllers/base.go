package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

import (
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	fmt.Printf("Dbdriver: %s\n", Dbdriver)
	fmt.Printf("DbUser: %s\n", DbUser)
	fmt.Printf("DbHost: %s\n", DbHost)
	fmt.Printf("DbName: %s\n", DbName)
	fmt.Printf("DbPort: %s\n", DbPort)
	fmt.Printf("DbPassword: %s\n", DbPassword)

	var err error

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=require password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	fmt.Printf("DBURL: %s\n", DBURL)

	server.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8099")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
