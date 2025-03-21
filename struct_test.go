package gouse

import (
	"fmt"
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

type Job struct {
	Title string
}

type Developer struct {
	Name    string
	Age     int
	Address string
	Title   string
}

func TestGetTag(t *testing.T) {
	type Person struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Address string `json:"address"`
	}

	expectedTags := []string{"name", "age", "address"}

	person := Person{
		Name:    "John",
		Age:     30,
		Address: "123 Main St",
	}

	tags := GetTag(person)

	if !Equal(tags, expectedTags) {
		t.Errorf("GetTagName() = %v, want %v", tags, expectedTags)
	}
}

func TestMergeStruct(t *testing.T) {
	tests := []struct {
		name   string
		input  []any
		output any
	}{
		{
			name:   "Merge Person and Job",
			input:  []any{Person{Name: "John", Age: 25, Address: "123 Main St"}, Job{Title: "Developer"}},
			output: map[string]any{"Name": "John", "Age": 25, "Address": "123 Main St", "Title": "Developer"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MergeStruct(test.input...)
			fmt.Println(result)
			fmt.Println(result.(map[string]any)["Name"])
			if !reflect.DeepEqual(result, test.output) {
				t.Errorf("Expected %+v, but got %+v", test.output, result)
			}
		})
	}
}

func TestRemoveStruct(t *testing.T) {
	tests := []struct {
		name           string
		inputStruct    any
		fieldsToRemove []string
		expectedResult any
	}{
		{
			name:           "Remove Address field",
			inputStruct:    Person{Name: "John", Age: 25, Address: "123 Main St"},
			fieldsToRemove: []string{"Address"},
			expectedResult: struct {
				Name string
				Age  int
			}{Name: "John", Age: 25},
		},
		{
			name:           "Remove multiple fields",
			inputStruct:    Person{Name: "Jane", Age: 30, Address: "456 Side St"},
			fieldsToRemove: []string{"Age", "Address"},
			expectedResult: struct{ Name string }{Name: "Jane"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := RmStruct(test.inputStruct, test.fieldsToRemove...)
			expectedResult := reflect.ValueOf(test.expectedResult).Interface()

			if !reflect.DeepEqual(result, expectedResult) {
				t.Errorf("Expected %+v, but got %+v", expectedResult, result)
			}
		})
	}
}

func TestAddStruct(t *testing.T) {
	tests := []struct {
		name           string
		inputStruct    any
		newFields      map[string]any
		expectedResult any
	}{
		{
			name:           "Add Email field",
			inputStruct:    Person{Name: "John", Age: 25, Address: "123 Main St"},
			newFields:      map[string]any{"Email": "john@example.com"},
			expectedResult: map[string]any{"Name": "John", "Age": 25, "Address": "123 Main St", "Email": "john@example.com"},
		},
		{
			name:           "Add multiple fields",
			inputStruct:    Person{Name: "Jane", Age: 30, Address: "456 Side St"},
			newFields:      map[string]any{"Email": "jane@example.com", "Phone": "555-1234"},
			expectedResult: map[string]any{"Name": "Jane", "Age": 30, "Address": "456 Side St", "Email": "jane@example.com", "Phone": "555-1234"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := AddStruct(test.inputStruct, test.newFields)

			resultMap := make(map[string]any)
			for _, field := range reflect.VisibleFields(reflect.TypeOf(result)) {
				resultMap[field.Name] = reflect.ValueOf(result).FieldByName(field.Name).Interface()
			}

			expectedResultMap := test.expectedResult.(map[string]any)

			if !reflect.DeepEqual(resultMap, expectedResultMap) {
				t.Errorf("Expected %+v, but got %+v", expectedResultMap, resultMap)
			}
		})
	}
}

