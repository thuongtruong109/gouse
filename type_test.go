package gouse

import (
	"errors"
	"reflect"
	"testing"

	"github.com/thuongtruong109/gouse/shared"
)

/* Testing type of variable */

func TestIsInt(t *testing.T) {
	var i int
	if !IsInt(i) {
		t.Error("IsInt(i) should return true")
	}
}

func TestIsUnInt(t *testing.T) {
	var i uint
	if !IsUnInt(i) {
		t.Error("IsUnInt(i) should return true")
	}
}

func TestIsFloat(t *testing.T) {
	var i float32
	if !IsFloat(i) {
		t.Error("IsFloat(i) should return true")
	}
}

func TestIsComplex(t *testing.T) {
	var i complex64
	if !IsComplex(i) {
		t.Error("IsComplex(i) should return true")
	}
}

func TestIsNumber(t *testing.T) {
	var i int
	if !IsNumber(i) {
		t.Error("IsNumber(i) should return true")
	}
}

func TestIsString(t *testing.T) {
	var i string
	if !IsString(i) {
		t.Error("IsString(i) should return true")
	}
}

func TestIsRune(t *testing.T) {
	var i rune
	if !IsRune(i) {
		t.Error("IsRune(i) should return true")
	}
}

func TestIsByte(t *testing.T) {
	var i byte
	if !IsByte(i) {
		t.Error("IsByte(i) should return true")
	}
}

func TestIsUintptr(t *testing.T) {
	var i uintptr
	if !IsUintptr(i) {
		t.Error("IsUintptr(i) should return true")
	}
}

func TestIsError(t *testing.T) {
	var i error = errors.New("test")
	if !IsError(i) {
		t.Error("IsError(i) should return true")
	}
}

func TestIsChannel(t *testing.T) {
	var i chan int
	if !IsChannel(i) {
		t.Error("IsChannel(i) should return true")
	}
}

func TestIsUnsafePointer(t *testing.T) {
	var i uintptr
	if !IsUnsafePointer(i) {
		t.Error("IsUnsafePointer(i) should return true")
	}
}

func TestIsPointer(t *testing.T) {
	var i *int
	if !IsPointer(i) {
		t.Error("IsPointer(i) should return true")
	}
}

func TestIsBool(t *testing.T) {
	var i bool
	if !IsBool(i) {
		t.Error("IsBool(i) should return true")
	}
}

func TestIsSlice(t *testing.T) {
	var i []int
	if !IsSlice(i) {
		t.Error("IsSlice(i) should return true")
	}
}

func TestIsMap(t *testing.T) {
	var i map[string]int
	if !IsMap(i) {
		t.Error("IsMap(i) should return true")
	}
}

func TestIsStruct(t *testing.T) {
	type testStruct struct {
		a int
	}
	var i = testStruct{
		a: 1,
	}
	if !IsStruct(i) {
		t.Error("IsStruct(i) should return true")
	}
}

func TestIsArray(t *testing.T) {
	var i [3]int
	if !IsArray(i) {
		t.Error("IsArray(i) should return true")
	}
}

func TestIsFunc(t *testing.T) {
	var i func()
	if !IsFunc(i) {
		t.Error("IsFunc(i) should return true")
	}
}

func TestIsNil(t *testing.T) {
	var i interface{}
	if !IsNil(i) {
		t.Error("IsNil(i) should return true")
	}
}

func TestIsEmpty(t *testing.T) {
	var i string
	if !IsEmpty(i) {
		t.Error("IsEmpty(i) should return true")
	}
}

func TestIsZero(t *testing.T) {
	var i int
	if !IsZero(i) {
		t.Error("IsZero(i) should return true")
	}
}

func TestUndefined(t *testing.T) {
	var i bool
	if !IsUndefined(i) {
		t.Error("IsUndefined(i) should return true")
	}
}

/* Testing type of form */

func TestIsUUID(t *testing.T) {
	tests := []shared.ITest{
		{
			Name:    "Test 1",
			Input:   "123e4567-e89b-12d3-a456-426614174000",
			Want:    true,
			WantErr: false,
		},
		{
			Name:    "Test 2",
			Input:   "123e4567-e89b-12d3-a456-42661417400",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 3",
			Input:   "",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 4",
			Input:   "123e4567-e89b-12d3-a456-4266141740001",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 5",
			Input:   "123e4567-e89b-12d3-a456-42661417400-",
			Want:    false,
			WantErr: true,
		},
	}

	for _, tt := range tests {
		isValid, err := IsUUID(tt.Input)
		if (err != nil) != tt.WantErr {
			t.Errorf("IsUUID(): %s error = %v, wantErr %v", tt.Name, err, tt.WantErr)
		}

		if (isValid != tt.Want) || !reflect.DeepEqual(isValid, tt.Want) {
			t.Errorf("IsUUID(): %s got = %v, want %v", tt.Name, isValid, tt.Want)
		}
	}
}

