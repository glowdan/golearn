package main

import (
	"net/http"
	"fmt"
	"html"
)

type Hello struct {}
type String string
func (h Hello) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	fmt.Fprint(w, "Hello!")
}
func (s String) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	fmt.Fprint(w, "Hello! String")
}

func main() {
	var h Hello
	http.HandleFunc("/s", func(w http.ResponseWriter,r *http.Request){
			fmt.Fprint(w, "Hello! String")
		})
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		})

	http.ListenAndServe("localhost:4000", h)
}

