package http_api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"recuperator/recuperator"
	"strconv"
)

func RegisterHandlers(r recuperator.Recuperator) {
	http.HandleFunc("/recuperator/getData", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("getData\n")
		d, _ := r.GetData()
		j, _ := json.Marshal(d)
		writer.Write(j)
	})

	http.HandleFunc("/recuperator/getTemperature", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("getTemperature\n")
		d, _ := r.GetTemperature()
		j, _ := json.Marshal(d)

		writer.Write(j)
	})

	http.HandleFunc("/recuperator/setTemperature", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("setTemperature\n")
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
		fmt.Printf("setFanSpeed\n")
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
}
