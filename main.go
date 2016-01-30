package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage (title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title : title, Body : body}, err
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Coucou %s !", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	fmt.Println(title)
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body )
}


func main() {
	// p1 := &Page{Title : "Coucou", Body : []byte("Ceci est mon body")}
	// p1.save()
	// fmt.Println(string(p1.Body))
	// p2, _ := loadPage("Coucou")
	// fmt.Println(string(p2.Body))

	// http.HandleFunc("/", handler)
	
	http.HandleFunc("/view/", viewHandler)

	http.ListenAndServe(":8080", nil)

}
