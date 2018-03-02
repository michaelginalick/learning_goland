package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "pathlogic"
)

type site struct {
	ID   int    `json:"id"`
	Name string `json:"string"`
}

type jSONData struct {
	SiteID       int       `json:"siteId"`
	ClientID     string    `json:"clientID"`
	ModifiedDate time.Time `json:"modifiedDate"`
}

type menu struct {
	ID       int `json:"id"`
	Schedule struct {
		ID int `json:"id"`
	} `json:"schedule"`
}

func openDB() *sql.DB {
	var db *sql.DB

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func siteShow(w http.ResponseWriter, r *http.Request) {
	var s site
	jd := new(jSONData)
	vars := mux.Vars(r)
	siteID := vars["siteId"]
	db := openDB()

	err := db.QueryRow("select id from sites where id = $1", siteID).Scan(&s.ID)

	if err != nil {
		log.Fatalf("Failed to redirect stderr to file: %v", err)
	}

	if err != nil {
		log.Printf("Failed to encode a JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	jd.SiteID = s.ID
	jd.ClientID = s.Name
	jd.ModifiedDate = time.Now()
	data, err := json.Marshal(jd)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(data))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/sites/{siteId}", siteShow).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}
