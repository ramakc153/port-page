package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	port := ":3000"
	resp, _ := http.Get("https://api.open-meteo.com/v1/forecast?latitude=-6.265011347092098&longitude=106.79644618569907&current=temperature_2m&timezone=Asia%2FBangkok&forecast_days=1")
	bodyStr, _ := io.ReadAll(resp.Body)
	fmt.Println(string(bodyStr))
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("server served on ", port)
	log.Println(http.ListenAndServe(port, nil))
}
