package main

import (
	"fmt"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Println(method)
	if method == "GET" {
		fmt.Println("GET method")
		w.WriteHeader(200)
	}
}

func main() {
	http.HandleFunc("/hello", SayHello)

	fmt.Println("HTTP run server")
	http.ListenAndServe("localhost:8080", nil)

}
