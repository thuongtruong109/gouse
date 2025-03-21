
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 String' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

## 1. Capitalize

Description: Capitalize the first letter of a string.<br>Input params: (s string)<br>

```go
func Capitalize() {
	var str = "hello world"
	println(gouse.Capitalize(str))
}
```

## 2. Camel case

Description: Convert string to Camel Case.<br>Input params: (s string)<br>

```go
func CamelCase() {
	var str = "convert_any-string TO-camelCase in Go! 123"
	println("Convert string to Camel Case: ", gouse.CamelCase(str))
}
```

## 3. Snake case

Description: Convert string to Snake Case.<br>Input params: (s string)<br>

```go
func SnakeCase() {
	var str1 = "hello world"
	var str2 = "ConvertThisString123"
	println("Convert string to Snake Case: ", gouse.SnakeCase(str1), gouse.SnakeCase(str2))
}
```

## 4. Kebab case

Description: Convert string to Kebab Case.<br>Input params: (s string)<br>

```go
func KebabCase() {
	var str = "hello world"
	println("Convert string to Kebab Case: ", gouse.KebabCase(str))
}
```

## 5. Is letter

Description: Check is letter character.<br>Input params: (s string)<br>

```go
func IsLetter() {
	var str = "hello world"
	println("Check is letter character: ", gouse.IsLetter(str[0]))
}
```

## 6. Is digit

Description: Check is number character.<br>Input params: (s string)<br>

```go
func IsDigit() {
	var str = "1hello world"
	println("Check is number character: ", gouse.IsDigit(str[0]))
}
```

## 7. Includes

Description: Check is contain substring in string.<br>Input params: (s string, substr string)<br>

```go
func Includes() {
	var str = "hello world, this is world"
	println("Check substring in string: ", gouse.Includes(str, "world"))
}
```

## 8. Is lower

Description: Check is lower string.<br>Input params: (s string)<br>

```go
func IsLower() {
	var str = "hELLO WORLD"
	println("Check is lower string: ", gouse.IsLower(str[0]))
}
```

## 9. Is upper

Description: Check is upper string.<br>Input params: (s string)<br>

```go
func IsUpper() {
	var str = "Hello world"
	println("Check is upper string: ", gouse.IsUpper(str[0]))
}
```

## 10. Split

Description: Split string by separator.<br>Input params: (s string, separator string)<br>

```go
func Split() {
	var str = "hello world"
	gouse.Println("Split string by separator: ", gouse.Split(str, "l"))
}
```

## 11. Words

Description: Split string by single word<br>Input params: (s string)<br>

```go
func Words() {
	var str = "hello world"
	println("Split string to array of words: ", gouse.Words(str))

}
```

## 12. Reverse

Description: Reverse string.<br>Input params: (s string)<br>

```go
func Reverse() {
	var str = "hello world"
	println("Reverse string: ", gouse.Reverse(str))
}
```

## 13. Lower

Description: Convert string to lower case.<br>Input params: (s string)<br>

```go
func Lower() {
	var str = "HELLO WORLD"
	println("Lower string (string): ", gouse.Lowers(str))
	println("Lower string (byte): ", gouse.Lower(str[0]))
	println("Lower first string: ", gouse.LowerFirst(str))
}
```

## 14. Upper

Description: Convert string to upper case.<br>Input params: (s string)<br>

```go
func Upper() {
	var str = "hello world"
	println("Upper string (string): ", gouse.Uppers(str))
	println("Upper string (byte): ", gouse.Upper(str[0]))
	println("Upper first string: ", gouse.UpperFirst(str))
}
```

## 15. Repeat

Description: Repeat string n times.<br>Input params: (struct interface{}, fieldName string, value interface{})<br>

```go
func Repeat() {
	var str = "hello world"
	println("Repeat string: ", gouse.Repeat(str, 3))
}
```

## 16. Truncate

Description: Truncate string to n characters or add suffix.<br>Input params: (s string, length int, suffix ...string)<br>

```go
func Truncate() {
	var str = "hello world"
	println("Truncate string (default): ", gouse.Truncate(str, 5))
	println("Truncate string (custom): ", gouse.Truncate(str, 5, "***"))
}
```

