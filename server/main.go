package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "log"
    "fmt"
    _ "github.com/lib/pq"
    "database/sql"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "password"
  dbname   = "pathlogic"
)


type Site struct {
  id int
  name string
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

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "hello word")
}


func SiteShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  siteId := vars["siteId"]
  db := openDB()
  row := db.QueryRow("select * from sites where id = $1", siteId)
  site := new(Site)

  err := row.Scan(&site.id, &site.name)
  if err == sql.ErrNoRows {
    http.NotFound(w, r)
    return
  } else if err != nil {
    fmt.Println(err)
    http.Error(w, http.StatusText(500), 500)
    return
  }

  fmt.Fprintf(w, "%d, %s", site.id, site.name)
}


func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", Index)
  router.HandleFunc("/sites/{siteId}", SiteShow)
  log.Fatal(http.ListenAndServe(":8080", router))

}