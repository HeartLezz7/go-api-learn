package main

import (
	"fmt"
	"net/http"
)

func SayHello(response http.ResponseWriter, request *http.Request) {
	method := request.Method
	fmt.Println(method)
	if method == "GET" {
		fmt.Println("GET method")
		response.WriteHeader(200)
		response.Write([]byte("welcome to go world"))
	}
}

func main() {
	http.HandleFunc("/hello", SayHello)

	fmt.Println("HTTP run server")
	http.ListenAndServe("localhost:8080", nil)

}
