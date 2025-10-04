package main

import (
	"fmt"
	"net/http"

	"BLOG_API/handlers"
)

//TODO: MAKE CI/CD DO NOT MAKE CHANGES WITHOUT THE WEBSITE RUNNING ON THE SERVER!!!

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", serveFiles)
	mux.HandleFunc("GET /blog", handlers.GetALlPost)
	mux.HandleFunc("POST /blog", handlers.CreatePost)
	mux.HandleFunc("POST /blog/update", handlers.UpdatePost)
	mux.HandleFunc("GET /blog/search/{id}", handlers.SearchPost)
	mux.HandleFunc("DELETE /blog/{id}", handlers.DeletePost)
	fmt.Printf("Listening...")
	http.ListenAndServe(":8080", mux)
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	p := "./frontend" + r.URL.Path
	fmt.Println(p)
	http.ServeFile(w, r, p)
}
