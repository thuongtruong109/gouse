# String

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

#### 1. SampleStringCapitalize

```go
func SampleStringCapitalize() {
	var str = "hello world"
	println(gouse.Capitalize(str))
}
```

#### 2. SampleStringCamelCase

```go
func SampleStringCamelCase() {
	var str = "convert_any-string TO-camelCase in Go! 123"
	println("Convert string to Camel Case: ", gouse.CamelCase(str))
}
```

#### 3. SampleStringSnakeCase

```go
func SampleStringSnakeCase() {
	var str1 = "hello world"
	var str2 = "ConvertThisString123"
	println("Convert string to Snake Case: ", gouse.SnakeCase(str1), gouse.SnakeCase(str2))
}
```

#### 4. SampleStringKebabCase

```go
func SampleStringKebabCase() {
	var str = "hello world"
	println("Convert string to Kebab Case: ", gouse.KebabCase(str))
}
```

#### 5. SampleStringIsLetter

```go
func SampleStringIsLetter() {
	var str = "hello world"
	println("Check is letter character: ", gouse.IsLetter(str[0]))
}
```

#### 6. SampleStringIsDigit

```go
func SampleStringIsDigit() {
	var str = "1hello world"
	println("Check is number character: ", gouse.IsDigit(str[0]))
}
```

#### 7. SampleStringIncludes

```go
func SampleStringIncludes() {
	var str = "hello world, this is world"
	println("Check substring in string: ", gouse.Includes(str, "world"))
}
```

#### 8. SampleStringIsLower

```go
func SampleStringIsLower() {
	var str = "hELLO WORLD"
	println("Check is lower string: ", gouse.IsLower(str[0]))
}
```

#### 9. SampleStringIsUpper

```go
func SampleStringIsUpper() {
	var str = "Hello world"
	println("Check is upper string: ", gouse.IsUpper(str[0]))
}
```

#### 10. SampleStringSplit

```go
func SampleStringSplit() {
	var str = "hello world"
	fmt.Println("Split string by separator: ", gouse.Split(str, "l"))
}
```

#### 11. SampleStringWords

```go
func SampleStringWords() {
	var str = "hello world"
	println("Split string to array of words: ", gouse.Words(str))

}
```

#### 12. SampleStringReverse

```go
func SampleStringReverse() {
	var str = "hello world"
	println("Reverse string: ", gouse.Reverse(str))
}
```

#### 13. SampleStringLower

```go
func SampleStringLower() {
	var str = "HELLO WORLD"
	println("Lower string (string): ", gouse.Lowers(str))
	println("Lower string (byte): ", gouse.Lower(str[0]))
	println("Lower first string: ", gouse.LowerFirst(str))
}
```

#### 14. SampleStringUpper

```go
func SampleStringUpper() {
	var str = "hello world"
	println("Upper string (string): ", gouse.Uppers(str))
	println("Upper string (byte): ", gouse.Upper(str[0]))
	println("Upper first string: ", gouse.UpperFirst(str))
}
```

#### 15. SampleStringRepeat

```go
func SampleStringRepeat() {
	var str = "hello world"
	println("Repeat string: ", gouse.Repeat(str, 3))
}
```

#### 16. SampleStringTruncate

```go
func SampleStringTruncate() {
	var str = "hello world"
	println("Truncate string (default): ", gouse.Truncate(str, 5))
	println("Truncate string (custom): ", gouse.Truncate(str, 5, "***"))
}
```

#### 17. SampleStringReplace

```go
func SampleStringReplace() {
	var str = "hello world, this is world"
	println("Replace string: ", gouse.Replace(str, "world", "golang"))
}
```

#### 18. SampleStringTrim

```go
func SampleStringTrim() {
	var str = "   hello world, this is world   "
	println("Trim string: ", gouse.Trim(str))
	println("Trim left string: ", gouse.LTrim(str))
	println("Trim right string: ", gouse.RTrim(str))
}
```

#### 19. SampleStringTrimBlank

```go
func SampleStringTrimBlank() {
	println("Trim blank string: ", gouse.TrimBlank("   hello world, this is world   "))
	println("Trim left blank string: ", gouse.TrimBlank("   hello world, this is world   \t"))
	println("Trim right blank string: ", gouse.TrimBlank("   hello world, this is world   \n"))
}
```

#### 20. SampleStringTrimPrefix

```go
func SampleStringTrimPrefix() {
	var str = "   hello world, this is world   "
	println("Trim prefix string: ", gouse.TrimPrefix(str, "   "))
	println("Trim suffix string: ", gouse.TrimSuffix(str, "   "))
}
```

#### 21. SampleStringTrimSuffix

```go
func SampleStringTrimSuffix() {
	var str = "   hello world, this is world   "
	println("Trim suffix string: ", gouse.TrimSuffix(str, "   "))
}
```

#### 22. SampleStringJoin

```go
func SampleStringJoin() {
	var str = []string{"hello", "world"}
	println("Join string: ", gouse.Join(str, "-"))
}
```

#### 23. SampleStringConcat

```go
func SampleStringConcat() {
	println("Concat string: ", gouse.Concat("hello", "world"))
}
```

#### 24. SampleStringSubStr

