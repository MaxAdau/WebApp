package main

import (
	"fmt"
	"net/http"
	"gopkg.in/mgo.v2/bson"
)

var db = Db{ip : "127.0.0.1"}


func main (){

	db.Connect("127.0.0.1")
	SetHandlers()
	
	http.ListenAndServe(":8080", nil)






	coll := GetCollection("test", "people")

	result := Person{}
	coll.Find(bson.M{"name" : "Ale"}).One(&result)

	fmt.Println(result)
}