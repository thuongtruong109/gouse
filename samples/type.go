package samples

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
)

/*
Description: Convert string to other type
*/
func TypeStringConvert() {
	println("Convert string to int: ", gouse.StringToInt("1"))
	fmt.Println("Convert string to float: ", gouse.StringToFloat("1.1"))
	println("Convert string to bool: ", gouse.StringToBool("true"))
	fmt.Println("Convert string to bytes: ", string(gouse.StringToBytes("1")), "->", gouse.StringToBytes("1"))
	fmt.Println("Convert strings to bytes: ", string(gouse.StringsToBytes([]string{"1", "2", "3"})), "->", gouse.StringsToBytes([]string{"1", "2", "3"}))
}

/*
Description: Convert other type to string
*/
func TypeCastToString() {
	println("Cast int to string: ", gouse.IntToString(1))
	fmt.Println("Cast float to string: ", gouse.FloatToString(1.1))
	println("Cast bool to string: ", gouse.BoolToString(true))
	println("Cast interface to string: ", gouse.ToString([]int{1, 2, 3}))
	println("Cast bytes to string: ", gouse.BytesToString([]byte{49, 50, 51}))
	println("Cast rune to string: ", gouse.RuneToString('a'))
}

/*
Description: Convert struct to string, map
*/
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

	println("Struct to string: ", gouse.StructToString(companyInfo))
	fmt.Println("Struct to map: ", gouse.StructToMap(companyInfo))
}

/*
Description: Check type of variable
*/
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

/*
Description: Check string is UUID
Input params: (id string)
*/
func TypeCheckUUID() {
	isValid, err := gouse.IsUUID("123e4567-e89b-12d3-a456-426614174000")
	if err != nil {
		println(err.Error())
	}
	println("Check is valid uuid: ", isValid)
}

/*
Description: Check string is gmail
Input params: (email string)
*/
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

	println("--- Check valid gmail ---")
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

/*
Description: Check string is yahoo
Input params: (email string)
*/
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

	println("--- Check valid yahoo ---")
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

/*
Description: Check string is outlook
Input params: (email string)
*/
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

	println("--- Check valid outlook ---")
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

/*
Description: Check string is education email
Input params: (email string)
*/
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

	println("--- Check valid education mail ---")
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

/*
Description: Check string is email with custom domain
Input params: (email string, domain string)
*/
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

	println("--- Check valid custom domain ---")
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

/*
Description: Check string is valid username recommended
Input params: (username string)
*/
func TypeCheckUsername() {
	usernames := []string{
		"okahehe",
		"missingat.sign",
		"#$%&*()",
	}

	println("--- Check valid username ---")
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

/*
Description: Check string is valid password recommended
Input params: (password string)
*/
func TypeCheckPassword() {
	passwords := []string{
		"okahehe",
		"missingat.sign",
		"#$%&*()",
		"Username01$",
	}

	println("--- Check valid password ---")
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

/*
Description: Check string is valid phone number
Input params: (phone string)
*/
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

	println("--- Check valid phone number ---")
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
