package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWritter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// fmt.Printf("%+v", resp)

	ioReader := resp.Body
	// ioWriter := os.Stdout
	ioWriter := logWritter{}

	written, err := io.Copy(ioWriter, ioReader)
	fmt.Printf("\n\n%+v", written)

}

func (logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
