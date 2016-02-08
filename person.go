package main

import (
	"fmt"
	"regexp"
	_ "gopkg.in/mgo.v2/bson"
	"net/url"

)

type Person struct {
	Name string
	Phone string
	Age int
	
}

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
	return fmt.Sprintf("Searching : %s", str)
}

func (p Person) Create(str string) string {
	session := db.Session.Copy()
	c := session.DB("toilet").C("restaurants")
	p = Person{}
	c.Find("").One(&p)
	// err := c.Find({"cuisine" : "Italian"})
	return "lol\n"
	// return fmt.Sprintf("Creating : %s", string(p))
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


type Persons []Person 