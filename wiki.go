package main

import (
  //"fmt"
  "io/ioutil"
  "net/http"
  "html/template"
)

func main() {
  http.HandleFunc("/view/", viewHandler)
  http.HandleFunc("/edit/", editHandler)
  http.HandleFunc("/save/", saveHandler)
  http.ListenAndServe(":8080", nil)
}

// Defining a struct to represent a single page
// The Body is a byte slice instead of text because the io library
// expects a byte slice when reading/writing
type Page struct {
  Title string
  Body []byte
}

// A save method for a page that writes the page to a file
// This has an explicit receiver so it can be called using page.save()
// This writes the file and returns the error returned by WriteFile
func (p *Page) save() error {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}

// Loads a page, taking the string title to find the file and load it
// back into memory, returning a pointer to a page literal
func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
  t, err := template.ParseFiles(tmpl + ".html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  err = t.Execute(w,p)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/view/"):]
  p, err := loadPage(title)
  if err != nil {
    http.Redirect(w, r, "/edit/"+title, http.StatusFound)
    return
  }
  renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/edit/"):]
  p, err := loadPage(title)
  if err != nil {
    p = &Page{Title: title}
  }
  renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/save/"):]
  body := r.FormValue("body")
  p := &Page{Title: title, Body: []byte(body)}
  err := p.save()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
