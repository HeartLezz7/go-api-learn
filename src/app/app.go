package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func SayHello(response http.ResponseWriter, request *http.Request) {
	method := request.Method
	fmt.Println(method)
	if method == "GET" {
		fmt.Println("GET method")
		response.WriteHeader(200)
		response.Write([]byte("welcome to go world"))
	} else if method == "POST" {
		fmt.Println("POST method")
		jsonByte, err := io.ReadAll(request.Body)
		fmt.Println("check request body", jsonByte)
		if err != nil {
			response.WriteHeader(500)
			response.Write([]byte("Access denind"))
		}
		var user User
		unMarshallError := json.Unmarshal(jsonByte, &user)
		fmt.Println(user, unMarshallError, "check user")
		if unMarshallError != nil {
			response.WriteHeader(500)
			response.Write([]byte("Convert error"))
			return
		}
		jsonMars, _ := json.Marshal(user) // create JSON(byte array)
		response.WriteHeader(201)
		response.Write(jsonMars)
	}
}

func main() {
	http.HandleFunc("/hello", SayHello)

	fmt.Println("HTTP run server")
	http.ListenAndServe("localhost:8080", nil)

}
