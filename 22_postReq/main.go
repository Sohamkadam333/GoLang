package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	const myUrl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`{
		"name":"soham",
		"age":23
	}`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		fmt.Println("Error ", err)
		return
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println("Content : ", string(content))

}
