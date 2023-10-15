package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/endpointB", func(w http.ResponseWriter, r *http.Request) {
		// Совершаем HTTP-запрос к ServiceA
		resp, err := http.Get("http://localhost:8080/endpointA")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "hi!ServiceA. Response from ServiceA: %s", body)
	})

	http.ListenAndServe(":8081", nil)
}
