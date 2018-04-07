package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/handlers"
	"./routes/siteShow"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	// "./dataBase/db"
	// "./structs/site"
	// "./structs/jsonData"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/sites/{siteId}", siteShow.SiteShow).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, router)))
}
