
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Struct' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Struct get tags

Description: Get tags from a struct<br>Input params: (struct interface{})<br>

```go
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
```

## 2. Struct add

Description: Add fields to a struct<br>Input params: (struct, newField interface{})<br>

```go
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
```

## 3. Struct clone

Description: Clone a struct<br>Input params: (struct interface{})<br>

```go
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
```

## 4. Struct get

Description: Get fields from a struct<br>Input params: (struct interface{}, fieldName string)<br>

```go
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
```

## 5. Struct has

Description: Check if a struct has a field<br>Input params: (struct interface{}, fieldName string)<br>

```go
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
```

## 6. Struct merge

Description: Merge two structs<br>Input params: (struct1, struct2 interface{})<br>

```go
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
```

## 7. Struct remove

Description: Remove fields from a struct<br>Input params: (struct interface{}, fieldName string)<br>

```go
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
```

## 8. Struct set

Description: Set update a field in a struct<br>Input params: (struct interface{}, fieldName string, value interface{})<br>

```go
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
```
