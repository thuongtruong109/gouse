package gouse

import "reflect"

func GetTag(structInstance any) []string {
	structType := reflect.TypeOf(structInstance)
	var tags []string

	for i := range structType.NumField() {
		f := structType.Field(i)
		tags = append(tags, f.Tag.Get("json"))
	}

	return tags
}

func MergeStruct(structs ...any) any {
	if len(structs) < 2 {
		panic("At least two structs are required to merge")
	}

	fieldValues := make(map[string]reflect.Value)

	for _, s := range structs {
		v := reflect.ValueOf(s)
		t := v.Type()

		for i := range v.NumField() {
			fieldName := t.Field(i).Name
			fieldValues[fieldName] = v.Field(i)
		}
	}

	resultMap := make(map[string]any)
	for fieldName, fieldValue := range fieldValues {
		resultMap[fieldName] = fieldValue.Interface()
	}

	return resultMap
}

func RmStruct(structInstance any, fields ...string) any {
	structValue := reflect.ValueOf(structInstance)
	structType := structValue.Type()

	excludedFields := make(map[string]bool)
	for _, field := range fields {
		excludedFields[field] = true
	}

	var newFields []reflect.StructField
	for i := range structValue.NumField() {
		fieldName := structType.Field(i).Name
		if !excludedFields[fieldName] {
			newFields = append(newFields, structType.Field(i))
		}
	}

	newStructType := reflect.StructOf(newFields)
	newStructValue := reflect.New(newStructType).Elem()

	for _, field := range newFields {
		fieldName := field.Name
		value := structValue.FieldByName(fieldName).Interface()
		newStructValue.FieldByName(fieldName).Set(reflect.ValueOf(value))
	}

	return newStructValue.Interface()
}

func AddStruct(structInstance any, fields map[string]any) any {
	structType := reflect.TypeOf(structInstance)

	existingValues := make(map[string]any)
	for i := range structType.NumField() {
		fieldName := structType.Field(i).Name
		existingValues[fieldName] = reflect.ValueOf(structInstance).Field(i).Interface()
	}

	for fieldName, value := range fields {
		existingValues[fieldName] = value
	}

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

	newStructValue := reflect.New(newStructType).Elem()

	for i := range newStructType.NumField() {
		fieldName := newStructType.Field(i).Name
		fieldValue := existingValues[fieldName]
		newStructValue.Field(i).Set(reflect.ValueOf(fieldValue))
	}

	return newStructValue.Interface()
}

func SetStruct(structInstance any, fieldName string, value any) {
	structValue := reflect.ValueOf(structInstance)

	if structValue.Kind() != reflect.Ptr || structValue.IsNil() {
		Println("Struct instance must be a non-nil pointer.")
		return
	}

	// Dereference the pointer to get the actual struct value
	structValue = structValue.Elem()

	field := structValue.FieldByName(fieldName)

	if !field.IsValid() || !field.CanSet() {
		Printf("Field %s is unexported or not found.\n", fieldName)
		return
	}

	if reflect.ValueOf(value).Type().AssignableTo(field.Type()) {
		field.Set(reflect.ValueOf(value))
	} else {
		Printf("Field %s type mismatch: expected %s, got %s\n", fieldName, field.Type(), reflect.ValueOf(value).Type())
	}
}

func GetStruct(structInstance any, fieldName string) any {
	structValue := reflect.ValueOf(structInstance)

	fieldValue := structValue.FieldByName(fieldName)

	return fieldValue.Interface()
}

func CloneStruct(structInstance any) any {
	structValue := reflect.ValueOf(structInstance)

	newStructValue := reflect.New(structValue.Type()).Elem()

	for i := range structValue.NumField() {
		fieldValue := structValue.Field(i)
		newStructValue.Field(i).Set(fieldValue)
	}

	return newStructValue.Interface()
}

func HasInStruct(structInstance any, fieldName string) bool {
	structType := reflect.TypeOf(structInstance)
	_, ok := structType.FieldByName(fieldName)
	return ok
}

func HasEmptyInStruct(structInstance any, fieldName string) bool {
	structValue := reflect.ValueOf(structInstance)
	fieldValue := structValue.FieldByName(fieldName)

	if !fieldValue.IsValid() {
		return true
	}

	return reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface())
}
