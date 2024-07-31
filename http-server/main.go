package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(responeWriter http.ResponseWriter, request *http.Request) {
	if (request.Method == "GET") {
		fmt.Printf("got / request\n")	
		io.WriteString(responeWriter, "This is my website!\n")	
	} else {
		responseString := fmt.Sprintf("Unsupported Http Method called: [%s] \n", request.Method)
		fmt.Println(responseString)
		io.WriteString(responeWriter, responseString)
	}
}

func getHello(responeWriter http.ResponseWriter, request *http.Request) {
	if (request.Method == "GET") {
		fmt.Printf("got /hello request\n")
		io.WriteString(responeWriter, "Hello, HTTP!\n")
	} else {
		responseString := fmt.Sprintf("Unsupported Http Method called: [%s] \n", request.Method)
		fmt.Println(responseString)
		io.WriteString(responeWriter, responseString)
	}
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	// Declares the variable err and assigns the result of http.ListenAndServe(":8080", nil) to it.
	err := http.ListenAndServe(":8080", nil)

  	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server on port 8080: %s\n", err)
		os.Exit(1)
	}
}
