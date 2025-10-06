package main

import (
	"fmt"
	"net/http"
	"os"

	"BLOG_API/handlers"
)

// TODO: MAKE CI/CD DO NOT MAKE CHANGES WITHOUT THE WEBSITE RUNNING ON THE SERVER!!!
// TODO: connect the blog docker container and the mysql container to a single network, the ip for the db go file uses is 172.35.0.3
func main() {
	frontendPath := os.Getenv("FRONTEND_PATH")
	fs := http.FileServer(http.Dir(frontendPath))
	http.Handle("/src/frontend", http.StripPrefix("/src/frontend", fs))
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
