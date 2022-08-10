package main

import (
	"julpin_1/handler"
	"log"
	"net/http"
)

func main() {

	// menambahkan handler ke server
	mux := http.NewServeMux()

	// static file server
	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/home", handler.HomeHandler)
	mux.HandleFunc("/ubah", handler.UpdateHandler)
	mux.HandleFunc("/del", handler.DeleteHandler)

	log.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