## 17. Replace

Description: Replace substring in string.<br>Input params: (s string, old string, new string)<br>

```go
func Replace() {
	var str = "hello world, this is world"
	println("Replace string: ", gouse.Replace(str, "world", "golang"))
}
```

## 18. Trim

Description: Trim left and right space in string.<br>Input params: (s string)<br>

```go
func Trim() {
	var str = "   hello world, this is world   "
	println("Trim string: ", gouse.Trim(str))
	println("Trim left string: ", gouse.LTrim(str))
	println("Trim right string: ", gouse.RTrim(str))
}
```

## 19. Trim blank

Description: Trim left and right blank in string.<br>Input params: (s string)<br>

```go
func TrimBlank() {
	println("Trim blank string: ", gouse.TrimBlank("   hello world, this is world   "))
	println("Trim left blank string: ", gouse.TrimBlank("   hello world, this is world   \t"))
	println("Trim right blank string: ", gouse.TrimBlank("   hello world, this is world   \n"))
}
```

## 20. Trim prefix

Description: Trim left and right prefix in string.<br>Input params: (s string, prefix string)<br>

```go
func TrimPrefix() {
	var str = "   hello world, this is world   "
	println("Trim prefix string: ", gouse.TrimPre(str, "   "))
	println("Trim suffix string: ", gouse.TrimSuf(str, "   "))
}
```

## 21. Trim suffix

Description: Trim left and right suffix in string.<br>Input params: (s string, suffix string)<br>

```go
func TrimSuffix() {
	var str = "   hello world, this is world   "
	println("Trim suffix string: ", gouse.TrimSuf(str, "   "))
}
```

## 22. Join

Description: Join string array to string.<br>Input params: (s []string, separator string)<br>

```go
func Join() {
	var str = []string{"hello", "world"}
	println("Join string: ", gouse.Join(str, "-"))
}
```

## 23. Concat

Description: Concat string.<br>Input params: (s ...string)<br>

```go
func Concat() {
	println("Concat string: ", gouse.Concat("hello", "world"))
}
```

## 24. Sub str

Description: Get substring from string.<br>Input params: (s string, start int, end ...int)<br>

```go
func SubStr() {
	var str = "hello world, this is world"
	println("Sub string: ", gouse.SubStr(str, 0, 5))
	println("Sub string: ", gouse.SubStr(str, 0, 1))
	println("Sub string (only start): ", gouse.SubStr(str, -5))
	println("Sub string (with negative index): ", gouse.SubStr(str, -5, -1))
}
```

## 25. Slice

Description: Slice string from start to end.<br>Input params: (s string, start int, end ...int)<br>

```go
func Slice() {
	var str = "hello world, this is world"
	println("Slice string: ", gouse.Slice(str, 0, 5))
	println("Slice string: ", gouse.Slice(str, 0, 1))
	println("Slice string (only start): ", gouse.Slice(str, -5))
	println("Slice string (not parameters): ", gouse.Slice(str))
	println("Slice string (with negative index): ", gouse.Slice(str, -5, -1))
}
```

## 26. Splice

Description: Splice string from start to end.<br>Input params: (s string, start int, end int, replace ...string)<br>

```go
func Splice() {
	var str = "helloworld, this is world"
	println("Splice string (default not replace): ", gouse.Splice(str, 0, 5))
	println("Splice string (with replace): ", gouse.Splice(str, 1, 5, "golang"))
	println("Splice string (with replace multiple): ", gouse.Splice(str, 1, 5, "golang1", "golang2"))
}
```

## 27. Starts with

Description: Check is start with substring in string.<br>Input params: (s string, subStr string)<br>

```go
func StartsWith() {
	var str = "hello world, this is world"
	println("Starts with: ", gouse.StartsWith(str, "hello"))
}
```

## 28. Ends with

Description: Check is end with substring in string.<br>Input params: (s string, subStr string)<br>

```go
func EndsWith() {
	var str = "hello world, this is world"
	println("Ends with: ", gouse.EndsWith(str, "world"))
}
```

