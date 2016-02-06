package main

import (
	"fmt"
	"net/http"
	"regexp"
)

type Handler struct {
	p Person
	c Car
}

func (h *Handler) WebAPI(w http.ResponseWriter, r *http.Request) {
	// Create a regexp matching APIObj Categories
	rule := regexp.MustCompile("^/(Person|Car)/")
	category := rule.FindString(r.URL.Path)

	fmt.Fprintf(w, "r.URL contains : %+v\n", *r.URL)

	// Call for the matched handler
	// Sending URL without category
	switch url := r.URL.Path[len(category):] ; category {
		case "" : fmt.Fprintf(w, "Category is not correct\n")
		case "/Person/" : fmt.Fprintf(w, h.p.Handler(url))
		case "/Car/" : fmt.Fprintf(w, h.c.Handler(url))

	}

}