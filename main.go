package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	/*
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
	*/
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Printf("%+v", resp)

	bs := make([]byte, 999)
	resp.Body.Read(bs)
	fmt.Println("\n\n", string(bs))
}
