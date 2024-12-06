
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 String' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

## 1. String capitalize



```go
func StringCapitalize() {
	var str = "hello world"
	println(gouse.Capitalize(str))
}
```

## 2. String camel case



```go
func StringCamelCase() {
	var str = "convert_any-string TO-camelCase in Go! 123"
	println("Convert string to Camel Case: ", gouse.CamelCase(str))
}
```

## 3. String snake case



```go
func StringSnakeCase() {
	var str1 = "hello world"
	var str2 = "ConvertThisString123"
	println("Convert string to Snake Case: ", gouse.SnakeCase(str1), gouse.SnakeCase(str2))
}
```

## 4. String kebab case



```go
func StringKebabCase() {
	var str = "hello world"
	println("Convert string to Kebab Case: ", gouse.KebabCase(str))
}
```

## 5. String is letter



```go
func StringIsLetter() {
	var str = "hello world"
	println("Check is letter character: ", gouse.IsLetter(str[0]))
}
```

## 6. String is digit



```go
func StringIsDigit() {
	var str = "1hello world"
	println("Check is number character: ", gouse.IsDigit(str[0]))
}
```

## 7. String includes



```go
func StringIncludes() {
	var str = "hello world, this is world"
	println("Check substring in string: ", gouse.Includes(str, "world"))
}
```

## 8. String is lower



```go
func StringIsLower() {
	var str = "hELLO WORLD"
	println("Check is lower string: ", gouse.IsLower(str[0]))
}
```

## 9. String is upper



```go
func StringIsUpper() {
	var str = "Hello world"
	println("Check is upper string: ", gouse.IsUpper(str[0]))
}
```

## 10. String split



```go
func StringSplit() {
	var str = "hello world"
	fmt.Println("Split string by separator: ", gouse.Split(str, "l"))
}
```

## 11. String words



```go
func StringWords() {
	var str = "hello world"
	println("Split string to array of words: ", gouse.Words(str))

}
```

## 12. String reverse



```go
func StringReverse() {
	var str = "hello world"
	println("Reverse string: ", gouse.Reverse(str))
}
```

## 13. String lower



```go
func StringLower() {
	var str = "HELLO WORLD"
	println("Lower string (string): ", gouse.Lowers(str))
	println("Lower string (byte): ", gouse.Lower(str[0]))
	println("Lower first string: ", gouse.LowerFirst(str))
}
```

## 14. String upper



```go
func StringUpper() {
	var str = "hello world"
	println("Upper string (string): ", gouse.Uppers(str))
	println("Upper string (byte): ", gouse.Upper(str[0]))
	println("Upper first string: ", gouse.UpperFirst(str))
}
```

## 15. String repeat



```go
func StringRepeat() {
	var str = "hello world"
	println("Repeat string: ", gouse.Repeat(str, 3))
}
```

## 16. String truncate



```go
func StringTruncate() {
	var str = "hello world"
	println("Truncate string (default): ", gouse.Truncate(str, 5))
	println("Truncate string (custom): ", gouse.Truncate(str, 5, "***"))
}
```

## 17. String replace



```go
func StringReplace() {
	var str = "hello world, this is world"
	println("Replace string: ", gouse.Replace(str, "world", "golang"))
}
```

## 18. String trim



```go
func StringTrim() {
	var str = "   hello world, this is world   "
	println("Trim string: ", gouse.Trim(str))
	println("Trim left string: ", gouse.LTrim(str))
	println("Trim right string: ", gouse.RTrim(str))
}
```

## 19. String trim blank



```go
func StringTrimBlank() {
	println("Trim blank string: ", gouse.TrimBlank("   hello world, this is world   "))
	println("Trim left blank string: ", gouse.TrimBlank("   hello world, this is world   \t"))
	println("Trim right blank string: ", gouse.TrimBlank("   hello world, this is world   \n"))
}
```

## 20. String trim prefix



```go
func StringTrimPrefix() {
	var str = "   hello world, this is world   "
	println("Trim prefix string: ", gouse.TrimPrefix(str, "   "))
	println("Trim suffix string: ", gouse.TrimSuffix(str, "   "))
}
```

