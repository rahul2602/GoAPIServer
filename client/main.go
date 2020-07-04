package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

// Data : Structure used to store JSON DATA from form
type Data struct {
	EventName          string
	WebsiteURL         string
	SessionID          string
	OriginalHeight     int
	OriginalWidth      int
	ActualHeight       int
	ActualWidth        int
	CopyAndPaste       map[string]bool
	FormCompletionTime int
}

func main() {
	http.HandleFunc("/", IndexPage)
	http.HandleFunc("/form", EventData)
	http.ListenAndServe(":3000", nil)
}

// IndexPage : To view Index page
func IndexPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
}

//EventData : To handle the post of incoming request
func EventData(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("err", err)
	}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Println("err", err)
	}
	log.Println(resp)
}
