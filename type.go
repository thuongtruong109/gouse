package gouse

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/google/uuid"
	"github.com/thuongtruong109/gouse/shared"
)

/* Check type of variable */

func IsInt(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "int")
}

func IsUnInt(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "uint")
}

func IsFloat(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "float")
}

func IsComplex(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "complex")
}

func IsNumber(v interface{}) bool {
	return IsInt(v) || IsFloat(v)
}

func IsString(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "string")
}

func IsRune(v interface{}) bool {
	_, ok := v.(rune)
	return ok
}

func IsByte(v interface{}) bool {
	_, ok := v.(byte)
	return ok
}

func IsUintptr(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "uintptr")
}

func IsError(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "error")
}

func IsChannel(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "chan")
}

func IsUnsafePointer(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Uintptr
}

func IsPointer(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "*")
}

func IsBool(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "bool")
}

func IsSlice(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "[]")
}

func IsMap(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "map")
}

func IsStruct(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Struct
}

func IsArray(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "[")
}

func IsFunc(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "func")
}

func IsNil(v interface{}) bool {
	return v == nil

	// 	v := reflect.ValueOf(value)
	// 	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
	// 		return v.IsNil()
	// 	}

	// 	return false
}

func IsEmpty(v interface{}) bool {
	if v == nil {
		return true
	}

	switch value := v.(type) {
	case string:
		return value == ""
	case int, int8, int16, int32, int64:
		return value == 0
	case uint, uint8, uint16, uint32, uint64:
		return value == 0
	case float32, float64:
		return value == 0
	case bool:
		return !value
	case []interface{}:
		return len(value) == 0
	case map[string]interface{}:
		return len(value) == 0
	default:
		return false
	}
}

func IsUndefined(v interface{}) bool {
	return v == nil || IsEmpty(v)
}

func IsZero(v interface{}) bool {
	return v == nil || IsEmpty(v)
}

/* Check type of form */

func IsUUID(input string) (bool, error) {
	if input == "" {
		return false, ERROR_REQUIRED_UUID
	}

	_, err := uuid.Parse(input)
	if err != nil {
		return false, ERROR_INVALID_UUID
	}

	return true, nil
}

func isMail(emailStr, domain string) (bool, error) {
	if Includes(emailStr, "@") {
		if !IsMatchReg(shared.EmailLenReg, emailStr) {
			return false, ERROR_EMAIL_LENGTH
		}

		split := strings.Split(emailStr, "@")
		if len(split) != 2 || !Includes(split[1], ".") {
			return false, ERROR_INVALID_EMAIL
		} else if !Includes(split[1], domain) {
			return false, ERROR_INVALID_EMAIL
		} else {
			return true, nil
		}
	} else {
		return false, ERROR_INVALID_EMAIL
	}
}

func IsGmail(emailStr string) (bool, error) {
	return isMail(emailStr, "gmail.com")
}

func IsYahoo(emailStr string) (bool, error) {
	return isMail(emailStr, "yahoo.com")
}

func IsOutlook(emailStr string) (bool, error) {
	return isMail(emailStr, "outlook.com")
}

func IsEdu(emailStr string) (bool, error) {
	return isMail(emailStr, "edu")
}

func IsEmail(emailStr, customDomain string) (bool, error) {
	return isMail(emailStr, customDomain)
}

func IsUsername(username string) (bool, error) {
	if !IsMatchReg(shared.UsernameReg, username) {
		return false, ERROR_INVALID_USERNAME
	}

	return true, nil
}

func IsPassword(password string) (bool, error) {
	if !IsMatchReg(shared.PasswordLenReg, password) {
		return false, ERROR_PASSWORD_LENGTH
	} else if !IsMatchReg(shared.PasswordLowerReg, password) {
		return false, ERROR_PASSWORD_EMPTY_LOWER
	} else if !IsMatchReg(shared.PasswordUpperReg, password) {
		return false, ERROR_PASSWORD_EMPTY_UPPER
	} else if !IsMatchReg(shared.PasswordDigitReg, password) {
		return false, ERROR_PASSWORD_EMPTY_DIGIT
	} else if !IsMatchReg(shared.PasswordSpecialReg, password) {
		return false, ERROR_PASSWORD_EMPTY_SPECIAL
	} else {
		return true, nil
	}
}

func IsPhone(phone string) (bool, error) {
	if !IsMatchReg(shared.PhoneReg, phone) {
		return false, ERROR_INVALID_PHONE
	}

	return true, nil
}

/* Convert type */

func StructToString(data interface{}) string {
	v := reflect.ValueOf(data)
	t := v.Type()

	var fields []string

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i).Interface()
		fields = append(fields, fmt.Sprintf("%s: %v", fieldName, fieldValue))
	}

	// replace with array.Join()
	var result string
	for i, v := range fields {
		result += v
		if i != len(fields)-1 {
			result += ", "
		}
	}

	return fmt.Sprintf("%s{%s}", t.Name(), result)
}

func StructToMap(data interface{}) map[string]interface{} {
	v := reflect.ValueOf(data)
	t := v.Type()

	result := make(map[string]interface{})

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i).Interface()
		result[fieldName] = fieldValue
	}

	return result
}

func StringToInt(data string) int {
	var result int
	fmt.Sscanf(data, "%d", &result)
	return result
}

func StringToFloat(data string) float64 {
	var result float64
	fmt.Sscanf(data, "%f", &result)
	return result
}

func StringToBool(data string) bool {
	var result bool
	fmt.Sscanf(data, "%t", &result)
	return result
}

func StringToBytes(data string) []byte {
	return []byte(data)
}

func StringsToBytes(data []string) []byte {
	var result []byte
	for _, v := range data {
		result = append(result, StringToBytes(v)...)
	}
	return result
}

// func ErrorToString(data error) string {
// 	return data.Error()
// }

func IntToString(data int) string {
	return fmt.Sprintf("%d", data)
}

func FloatToString(data float64) string {
	return fmt.Sprintf("%f", data)
}

func BoolToString(data bool) string {
	return fmt.Sprintf("%t", data)
}

func ToString(data interface{}) string {
	return fmt.Sprintf("%v", data)
}

func BytesToString(data []byte) string {
	return string(data)
}

func RuneToString(data rune) string {
	return string(data)
}

// MapAsString is not yet in example
func MapAsString[T string | []string](data map[string]T) string {
	var result string

	for key, value := range data {
		result += fmt.Sprintf("%s: %s\n", key, value)
	}

	return result
}

// func formatSlice[T string | []string](values T) string {
// 	return "\"" + fmt.Sprintf("%s", values) + "\""
// }
