package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "log"
    "fmt"
    _ "github.com/lib/pq"
    "database/sql"
    "encoding/json"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "password"
  dbname   = "pathlogic"
)


type Site struct {
  Id int `json:"id"`
  Name string `json:"string"`
}

type JSONData struct {
  SiteID int `json:"siteId"`
}

func openDB() *sql.DB {
  var db *sql.DB

  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }

  if err = db.Ping(); err != nil {
    log.Fatal(err)
  }
  fmt.Println("Successfully connected!")
  return db
}

func SiteShow(w http.ResponseWriter, r *http.Request) {
  var s Site
  jd := new(JSONData)
  vars := mux.Vars(r)
  siteId := vars["siteId"]
  db := openDB()
  
  err := db.QueryRow("select id from sites where id = $1", siteId).Scan(&s.Id)

   if err != nil {
      log.Fatalf("Failed to redirect stderr to file: %v", err)
    }

    if err != nil {
      log.Printf("Failed to encode a JSON response: %v", err)
      w.WriteHeader(http.StatusInternalServerError)
      return
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    jd.SiteID = s.Id
    data, err := json.Marshal(jd)

    if err != nil {
        fmt.Println(err)
    }
    fmt.Fprintf(w, string(data))
}


func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/sites/{siteId}", SiteShow).Methods("GET")
  log.Fatal(http.ListenAndServe(":8080", router))

}