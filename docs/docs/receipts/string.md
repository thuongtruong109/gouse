
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 String' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

## 1. String capitalize

Description: Capitalize the first letter of a string.<br>Input params: (s string)<br>

```go
func StringCapitalize() {
	var str = "hello world"
	println(gouse.Capitalize(str))
}
```

## 2. String camel case

Description: Convert string to Camel Case.<br>Input params: (s string)<br>

```go
func StringCamelCase() {
	var str = "convert_any-string TO-camelCase in Go! 123"
	println("Convert string to Camel Case: ", gouse.CamelCase(str))
}
```

## 3. String snake case

Description: Convert string to Snake Case.<br>Input params: (s string)<br>

```go
func StringSnakeCase() {
	var str1 = "hello world"
	var str2 = "ConvertThisString123"
	println("Convert string to Snake Case: ", gouse.SnakeCase(str1), gouse.SnakeCase(str2))
}
```

## 4. String kebab case

Description: Convert string to Kebab Case.<br>Input params: (s string)<br>

```go
func StringKebabCase() {
	var str = "hello world"
	println("Convert string to Kebab Case: ", gouse.KebabCase(str))
}
```

## 5. String space case

Description: Convert string to Space Case.<br>Input params: (s string)<br>

```go
func StringSpaceCase() {
	var str = "hello world"
	println("Convert string to Space Case: ", gouse.SpaceCase(str))
}
```

## 6. String custom case

Description: Convert string to custom.<br>Input params: (s string, separator string)<br>

```go
func StringCustomCase() {
	var str = "hello world"
	println("Convert string to Custom Case: ", gouse.CustomCase(str, "%"))
}
```

## 7. String is letter

Description: Check is letter character.<br>Input params: (s string)<br>

```go
func StringIsLetter() {
	var str = "hello world"
	println("Check is letter character: ", gouse.IsLetter(str[0]))
}
```

## 8. String is digit

Description: Check is number character.<br>Input params: (s string)<br>

```go
func StringIsDigit() {
	var str = "1hello world"
	println("Check is number character: ", gouse.IsDigit(str[0]))
}
```

## 9. String includes

Description: Check is contain substring in string.<br>Input params: (s string, substr string)<br>

```go
func StringIncludes() {
	var str = "hello world, this is world"
	println("Check substring in string: ", gouse.Includes(str, "world"))
}
```

## 10. String is lower

Description: Check is lower string.<br>Input params: (s string)<br>

```go
func StringIsLower() {
	var str = "hELLO WORLD"
	println("Check is lower string: ", gouse.IsLower(str[0]))
}
```

## 11. String is upper

Description: Check is upper string.<br>Input params: (s string)<br>

```go
func StringIsUpper() {
	var str = "Hello world"
	println("Check is upper string: ", gouse.IsUpper(str[0]))
}
```

## 12. String split

Description: Split string by separator.<br>Input params: (s string, separator string)<br>

```go
func StringSplit() {
	var str = "hello world"
	gouse.Println("Split string by separator: ", gouse.Split(str, "l"))
}
```

## 13. String words

Description: Split string by single word<br>Input params: (s string)<br>

```go
func StringWords() {
	var str = "hello world"
	println("Split string to array of words: ", gouse.Words(str))

}
```

## 14. String reverse

Description: Reverse string.<br>Input params: (s string)<br>

```go
func StringReverse() {
	var str = "hello world"
	println("Reverse string: ", gouse.Reverse(str))
}
```

## 15. String lower

Description: Convert string to lower case.<br>Input params: (s string)<br>

```go
func StringLower() {
	var str = "HELLO WORLD"
	println("Lower string (string): ", gouse.Lowers(str))
	println("Lower string (byte): ", gouse.Lower(str[0]))
	println("Lower first string: ", gouse.LowerFirst(str))
}
```

## 16. String upper

Description: Convert string to upper case.<br>Input params: (s string)<br>

```go
func StringUpper() {
	var str = "hello world"
	println("Upper string (string): ", gouse.Uppers(str))
	println("Upper string (byte): ", gouse.Upper(str[0]))
	println("Upper first string: ", gouse.UpperFirst(str))
}
```

