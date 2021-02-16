package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	/*
		resp, err := http.Get("http://google.com")
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Printf("%+v", resp)
		ioReadCloser := resp.Body */
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
