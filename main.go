package main

import (
	_ "fmt"
	"net/http"
	_	"gopkg.in/mgo.v2/bson"
)

var db = Db{ip : "127.0.0.1"}


func main (){

	db.Connect("127.0.0.1")
	defer db.Session.Close()


// Used to insert datas
	// c := db.Session.DB("WebApp").C("person")
	// c.Insert(
	// &Person{Name: "Jesus", Age:33, Phone:"0666666666"},
	// &Person{Name: "Mahomet", Age:99, Phone:"0777777777"},
	// &Person{Name: "Santa"},
	// &Person{Name: "Myself", Age:27, Phone:"0123456789"},
	// &Person{Name: "Jean Edouart de St Gilbert", Age:2, Phone:"0987654321"},
	// &Person{Name: "Sadam", Age:56, Phone:"010504080506"},
	// &Person{Name: "Jules Cesar", Phone:"0000050000"},
	// &Person{Name: "Batman",  Phone:"06786663466"})



	mux := http.NewServeMux()

	h := &Handler {Person{}, Person{}}

	// trouver un moyen pour envoyer tout ce qui vient de /Person*
	mux.HandleFunc("/", h.WebAPI)
	http.ListenAndServe(":8080", mux)


	// fmt.Println(result)
}