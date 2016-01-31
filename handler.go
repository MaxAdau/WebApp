package main

import (
	"fmt" // just for debug
	"net/http"
	)

func SetHandlers() {
	http.HandleFunc("/Person", PersonHandler)
	http.HandleFunc("/", Index)
}

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Coucou %s !", r.URL.Path[1:])
	// Appel a Person.sdfasdf(r.URL.Path[1:])
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Coucou %s !", r.URL.Path[1:])

}