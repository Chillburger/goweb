package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

// this will save the body to a text file, using the page Title as the file Name
// return an error value because that is the return value of WriteFile
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// constructs the filename from the title parameter, reads the file's contents into a new variables
// body and returns a new Page Literal
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// of the type http.HandlerFunc, takes an http.ResponseWriter, and http.Request
// http.ResponseWriter value assambles HTTP server's response, by writing to it
// htt.Request is a data structure that represents the client HTTP request r.URL.Path
// is the path component of the request URL, trailing 1: means create a subslice of Path from the
// first character to the end
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi There, I love %s!", r.URL.Path[1:])
}

func main() {
	// handle all requests to the web root with handler
	http.HandleFunc("/", handler)
	// specifies that server should listen on 8080 on any inerface
	http.ListenAndServe(":8080", nil)
}
