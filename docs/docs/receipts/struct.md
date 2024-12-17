
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Struct' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

## 1. Struct add

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

	gouse.Printf("Struct after adding fields: %+v\n", result)
}
```

## 2. Struct clone

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

	gouse.Printf("Original: %+v\n", person)

	clone := gouse.CloneStruct(person)

	updateClone := clone.(Clone_Person)
	updateClone.Name = "Updated Name"
	gouse.Printf("Clone: %+v\n", updateClone)
}
```

## 3. Struct get

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
	gouse.Printf("Name: %s\n", name)
}
```

## 4. Struct has

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
	gouse.Printf("Has: %+v\n", has)

	hasEmpty := gouse.HasEmptyInStruct(person, "Email")
	gouse.Printf("Has empty: %+v\n", hasEmpty)
}
```

## 5. Struct merge

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

	gouse.Printf("Struct after merged: %+v\n", merged)

	gouse.Println("Name:", merged.(map[string]interface{})["Name"])

	gouse.Println("City:", merged.(map[string]interface{})["City"])
}
```

## 6. Struct remove

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

	gouse.Printf("Struct after removed field: %+v\n", gouse.RemoveStruct(person, "Email"))
}
```

## 7. Struct set

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

	gouse.Printf("Struct after setting field: %+v\n", person)
}
```
