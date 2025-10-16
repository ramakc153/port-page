package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

type Response struct {
	Message string
	Lat     float32 `json:"lat"`
	Lon     float32 `json:"lon"`
	City    string  `json:"city"`
}

func main() {
	port := ":3000"

	// resp, _ := http.Get("https://api.open-meteo.com/v1/forecast?latitude=-6.265011347092098&longitude=106.79644618569907&current=temperature_2m&timezone=Asia%2FBangkok&forecast_days=1")
	// bodyStr, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(bodyStr))
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/location", getLocation)
	log.Println("server served on ", port)
	log.Println(http.ListenAndServe(port, nil))

}

func getLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	queryParams := r.URL.Query()
	ip := queryParams.Get("clientIp")
	fmt.Println()
	if ip == "" || net.ParseIP(ip) == nil {

		response := map[string]string{
			"message": "invalid IP",
			"lat":     "",
			"lon":     "",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	urlBuilder := fmt.Sprintf("http://ip-api.com/json/%s", ip)
	resp, err := http.Get(urlBuilder)
	if err != nil {
		errStr := fmt.Sprintf("error when getting ip %s", err.Error())
		response := map[string]string{
			"message": errStr,
			"lat":     "",
			"lon":     "",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var response Response
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		panic(err)
	}
	response.Message = "success"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return
}
