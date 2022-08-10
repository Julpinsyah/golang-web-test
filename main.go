package main

import (
	"julpin_1/handler"
	"log"
	"net/http"
)

func main() {

	//inisialisasi mux server
	mux := http.NewServeMux()

	//menambahkan static file ke server
	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// menambahkan handler ke server
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
