
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Type' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Type string convert

Description: Convert string to other type<br>

```go
func TypeStringConvert() {
	println("Convert string to int: ", gouse.Str2Int("1"))
	fmt.Println("Convert string to float: ", gouse.Str2Float("1.1"))
	println("Convert string to bool: ", gouse.Str2Bool("true"))
	fmt.Println("Convert string to bytes: ", string(gouse.Str2Bytes("1")), "->", gouse.Str2Bytes("1"))
	fmt.Println("Convert strings to bytes: ", string(gouse.Strs2Bytes([]string{"1", "2", "3"})), "->", gouse.Strs2Bytes([]string{"1", "2", "3"}))
}
```

## 2. Type cast to string

Description: Convert other type to string<br>

```go
func TypeCastToString() {
	println("Cast int to string: ", gouse.Int2Str(1))
	fmt.Println("Cast float to string: ", gouse.Float2Str(1.1))
	println("Cast bool to string: ", gouse.Bool2Str(true))
	println("Cast interface to string: ", gouse.ToStr([]int{1, 2, 3}))
	println("Cast bytes to string: ", gouse.Bytes2Str([]byte{49, 50, 51}))
	println("Cast rune to string: ", gouse.Rune2Str('a'))
}
```

## 3. Type struct convert

Description: Convert struct to string, map<br>

```go
func TypeStructConvert() {
	type CompanyInfo struct {
		Company string
		Address string
		Contact string
	}

	companyInfo := CompanyInfo{
		Company: "GeeksforGeeks",
		Address: "Noida",
		Contact: "+91 9876543210",
	}

	println("Struct to string: ", gouse.Struct2Str(companyInfo))
	fmt.Println("Struct to map: ", gouse.Struct2Map(companyInfo))
}
```

## 4. Type check

Description: Check type of variable<br>

```go
func TypeCheck() {
	println("Check type is int: ", gouse.IsInt(1))
	println("Check type is uint: ", gouse.IsUnInt(-1))
	println("Check type is float: ", gouse.IsFloat(1.1))
	println("Check type is complex: ", gouse.IsComplex(1.1))
	println("Check type is number: ", gouse.IsNumber(1.23))
	println("Check type is string: ", gouse.IsString("1"))
	println("Check type is rune: ", gouse.IsRune('a'))
	println("Check type is byte: ", gouse.IsByte(byte('a')))
	println("Check type is uintptr: ", gouse.IsUintptr(uintptr(1)))
	println("Check type is bool: ", gouse.IsBool(true))
	println("Check type is slice: ", gouse.IsSlice([]int{1, 2, 3}))
	println("Check type is map: ", gouse.IsMap(map[string]int{"a": 1}))
	println("Check type is struct: ", gouse.IsStruct(struct{}{}))
	println("Check type is error: ", gouse.IsError(new(error)))
	println("Check type is channel: ", gouse.IsChannel(make(chan int)))
	println("Check type is array: ", gouse.IsArray([3]int{1, 2, 3}))
	println("Check type is unsafe pointer: ", gouse.IsUnsafePointer(new(*int)))
	println("Check type is pointer: ", gouse.IsPointer(new(int)))
	println("Check type is func: ", gouse.IsFunc(func() {}))
	println("Check type is nil: ", gouse.IsNil(nil))
	println("Check type is empty: ", gouse.IsEmpty(""))
	println("Check type is empty: ", gouse.IsEmpty(0))
	println("Check type is zero: ", gouse.IsZero(0))
}
```

## 5. Type check u u i d

Description: Check string is UUID<br>Input params: (id string)<br>

```go
func TypeCheckUUID() {
	isValid, err := gouse.IsUUID("123e4567-e89b-12d3-a456-426614174000")
	if err != nil {
		println(err.Error())
	}
	println("Check is valid uuid: ", isValid)
}
```

## 6. Type check gmail

Description: Check string is gmail<br>Input params: (email string)<br>

```go
func TypeCheckGmail() {
	var emails = []string{
		"okahehe@gmail.com",
		"test@example.com",
		"short@ex.com",
		"toolong@exampledomain.com",
		"missingat.sign.com",
		"invalid@.domain.com",
		"user2@yahoo.com",
		"bean@outlook.com",
		"ititiu19228@student.hcmiu.edu.vn",
	}

	for _, email := range emails {
		isEmail, err := gouse.IsGmail(email)
		if err != nil {
			fmt.Printf("%s - %s\n", email, err)
		}

		if isEmail {
			fmt.Printf("%s - valid gmail\n", email)
		}
	}
}
```

## 7. Type check yahoo

Description: Check string is yahoo<br>Input params: (email string)<br>

