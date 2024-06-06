package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("Running... ")
	defer fmt.Println("...ended")
	http.HandleFunc("/", index)
	http.ListenAndServe("localhost:80", nil)
}

var indexTemplate = template.Must(
	template.ParseFiles("templates/index.tmpl"),
)

type Index struct {
	Title, Body string
	Links       []Link
}
type Link struct {
	URL, Title string
}

func index(res http.ResponseWriter, req *http.Request) {
	/*data := &Index{
		Title: "ken-morel",
		Body:  "Welcom newbies",
	}
	if err := indexTemplate.Execute(res, data); err != nil {
		log.Println(err)
	}*/
	res.WriteHeader(200)
	res.Write([]byte("hello world"))
}
