package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Posting Form")

	const myUrl = "http://localhost:8000/postform"

	data := url.Values{}

	data.Add("Name", "Soham")
	data.Add("Age", "23")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		fmt.Println("Error ", err)
		return
	}

	content, _ := io.ReadAll(response.Body)

	fmt.Println("Content = ", string(content))
}
