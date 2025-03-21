package gouse

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func IsInt(v any) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "int")
}

func IsUnInt(v any) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "uint")
}

func IsFloat(v any) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "float")
}

func IsComplex(v any) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "complex")
}

func IsNumber(v any) bool {
	return IsInt(v) || IsFloat(v)
}

func IsString(v any) bool {
	return strings.Contains(Sprintf("%T", v), "string")
}

func IsRune(v any) bool {
	_, ok := v.(rune)
	return ok
}

func IsByte(v any) bool {
	_, ok := v.(byte)
	return ok
}

// func IsUintptr(v any) bool {
// 	return strings.Contains(Sprintf("%T", v), "uintptr")
// }

func IsUintptr(v any) bool {
	_, ok := v.(uintptr)
	return ok
}

func IsError(v any) bool {
	return strings.Contains(Sprintf("%T", v), "error")
}

func IsChannel(v any) bool {
	return strings.Contains(Sprintf("%T", v), "chan")
}

func IsUnsafePointer(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Uintptr
}

func IsPointer(v any) bool {
	return strings.Contains(Sprintf("%T", v), "*")
}

func IsBool(v any) bool {
	return strings.Contains(Sprintf("%T", v), "bool")
}

func IsSlice(v any) bool {
	return strings.Contains(Sprintf("%T", v), "[]")
}

func IsMap(v any) bool {
	return strings.Contains(Sprintf("%T", v), "map")
}

func IsStruct(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Struct
}

func IsArray(v any) bool {
	return strings.Contains(Sprintf("%T", v), "[")
}

func IsFunc(v any) bool {
	return strings.Contains(Sprintf("%T", v), "func")
}

func IsNil(v any) bool {
	return v == nil

	// 	v := reflect.ValueOf(value)
	// 	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
	// 		return v.IsNil()
	// 	}

	// 	return false
}

func IsEmpty(v any) bool {
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
	case []any:
		return len(value) == 0
	case map[string]any:
		return len(value) == 0
	default:
		return false
	}
}

func IsUndefined(v any) bool {
	return v == nil || IsEmpty(v)
}

func IsZero(v any) bool {
	return v == nil || IsEmpty(v)
}

func IsUUID(input string) (bool, error) {
	if input == "" {
		return false, ErrorRequiredUuid
	}

	_, err := uuid.Parse(input)
	if err != nil {
		return false, ErrorInvalidUuid
	}

	return true, nil
}

func _isMail(emailStr, domain string) (bool, error) {
	if Includes(emailStr, "@") {
		if !IsMatchReg(EmailLenReg, emailStr) {
			return false, ErrorEmailLength
		}

		split := strings.Split(emailStr, "@")
		if len(split) != 2 || !Includes(split[1], ".") {
			return false, ErrorInvalidEmail
		} else if !Includes(split[1], domain) {
			return false, ErrorInvalidEmail
		} else {
			return true, nil
		}
	} else {
		return false, ErrorInvalidEmail
	}
}

func IsGmail(emailStr string) (bool, error) {
	return _isMail(emailStr, "gmail.com")
}

func IsYahoo(emailStr string) (bool, error) {
	return _isMail(emailStr, "yahoo.com")
}

func IsOutlook(emailStr string) (bool, error) {
	return _isMail(emailStr, "outlook.com")
}

func IsEdu(emailStr string) (bool, error) {
	return _isMail(emailStr, "edu")
}

func IsEmail(emailStr, customDomain string) (bool, error) {
	return _isMail(emailStr, customDomain)
}

func IsUsername(username string) (bool, error) {
	if !IsMatchReg(UsernameReg, username) {
		return false, ErrorInvalidUsername
	}

	return true, nil
}

func IsPassword(password string) (bool, error) {
	if !IsMatchReg(PasswordLenReg, password) {
		return false, ErrorPasswordLength
	} else if !IsMatchReg(PasswordLowerReg, password) {
		return false, ErrorPasswordEmptyLower
	} else if !IsMatchReg(PasswordUpperReg, password) {
		return false, ErrorPasswordEmptyUpper
	} else if !IsMatchReg(PasswordDigitReg, password) {
		return false, ErrorPasswordEmptyDigit
	} else if !IsMatchReg(PasswordSpecialReg, password) {
		return false, ErrorPasswordEmptySpecial
	} else {
		return true, nil
	}
}

func IsPhone(phone string) (bool, error) {
	if !IsMatchReg(PhoneReg, phone) {
		return false, ErrorInvalidPhone
	}

	return true, nil
}

func Struct2Str(data any) string {
	v := reflect.ValueOf(data)
	t := v.Type()

	var fields []string

	for i := range v.NumField() {
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i).Interface()
		fields = append(fields, Sprintf("%s: %v", fieldName, fieldValue))
	}

	// replace with array.Join()
	var result string
	for i, v := range fields {
		result += v
		if i != len(fields)-1 {
			result += ", "
		}
	}

	return Sprintf("%s{%s}", t.Name(), result)
}

func Struct2Map(data any) map[string]any {
	v := reflect.ValueOf(data)
	t := v.Type()

	result := make(map[string]any)

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i).Interface()
		result[fieldName] = fieldValue
	}

	return result
}

func Str2Int(data string) int {
	var result int
	fmt.Sscanf(data, "%d", &result)
	return result
}

func Str2Float(data string) float64 {
	var result float64
	fmt.Sscanf(data, "%f", &result)
	return result
}

func Str2Bool(data string) bool {
	var result bool
	fmt.Sscanf(data, "%t", &result)
	return result
}

func Str2Bytes(data string) []byte {
	return []byte(data)
}

func Strs2Bytes(data []string) []byte {
	var result []byte
	for _, v := range data {
		result = append(result, Str2Bytes(v)...)
	}
	return result
}

// func Error2Str(data error) string {
// 	return data.Error()
// }

func Int2Str(data int) string {
	return Sprintf("%d", data)
}

func Float2Str(data float64) string {
	return Sprintf("%f", data)
}

func Bool2Str(data bool) string {
	return Sprintf("%t", data)
}

func ToStr(data any) string {
	switch v := data.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'f', 2, 64)
	case bool:
		return strconv.FormatBool(v)
	default:
		return fmt.Sprintf("%v", data)
	}
}

func Bytes2Str(data []byte) string {
	return string(data)
}

func Rune2Str(data rune) string {
	return string(data)
}

func Map2Str[T string | []string](data map[string]T) string {
	var result string

	for key, value := range data {
		result += Sprintf("%s: %s\n", key, value)
	}

	return result
}

func Time2Str(t time.Time) string {
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()

	hourStr := strconv.Itoa(hour)
	minuteStr := strconv.Itoa(minute)
	secondStr := strconv.Itoa(second)

	if hour < 10 {
		hourStr = Sprintf("0%s", hourStr)
	}
	if minute < 10 {
		minuteStr = Sprintf("0%s", minuteStr)
	}
	if second < 10 {
		secondStr = Sprintf("0%s", secondStr)
	}

	return Sprintf("%s:%s:%s", hourStr, minuteStr, secondStr)
}