## 17. String repeat

Description: Repeat string n times.<br>Input params: (struct interface{}, fieldName string, value interface{})<br>

```go
func StringRepeat() {
	var str = "hello world"
	println("Repeat string: ", gouse.Repeat(str, 3))
}
```

## 18. String truncate

Description: Truncate string to n characters or add suffix.<br>Input params: (s string, length int, suffix ...string)<br>

```go
func StringTruncate() {
	var str = "hello world"
	println("Truncate string (default): ", gouse.Truncate(str, 5))
	println("Truncate string (custom): ", gouse.Truncate(str, 5, "***"))
}
```

## 19. String replace

Description: Replace substring in string.<br>Input params: (s string, old string, new string)<br>

```go
func StringReplace() {
	var str = "hello world, this is world"
	println("Replace string: ", gouse.Replace(str, "world", "golang"))
}
```

## 20. String trim

Description: Trim left and right space in string.<br>Input params: (s string)<br>

```go
func StringTrim() {
	var str = "   hello world, this is world   "
	println("Trim string: ", gouse.Trim(str))
	println("Trim left string: ", gouse.LTrim(str))
	println("Trim right string: ", gouse.RTrim(str))
}
```

## 21. String trim blank

Description: Trim left and right blank in string.<br>Input params: (s string)<br>

```go
func StringTrimBlank() {
	println("Trim blank string: ", gouse.TrimBlank("   hello world, this is world   "))
	println("Trim left blank string: ", gouse.TrimBlank("   hello world, this is world   \t"))
	println("Trim right blank string: ", gouse.TrimBlank("   hello world, this is world   \n"))
}
```

## 22. String trim prefix

Description: Trim left and right prefix in string.<br>Input params: (s string, prefix string)<br>

```go
func StringTrimPrefix() {
	var str = "   hello world, this is world   "
	println("Trim prefix string: ", gouse.TrimPre(str, "   "))
	println("Trim suffix string: ", gouse.TrimSuf(str, "   "))
}
```

## 23. String trim suffix

Description: Trim left and right suffix in string.<br>Input params: (s string, suffix string)<br>

```go
func StringTrimSuffix() {
	var str = "   hello world, this is world   "
	println("Trim suffix string: ", gouse.TrimSuf(str, "   "))
}
```

## 24. String join

Description: Join string array to string.<br>Input params: (s []string, separator string)<br>

```go
func StringJoin() {
	var str = []string{"hello", "world"}
	println("Join string: ", gouse.Join(str, "-"))
}
```

## 25. String concat

Description: Concat string.<br>Input params: (s ...string)<br>

```go
func StringConcat() {
	println("Concat string: ", gouse.Concat("hello", "world"))
}
```

## 26. String sub str

Description: Get substring from string.<br>Input params: (s string, start int, end ...int)<br>

```go
func StringSubStr() {
	var str = "hello world, this is world"
	println("Sub string: ", gouse.SubStr(str, 0, 5))
	println("Sub string: ", gouse.SubStr(str, 0, 1))
	println("Sub string (only start): ", gouse.SubStr(str, -5))
	println("Sub string (with negative index): ", gouse.SubStr(str, -5, -1))
}
```

## 27. String slice

Description: Slice string from start to end.<br>Input params: (s string, start int, end ...int)<br>

```go
func StringSlice() {
	var str = "hello world, this is world"
	println("Slice string: ", gouse.Slice(str, 0, 5))
	println("Slice string: ", gouse.Slice(str, 0, 1))
	println("Slice string (only start): ", gouse.Slice(str, -5))
	println("Slice string (not parameters): ", gouse.Slice(str))
	println("Slice string (with negative index): ", gouse.Slice(str, -5, -1))
}
```

## 28. String splice

Description: Splice string from start to end.<br>Input params: (s string, start int, end int, replace ...string)<br>

