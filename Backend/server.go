package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func writeState(w http.ResponseWriter, r *http.Request, data map[string]DeviceSettings) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error reading request body: %v", err)
		return
	}

	// Close the request body to avoid resource leaks
	defer r.Body.Close()

	// Create a new UpdateSync instance to decode JSON into
	var update UpdateSync

	// Unmarshal the JSON into the UpdateSync struct
	if err := json.Unmarshal(body, &update); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding JSON: %v", err)
		return
	}

	// Successfully parsed the JSON into the struct
	// You can now use the update variable to access the parsed data

	data[update.Key] = update.Data

	w.WriteHeader(http.StatusOK)
}

func toJsonString(val DeviceSettings) (string, error) {
	result, err := json.Marshal(val)
	if err != nil {
		return "", err
	}

	return string(result), err
}

func readState(w http.ResponseWriter, r *http.Request, data map[string]DeviceSettings) {
	deviceId := r.URL.Query().Get("deviceId")

	result, ok := data[deviceId]

	if !ok {
		decode, err := toJsonString(DeviceSettings{IsOn: false})

		if err != nil {
			w.WriteHeader(http.StatusPartialContent)
			fmt.Fprintf(w, "Could not find key %s, failure occured: %v", deviceId, err)
		} else {
			w.WriteHeader(http.StatusPartialContent)
			fmt.Fprintf(w, "%s", decode)
		}

		return
	}

	stringRes, err := toJsonString(result)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to parse device with id %s", deviceId)
	} else {
		fmt.Fprintf(w, "%s", stringRes)
	}
}

type DeviceSettings struct {
	IsOn bool `json:"isOn"`
}

type UpdateSync struct {
	Key  string         `json:"key"`
	Data DeviceSettings `json:"deviceSettings"`
}

func main() {
	data := make(map[string]DeviceSettings)

	data["0"] = DeviceSettings{IsOn: true}
	data["1"] = DeviceSettings{IsOn: true}
	data["2"] = DeviceSettings{IsOn: true}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,
			`Use /writeState to write device state. Input JSON {Key: 'deviceId', deviceState: {...}}
		readState to read device state. Use header 'deviceId' to select the device.`)
	})

	http.HandleFunc("/writeState", func(w http.ResponseWriter, r *http.Request) {
		writeState(w, r, data)
	})

	http.HandleFunc("/readState", func(w http.ResponseWriter, r *http.Request) {
		readState(w, r, data)
	})

	http.ListenAndServe(":8080", nil)
}
