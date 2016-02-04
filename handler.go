package main

import (
	"fmt" // just for debug
	"net/http"
	"html"
	// "regexp" // used to check the URL validity
	// "errors" // used to add errors
	)

type Handler struct {

}

func init() {
	//TODO : get "/person*"
	// http.HandleFunc("/person", PersonHandler)
	// http.HandleFunc("/", Index)
}

func (h *Handler) PersonHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.URL)
		// TODO : pourquoi rien ne s' affiche ?

	fmt.Printf("%+v\n", r.URL)
	fmt.Fprintf(w, "Scheme :%s\n", html.EscapeString(r.URL.Scheme))
	fmt.Fprintf(w, "Opaque :%s\n", html.EscapeString(r.URL.Opaque))
	// fmt.Fprintln(w, "User :%+v\n", html.EscapeString(r.?tring()))
	fmt.Fprintf(w, "Host :%s\n", html.EscapeString(r.URL.Host))
	fmt.Fprintf(w, "Path :%s\n", html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, "RawPath :%s\n", html.EscapeString(r.URL.RawPath))
	fmt.Fprintf(w, "RawQuery :%s\n", html.EscapeString(r.URL.RawQuery))
	fmt.Fprintf(w, "Fragment :%s\n", html.EscapeString(r.URL.Fragment))

	//check avec regexp.MustCompile(un truc)










	/*coll := GetCollection("test", "people")
	result := Person{}
	coll.Find(bson.M{"name" : "Ale"}).One(&result)
	*/
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Coucou %s !", r.URL.Path[1:])

}