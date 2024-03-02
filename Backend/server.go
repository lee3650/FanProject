package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func writeState(w http.ResponseWriter, r *http.Request, data map[string]DeviceSettings) {
    request = // TODO! 
}

func readState(w http.ResponseWriter, r *http.Request, data map[string]DeviceSettings) {
	deviceId := r.URL.Query().Get("deviceId")
	fmt.Fprintf(w, "Welcome to my website!")
	fmt.Fprintf(w, "Read device id: %s", deviceId)
}

type DeviceSettings struct {
	IsOn bool
}

type UpdateSync struct {
	Key  string
	Data DeviceSettings
}

func main() {
	data := make(map[string]DeviceSettings)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		deviceId := r.URL.Query().Get("deviceId")
		fmt.Fprintf(w, "Welcome to my website!")
		fmt.Fprintf(w, "Read device id: %s", deviceId)
	})

	http.HandleFunc("/writeState", func(w http.ResponseWriter, r *http.Request) {
		writeState(w, r, data)
	})

	http.HandleFunc("/readState", func(w http.ResponseWriter, r *http.Request) {
		readState(w, r, data)
	})

	http.ListenAndServe(":8080", nil)
}
