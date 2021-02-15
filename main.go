package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

/*
func main() {
	alex := person{
		firstName:"Alex",
		lastName:"Jaramillo",
	}
	fmt.Println(alex)
} */
func main() {
	var alex person
	alex.print()

	alex = person{
		firstName: "Alex",
		lastName:  "Jaramillo",
		contactInfo: contactInfo{
			email: "pending",
		},
	}

	alex.print()

	alex.lastName = "Muñoz"
	alex.contactInfo.email = "contacto@alexanderjaramillo.com"
	alex.contactInfo.zip = 040062
}

func (p person) print() {
	fmt.Printf("%+v\n", p) // plus + to show keys
}
