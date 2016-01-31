package main

import (
	"fmt" // just for debug
	"net/http"
	// "regexp" // used to check the URL validity
	// "errors" // used to add errors
	)

func SetHandlers() {
	//TODO : get "/person*"
	http.HandleFunc("/person", PersonHandler)
	// http.HandleFunc("/", Index)
}


func PersonHandler(w http.ResponseWriter, r *http.Request) {

	var _ = PersonIndex(r.URL.Path[:])

	// fmt.Fprintf(w, "Result %s !", result)
	// Appel a Person.sdfasdf(r.URL.Path[1:])
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Coucou %s !", r.URL.Path[1:])

}