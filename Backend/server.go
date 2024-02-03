package main 

import (
	"net/http"
	"fmt"
)

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website!")
    })

    http.ListenAndServe(":80", nil)
}