```go
func SampleStringSubStr() {
	var str = "hello world, this is world"
	println("Sub string: ", gouse.SubStr(str, 0, 5))
	println("Sub string: ", gouse.SubStr(str, 0, 1))
	println("Sub string (only start): ", gouse.SubStr(str, -5))
	println("Sub string (with negative index): ", gouse.SubStr(str, -5, -1))
}
```

#### 25. SampleStringSlice

```go
func SampleStringSlice() {
	var str = "hello world, this is world"
	println("Slice string: ", gouse.Slice(str, 0, 5))
	println("Slice string: ", gouse.Slice(str, 0, 1))
	println("Slice string (only start): ", gouse.Slice(str, -5))
	println("Slice string (not parameters): ", gouse.Slice(str))
	println("Slice string (with negative index): ", gouse.Slice(str, -5, -1))
}
```

#### 26. SampleStringSplice

```go
func SampleStringSplice() {
	var str = "helloworld, this is world"
	println("Splice string (default not replace): ", gouse.Splice(str, 0, 5))
	println("Splice string (with replace): ", gouse.Splice(str, 1, 5, "golang"))
	println("Splice string (with replace multiple): ", gouse.Splice(str, 1, 5, "golang1", "golang2"))
}
```

#### 27. SampleStringStartsWith

```go
func SampleStringStartsWith() {
	var str = "hello world, this is world"
	println("Starts with: ", gouse.StartsWith(str, "hello"))
}
```

#### 28. SampleStringEndsWith

```go
func SampleStringEndsWith() {
	var str = "hello world, this is world"
	println("Ends with: ", gouse.EndsWith(str, "world"))
}
```

#### 29. SampleStringEscape

```go
func SampleStringEscape() {
	var str = "This is a <b>bold</b> statement & \"quote\" – © ®"
	println("Escape string: ", gouse.Escape(str))
}
```

#### 30. SampleStringUnescape

```go
func SampleStringUnescape() {
	var str = "This is a &lt;b&gt;bold&lt;/b&gt; statement &amp; &quot;quote&quot; – © ®"
	println("Unescape string: ", gouse.Unescape(str))
}
```

#### 31. SampleStringPad

```go
func SampleStringPad() {
	var str = "hello world"
	println("Pad-left string: ", gouse.PadStart(str, 20, '$'))
	println("Pad-right string: ", gouse.PadEnd(str, 20, '@'))
}
```

#### 32. SampleStringCount

```go
func SampleStringCount() {
	var str = "hello world wo wo"
	println("Count words/substr in string (default): ", gouse.Count(str))
	println("Count words/substr in string (with char): ", gouse.Count(str, "wo"))
}
```

#### 33. SampleStringLines

```go
func SampleStringLines() {
	var str = "hello world\nwo wo"
	println("Count lines of string: ", gouse.Lines(str))
}
```

#### 34. SampleStringIndex

```go
func SampleStringIndex() {
	var str = "hello world, this is world"

	f1, l1 := gouse.IndexSubStr(str, "l")
	fmt.Printf("First index start at: %d, end at: %d\n", f1, l1)

	f2, l2 := gouse.IndexSubStr(str, "world")
	fmt.Printf("First index start at: %d, end at: %d\n", f2, l2)

	f3 := gouse.FIndex(str, "l")
	fmt.Println("First index start at: ", f3)

	l3 := gouse.LIndex(str, "world")
	fmt.Println("Last index start at: ", l3)

	f4, l4 := gouse.IndexSubStr(str, "oo")
	if f4 == -1 || l4 == -1 {
		fmt.Println("Not found")
	}

	if gouse.FIndex(str, "oo") == -1 {
		fmt.Println("Not found")
	}

	if gouse.LIndex(str, "oo") == -1 {
		fmt.Println("Not found")
	}
}
```

#### 35. SampleStringRandom

```go
func SampleStringRandom() {
	println("Random chain string: ", gouse.RandStr(10))

	println("Random chain number: ", gouse.RandDigit(6))
}
```

#### 36. SampleStringAt

```go
func SampleStringAt() {
	var str = "hello world"
	println("At string: ", gouse.At(str, 1))
	println("At string: ", gouse.At(str, -5))
}
```

#### 37. SampleStringCodePointAt

```go
func SampleStringCodePointAt() {
	var str = "hello world"
	println("Code point at string: ", gouse.CodePointAt(str, 1))
	println("Code point at string: ", gouse.CodePointAt(str, -5))
}
```

#### 38. SampleStringCodePoint

```go
func SampleStringCodePoint() {
	asciiValues := gouse.CodePoint("hello world")

	print("Code point string: ")
	for _, asciiValue := range asciiValues {
		fmt.Printf("%d ", asciiValue)
	}
}
```

#### 39. SampleStringFromCodePointAt

```go
func SampleStringFromCodePointAt() {
	println("From code point at string: ", gouse.FromCodePointAt(9733))
	println("From code point at string: ", gouse.FromCodePointAt(9731))
}
```

#### 40. SampleStringFromCodePoint

```go
func SampleStringFromCodePoint() {
	println("From code point string: ", gouse.FromCodePoint(104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100))
}
```
