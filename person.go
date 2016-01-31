package main

import (
	"fmt" // Pour debug
)

type Person struct {
	name string
	Phone string
	age int
	
}

func PersonIndex(str string) string{
	
	// TODO : pourquoi rien ne s' affiche ?
	fmt.Println("Coucou %s !", str)


	// TODO : switch entres les differents operations
		




	/*coll := GetCollection("test", "people")
	result := Person{}
	coll.Find(bson.M{"name" : "Ale"}).One(&result)
*/	return "coucou"
}

type Persons []Person