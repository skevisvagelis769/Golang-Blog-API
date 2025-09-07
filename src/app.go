package main

import (
	"net/http"

	"BLOG_API/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /blog", handlers.GetALlPost)
	mux.HandleFunc("POST /blog", handlers.CreatePost)
	http.ListenAndServe(":8080", mux)
}
