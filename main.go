package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Response struct
type Response struct {
	Chars []Chars `json:"results"`
}

// Chars Struct
type Chars struct {
	Name   string `json:name`
	Image  string `json:image`
	Gender string `json:gender`
	Status string `json:status`
}

// CharsPage Struct
type CharsPage struct {
	Title string
	Items []Chars
}

func handler(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("CHAR_NAME")
	status := os.Getenv("CHAR_STATUS")
	req, _ := http.NewRequest("GET", "https://rickandmortyapi.com/api/character/", nil)
	q := req.URL.Query()
	q.Add("name", name)
	q.Add("status", status)
	req.URL.RawQuery = q.Encode()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Get(req.URL.String())

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	p := CharsPage{Title: "This is the " + name + "iest page", Items: responseObject.Chars}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func mkslice(a []Chars, start, end int) []Chars {
	return a[start:end]
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
