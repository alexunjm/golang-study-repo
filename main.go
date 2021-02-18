package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.Version())

	// Itoa doesn't return any errors
	// So, you don't have to handle the errors for it

	s := strconv.Itoa(42)

	fmt.Println(s)

	// Atoi returns an error value
	// So, you should always check it

	// go run main.go 39
	n, err := strconv.Atoi(os.Args[1])

	fmt.Println("Converted number    :", n)
	fmt.Println("Returned error value:", err)

}