## 29. Escape

Description: Convert string to HTML entities.<br>Input params: (s string)<br>

```go
func Escape() {
	var str = "This is a <b>bold</b> statement & \"quote\" – © ®"
	println("Escape string: ", gouse.Esc(str))
}
```

## 30. Unescape

Description: Convert HTML entities to string.<br>Input params: (s string)<br>

```go
func Unescape() {
	var str = "This is a &lt;b&gt;bold&lt;/b&gt; statement &amp; &quot;quote&quot; – © ®"
	println("Unescape string: ", gouse.Unesc(str))
}
```

## 31. Pad

Description: Add padding to string from left or right.<br>Input params: (s string, length int, padStr string)<br>

```go
func Pad() {
	var str = "hello world"
	println("Pad-left string: ", gouse.LPad(str, 20, '$'))
	println("Pad-right string: ", gouse.RPad(str, 20, '@'))
}
```

## 32. Count

Description: Count words/substr in string.<br>Input params: (s string, substr ...string)<br>

```go
func Count() {
	var str = "hello world wo wo"
	println("Count words/substr in string (default): ", gouse.Count(str))
	println("Count words/substr in string (with char): ", gouse.Count(str, "wo"))
}
```

## 33. Lines

Description: Count lines in string.<br>Input params: (s string)<br>

```go
func Lines() {
	var str = "hello world\nwo wo"
	println("Count lines of string: ", gouse.Lines(str))
}
```

## 34. Index

Description: Get index of substring in string.<br>Input params: (s string, substr string)<br>

```go
func Index() {
	var str = "hello world, this is world"

	f1, l1 := gouse.IdxSubStr(str, "l")
	gouse.Printf("First index start at: %d, end at: %d\n", f1, l1)

	f2, l2 := gouse.IdxSubStr(str, "world")
	gouse.Printf("First index start at: %d, end at: %d\n", f2, l2)

	f3 := gouse.FirstIdx(str, "l")
	gouse.Println("First index start at: ", f3)

	l3 := gouse.LastIdx(str, "world")
	gouse.Println("Last index start at: ", l3)

	f4, l4 := gouse.IdxSubStr(str, "oo")
	if f4 == -1 || l4 == -1 {
		gouse.Println("Not found")
	}

	if gouse.FirstIdx(str, "oo") == -1 {
		gouse.Println("Not found")
	}

	if gouse.LastIdx(str, "oo") == -1 {
		gouse.Println("Not found")
	}
}
```

## 35. Random chain

Description: Random chain string.<br>Input params: (length int)<br>

```go
func RandomChain() {
	println("Random chain string: ", gouse.RandStr(10))
	println("Random chain number: ", gouse.RandDigit(6))
}
```

## 36. At

Description: Get character at index in string.<br>Input params: (s string, index int)<br>

```go
func At() {
	var str = "hello world"
	println("At string: ", gouse.At(str, 1))
	println("At string: ", gouse.At(str, -5))
}
```

## 37. Code point at

Description: Get code point at index in string.<br>Input params: (s string, index int)<br>

```go
func CodePointAt() {
	var str = "hello world"
	println("Code point at string: ", gouse.CodePointAt(str, 1))
	println("Code point at string: ", gouse.CodePointAt(str, -5))
}
```

## 38. Code point

Description: Get code point of string.<br>Input params: (s string)<br>

```go
func CodePoint() {
	asciiValues := gouse.CodePoint("hello world")

	print("Code point string: ")
	for _, asciiValue := range asciiValues {
		gouse.Printf("%d ", asciiValue)
	}
}
```

## 39. From code point at

Description: Transform code point to string.<br>Input params: (codePoint int)<br>

```go
func FromCodePointAt() {
	println("From code point at string: ", gouse.FromCodePointAt(9733))
	println("From code point at string: ", gouse.FromCodePointAt(9731))
}
```

## 40. From code point

Description: Transform multi code points to string.<br>Input params: (codePoints ...int)<br>

```go
func FromCodePoint() {
	println("From code point string: ", gouse.FromCodePoint(104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100))
}
```
