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

// NOTE: packages
// func main() {
// 	log.Println("Hello")
//
// 	var myVar helpers.SomeType
// 	myVar.TypeName = "Some name"
// 	log.Println(myVar.TypeName)
// }

// NOTE: channels

// const numPool = 1000
//
// func CalculateValue(intChan chan int) {
// 	randomNumber := helpers.RandomNumber(numPool)
//
// 	intChan <- randomNumber
// }
//
// func main() {
// 	intChan := make(chan int)
//
// 	defer close(intChan)
//
// 	go CalculateValue(intChan)
//
// 	num := <-intChan
//
// 	log.Println(num)
// }

// NOTE: Json
// type Person struct {
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// 	HairColor string `json:"hair_color"`
// 	HasDog    bool   `json:"has_dog"`
// }
//
// func main() {
// 	myJson := `
// [
//   {
//     "first_name": "Clark",
//     "last_name": "Kent",
//     "hair_color": "black",
//     "has_dog": true
//   },
//   {
//     "first_name": "Bruce",
//     "last_name": "Wayne",
//     "hair_color": "black",
//     "has_dog": false
//   }
// ]`
//
// 	var persons []Person
//
// 	err := json.Unmarshal([]byte(myJson), &persons)
// 	if err != nil {
// 		log.Println("Error unmarshalling json", err)
// 	}
//
// 	log.Printf("persons: %v", persons)
//
// 	// write json from a struct
//
// 	var heroes []Person
//
// 	wally := Person{
// 		FirstName: "Wally",
// 		LastName:  "West",
// 		HairColor: "red",
// 		HasDog:    false,
// 	}
//
// 	diana := Person{
// 		FirstName: "Diana",
// 		LastName:  "Prince",
// 		HairColor: "black",
// 		HasDog:    false,
// 	}
//
// 	heroes = append(heroes, wally)
// 	heroes = append(heroes, diana)
//
// 	newJson, err := json.MarshalIndent(heroes, "", "  ")
// 	if err != nil {
// 		log.Println("Error marshaling", err)
// 	}
//
// 	fmt.Println("newJSon:", string(newJson))
// }

// NOTE: Testing
func main() {
	result, err := divide(100.0, 0.0)
	if err != nil {
		log.Println(err)

		return
	}

	log.Println("Result of division is", result)
}

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("cannot divide by 0")
	}

	result = x / y

	return result, nil
}
