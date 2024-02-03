package main 

import (
	"net/http"
	"fmt"
)

func set(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, fmt.Sprintf("state: %s", r.FormValue("state"))) 
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, fmt.Sprintf("device id: %s", r.FormValue("deviceId")))
}

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website!")
    })

    http.HandleFunc("/set", set)

    http.HandleFunc("/get", get)

    http.ListenAndServe(":80", nil)
}
