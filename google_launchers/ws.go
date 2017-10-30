package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"../google"
	"html/template"
)

func main() {
	http.HandleFunc("/fakeGoogleSearch", handleSearch) // init a new API
	fmt.Println("serving on http://localhost:8080/fakeGoogleSearch")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handleSearch(w http.ResponseWriter, req *http.Request) {
	log.Println("serving", req.URL)
	start := time.Now()
	results, err := google.SearchReplicated("golang", 104*time.Millisecond) // run a new search
	elapsed := time.Since(start)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := google.Response{Results: results, Elapsed: elapsed}

	responseTemplate.Execute(w, resp) // build and print to screen the results
	fmt.Println(resp, err)
}

var responseTemplate = template.Must(template.New("results").Parse(
`<html><head/><body>
  <ol>{{range .Results}}
    <li>{{.Title}} - <a href="{{.URL}}">{{.URL}}</a></li>
  {{end}}</ol>
  <p>{{len .Results}} results in {{.Elapsed}}</p>
</body></html>`))