```go
func StringSplice() {
	var str = "helloworld, this is world"
	println("Splice string (default not replace): ", gouse.Splice(str, 0, 5))
	println("Splice string (with replace): ", gouse.Splice(str, 1, 5, "golang"))
	println("Splice string (with replace multiple): ", gouse.Splice(str, 1, 5, "golang1", "golang2"))
}
```

## 29. String starts with

Description: Check is start with substring in string.<br>Input params: (s string, subStr string)<br>

```go
func StringStartsWith() {
	var str = "hello world, this is world"
	println("Starts with: ", gouse.StartsWith(str, "hello"))
}
```

## 30. String ends with

Description: Check is end with substring in string.<br>Input params: (s string, subStr string)<br>

```go
func StringEndsWith() {
	var str = "hello world, this is world"
	println("Ends with: ", gouse.EndsWith(str, "world"))
}
```

## 31. String escape

Description: Convert string to HTML entities.<br>Input params: (s string)<br>

```go
func StringEscape() {
	var str = "This is a <b>bold</b> statement & \"quote\" – © ®"
	println("Escape string: ", gouse.Esc(str))
}
```

## 32. String unescape

Description: Convert HTML entities to string.<br>Input params: (s string)<br>

```go
func StringUnescape() {
	var str = "This is a &lt;b&gt;bold&lt;/b&gt; statement &amp; &quot;quote&quot; – © ®"
	println("Unescape string: ", gouse.Unesc(str))
}
```

## 33. String pad

Description: Add padding to string from left or right.<br>Input params: (s string, length int, padStr string)<br>

```go
func StringPad() {
	var str = "hello world"
	println("Pad-left string: ", gouse.LPad(str, 20, '$'))
	println("Pad-right string: ", gouse.RPad(str, 20, '@'))
}
```

## 34. String count

Description: Count words/substr in string.<br>Input params: (s string, substr ...string)<br>

```go
func StringCount() {
	var str = "hello world wo wo"
	println("Count words/substr in string (default): ", gouse.Count(str))
	println("Count words/substr in string (with char): ", gouse.Count(str, "wo"))
}
```

## 35. String lines

Description: Count lines in string.<br>Input params: (s string)<br>

```go
func StringLines() {
	var str = "hello world\nwo wo"
	println("Count lines of string: ", gouse.Lines(str))
}
```

## 36. String index

Description: Get index of substring in string.<br>Input params: (s string, substr string)<br>

```go
func StringIndex() {
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

## 37. String random chain

Description: Random chain string.<br>Input params: (length int)<br>

```go
func StringRandomChain() {
	println("Random chain string: ", gouse.RandStr(10))
	println("Random chain number: ", gouse.RandDigit(6))
}
```

## 38. String at

Description: Get character at index in string.<br>Input params: (s string, index int)<br>

```go
func StringAt() {
	var str = "hello world"
	println("At string: ", gouse.At(str, 1))
	println("At string: ", gouse.At(str, -5))
}
```

## 39. String code point at

Description: Get code point at index in string.<br>Input params: (s string, index int)<br>

```go
func StringCodePointAt() {
	var str = "hello world"
	println("Code point at string: ", gouse.CodePointAt(str, 1))
	println("Code point at string: ", gouse.CodePointAt(str, -5))
}
```

## 40. String code point

Description: Get code point of string.<br>Input params: (s string)<br>

```go
func StringCodePoint() {
	asciiValues := gouse.CodePoint("hello world")

	print("Code point string: ")
	for _, asciiValue := range asciiValues {
		gouse.Printf("%d ", asciiValue)
	}
}
```

## 41. String from code point at

Description: Transform code point to string.<br>Input params: (codePoint int)<br>

```go
func StringFromCodePointAt() {
	println("From code point at string: ", gouse.FromCodePointAt(9733))
	println("From code point at string: ", gouse.FromCodePointAt(9731))
}
```

## 42. String from code point

Description: Transform multi code points to string.<br>Input params: (codePoints ...int)<br>

```go
func StringFromCodePoint() {
	println("From code point string: ", gouse.FromCodePoint(104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100))
}
```
