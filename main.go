package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"html/template"
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

func renderHTML(w http.ResponseWriter, p *Page, tplt string) {

	t, err := template.ParseFiles(tplt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	fmt.Println(title)
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderHTML(w, p, "view.html")
	// t, _ := template.ParseFiles("view.html")
	// t.Execute(w, p)
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body )
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil{
		p = &Page{Title : title}
	}

	renderHTML(w, p, "edit.html")
	// t, _ := template.ParseFiles("edit.html")
	// t.Execute(w, p)
/*	fmt.Fprintf (w,"<h1>Editing %s</h1>"+
        "<form action=\"/save/%s\" method=\"POST\">"+
        "<textarea name=\"body\">%s</textarea><br>"+
        "<input type=\"submit\" value=\"Save\">"+
        "</form>", p.Title, p.Title, p.Body)*/
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title : title, Body : []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

func main() {
	// p1 := &Page{Title : "Coucou", Body : []byte("Ceci est mon body")}
	// p1.save()
	// fmt.Println(string(p1.Body))
	// p2, _ := loadPage("Coucou")
	// fmt.Println(string(p2.Body))

	// http.HandleFunc("/", handler)
	
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	http.ListenAndServe(":8080", nil)

}
