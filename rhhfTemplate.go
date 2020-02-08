package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type NewsAggPage struct {
	Title   string
	Article []Article
}

type Article struct {
	Text  string
	Video string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "RHHF Blog", Article: []Article{
		Article{Text: "Feb 7, 2020: Goony Chonga with her hit, No Effort", Video: "https://www.youtube.com/embed/ZabFoeD5n5s"},
		Article{Text: "Feb 6, 2020: Let's appreciate Princess Nokia for her new single, G.O.A.T.", Video: "https://www.youtube.com/embed/qa5HMtVqR6o"},
		Article{Text: "Feb 5, 2020: Let's check out Mestiza, this is her hit 'En Las Nubes'", Video: "https://www.youtube.com/embed/3drucBTcLuI"},
	}}

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	//http.ListenAndServe(":8000", nil)
	err := http.ListenAndServeTLS(":4055", "./freshmintrecords_com.crt", "./freshmintrecords.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
