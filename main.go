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
