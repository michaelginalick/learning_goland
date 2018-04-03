package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	_ "github.com/lib/pq"
	"github.com/server/structs/site"
	"github.com/server/structs/jsonData"
	"github.com/server/db"
)


func siteShow(w http.ResponseWriter, r *http.Request) {
	s := new(site.Site)
	jd := new(jsonData.JsonData)

	vars := mux.Vars(r)
	siteID := vars["siteId"]
	db := db.OpenDB()
	

	err := db.QueryRow("select id from sites where id = $1", siteID).Scan(&s.ID)

	if err != nil {
		log.Fatalf("Failed to redirect stderr to file: %v", err)
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
