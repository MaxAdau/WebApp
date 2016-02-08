package main

import (
	"fmt"
	"regexp"
	"gopkg.in/mgo.v2/bson"
	"net/url"

)

type Person struct {
	Name string
	Phone string
	Age int
	
}

type Persons []Person 

func (p Person) Handler(url *url.URL) string {
	// Todo : rajouter les arguments en fin de regexp
	rule := regexp.MustCompile("/(search|create|read|update|delete)")
	verb := rule.FindString(url.Path)

	// Switch on verb. Sending  str without the verb part
	switch  verb {
		case "" : return fmt.Sprintf("APIObj [Person] has no such verb\n")
		case "/search" 	: return p.Search(url.RawQuery)
		case "/create" 	: return p.Create(url.RawQuery)
		case "/read" 	: return p.Read(url.RawQuery)
		case "/update" 	: return p.Update(url.RawQuery)
		case "/delete" 	: return p.Delete(url.RawQuery)
	}
	return url.RawQuery
}

func (p Person) Search(str string) string {
	// ParseQuery
	_, err := url.ParseQuery(str)
	if err !=  nil {
		return fmt.Sprintf("Error while Parsing query %s\n", str)
	}

	// Check word validity --
	//
	//

	// Create db Session and collection
	session := db.Session.Copy()
	c := session.DB("WebApp").C("person")
	// Find in collection

	var persons Persons
	err = c.Find(bson.M{"name": "Santa"}).All(&persons)

	fmt.Printf("Query = %+v", persons)
	// Format answer. Can have more than one !
	return fmt.Sprintf("Searching : %s", str)
}

func (p Person) Create(str string) string {

	return fmt.Sprintf("Creating : %s", str)
}

func (p Person) Read(str string) string {
	return fmt.Sprintf("Reading : %s", str)
}
func (p Person) Update(str string) string {
	return fmt.Sprintf("Updating : %s", str)
}
func (p Person) Delete(str string) string {
	return fmt.Sprintf("Deleting : %s", str)
}

