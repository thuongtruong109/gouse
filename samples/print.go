package samples

import "github.com/thuongtruong109/gouse"

func Print() {
	gouse.Println("Hello,", "this is custom fmt", 123, true)
	gouse.Print("This is Print: ")
	gouse.Println("No newline above!")
	gouse.Printf("Formatted: number=%d, text=%s, float=%f, bool=%t\n", 42, "Golang", 3.14159, true)
}

func Sprint() {
	str1 := gouse.Sprint("Sprint result:", 42, "text", true)
	gouse.Println("Sprint:", str1)

	str2 := gouse.Sprintln("Sprintln result:", 123, "another line", false)
	gouse.Print("Sprintln:", str2)

	str3 := gouse.Sprintf("Sprintf formatted: %d %s %.2f %t", 100, "GoLang", 1.23, false)
	gouse.Println("Sprintf:", str3)
}
