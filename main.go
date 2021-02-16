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

func main() {
	var alex person
	alex.print("empty or default person struct\n")

	alex = person{
		firstName: "Alex",
		lastName:  "Jaramillo",
		contactInfo: contactInfo{
			email: "contacto@alexanderjaramillo.com",
		},
	}
	alex.print("before update without pointer\n")

	alex.updateNameCheck("Jhon")
	alex.print("after update without pointer\n")

	alex.updateName("Jhon")
	alex.print("after update using pointer\n")

}

func (p person) print(prefix string) {
	fmt.Printf("%s %+v\n", prefix, p) // plus + to show keys
}

// creates a p person as a new struct on RAM
// then, update last created inside func.
// The source struct expected to update does not change
func (p person) updateNameCheck(newFirstName string) {
	p.firstName = newFirstName
}

// receiver says that p is a pointer to RAM address location,
// and the value of address is person struct type
func (p *person) updateName(newFirstName string) {
	// *p takes the value of RAM address (the person)
	(*p).firstName = newFirstName
}
