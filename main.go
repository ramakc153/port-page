package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":3000"
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("server served on ", port)
	log.Println(http.ListenAndServe(port, nil))
}
