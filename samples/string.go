package samples

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
)

/*
Description: Capitalize the first letter of a string.
Input params: (s string)
*/
func StringCapitalize() {
	var str = "hello world"
	println(gouse.Capitalize(str))
}

/*
Description: Convert string to Camel Case.
Input params: (s string)
*/
func StringCamelCase() {
	var str = "convert_any-string TO-camelCase in Go! 123"
	println("Convert string to Camel Case: ", gouse.CamelCase(str))
}

/*
Description: Convert string to Snake Case.
Input params: (s string)
*/
func StringSnakeCase() {
	var str1 = "hello world"
	var str2 = "ConvertThisString123"
	println("Convert string to Snake Case: ", gouse.SnakeCase(str1), gouse.SnakeCase(str2))
}

/*
Description: Convert string to Kebab Case.
Input params: (s string)
*/
func StringKebabCase() {
	var str = "hello world"
	println("Convert string to Kebab Case: ", gouse.KebabCase(str))
}

/*
Description: Convert string to Space Case.
Input params: (s string)
*/
func StringSpaceCase() {
	var str = "hello world"
	println("Convert string to Space Case: ", gouse.SpaceCase(str))
}

/*
Description: Convert string to custom.
Input params: (s string, separator string)
*/
func StringCustomCase() {
	var str = "hello world"
	println("Convert string to Custom Case: ", gouse.CustomCase(str, "%"))
}

/*
Description: Check is letter character.
Input params: (s string)
*/
func StringIsLetter() {
	var str = "hello world"
	println("Check is letter character: ", gouse.IsLetter(str[0]))
}

/*
Description: Check is number character.
Input params: (s string)
*/
func StringIsDigit() {
	var str = "1hello world"
	println("Check is number character: ", gouse.IsDigit(str[0]))
}

/*
Description: Check is contain substring in string.
Input params: (s string, substr string)
*/
func StringIncludes() {
	var str = "hello world, this is world"
	println("Check substring in string: ", gouse.Includes(str, "world"))
}

/*
Description: Check is lower string.
Input params: (s string)
*/
func StringIsLower() {
	var str = "hELLO WORLD"
	println("Check is lower string: ", gouse.IsLower(str[0]))
}

/*
Description: Check is upper string.
Input params: (s string)
*/
func StringIsUpper() {
	var str = "Hello world"
	println("Check is upper string: ", gouse.IsUpper(str[0]))
}

/*
Description: Split string by separator.
Input params: (s string, separator string)
*/
func StringSplit() {
	var str = "hello world"
	fmt.Println("Split string by separator: ", gouse.Split(str, "l"))
}

/*
Description: Split string by single word
Input params: (s string)
*/
func StringWords() {
	var str = "hello world"
	println("Split string to array of words: ", gouse.Words(str))

}

/*
Description: Reverse string.
Input params: (s string)
*/
func StringReverse() {
	var str = "hello world"
	println("Reverse string: ", gouse.Reverse(str))
}

/*
Description: Convert string to lower case.
Input params: (s string)
*/
func StringLower() {
	var str = "HELLO WORLD"
	println("Lower string (string): ", gouse.Lowers(str))
	println("Lower string (byte): ", gouse.Lower(str[0]))
	println("Lower first string: ", gouse.LowerFirst(str))
}

/*
Description: Convert string to upper case.
Input params: (s string)
*/
func StringUpper() {
	var str = "hello world"
	println("Upper string (string): ", gouse.Uppers(str))
	println("Upper string (byte): ", gouse.Upper(str[0]))
	println("Upper first string: ", gouse.UpperFirst(str))
}

/*
Description: Repeat string n times.
Input params: (struct interface{}, fieldName string, value interface{})
*/
func StringRepeat() {
	var str = "hello world"
	println("Repeat string: ", gouse.Repeat(str, 3))
}

/*
Description: Truncate string to n characters or add suffix.
Input params: (s string, length int, suffix ...string)
*/
func StringTruncate() {
	var str = "hello world"
	println("Truncate string (default): ", gouse.Truncate(str, 5))
	println("Truncate string (custom): ", gouse.Truncate(str, 5, "***"))
}

/*
Description: Replace substring in string.
Input params: (s string, old string, new string)
*/
func StringReplace() {
	var str = "hello world, this is world"
	println("Replace string: ", gouse.Replace(str, "world", "golang"))
}

/*
Description: Trim left and right space in string.
Input params: (s string)
*/
func StringTrim() {
	var str = "   hello world, this is world   "
	println("Trim string: ", gouse.Trim(str))
	println("Trim left string: ", gouse.LTrim(str))
	println("Trim right string: ", gouse.RTrim(str))
}

/*
Description: Trim left and right blank in string.
Input params: (s string)
*/
func StringTrimBlank() {
	println("Trim blank string: ", gouse.TrimBlank("   hello world, this is world   "))
	println("Trim left blank string: ", gouse.TrimBlank("   hello world, this is world   \t"))
	println("Trim right blank string: ", gouse.TrimBlank("   hello world, this is world   \n"))
}

/*
Description: Trim left and right prefix in string.
Input params: (s string, prefix string)
*/
func StringTrimPrefix() {
	var str = "   hello world, this is world   "
	println("Trim prefix string: ", gouse.TrimPre(str, "   "))
	println("Trim suffix string: ", gouse.TrimSuf(str, "   "))
}

/*
Description: Trim left and right suffix in string.
Input params: (s string, suffix string)
*/
func StringTrimSuffix() {
	var str = "   hello world, this is world   "
	println("Trim suffix string: ", gouse.TrimSuf(str, "   "))
}

/*
Description: Join string array to string.
Input params: (s []string, separator string)
*/
func StringJoin() {
	var str = []string{"hello", "world"}
	println("Join string: ", gouse.Join(str, "-"))
}

