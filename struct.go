package gouse

import "reflect"

func GetTagName(structInstance interface{}) []string {
	structType := reflect.TypeOf(structInstance)
	var tags []string

	for i := 0; i < structType.NumField(); i++ {
		f := structType.Field(i)
		tags = append(tags, f.Tag.Get("json"))
	}

	return tags
}

// Usage:

// type User struct {
// 	Id   int    `json:"id"`
// 	Name string `json:"name"`
// }

// func main() {
// 	u := User{1, "Tom"}
// 	fmt.Println(GetTagName(u))
// }

// func Merge(structs ...interface{}) interface{} {
// 	if len(structs) < 2 {
// 		panic("At least two structs are required to merge")
// 	}

// 	fieldValues := make(map[string]interface{})

// 	// Iterate through the input structs
// 	for _, structInstance := range structs {
// 		structValue := reflect.ValueOf(structInstance)

// 		// Iterate through the fields of the struct
// 		for i := 0; i < structValue.NumField(); i++ {
// 			field := structValue.Type().Field(i)
// 			fieldName := field.Name

// 			// Store field value in the map
// 			fieldValues[fieldName] = structValue.Field(i).Interface()
// 		}
// 	}

// 	// Create a slice of reflect.StructField for the new struct
// 	var newFields []reflect.StructField

// 	// Iterate through the fields of the new struct
// 	for fieldName, value := range fieldValues {
// 		// Add fields to the new struct dynamically
// 		newFields = append(newFields, reflect.StructField{
// 			Name: fieldName,
// 			Type: reflect.TypeOf(value),
// 		})
// 	}

// 	// Create a new struct type
// 	newStructType := reflect.StructOf(newFields)

// 	// Create a new instance of the struct
// 	newStructValue := reflect.New(newStructType).Elem()

// 	// Set field values in the new struct
// 	for _, field := range newFields {
// 		fieldName := field.Name
// 		value := fieldValues[fieldName]
// 		newStructValue.FieldByName(fieldName).Set(reflect.ValueOf(value))
// 	}

// 	return newStructValue.Interface()
// }

func MergeStruct(structs ...interface{}) interface{} {
	if len(structs) < 2 {
		panic("At least two structs are required to merge")
	}

	// Create a map to store field values
	fieldValues := make(map[string]reflect.Value)

	// Iterate over each struct
	for _, s := range structs {
		v := reflect.ValueOf(s)
		t := v.Type()

		// Iterate over fields of each struct
		for i := 0; i < v.NumField(); i++ {
			fieldName := t.Field(i).Name
			fieldValues[fieldName] = v.Field(i)
		}
	}

	// Create a new map with merged fields
	resultMap := make(map[string]interface{})
	for fieldName, fieldValue := range fieldValues {
		resultMap[fieldName] = fieldValue.Interface()
	}

	return resultMap
}

func RemoveStruct(structInstance interface{}, fields ...string) interface{} {
	structValue := reflect.ValueOf(structInstance)
	structType := structValue.Type()

	// Create a map to store excluded field names
	excludedFields := make(map[string]bool)
	for _, field := range fields {
		excludedFields[field] = true
	}

	// Create a new list of struct fields excluding the specified fields
	var newFields []reflect.StructField
	for i := 0; i < structValue.NumField(); i++ {
		fieldName := structType.Field(i).Name
		if !excludedFields[fieldName] {
			newFields = append(newFields, structType.Field(i))
		}
	}

	newStructType := reflect.StructOf(newFields)
	newStructValue := reflect.New(newStructType).Elem()

	// Copy values from the original struct to the new struct
	for _, field := range newFields {
		fieldName := field.Name
		value := structValue.FieldByName(fieldName).Interface()
		newStructValue.FieldByName(fieldName).Set(reflect.ValueOf(value))
	}

	return newStructValue.Interface()
}

func AddStruct(structInstance interface{}, fields map[string]interface{}) interface{} {
	structType := reflect.TypeOf(structInstance)

	// Create a map to store existing field values
	existingValues := make(map[string]interface{})
	for i := 0; i < structType.NumField(); i++ {
		fieldName := structType.Field(i).Name
		existingValues[fieldName] = reflect.ValueOf(structInstance).Field(i).Interface()
	}

	// Add new fields to the map
	for fieldName, value := range fields {
		existingValues[fieldName] = value
	}

	// Create a new struct type with the combined fields
	newStructType := reflect.StructOf(
		func() []reflect.StructField {
			var fields []reflect.StructField
			for fieldName, value := range existingValues {
				fields = append(fields, reflect.StructField{
					Name: fieldName,
					Type: reflect.TypeOf(value),
				})
			}
			return fields
		}(),
	)

	// Create a new struct instance of the new type
	newStructValue := reflect.New(newStructType).Elem()

	// Set field values for the new instance
	for i := 0; i < newStructType.NumField(); i++ {
		fieldName := newStructType.Field(i).Name
		fieldValue := existingValues[fieldName]
		newStructValue.Field(i).Set(reflect.ValueOf(fieldValue))
	}

	return newStructValue.Interface()
}

func SetStruct(structInstance interface{}, fieldName string, value interface{}) {
	// Get the reflect.Value of the struct instance
	structValue := reflect.ValueOf(structInstance)

	// Check if the structInstance is a pointer
	if structValue.Kind() != reflect.Ptr || structValue.IsNil() {
		Println("Struct instance must be a non-nil pointer.")
		return
	}

	// Dereference the pointer to get the actual struct value
	structValue = structValue.Elem()

	// Get the field by name
	field := structValue.FieldByName(fieldName)

	// Check if the field is valid and exported
	if !field.IsValid() || !field.CanSet() {
		Printf("Field %s is unexported or not found.\n", fieldName)
		return
	}

	// Check if the types match before setting the value
	if reflect.ValueOf(value).Type().AssignableTo(field.Type()) {
		field.Set(reflect.ValueOf(value))
	} else {
		Printf("Field %s type mismatch: expected %s, got %s\n", fieldName, field.Type(), reflect.ValueOf(value).Type())
	}
}

func GetStruct(structInstance interface{}, fieldName string) interface{} {
	structValue := reflect.ValueOf(structInstance)

	// Get the field value
	fieldValue := structValue.FieldByName(fieldName)

	return fieldValue.Interface()
}

func CloneStruct(structInstance interface{}) interface{} {
	structValue := reflect.ValueOf(structInstance)

	// Create a new struct instance
	newStructValue := reflect.New(structValue.Type()).Elem()

	// Copy values from the original struct to the new struct
	for i := 0; i < structValue.NumField(); i++ {
		fieldValue := structValue.Field(i)
		newStructValue.Field(i).Set(fieldValue)
	}

	return newStructValue.Interface()
}

func HasInStruct(structInstance interface{}, fieldName string) bool {
	structType := reflect.TypeOf(structInstance)
	_, ok := structType.FieldByName(fieldName)
	return ok
}

func HasEmptyInStruct(structInstance interface{}, fieldName string) bool {
	structValue := reflect.ValueOf(structInstance)
	fieldValue := structValue.FieldByName(fieldName)

	if !fieldValue.IsValid() {
		return true
	}

	return reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface())
}
