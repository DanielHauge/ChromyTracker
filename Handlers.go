package main

import (
	"net/http"
	"bytes"

	_ "github.com/lib/pq"
	"gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
	"io"
	"log"
	"encoding/json"
	"database/sql"
	"fmt"
	"time"
)

var DB *sql.DB

func setheader (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin ,Accept, Content-Type, Content-Length, Accept-Encoding")
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	} else { w.Header().Set("Access-Control-Allow-Origin", "*")}
}

func Index(w http.ResponseWriter, r *http.Request) {
	setheader(w,r)

	var buffer bytes.Buffer
	buffer.WriteString("# Chromy-Tracker API!\n")


	w.Write(blackfriday.Run([]byte(buffer.String())))
}

func SaveTask(w http.ResponseWriter, r *http.Request){
	setheader(w,r)
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
		log.Printf(err.Error())
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var Task Task

	if err := json.Unmarshal(body, &Task); err != nil {
		log.Println("Couldn't marshal")
	}



	stmt, err := DB.Prepare("INSERT INTO Tasks (navn, type, lastdone, midtime, crittime) VALUES (?, ?, ?, ?, ?)")
	if err !=nil{
		log.Println("It was prepare")
		log.Println(err.Error())
	}

	_, err = stmt.Exec(Task.Navn, Task.Type, Task.LastDone, Task.MidTime, Task.CritTime)
	if err != nil {
		log.Println("It was execute")
		log.Println(err.Error())
	}

}


func GetAll(w http.ResponseWriter, r *http.Request){
	setheader(w,r)
	log.Println("I GOT GETALL!")
	results := []Task{}

	//connStr := "postgres://postgres:1702941837@192.168.33.10"

	DB.Ping()


	rows, err := DB.Query("SELECT * FROM Tasks;")
	for rows.Next(){
		var id int
		var navn string
		var Type string
		var LastDone time.Time
		var Midtime time.Time
		var Crittime time.Time


		if err = rows.Scan(&id ,&navn, &Type, &LastDone, &Midtime, &Crittime); err != nil{
			log.Println("IT WAS ROW SCAN!")
			log.Fatal(err)
		}

		results = append(results, Task{id,navn, Type, LastDone, Midtime, Crittime})

	}
	if err!=nil{
	log.Println("It was something with QUERY!")
	}

	msgs, err := json.Marshal(results); if err != nil{ log.Println("IT WAS MARSHAL!") }

	fmt.Fprint(w, string(msgs))


}


