package siteShow

import(
	"net/http"
	"fmt"
	"log"
	"../../dataBase/db"
	"../../structs/site"
	"../../structs/jsonData"
	"time"
	"encoding/json"
	"github.com/gorilla/mux"
)


func SiteShow(w http.ResponseWriter, r *http.Request) {
	s := new(site.Site)
	jd := new(jsonData.JsonData)

	vars := mux.Vars(r)
	siteID := vars["siteId"]
	db := db.OpenDB()

	err := db.QueryRow("select id, name from sites where id = $1", siteID).Scan(&s.ID, &s.Name)

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