```go
func TypeCheckYahoo() {
	var emails = []string{
		"okahehe@gmail.com",
		"test@example.com",
		"short@ex.com",
		"toolong@exampledomain.com",
		"missingat.sign.com",
		"invalid@.domain.com",
		"user2@yahoo.com",
		"bean@outlook.com",
		"ititiu19228@student.hcmiu.edu.vn",
	}

	for _, email := range emails {
		isEmail, err := gouse.IsYahoo(email)
		if err != nil {
			fmt.Printf("%s - %s\n", email, err)
		}

		if isEmail {
			fmt.Printf("%s - valid yahoo email\n", email)
		}
	}
}
```

## 8. Type check outlook

Description: Check string is outlook<br>Input params: (email string)<br>

```go
func TypeCheckOutlook() {
	var emails = []string{
		"okahehe@gmail.com",
		"test@example.com",
		"short@ex.com",
		"toolong@exampledomain.com",
		"missingat.sign.com",
		"invalid@.domain.com",
		"user2@yahoo.com",
		"bean@outlook.com",
		"ititiu19228@student.hcmiu.edu.vn",
	}

	for _, email := range emails {
		isEmail, err := gouse.IsOutlook(email)
		if err != nil {
			fmt.Printf("%s - %s\n", email, err)
		}

		if isEmail {
			fmt.Printf("%s - valid outlook email\n", email)
		}
	}
}
```

## 9. Type check edu

Description: Check string is education email<br>Input params: (email string)<br>

```go
func TypeCheckEdu() {
	var emails = []string{
		"okahehe@gmail.com",
		"test@example.com",
		"short@ex.com",
		"toolong@exampledomain.com",
		"missingat.sign.com",
		"invalid@.domain.com",
		"user2@yahoo.com",
		"bean@outlook.com",
		"ititiu19228@student.hcmiu.edu.vn",
	}

	for _, email := range emails {
		isEmail, err := gouse.IsEdu(email)
		if err != nil {
			fmt.Printf("%s - %s\n", email, err)
		}

		if isEmail {
			fmt.Printf("%s - valid education email\n", email)
		}
	}
}
```

## 10. Type check email

Description: Check string is email with custom domain<br>Input params: (email string, domain string)<br>

```go
func TypeCheckEmail() {
	var emails = []string{
		"okahehe@gmail.com",
		"test@example.com",
		"short@ex.com",
		"toolong@exampledomain.com",
		"missingat.sign.com",
		"invalid@.domain.com",
		"user2@yahoo.com",
		"bean@outlook.com",
		"ititiu19228@student.hcmiu.edu.vn",
	}

	for _, email := range emails {
		isEmail, err := gouse.IsEmail(email, "edu.vn")
		if err != nil {
			fmt.Printf("%s - %s\n", email, err)
		}

		if isEmail {
			fmt.Printf("%s - valid custom domain email\n", email)
		}
	}
}
```

## 11. Type check username

Description: Check string is valid username recommended<br>Input params: (username string)<br>

```go
func TypeCheckUsername() {
	usernames := []string{
		"okahehe",
		"missingat.sign",
		"#$%&*()",
	}

	for _, username := range usernames {
		isUsername, err := gouse.IsUsername(username)
		if err != nil {
			fmt.Printf("%s - %s\n", username, err)
		}

		if isUsername {
			fmt.Printf("%s - valid username\n", username)
		}
	}
}
```

## 12. Type check password

Description: Check string is valid password recommended<br>Input params: (password string)<br>

```go
func TypeCheckPassword() {
	passwords := []string{
		"okahehe",
		"missingat.sign",
		"#$%&*()",
		"Username01$",
	}

	for _, password := range passwords {
		isPassword, err := gouse.IsPassword(password)
		if err != nil {
			fmt.Printf("%s - %s\n", password, err)
		}

		if isPassword {
			fmt.Printf("%s - valid password\n", password)
		}
	}
}
```

## 13. Type check phone

Description: Check string is valid phone number<br>Input params: (phone string)<br>

```go
func TypeCheckPhone() {
	//  Note: Phone format syntax: +<country_calling_code> (<area_Prefix_mobile_code>) <phone_number>
	// Reference at https://en.wikipedia.org/wiki/List_of_mobile_telephone_prefixes_by_country#:~:text=Property%20Value%20%20Country%20or%20unrecognized%20territory%20,73%20%20%20Etisalat%20%20%20www.etisalat.af%20
	phoneNumbers := []string{
		"+1 (123) 456-7890",
		"+44 (20) 1234-5678",
		"+81 (3) 1234-5678",
		"InvalidPhoneNumber",
		"+1 (123) 45-67890",
		"+84 (3) 668-22796",
	}

	for _, phoneNumber := range phoneNumbers {
		isPhoneNumber, err := gouse.IsPhone(phoneNumber)
		if err != nil {
			fmt.Printf("%s - %s\n", phoneNumber, err)
		}

		if isPhoneNumber {
			fmt.Printf("%s - valid phone number\n", phoneNumber)
		}
	}
}
```
