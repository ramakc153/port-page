package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("server served on :80")
	log.Println(http.ListenAndServe(":80", nil))
}