func TestIsGmail(t *testing.T) {
	tests := []shared.ITest{
		{
			Name:    "Test 1",
			Input:   "example@gmail.com",
			Want:    true,
			WantErr: false,
		},
		{
			Name:    "Test 2",
			Input:   "example",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 3",
			Input:   "example@",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 4",
			Input:   "example@.com",
			Want:    false,
			WantErr: true,
		},
	}

	for _, tt := range tests {
		isValid, err := IsGmail(tt.Input)
		if (err != nil) != tt.WantErr {
			t.Errorf("IsGmail(): %s error = %v, wantErr %v", tt.Name, err, tt.WantErr)
		}

		if (isValid != tt.Want) || !reflect.DeepEqual(isValid, tt.Want) {
			t.Errorf("IsGmail(): %s got = %v, want %v", tt.Name, isValid, tt.Want)
		}
	}
}

func TestIsYahoo(t *testing.T) {
	tests := []shared.ITest{
		{
			Name:    "Test 1",
			Input:   "example@yahoo.com",
			Want:    true,
			WantErr: false,
		},
		{
			Name:    "Test 2",
			Input:   "example",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 3",
			Input:   "example@.com",
			Want:    false,
			WantErr: true,
		},
	}

	for _, tt := range tests {
		isValid, err := IsYahoo(tt.Input)
		if (err != nil) != tt.WantErr {
			t.Errorf("IsYahoo(): %s error = %v, wantErr %v", tt.Name, err, tt.WantErr)
		}

		if (isValid != tt.Want) || !reflect.DeepEqual(isValid, tt.Want) {
			t.Errorf("IsYahoo(): %s got = %v, want %v", tt.Name, isValid, tt.Want)
		}
	}
}

func TestIsOutlook(t *testing.T) {
	tests := []shared.ITest{
		{
			Name:    "Test 1",
			Input:   "example@outlook.com",
			Want:    true,
			WantErr: false,
		},
		{
			Name:    "Test 2",
			Input:   "example",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 3",
			Input:   "example@.com",
			Want:    false,
			WantErr: true,
		},
	}

	for _, tt := range tests {
		isValid, err := IsOutlook(tt.Input)
		if (err != nil) != tt.WantErr {
			t.Errorf("IsOutlook(): %s error = %v, wantErr %v", tt.Name, err, tt.WantErr)
		}

		if (isValid != tt.Want) || !reflect.DeepEqual(isValid, tt.Want) {
			t.Errorf("IsOutlook(): %s got = %v, want %v", tt.Name, isValid, tt.Want)
		}
	}
}

func TestIsEdu(t *testing.T) {
	tests := []shared.ITest{
		{
			Name:    "Test 1",
			Input:   "example@edu.vn",
			Want:    true,
			WantErr: false,
		},
		{
			Name:    "Test 2",
			Input:   "",
			Want:    false,
			WantErr: true,
		},
	}

	for _, tt := range tests {
		isValid, err := IsEdu(tt.Input)
		if (err != nil) != tt.WantErr {
			t.Errorf("IsEdu(): %s error = %v, wantErr %v", tt.Name, err, tt.WantErr)
		}

		if (isValid != tt.Want) || !reflect.DeepEqual(isValid, tt.Want) {
			t.Errorf("IsEdu(): %s got = %v, want %v", tt.Name, isValid, tt.Want)
		}
	}
}

func TestIsEmail(t *testing.T) {
	tests := []shared.ITest{
		{
			Name:    "Test 1",
			Input:   "user01@example.com",
			Want:    true,
			WantErr: false,
		},
		{
			Name:    "Test 2",
			Input:   "",
			Want:    false,
			WantErr: true,
		},
	}

	for _, tt := range tests {
		isValid, err := IsEmail(tt.Input, "example.com")
		if (err != nil) != tt.WantErr {
			t.Errorf("IsEmail(): %s error = %v, wantErr %v", tt.Name, err, tt.WantErr)
		}

		if (isValid != tt.Want) || !reflect.DeepEqual(isValid, tt.Want) {
			t.Errorf("IsEmail(): %s got = %v, want %v", tt.Name, isValid, tt.Want)
		}
	}
}

func TestUsername(t *testing.T) {
	tests := []shared.ITest{
		{
			Name:    "Test 1",
			Input:   "example",
			Want:    true,
			WantErr: false,
		},
		{
			Name:    "Test 2",
			Input:   "example@",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 3",
			Input:   "example@.",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 4",
			Input:   "example@.com",
			Want:    false,
			WantErr: true,
		},
	}

	for _, tt := range tests {
		isValid, err := IsUsername(tt.Input)
		if (err != nil) != tt.WantErr {
			t.Errorf("IsUsername(): %s error = %v, wantErr %v", tt.Name, err, tt.WantErr)
		}

		if (isValid != tt.Want) || !reflect.DeepEqual(isValid, tt.Want) {
			t.Errorf("IsUsername(): %s got = %v, want %v", tt.Name, isValid, tt.Want)
		}
	}
}

