package samples

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
)

/*
Description: Get tags from a struct
Input params: (struct interface{})
*/
func StructGetTags() {
	type Person struct {
		Name     string `json:"name"`
		Age      int    `json:"age"`
		Location string `json:"location"`
	}

	p := Person{
		Name:     "Alice",
		Age:      30,
		Location: "Wonderland",
	}

	tags := gouse.GetTags(p)

	fmt.Println("JSON Tags:", tags)
}

/*
Description: Add fields to a struct
Input params: (struct, newField interface{})
*/
func StructAdd() {
	type Add_Person struct {
		Name  string
		Age   int
		Email string
	}

	person := Add_Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	newFields := map[string]interface{}{
		"Address": "123 Main St",
		"Phone":   "555-1234",
	}
	result := gouse.AddStruct(person, newFields)

	fmt.Printf("Struct after adding fields: %+v\n", result)
}

/*
Description: Clone a struct
Input params: (struct interface{})
*/
func StructClone() {
	type Clone_Person struct {
		Name  string
		Age   int
		Email string
	}

	person := Clone_Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	fmt.Printf("Original: %+v\n", person)

	clone := gouse.CloneStruct(person)

	updateClone := clone.(Clone_Person)
	updateClone.Name = "Updated Name"
	fmt.Printf("Clone: %+v\n", updateClone)
}

/*
Description: Get fields from a struct
Input params: (struct interface{}, fieldName string)
*/
func StructGet() {
	type Get_Person struct {
		Name  string
		Age   int
		Email string
	}

	person := Get_Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	name := gouse.GetStruct(person, "Name")
	fmt.Printf("Name: %s\n", name)
}

/*
Description: Check if a struct has a field
Input params: (struct interface{}, fieldName string)
*/
func StructHas() {
	type Has_Person struct {
		Name  string
		Age   int
		Email string
	}

	person := Has_Person{
		Name:  "Example",
		Age:   40,
		Email: "",
	}

	has := gouse.HasInStruct(person, "Email")
	fmt.Printf("Has: %+v\n", has)

	hasEmpty := gouse.HasEmptyInStruct(person, "Email")
	fmt.Printf("Has empty: %+v\n", hasEmpty)
}

/*
Description: Merge two structs
Input params: (struct1, struct2 interface{})
*/
func StructMerge() {
	type Merge_Person struct {
		Name  string
		Age   int
		Email string
	}

	type Merge_Address struct {
		City    string
		Street  string
		ZipCode string
	}

	person := Merge_Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	address := Merge_Address{
		City:    "New York",
		Street:  "123 Main St",
		ZipCode: "10001",
	}

	merged := gouse.MergeStruct(person, address)

	fmt.Printf("Struct after merged: %+v\n", merged)

	fmt.Println("Name:", merged.(map[string]interface{})["Name"])

	fmt.Println("City:", merged.(map[string]interface{})["City"])
}

/*
Description: Remove fields from a struct
Input params: (struct interface{}, fieldName string)
*/
func StructRemove() {
	type Remove_Person struct {
		Name  string
		Age   int
		Email string
	}

	person := Remove_Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	fmt.Printf("Struct after removed field: %+v\n", gouse.RmStruct(person, "Email"))
}

/*
Description: Set update a field in a struct
Input params: (struct interface{}, fieldName string, value interface{})
*/
func StructSet() {
	type Set_Person struct {
		Name  string
		Age   int
		Email string
	}

	person := &Set_Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	gouse.SetStruct(person, "Name", "Updated Name")

	fmt.Printf("Struct after setting field: %+v\n", person)
}
