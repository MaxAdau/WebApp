package main

import (
	"fmt"
	"regexp"
)

type Person struct {
	Name string
	Phone string
	Age int
	
}

func (p Person) Handler(str string) string {
	// Todo : rajouter les arguments en fin de regexp
	rule := regexp.MustCompile("^(Search|Create|Read|Update|Delete)/")
	verb := rule.FindString(str)

	// Switch on verb. Sending  str without the verb part
	switch str = str[len(verb):] ; verb {
		case "" : return fmt.Sprintf("APIObj [Person] has no such verb\n")
		case "Search/" 	: return p.Search(str)
		case "Create/" 	: return p.Create(str)
		case "Read/" 	: return p.Read(str)
		case "Update/" 	: return p.Update(str)
		case "Delete/" 	: return p.Delete(str)
	}
	return str
}

func (p Person) Search(str string) string {
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


type Persons []Person 