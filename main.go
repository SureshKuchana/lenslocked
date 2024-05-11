package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, _ *http.Request){
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome 1</h1>")
}

func contactHandler(w http.ResponseWriter, _ *http.Request){
	fmt.Fprint(w, `<h1>Contact Page</h1>
	<p>To get in touch, email me <a href='mailto:kuchana123.suresh@gmail.com'>suresh</a></p>`)
}

func main(){
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on http://localhost:3000...")
	http.ListenAndServe(":3000", r)
}