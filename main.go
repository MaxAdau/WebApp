package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"html/template"
	"regexp"
	// "errors"
	"github.com/MaxAdau/lib"
)

// Declare a Strcut Page with two fields
type Page struct {
	Title string
	Body []byte
}

// Save a Page into a file
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// Return an *Page from a title
func loadPage (title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title : title, Body : body}, err
}

// Print the URL PATH. Currently unused
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Coucou %s !", r.URL.Path[1:])
}


// Declare a viariabla named template. This Must fonction throw a panic error
// if thoses filenames are not found
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// Render HTML from the page name
func renderTemplate(w http.ResponseWriter, p *Page, tplt string) {

	err := templates.ExecuteTemplate(w, tplt, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Load a page , link to the /view/ URL
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/view/"):]
	/*title, err := getTitle(w, r)
	if err != nil{
		return
	}*/
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, p, "view.html")
	// t, _ := template.ParseFiles("view.html")
	// t.Execute(w, p)
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body )
}

// Edit a page, linked to the /edit/ URL
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
/*	title, err := getTitle(w, r)
	if err != nil{
		return
	}*/
	p, err := loadPage(title)
	if err != nil{
		p = &Page{Title : title}
	}

	renderTemplate(w, p, "edit.html")
	// t, _ := template.ParseFiles("edit.html")
	// t.Execute(w, p)
/*	fmt.Fprintf (w,"<h1>Editing %s</h1>"+
        "<form action=\"/save/%s\" method=\"POST\">"+
        "<textarea name=\"body\">%s</textarea><br>"+
        "<input type=\"submit\" value=\"Save\">"+
        "</form>", p.Title, p.Title, p.Body)*/
}

// Save a Page, linked to the /save/ URL
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title : title, Body : []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

// Handle view, edit and save Handlers
// take a *Handler function on parameter, check if the title page is valid
// through valid_path variable
// Call the Handler func with the title as parameter
// https://golang.org/ref/spec#Function_literals for more details
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m ==nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

// Declare a viariable that check the regexp. MustCompile Panic if regexp
// is not fullfiled
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

/*func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m ==nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil
}*/

func main() {
	// p1 := &Page{Title : "Coucou", Body : []byte("Ceci est mon body")}
	// p1.save()
	// fmt.Println(string(p1.Body))
	// p2, _ := loadPage("Coucou")
	// fmt.Println(string(p2.Body))

	// http.HandleFunc("/", handler)
	
	/*http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))*/

	// http.ListenAndServe(":8080", nil)
	lib.PrintHello()

}
