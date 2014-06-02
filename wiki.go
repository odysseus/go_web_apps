package main

import (
  "fmt"
  "io/ioutil"
)

func main() {
  p1 := &Page{Title: "TestPage", Body: []byte("This is a simple page.")}
  p1.save()
  p2, _ := loadPage("TestPage")
  fmt.Println(string(p2.Body))
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
