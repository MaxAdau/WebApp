package main

import (
	_ "fmt"
	"net/http"
	_	"gopkg.in/mgo.v2/bson"
)

var db = Db{ip : "127.0.0.1"}


func main (){

	db.Connect("127.0.0.1")
	mux := http.NewServeMux()



	h := &Handler {}

	// trouver un moyen pour envoyer tout ce qui vient de /Person*
	mux.HandleFunc("/", h.WebAPI)

	http.ListenAndServe(":8080", mux)


	// fmt.Println(result)
}