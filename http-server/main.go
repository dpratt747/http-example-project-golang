package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"encoding/json"
	"strconv"
)

type user_struct struct {
    Name string
	Age int
}

func decodeJsonUserAndRespond(responseWriter http.ResponseWriter, request http.Request) {
	// Example of Json decoding
	decoder := json.NewDecoder(request.Body)

	// Will error if any fields appear in the Json that do not map to the Struct
	decoder.DisallowUnknownFields()

	var user user_struct

	// Stores the result of decode to the value user, if it fails then the errors are stored in the val err
	err := decoder.Decode(&user)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to decode incomming json %s\n", err)
		fmt.Println(errorMessage)
		http.Error(responseWriter, errorMessage, http.StatusUnprocessableEntity)
	} else {
		responseString := fmt.Sprintf("Input user age is [%s], and name is [%s]\n", strconv.Itoa(user.Age), user.Name)
		fmt.Println(responseString)
		io.WriteString(responseWriter, responseString)
	}
}

func postUser(responseWriter http.ResponseWriter, request *http.Request) {

	errorResponseString := fmt.Sprintf("Unsupported Http Method called: [%s] \n", request.Method)

	switch request.Method {
		case http.MethodGet:
			fmt.Println(errorResponseString)
			http.Error(responseWriter, errorResponseString, http.StatusMethodNotAllowed)
		case http.MethodPost:
			fmt.Printf("POST / request\n")	
			decodeJsonUserAndRespond(responseWriter, *request)
		case http.MethodPut:
			fmt.Println(errorResponseString)
			http.Error(responseWriter, errorResponseString, http.StatusMethodNotAllowed)
		case http.MethodDelete:
			fmt.Println(errorResponseString)
			http.Error(responseWriter, errorResponseString, http.StatusMethodNotAllowed)
		default:
			fmt.Println(errorResponseString)
			http.Error(responseWriter, errorResponseString, http.StatusMethodNotAllowed)
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

func getRoot(responseWriter http.ResponseWriter, request *http.Request) {
	if (request.Method == "GET") {
		fmt.Printf("got / request\n")	
		io.WriteString(responseWriter, "This is my website!\n")	
	} else {
		errorResponseString := fmt.Sprintf("Unsupported Http Method called: [%s] \n", request.Method)
		fmt.Println(errorResponseString)
		http.Error(responseWriter, errorResponseString, http.StatusMethodNotAllowed)

	}
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/user", postUser)
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
