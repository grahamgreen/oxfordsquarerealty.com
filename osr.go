package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("osr.html")
	t.Execute(w, nil)
}

func main() {
	var port int
	flag.IntVar(&port, "port", 9000, "port to listen on")
	flag.Parse()

	listenString := "127.0.0.1:" + strconv.Itoa(port)
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	log.Fatal(http.ListenAndServe(listenString, r))
}
