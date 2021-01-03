package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"recuperator/recuperator"
	"strconv"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getListenAddress() string {
	port := getEnv("PORT", "18080")
	return ":" + port
}

func main() {
	recuperatorIpAddress := getEnv("RECUPERATOR_IP", "192.168.0.71")
	recuperatorLogin := getEnv("RECUPERATOR_LOGIN", "admin")
	recuperatorPass := getEnv("RECUPERATOR_PASS", "admin")

	fmt.Printf("--------------------------------------------------\n")
	fmt.Printf("HouseProxy at " + getListenAddress() + "\n")
	fmt.Printf("Recuperator IP: " + recuperatorIpAddress + "\n")
	fmt.Printf("--------------------------------------------------\n")

	r := recuperator.NewClient(recuperatorIpAddress, recuperatorLogin, recuperatorPass)

	http.HandleFunc("/recuperator/getData", func(writer http.ResponseWriter, request *http.Request) {
		d, _ := r.GetData()
		j, _ := json.Marshal(d)
		writer.Write(j)
	})

	http.HandleFunc("/recuperator/getTemperature", func(writer http.ResponseWriter, request *http.Request) {
		d, _ := r.GetTemperature()
		j, _ := json.Marshal(d)

		writer.Write(j)
	})

	http.HandleFunc("/recuperator/setTemperature", func(writer http.ResponseWriter, request *http.Request) {
		keys := request.URL.Query()
		i, err := strconv.Atoi(keys.Get("value"))
		if err != nil {
			i = 22
		}
		d, _ := r.SetTemperature(i)
		j, _ := json.Marshal(d)
		writer.Write(j)
	})

	http.HandleFunc("/recuperator/setFanSpeed", func(writer http.ResponseWriter, request *http.Request) {
		keys := request.URL.Query()
		i, err := strconv.Atoi(keys.Get("value"))
		if err != nil {
			i = 1
		}
		d, _ := r.SetFanSpeed(i)
		j, _ := json.Marshal(d)
		writer.Write(j)
	})

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("HomeProxy"))
	})

	if err := http.ListenAndServe(getListenAddress(), nil); err != nil {
		panic(err)
	}
}
