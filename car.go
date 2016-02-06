package main

import (
	"fmt"
	"regexp"
)

type Car struct {
	Name string
	Model string
	Age int
	
}

func (c Car) Handler(str string) string {
	// Todo : rajouter les arguments en fin de regexp
	rule := regexp.MustCompile("^(Search|Create|Read|Update|Delete)/")
	verb := rule.FindString(str)

	// Switch on verb. Sending  str without the verb part
	switch str = str[len(verb):] ; verb {
		case "" : return fmt.Sprintf("APIObj [Car] has no such verb\n")
		case "Search/" 	: return c.Search(str)
		case "Create/" 	: return c.Create(str)
		case "Read/" 	: return c.Read(str)
		case "Update/" 	: return c.Update(str)
		case "Delete/" 	: return c.Delete(str)
	}
	return str
}

func (c Car) Search(str string) string {
	return fmt.Sprintf("Searching : %s", str)
}

func (c Car) Create(str string) string {
	return fmt.Sprintf("Creating : %s", str)
}

func (c Car) Read(str string) string {
	return fmt.Sprintf("Reading : %s", str)
}
func (c Car) Update(str string) string {
	return fmt.Sprintf("Updating : %s", str)
}
func (c Car) Delete(str string) string {
	return fmt.Sprintf("Deleting : %s", str)
}


type Cars []Car 