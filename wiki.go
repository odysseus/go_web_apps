package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

func main() {
  //p1 := &Page{Title: "TestPage", Body: []byte("This is a simple page.")}
  //p1.save()
  //p2, _ := loadPage("TestPage")
  //fmt.Println(string(p2.Body))

  //http.HandleFunc("/", handler)
  //http.ListenAndServe(":8080", nil)

  http.HandleFunc("/view/", viewHandler)
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

// A simple web server/request handler
// when any request is sent it prints "Hi there, I love ____" where
// _____ is anything after the / or "the world!" if the request is for the root
func handler(w http.ResponseWriter, r *http.Request) {
  path := r.URL.Path[1:]
  if path == "" { path = "the world!" }
  fmt.Fprintf(w, "Hi there, I love %s", path)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/view/"):]
  p, _ := loadPage(title)
  fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
