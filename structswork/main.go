
package main

import "fmt"

// When passing in a slice to a receiver, and a particular value of the slice is updated, since a slice data structure
// is really a pointer to the actual array values, the ACTUAL value is updated (not by reference, but by value). This is
// one of the many gotchas with golang pointers (lookup value type vs reference type data structures)
//
// Value Types:
//             - int
//             - float
//             - string
//             - bool
//             - struct
//
// Reference Types:
//             - slice
//             - map
//             - channel
//             - pointer
//             - function

type contactInfo struct {
  email string
  zipcode int
}

type person struct {
  firstName string
  lastName string
  contactInfo // You do not need to specify the field name "contactInfo" here if you don't want to, will apply it by default
}

func (p person) print() {
  fmt.Printf("%+v ", p)
}

// receiver function is for a memory address of a person struct
func (pointerToPerson *person) updateName(newFirstName string, newLastName string) {
  // Parenthesis and Asterisk operator below is saying basically "give me the value this memory address is pointing at"
  (*pointerToPerson).firstName = newFirstName
  (*pointerToPerson).lastName = newLastName
}

func main()  {
  // Three different ways of declaring and instantiating structs
  alex := person{firstName: "Alex", lastName: "Anderson", contactInfo: contactInfo{email: "aanderson@mail.com", zipcode: 12345}}
  steve := person{"Steve", "Stevenson", contactInfo{"sstevenson@mail.com", 54321}}
  var jacob person

  fmt.Println(alex)
  fmt.Println(steve)

  fmt.Println(jacob)
  fmt.Printf("%+v \n", jacob)

  // Ampersand operator below is saying basically "give me the memory address of the value this variable is pointing to"
  // This however is the explicit way of defining a pointer. In the below syntax, golang is automatically turning it into
  // a memory address for you and assumes that is what you want...
  // jacobPointer := &jacob
  // jacobPointer.updateName("Jacob", "Jacobson")
  jacob.updateName("Jacob", "Jacobson")
  jacob.contactInfo = contactInfo{"jjacobson@mail.com", 15243}
  jacob.print()
}
