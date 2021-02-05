package main

import (
	"fmt"
	"html/template"
	"net/http"
	"io/ioutil"
)

var posts = makeRange(1, 5)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ukrf", ukrfHandler)
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":5051", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/main.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "main", nil)
}

func ukrfHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/ukrf.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "ukrf", posts)
}

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

type Page struct {
    Title string
    Body []byte
}

func loadPage(title string) (*Page, error) {
    filename := "templates/" + title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	t, _ := template.ParseFiles("templates/view.html", "templates/header.html")
    if err != nil {
        fmt.Fprintf(w, err.Error())
    }
    t.ExecuteTemplate(w, "view", p)
}