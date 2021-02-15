package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {
	alex := person{"Alex", "Jaramillo"}
	fmt.Println(alex)
}