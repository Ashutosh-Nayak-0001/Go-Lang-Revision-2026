package main

import (
	"encoding/json"
	"fmt"
)

// type Person struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// 	City string `json:"city"`
// }

type Address struct {
    Street string `json:"street"`
    City   string `json:"city"`
}

type Person struct {
    Name    string  `json:"name"`
    Age     int     `json:"age"`
    Address Address `json:"address"` // ✅ nested struct
}

func main() {

	// person := Person{
	// 	 Name :"Ashutosh",
	// 	 Age:25,
	// 	 City :"Hydrabad",
	// }
	// jsonData , err := json.Marshal(person)

	// if err != nil {
	// 	fmt.Println("Error:",err)
	// }

	// fmt.Println(string(jsonData))

// 	person := Person{
//     Name: "Rahul",
//     Age:  25,
//     Address: Address{
//         Street: "MG Road",
//         City:   "Hyderabad",
//     },
// }

// jsonData, _ := json.Marshal(person)
// fmt.Println(string(jsonData))
// {"name":"Rahul","age":25,"address":{"street":"MG Road","city":"Hyderabad"}}

	// jsonString := `{"name":"Rahul","age":25,"city":"Hyderabad"}`

	// var person Person

	// // convert to Go struct
	// err := json.Unmarshal([]byte(jsonString), &person)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// fmt.Println(person.Name) // Rahul
	// fmt.Println(person.Age)  // 25
	// fmt.Println(person.City) // Hyderabad
	// fmt.Println(person)

	jsonString := `{"name":"Rahul","age":25,"address":{"street":"MG Road","city":"Hyderabad"}}`
var person Person
json.Unmarshal([]byte(jsonString), &person)
fmt.Println(person) // Hyderabad

}