func TestSetStruct(t *testing.T) {
	tests := []struct {
		name           string
		inputStruct    any
		fieldName      string
		value          any
		expectedResult any
	}{
		{
			name:           "Set Name field",
			inputStruct:    &Person{Name: "John", Age: 25, Address: "123 Main St"},
			fieldName:      "Name",
			value:          "Jane",
			expectedResult: &Person{Name: "Jane", Age: 25, Address: "123 Main St"},
		},
		{
			name:           "Set Age field",
			inputStruct:    &Person{Name: "Bob", Age: 30, Address: "456 Side St"},
			fieldName:      "Age",
			value:          35,
			expectedResult: &Person{Name: "Bob", Age: 35, Address: "456 Side St"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			SetStruct(test.inputStruct, test.fieldName, test.value)
			result := test.inputStruct

			if !reflect.DeepEqual(result, test.expectedResult) {
				t.Errorf("Expected %+v, but got %+v", test.expectedResult, result)
			}
		})
	}
}

func TestGetStruct(t *testing.T) {
	tests := []struct {
		name           string
		inputStruct    any
		fieldName      string
		expectedResult any
	}{
		{
			name:           "Get Name field",
			inputStruct:    Person{Name: "John", Age: 25, Address: "123 Main St"},
			fieldName:      "Name",
			expectedResult: "John",
		},
		{
			name:           "Get Age field",
			inputStruct:    Person{Name: "Bob", Age: 30, Address: "456 Side St"},
			fieldName:      "Age",
			expectedResult: 30,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GetStruct(test.inputStruct, test.fieldName)

			if !reflect.DeepEqual(result, test.expectedResult) {
				t.Errorf("Expected %+v, but got %+v", test.expectedResult, result)
			}
		})
	}
}

func TestCloneStruct(t *testing.T) {
	tests := []struct {
		name           string
		inputStruct    any
		expectedResult any
	}{
		{
			name:           "Clone Person struct",
			inputStruct:    Person{Name: "John", Age: 25, Address: "123 Main St"},
			expectedResult: Person{Name: "John", Age: 25, Address: "123 Main St"},
		},
		{
			name:           "Clone with different values",
			inputStruct:    Person{Name: "Alice", Age: 30, Address: "456 Side St"},
			expectedResult: Person{Name: "Alice", Age: 30, Address: "456 Side St"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CloneStruct(test.inputStruct)
			expectedResult := test.expectedResult

			if !reflect.DeepEqual(result, expectedResult) {
				t.Errorf("Expected %+v, but got %+v", expectedResult, result)
			}
		})
	}
}

func TestHasInStruct(t *testing.T) {
	tests := []struct {
		name           string
		inputStruct    any
		fieldName      string
		expectedResult bool
	}{
		{
			name:           "Has Name field",
			inputStruct:    Person{Name: "John", Age: 25, Address: "123 Main St"},
			fieldName:      "Name",
			expectedResult: true,
		},
		{
			name:           "Does not have Email field",
			inputStruct:    Person{Name: "Alice", Age: 30, Address: "456 Side St"},
			fieldName:      "Email",
			expectedResult: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HasInStruct(test.inputStruct, test.fieldName)
			expectedResult := test.expectedResult

			if result != expectedResult {
				t.Errorf("Expected %v, but got %v", expectedResult, result)
			}
		})
	}
}

func TestHasEmptyInStruct(t *testing.T) {
	tests := []struct {
		name           string
		inputStruct    any
		fieldName      string
		expectedResult bool
	}{
		{
			name:           "Name field is not empty",
			inputStruct:    Person{Name: "John", Age: 25, Address: "123 Main St"},
			fieldName:      "Name",
			expectedResult: false,
		},
		{
			name:           "Email field is empty",
			inputStruct:    Person{Name: "Alice", Age: 30, Address: "456 Side St"},
			fieldName:      "Email",
			expectedResult: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HasEmptyInStruct(test.inputStruct, test.fieldName)
			expectedResult := test.expectedResult

			if result != expectedResult {
				t.Errorf("Expected %v, but got %v", expectedResult, result)
			}
		})
	}
}
