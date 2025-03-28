
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Print' />


```go
import (
	"fmt"
)
```

## 1. Print

Description: Print, Println, Printf<br>Input params: ...[](interface{})<br>

```go
func Print() {
	fmt.Println("Hello,", "this is custom fmt", 123, true)
	fmt.Print("This is Print: ")
	fmt.Println("No newline above!")
	fmt.Printf("Formatted: number=%d, text=%s, float=%f, bool=%t\n", 42, "Golang", 3.14159, true)
}
```

## 2. Sprint

Description: Sprint, Sprintln, Sprintf<br>Input params: ...[](interface{})<br>

```go
func Sprint() {
	str1 := fmt.Sprint("Sprint result:", 42, "text", true)
	fmt.Println("Sprint:", str1)

	str2 := fmt.Sprintln("Sprintln result:", 123, "another line", false)
	fmt.Print("Sprintln:", str2)

	str3 := fmt.Sprintf("Sprintf formatted: %d %s %.2f %t", 100, "GoLang", 1.23, false)
	fmt.Println("Sprintf:", str3)
}
```
