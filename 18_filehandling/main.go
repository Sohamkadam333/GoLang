package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("First File handling Code")

	File, err := os.Create("./Demo.txt")

	if err != nil {
		panic(err)
	}

	_, err = io.WriteString(File, "Hello Demo File")
	checkNilErr(err)

	defer File.Close()

	readFile("./Demo.txt")
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filePath string) {
	dataByte, err := ioutil.ReadFile(filePath)
	checkNilErr(err)

	fmt.Printf("File Data :\n%v", string(dataByte))
}
