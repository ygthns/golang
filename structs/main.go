package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p person) print() {
	fmt.Println(p)
}

//This is a type description, it means we're working with a pointer to a person.
func (p *person) updateName(newName string) {
	(*p).firstName = newName

}

func main() {

	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	// alex := person{"Alex", "Anderson"}
	// fmt.Println(alex)

	//&variable --> give me the memory address of the value this variable is pointing at
	//*pointer --> give me the value this memory address is pointing at

	jim := person{
		firstName: "Jim",
		lastName:  "Halpert",
		contact: contactInfo{
			email:   "jimhalpert98@gmail.com",
			zipCode: 94000,
		},
	}
	jim.print()
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

//Value Type: int, float, string, bool, structs-->(we need to use pointers for these)
//Reference Types: slices, maps, channels, pointers, functions-->(don't worry about pointers)