func TestIsPassword(t *testing.T) {
	tests := []shared.ITest{
		{
			Name:    "Test 1",
			Input:   "12345678",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 2",
			Input:   "1234567",
			Want:    false,
			WantErr: true,
		},
	}

	for _, tt := range tests {
		isValid, err := IsPassword(tt.Input)
		if (err != nil) != tt.WantErr {
			t.Errorf("IsPassword(): %s error = %v, wantErr %v", tt.Name, err, tt.WantErr)
		}

		if (isValid != tt.Want) || !reflect.DeepEqual(isValid, tt.Want) {
			t.Errorf("IsPassword(): %s got = %v, want %v", tt.Name, isValid, tt.Want)
		}
	}
}

func TestIsPhone(t *testing.T) {
	tests := []shared.ITest{
		{
			Name:    "Test 1",
			Input:   "12345678",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 2",
			Input:   "1234567",
			Want:    false,
			WantErr: true,
		},
	}

	for _, tt := range tests {
		isValid, err := IsPhone(tt.Input)
		if (err != nil) != tt.WantErr {
			t.Errorf("IsPhone(): %s error = %v, wantErr %v", tt.Name, err, tt.WantErr)
		}

		if (isValid != tt.Want) || !reflect.DeepEqual(isValid, tt.Want) {
			t.Errorf("IsPhone(): %s got = %v, want %v", tt.Name, isValid, tt.Want)
		}
	}
}

/* Testing convert type */

func TestStructToString(t *testing.T) {
	type testStruct struct {
		One   int
		Two   string
		Three bool
	}

	data := testStruct{1, "two", true}

	result := StructToString(data)
	expected := "testStruct{One: 1, Two: two, Three: true}"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestStructToMap(t *testing.T) {
	type testStruct struct {
		One   int
		Two   string
		Three bool
	}

	data := testStruct{1, "two", true}

	result := StructToMap(data)
	expected := map[string]interface{}{
		"One":   1,
		"Two":   "two",
		"Three": true,
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestStringToInt(t *testing.T) {
	data := "123"
	result := StringToInt(data)
	expected := 123

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestStringToFloat(t *testing.T) {
	data := "123.456"
	result := StringToFloat(data)
	expected := 123.456

	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestStringToBool(t *testing.T) {
	data := "true"
	result := StringToBool(data)
	expected := true

	if result != expected {
		t.Errorf("Expected %t, got %t", expected, result)
	}
}

func TestStringToBytes(t *testing.T) {
	data := "test"
	result := StringToBytes(data)
	expected := []byte{116, 101, 115, 116}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestStringsToBytes(t *testing.T) {
	data := []string{"test", "test"}
	result := StringsToBytes(data)
	expected := []byte{116, 101, 115, 116, 116, 101, 115, 116}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestIntToString(t *testing.T) {
	data := 123
	result := IntToString(data)
	expected := "123"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestFloatToString(t *testing.T) {
	data := 123.456
	result := FloatToString(data)
	expected := "123.456000"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestBoolToString(t *testing.T) {
	data := true
	result := BoolToString(data)
	expected := "true"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestBytesToString(t *testing.T) {
	data := []byte{116, 101, 115, 116}
	result := BytesToString(data)
	expected := "test"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestRuneToString(t *testing.T) {
	data := 't'
	result := RuneToString(data)
	expected := "t"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestMapAsString(t *testing.T) {
	data := map[string]string{
		"One":   "one",
		"Two":   "two",
		"Three": "three",
	}
	result := MapAsString(data)
	expected := "One: one\nTwo: two\nThree: three\n"

	resultType := reflect.TypeOf(result)
	expectedType := reflect.TypeOf(expected)

	if resultType != expectedType {
		t.Errorf("Expected %s, got %s", expectedType, resultType)
	}
}

func TestToString(t *testing.T) {
	type testStruct struct {
		One   int
		Two   string
		Three bool
	}

	data := testStruct{1, "two", true}
	result := ToString(data)
	expected := "{1 two true}"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

// func TestStringToRune(t *testing.T) {
// 	data := "t"
// 	result := StringToRune(data)
// 	expected := 't'

// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }

// func TestStringToTime(t *testing.T) {
// 	data := "2015-01-01 00:00:00"
// 	result := StringToTime(data)
// 	expected := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)

// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }

// func TestTimeToString(t *testing.T) {
// 	data := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
// 	result := TimeToString(data)
// 	expected := "2015-01-01 00:00:00"

// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }

// func TestStringToDuration(t *testing.T) {
// 	data := "1s"
// 	result := StringToDuration(data)
// 	expected := time.Second

// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }

// func TestDurationToString(t *testing.T) {
// 	data := time.Second
// 	result := DurationToString(data)
// 	expected := "1s"

// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }

// func TestStringToByteSize(t *testing.T) {
// 	data := "1KB"
// 	result := StringToByteSize(data)
// 	expected := ByteSize(1024)

// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }

// func TestByteSizeToString(t *testing.T) {
// 	data := ByteSize(1024)
// 	result := ByteSizeToString(data)
// 	expected := "1KB"

// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }

// func TestStringToComplex(t *testing.T) {
// 	data := "1+1i"
// 	result := StringToComplex(data)
// 	expected := complex(1, 1)

// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }

// func TestComplexToString(t *testing.T) {
// 	data := complex(1, 1)
// 	result := ComplexToString(data)
// 	expected := "1+1i"

// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }