package main

import (
	"errors"
	"log"
)

// NOTE: Basics
// func main() {
// 	fmt.Println("Hello World")
//
// 	var whatToSay string
// 	var i int
//
// 	whatToSay = "Goodby, cruel world"
//
// 	fmt.Println(whatToSay)
//
// 	i = 7
//
// 	fmt.Println("i is set to", i)
//
// 	whatWasSaid, theOtherThingThatSaid := saySomething()
//
// 	fmt.Println("The function returned", whatWasSaid, theOtherThingThatSaid)
// }
//
// func saySomething() (string, string) {
// 	return "something", "else"
// }

// NOTE: Pointers
// func main() {
// 	myString := "Green"
//
// 	log.Println("myString is set to", myString)
//
// 	changeUsingPointer(&myString)
//
// 	log.Println("after func call myString is set to", myString)
// }
//
// func changeUsingPointer(s *string) {
// 	log.Println("s is set to", *s)
// 	newValue := "Red"
// 	*s = newValue
// }

// // NOTE: structs
// type User struct {
// 	Age         int
// 	BirthDate   time.Time
// 	FirstName   string
// 	LastName    string
// 	PhoneNumber string
// }
//
// func main() {
// 	user := User{
// 		FirstName:   "Nikola",
// 		LastName:    "Dev",
// 		Age:         22,
// 		PhoneNumber: "3117742523",
// 	}
//
// 	log.Println(user.FirstName, user.LastName, "BirthDate", user.BirthDate)
// }

// NOTE: methods
// type myStruct struct {
// 	FirstName string
// }
//
// func (m *myStruct) printFirstName() string {
// 	return m.FirstName
// }
//
// func main() {
// 	var myVar myStruct
// 	myVar.FirstName = "John"
//
// 	myVar2 := myStruct{
// 		FirstName: "Mary",
// 	}
//
// 	log.Println("myVar is set to", myVar.printFirstName())
// 	log.Println("myVar2 is set to", myVar2.printFirstName())
// }

// NOTE: maps
// type User struct {
// 	FirstName string
// 	LastName  string
// }
//
// func main() {
// 	myMap := make(map[string]User)
//
// 	me := User{
// 		FirstName: "Nikola",
// 		LastName:  "Dev",
// 	}
//
// 	myMap["me"] = me
//
// 	log.Println(myMap["me"])
// }

// NOTE: Slices
// func main() {
// 	names := []string{"steven", "nikola", "cat"}
//
// 	log.Println(names)
// }

// NOTE: conditionals
// func main() {
// 	cat := "czat"
//
// 	if cat == "cat" {
// 		log.Println("Cat is cat")
// 	} else {
// 		log.Println("Cat is not cat")
// 	}
// }
