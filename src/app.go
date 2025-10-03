package main

import (
	"fmt"
	"net/http"

	"BLOG_API/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /blog", handlers.GetALlPost)
	mux.HandleFunc("POST /blog", handlers.CreatePost)
	mux.HandleFunc("POST /blog/update", handlers.UpdatePost)
	mux.HandleFunc("POST /blog/search/{id}", handlers.SearchPost)
	fmt.Printf("Listening...")
	http.ListenAndServe(":8080", mux)
}
