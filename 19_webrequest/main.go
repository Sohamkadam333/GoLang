package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// const url = "http://npav.net/"

func main() {
	fmt.Println("Handling First Webrequest")

	args := os.Args // get all args

	fmt.Println(args[1])

	response, err := http.Get(args[1])

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() // callers responsibility to close the connection

	dataBytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response = %v\n", string(dataBytes))

	fmt.Scanln()

}
