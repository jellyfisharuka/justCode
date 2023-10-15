package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/endpointA", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello my dear friend")
	})

	http.ListenAndServe(":8080", nil)
}
