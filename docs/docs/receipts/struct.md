
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Struct' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Struct add



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

## 2. Struct clone



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

## 3. Struct get



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

## 4. Struct has



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

## 5. Struct merge



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

## 6. Struct remove



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

	fmt.Printf("Struct after removed field: %+v\n", gouse.RemoveStruct(person, "Email"))
}
```

## 7. Struct set



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
