package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
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
	fmt.Println(alex)
	fmt.Printf("%+v\n\n", alex) // plus + to show keys

	alex = person{
		firstName: "Alex",
		lastName:  "Jaramillo",
		contact: contactInfo{
			email: "",
		},
	}

	// fmt.Println(alex)
	fmt.Printf("%v\n", alex)

	alex.lastName = "Muñoz"
	alex.contact.email = "contacto@alexanderjaramillo.com"
	alex.contact.zip = 040062
	fmt.Printf("%+v\n", alex) // plus + to show keys
}
