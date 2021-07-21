package main

import (
	"fmt"
	"io"
	"net/http"
)

type LogWriter struct {
}

func (lw LogWriter) Write(p []byte) (int, error) {
	fmt.Println("Writing ", len(p))
	return fmt.Println(string(p))
}

func main() {
	resp, _ := http.Get("http://www.google.com")
	// resb := make([]byte, 999999)
	// v, err := resp.Body.Read(resb)
	// fmt.Println(string(resb), v, err)

	var l LogWriter
	io.Copy(l, resp.Body)
}
