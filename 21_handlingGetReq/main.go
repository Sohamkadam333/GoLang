package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Handling GET Request")
	PerformGetReq()
}

func PerformGetReq() {
	const myUrl string = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return
	}
	defer response.Body.Close()

	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	fmt.Printf("Response = %s\n", string(dataBytes))
	fmt.Printf("Response content length = %d\n", response.ContentLength)
	fmt.Printf("Response status code = %d\n", response.StatusCode)

	// Read content using Strings Builder

	var responseString strings.Builder
	content, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	byteCount, err := responseString.Write(content)
	if err != nil {
		fmt.Printf("Error writing content to string builder: %v\n", err)
		return
	}

	fmt.Println("ByteCount is:", byteCount)
	fmt.Println(responseString.String())
}
