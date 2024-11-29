
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.25rem 0.75rem 0.25rem 0;' type='info' text='🔖 String' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

### <Badge style='font-size: 1.1rem;' type='tip' text='1. sample string capitalize' />



```go
func SampleStringCapitalize() {
	var str = "hello world"
	println(gouse.Capitalize(str))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='2. sample string camel case' />



```go
func SampleStringCamelCase() {
	var str = "convert_any-string TO-camelCase in Go! 123"
	println("Convert string to Camel Case: ", gouse.CamelCase(str))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='3. sample string snake case' />



```go
func SampleStringSnakeCase() {
	var str1 = "hello world"
	var str2 = "ConvertThisString123"
	println("Convert string to Snake Case: ", gouse.SnakeCase(str1), gouse.SnakeCase(str2))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='4. sample string kebab case' />



```go
func SampleStringKebabCase() {
	var str = "hello world"
	println("Convert string to Kebab Case: ", gouse.KebabCase(str))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='5. sample string is letter' />



```go
func SampleStringIsLetter() {
	var str = "hello world"
	println("Check is letter character: ", gouse.IsLetter(str[0]))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='6. sample string is digit' />



```go
func SampleStringIsDigit() {
	var str = "1hello world"
	println("Check is number character: ", gouse.IsDigit(str[0]))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='7. sample string includes' />



```go
func SampleStringIncludes() {
	var str = "hello world, this is world"
	println("Check substring in string: ", gouse.Includes(str, "world"))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='8. sample string is lower' />



```go
func SampleStringIsLower() {
	var str = "hELLO WORLD"
	println("Check is lower string: ", gouse.IsLower(str[0]))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='9. sample string is upper' />



```go
func SampleStringIsUpper() {
	var str = "Hello world"
	println("Check is upper string: ", gouse.IsUpper(str[0]))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='10. sample string split' />



```go
func SampleStringSplit() {
	var str = "hello world"
	fmt.Println("Split string by separator: ", gouse.Split(str, "l"))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='11. sample string words' />



```go
func SampleStringWords() {
	var str = "hello world"
	println("Split string to array of words: ", gouse.Words(str))

}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='12. sample string reverse' />



```go
func SampleStringReverse() {
	var str = "hello world"
	println("Reverse string: ", gouse.Reverse(str))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='13. sample string lower' />



```go
func SampleStringLower() {
	var str = "HELLO WORLD"
	println("Lower string (string): ", gouse.Lowers(str))
	println("Lower string (byte): ", gouse.Lower(str[0]))
	println("Lower first string: ", gouse.LowerFirst(str))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='14. sample string upper' />



```go
func SampleStringUpper() {
	var str = "hello world"
	println("Upper string (string): ", gouse.Uppers(str))
	println("Upper string (byte): ", gouse.Upper(str[0]))
	println("Upper first string: ", gouse.UpperFirst(str))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='15. sample string repeat' />



```go
func SampleStringRepeat() {
	var str = "hello world"
	println("Repeat string: ", gouse.Repeat(str, 3))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='16. sample string truncate' />



```go
func SampleStringTruncate() {
	var str = "hello world"
	println("Truncate string (default): ", gouse.Truncate(str, 5))
	println("Truncate string (custom): ", gouse.Truncate(str, 5, "***"))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='17. sample string replace' />



```go
func SampleStringReplace() {
	var str = "hello world, this is world"
	println("Replace string: ", gouse.Replace(str, "world", "golang"))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='18. sample string trim' />



```go
func SampleStringTrim() {
	var str = "   hello world, this is world   "
	println("Trim string: ", gouse.Trim(str))
	println("Trim left string: ", gouse.LTrim(str))
	println("Trim right string: ", gouse.RTrim(str))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='19. sample string trim blank' />



```go
func SampleStringTrimBlank() {
	println("Trim blank string: ", gouse.TrimBlank("   hello world, this is world   "))
	println("Trim left blank string: ", gouse.TrimBlank("   hello world, this is world   \t"))
	println("Trim right blank string: ", gouse.TrimBlank("   hello world, this is world   \n"))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='20. sample string trim prefix' />



```go
func SampleStringTrimPrefix() {
	var str = "   hello world, this is world   "
	println("Trim prefix string: ", gouse.TrimPrefix(str, "   "))
	println("Trim suffix string: ", gouse.TrimSuffix(str, "   "))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='21. sample string trim suffix' />



```go
func SampleStringTrimSuffix() {
	var str = "   hello world, this is world   "
	println("Trim suffix string: ", gouse.TrimSuffix(str, "   "))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='22. sample string join' />



```go
func SampleStringJoin() {
	var str = []string{"hello", "world"}
	println("Join string: ", gouse.Join(str, "-"))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='23. sample string concat' />



```go
func SampleStringConcat() {
	println("Concat string: ", gouse.Concat("hello", "world"))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='24. sample string sub str' />



```go
func SampleStringSubStr() {
	var str = "hello world, this is world"
	println("Sub string: ", gouse.SubStr(str, 0, 5))
	println("Sub string: ", gouse.SubStr(str, 0, 1))
	println("Sub string (only start): ", gouse.SubStr(str, -5))
	println("Sub string (with negative index): ", gouse.SubStr(str, -5, -1))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='25. sample string slice' />



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

### <Badge style='font-size: 1.1rem;' type='tip' text='26. sample string splice' />



```go
func SampleStringSplice() {
	var str = "helloworld, this is world"
	println("Splice string (default not replace): ", gouse.Splice(str, 0, 5))
	println("Splice string (with replace): ", gouse.Splice(str, 1, 5, "golang"))
	println("Splice string (with replace multiple): ", gouse.Splice(str, 1, 5, "golang1", "golang2"))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='27. sample string starts with' />



```go
func SampleStringStartsWith() {
	var str = "hello world, this is world"
	println("Starts with: ", gouse.StartsWith(str, "hello"))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='28. sample string ends with' />



```go
func SampleStringEndsWith() {
	var str = "hello world, this is world"
	println("Ends with: ", gouse.EndsWith(str, "world"))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='29. sample string escape' />



```go
func SampleStringEscape() {
	var str = "This is a <b>bold</b> statement & \"quote\" – © ®"
	println("Escape string: ", gouse.Escape(str))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='30. sample string unescape' />



```go
func SampleStringUnescape() {
	var str = "This is a &lt;b&gt;bold&lt;/b&gt; statement &amp; &quot;quote&quot; – © ®"
	println("Unescape string: ", gouse.Unescape(str))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='31. sample string pad' />



```go
func SampleStringPad() {
	var str = "hello world"
	println("Pad-left string: ", gouse.PadStart(str, 20, '$'))
	println("Pad-right string: ", gouse.PadEnd(str, 20, '@'))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='32. sample string count' />



```go
func SampleStringCount() {
	var str = "hello world wo wo"
	println("Count words/substr in string (default): ", gouse.Count(str))
	println("Count words/substr in string (with char): ", gouse.Count(str, "wo"))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='33. sample string lines' />



```go
func SampleStringLines() {
	var str = "hello world\nwo wo"
	println("Count lines of string: ", gouse.Lines(str))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='34. sample string index' />



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

### <Badge style='font-size: 1.1rem;' type='tip' text='35. sample string random' />



```go
func SampleStringRandom() {
	println("Random chain string: ", gouse.RandStr(10))

	println("Random chain number: ", gouse.RandDigit(6))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='36. sample string at' />



```go
func SampleStringAt() {
	var str = "hello world"
	println("At string: ", gouse.At(str, 1))
	println("At string: ", gouse.At(str, -5))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='37. sample string code point at' />



```go
func SampleStringCodePointAt() {
	var str = "hello world"
	println("Code point at string: ", gouse.CodePointAt(str, 1))
	println("Code point at string: ", gouse.CodePointAt(str, -5))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='38. sample string code point' />



```go
func SampleStringCodePoint() {
	asciiValues := gouse.CodePoint("hello world")

	print("Code point string: ")
	for _, asciiValue := range asciiValues {
		fmt.Printf("%d ", asciiValue)
	}
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='39. sample string from code point at' />



```go
func SampleStringFromCodePointAt() {
	println("From code point at string: ", gouse.FromCodePointAt(9733))
	println("From code point at string: ", gouse.FromCodePointAt(9731))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='40. sample string from code point' />



```go
func SampleStringFromCodePoint() {
	println("From code point string: ", gouse.FromCodePoint(104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100))
}
```