/*
Description: Concat string.
Input params: (s ...string)
*/
func StringConcat() {
	println("Concat string: ", gouse.Concat("hello", "world"))
}

/*
Description: Get substring from string.
Input params: (s string, start int, end ...int)
*/
func StringSubStr() {
	var str = "hello world, this is world"
	println("Sub string: ", gouse.SubStr(str, 0, 5))
	println("Sub string: ", gouse.SubStr(str, 0, 1))
	println("Sub string (only start): ", gouse.SubStr(str, -5))
	println("Sub string (with negative index): ", gouse.SubStr(str, -5, -1))
}

/*
Description: Slice string from start to end.
Input params: (s string, start int, end ...int)
*/
func StringSlice() {
	var str = "hello world, this is world"
	println("Slice string: ", gouse.Slice(str, 0, 5))
	println("Slice string: ", gouse.Slice(str, 0, 1))
	println("Slice string (only start): ", gouse.Slice(str, -5))
	println("Slice string (not parameters): ", gouse.Slice(str))
	println("Slice string (with negative index): ", gouse.Slice(str, -5, -1))
}

/*
Description: Splice string from start to end.
Input params: (s string, start int, end int, replace ...string)
*/
func StringSplice() {
	var str = "helloworld, this is world"
	println("Splice string (default not replace): ", gouse.Splice(str, 0, 5))
	println("Splice string (with replace): ", gouse.Splice(str, 1, 5, "golang"))
	println("Splice string (with replace multiple): ", gouse.Splice(str, 1, 5, "golang1", "golang2"))
}

/*
Description: Check is start with substring in string.
Input params: (s string, subStr string)
*/
func StringStartsWith() {
	var str = "hello world, this is world"
	println("Starts with: ", gouse.StartsWith(str, "hello"))
}

/*
Description: Check is end with substring in string.
Input params: (s string, subStr string)
*/
func StringEndsWith() {
	var str = "hello world, this is world"
	println("Ends with: ", gouse.EndsWith(str, "world"))
}

/*
Description: Convert string to HTML entities.
Input params: (s string)
*/
func StringEscape() {
	var str = "This is a <b>bold</b> statement & \"quote\" – © ®"
	println("Escape string: ", gouse.Esc(str))
}

/*
Description: Convert HTML entities to string.
Input params: (s string)
*/
func StringUnescape() {
	var str = "This is a &lt;b&gt;bold&lt;/b&gt; statement &amp; &quot;quote&quot; – © ®"
	println("Unescape string: ", gouse.Unesc(str))
}

/*
Description: Add padding to string from left or right.
Input params: (s string, length int, padStr string)
*/
func StringPad() {
	var str = "hello world"
	println("Pad-left string: ", gouse.LPad(str, 20, '$'))
	println("Pad-right string: ", gouse.RPad(str, 20, '@'))
}

/*
Description: Count words/substr in string.
Input params: (s string, substr ...string)
*/
func StringCount() {
	var str = "hello world wo wo"
	println("Count words/substr in string (default): ", gouse.Count(str))
	println("Count words/substr in string (with char): ", gouse.Count(str, "wo"))
}

/*
Description: Count lines in string.
Input params: (s string)
*/
func StringLines() {
	var str = "hello world\nwo wo"
	println("Count lines of string: ", gouse.Lines(str))
}

/*
Description: Get index of substring in string.
Input params: (s string, substr string)
*/
func StringIndex() {
	var str = "hello world, this is world"

	f1, l1 := gouse.IdxSubStr(str, "l")
	fmt.Printf("First index start at: %d, end at: %d\n", f1, l1)

	f2, l2 := gouse.IdxSubStr(str, "world")
	fmt.Printf("First index start at: %d, end at: %d\n", f2, l2)

	f3 := gouse.FirstIdx(str, "l")
	fmt.Println("First index start at: ", f3)

	l3 := gouse.LastIdx(str, "world")
	fmt.Println("Last index start at: ", l3)

	f4, l4 := gouse.IdxSubStr(str, "oo")
	if f4 == -1 || l4 == -1 {
		fmt.Println("Not found")
	}

	if gouse.FirstIdx(str, "oo") == -1 {
		fmt.Println("Not found")
	}

	if gouse.LastIdx(str, "oo") == -1 {
		fmt.Println("Not found")
	}
}

/*
Description: Random chain string.
Input params: (length int)
*/
func StringRandomChain() {
	println("Random chain string: ", gouse.RandStr(10))
	println("Random chain number: ", gouse.RandDigit(6))
}

/*
Description: Get character at index in string.
Input params: (s string, index int)
*/
func StringAt() {
	var str = "hello world"
	println("At string: ", gouse.At(str, 1))
	println("At string: ", gouse.At(str, -5))
}

/*
Description: Get code point at index in string.
Input params: (s string, index int)
*/
func StringCodePointAt() {
	var str = "hello world"
	println("Code point at string: ", gouse.CodePointAt(str, 1))
	println("Code point at string: ", gouse.CodePointAt(str, -5))
}

/*
Description: Get code point of string.
Input params: (s string)
*/
func StringCodePoint() {
	asciiValues := gouse.CodePoint("hello world")

	print("Code point string: ")
	for _, asciiValue := range asciiValues {
		fmt.Printf("%d ", asciiValue)
	}
}

/*
Description: Transform code point to string.
Input params: (codePoint int)
*/
func StringFromCodePointAt() {
	println("From code point at string: ", gouse.FromCodePointAt(9733))
	println("From code point at string: ", gouse.FromCodePointAt(9731))
}

/*
Description: Transform multi code points to string.
Input params: (codePoints ...int)
*/
func StringFromCodePoint() {
	println("From code point string: ", gouse.FromCodePoint(104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100))
}
