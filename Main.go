package main

import (
	"log"
	"net/http"

	"database/sql"
)

// go get gopkg.in/mgo.v2
// go get “github.com/gorilla/mux”
// go get github.com/rs/cors
// go get -u gopkg.in/russross/blackfriday.v2
// go get github.com/lib/pq



func main() {

	connStr := "user=postgres sslmode=disable password=1702941837 host=192.168.33.10"
	DB, _ = sql.Open("postgres", connStr)

	defer DB.Close()

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":9191", router))

}

