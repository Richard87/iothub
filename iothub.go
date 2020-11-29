package main

import (
	"fmt"
	"net/http"
)

func main() {

}

func init() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/open_gate", openGate)

	e := http.ListenAndServe(":8080", nil)
	if e != nil{
		panic(e)
	}
}

func openGate(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "{\"message\": \"opening gate\"}")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "<html><body>Hello, World!</body></html>")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}