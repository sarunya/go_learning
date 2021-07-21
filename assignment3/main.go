package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	readFromFile(args[1])
}

func readFromFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Caught error on opening file ", err)
		return
	}
	resp, err := io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Println("Caught error on opening file ", err)
		return
	}
	fmt.Println(resp)
}