## 21. String trim suffix



```go
func StringTrimSuffix() {
	var str = "   hello world, this is world   "
	println("Trim suffix string: ", gouse.TrimSuffix(str, "   "))
}
```

## 22. String join



```go
func StringJoin() {
	var str = []string{"hello", "world"}
	println("Join string: ", gouse.Join(str, "-"))
}
```

## 23. String concat



```go
func StringConcat() {
	println("Concat string: ", gouse.Concat("hello", "world"))
}
```

## 24. String sub str



```go
func StringSubStr() {
	var str = "hello world, this is world"
	println("Sub string: ", gouse.SubStr(str, 0, 5))
	println("Sub string: ", gouse.SubStr(str, 0, 1))
	println("Sub string (only start): ", gouse.SubStr(str, -5))
	println("Sub string (with negative index): ", gouse.SubStr(str, -5, -1))
}
```

## 25. String slice



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

## 26. String splice



```go
func StringSplice() {
	var str = "helloworld, this is world"
	println("Splice string (default not replace): ", gouse.Splice(str, 0, 5))
	println("Splice string (with replace): ", gouse.Splice(str, 1, 5, "golang"))
	println("Splice string (with replace multiple): ", gouse.Splice(str, 1, 5, "golang1", "golang2"))
}
```

## 27. String starts with



```go
func StringStartsWith() {
	var str = "hello world, this is world"
	println("Starts with: ", gouse.StartsWith(str, "hello"))
}
```

## 28. String ends with



```go
func StringEndsWith() {
	var str = "hello world, this is world"
	println("Ends with: ", gouse.EndsWith(str, "world"))
}
```

## 29. String escape



```go
func StringEscape() {
	var str = "This is a <b>bold</b> statement & \"quote\" – © ®"
	println("Escape string: ", gouse.Escape(str))
}
```

## 30. String unescape



```go
func StringUnescape() {
	var str = "This is a &lt;b&gt;bold&lt;/b&gt; statement &amp; &quot;quote&quot; – © ®"
	println("Unescape string: ", gouse.Unescape(str))
}
```

## 31. String pad



```go
func StringPad() {
	var str = "hello world"
	println("Pad-left string: ", gouse.PadStart(str, 20, '$'))
	println("Pad-right string: ", gouse.PadEnd(str, 20, '@'))
}
```

## 32. String count



```go
func StringCount() {
	var str = "hello world wo wo"
	println("Count words/substr in string (default): ", gouse.Count(str))
	println("Count words/substr in string (with char): ", gouse.Count(str, "wo"))
}
```

## 33. String lines



```go
func StringLines() {
	var str = "hello world\nwo wo"
	println("Count lines of string: ", gouse.Lines(str))
}
```

## 34. String index



```go
func StringIndex() {
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

## 35. String random



```go
func StringRandom() {
	println("Random chain string: ", gouse.RandStr(10))

	println("Random chain number: ", gouse.RandDigit(6))
}
```

## 36. String at



```go
func StringAt() {
	var str = "hello world"
	println("At string: ", gouse.At(str, 1))
	println("At string: ", gouse.At(str, -5))
}
```

## 37. String code point at



```go
func StringCodePointAt() {
	var str = "hello world"
	println("Code point at string: ", gouse.CodePointAt(str, 1))
	println("Code point at string: ", gouse.CodePointAt(str, -5))
}
```

## 38. String code point



```go
func StringCodePoint() {
	asciiValues := gouse.CodePoint("hello world")

	print("Code point string: ")
	for _, asciiValue := range asciiValues {
		fmt.Printf("%d ", asciiValue)
	}
}
```

## 39. String from code point at



```go
func StringFromCodePointAt() {
	println("From code point at string: ", gouse.FromCodePointAt(9733))
	println("From code point at string: ", gouse.FromCodePointAt(9731))
}
```

## 40. String from code point



```go
func StringFromCodePoint() {
	println("From code point string: ", gouse.FromCodePoint(104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100))
}
```
