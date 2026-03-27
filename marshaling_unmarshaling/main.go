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

// type Address struct {
// 	Street string `json:"street"`
// 	City   string `json:"city"`
// }

// type Person struct {
// 	Name    string  `json:"name"`
// 	Age     int     `json:"age"`
// 	Address Address `json:"address"` // ✅ nested struct
// }

// type MyTime struct {
// 	time.Time
// }

// // Custom marshal — output as Unix timestamp
// func (t MyTime) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(t.Unix())
// }

// // Custom unmarshal — read Unix timestamp back
// func (t *MyTime) UnmarshalJSON(data []byte) error {
// 	var unix int64
// 	if err := json.Unmarshal(data, &unix); err != nil {
// 		return err
// 	}
// 	t.Time = time.Unix(unix, 0)
// 	return nil
// }

type Book struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

type Shoes struct {
	Brand   string `json:"brand"`
	Size    int    `json:"size"`
	InStock bool   `json:"in_stock"`
}

type Movie struct{
	Title string `json:"movie_title"`
	Year  string  `json:"release_year"`
	Rating string  `json:"rating,omitempty"`
}

type Students struct{
	Name string `json:"name"`
	Grade string `json:"grade"`
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

	// jsonString := `{"name":"Rahul","age":25,"address":{"street":"MG Road","city":"Hyderabad"}}`
	// var person Person
	// json.Unmarshal([]byte(jsonString), &person)
	// fmt.Println(person) // Hyderabad

	//Q1- Books := Book{Title: "book title",
	// 	Author: "book author",
	// 	Price:  45.89,
	// }

	// bookjson, err := json.Marshal(Books)
	// if err == nil {
	// 	fmt.Println("egerg")
	// }
	// fmt.Println(string(bookjson))

	//Q2- jsonString := `{"brand":"Nike","size":42,"in_stock":true}`
	// var shoes Shoes
	// json.Unmarshal([]byte(jsonString), &shoes)
	// fmt.Println(shoes)

	// Q3 - movie := Movie{
	// 	Title: "Kick",
	// 	Year: "20/02/2002",
	// 	Rating : "5star",
	// }

	// output,_ := json.Marshal(movie)
	// fmt.Println(string(output))

	//q4 jsonString := []Students{
	// 	{Name:"ashutosh",Grade:"A"},
	// 	{Name:"kanha",Grade:"B"},
	// 	{Name:"Jhon",Grade:"c"},
	// }

	// output,_ := json.Marshal(jsonString)
	// fmt.Println(string(output))

}
