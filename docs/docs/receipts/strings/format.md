# Format

## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringSplit

```go
func SampleStringSplit() {
	var str = "hello world"
	fmt.Println("Split string by separator: ", strings.Split(str, "l"))
}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringWords

```go
func SampleStringWords() {
	var str = "hello world"
	println("Split string to array of words: ", strings.Words(str))

}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringReverse

```go
func SampleStringReverse() {
	var str = "hello world"
	println("Reverse string: ", strings.Reverse(str))
}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringLower

```go
func SampleStringLower() {
	var str = "HELLO WORLD"
	println("Lower string (string): ", strings.Lowers(str))
	println("Lower string (byte): ", strings.Lower(str[0]))
	println("Lower first string: ", strings.LowerFirst(str))
}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringUpper

```go
func SampleStringUpper() {
	var str = "hello world"
	println("Upper string (string): ", strings.Uppers(str))
	println("Upper string (byte): ", strings.Upper(str[0]))
	println("Upper first string: ", strings.UpperFirst(str))
}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringRepeat

```go
func SampleStringRepeat() {
	var str = "hello world"
	println("Repeat string: ", strings.Repeat(str, 3))
}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringTruncate

```go
func SampleStringTruncate() {
	var str = "hello world"
	println("Truncate string (default): ", strings.Truncate(str, 5))
	println("Truncate string (custom): ", strings.Truncate(str, 5, "***"))
}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringReplace

```go
func SampleStringReplace() {
	var str = "hello world, this is world"
	println("Replace string: ", strings.Replace(str, "world", "golang"))
}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringTrim

```go
func SampleStringTrim() {
	var str = "   hello world, this is world   "
	println("Trim string: ", strings.Trim(str))
	println("Trim left string: ", strings.LTrim(str))
	println("Trim right string: ", strings.RTrim(str))
}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringTrimBlank

```go
func SampleStringTrimBlank() {
	println("Trim blank string: ", strings.TrimBlank("   hello world, this is world   "))
	println("Trim left blank string: ", strings.TrimBlank("   hello world, this is world   \t"))
	println("Trim right blank string: ", strings.TrimBlank("   hello world, this is world   \n"))
}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringTrimPrefix

```go
func SampleStringTrimPrefix() {
	var str = "   hello world, this is world   "
	println("Trim prefix string: ", strings.TrimPrefix(str, "   "))
	println("Trim suffix string: ", strings.TrimSuffix(str, "   "))
}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringTrimSuffix

```go
func SampleStringTrimSuffix() {
	var str = "   hello world, this is world   "
	println("Trim suffix string: ", strings.TrimSuffix(str, "   "))
}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringJoin

```go
func SampleStringJoin() {
	var str = []string{"hello", "world"}
	println("Join string: ", strings.Join(str, "-"))
}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringConcat

```go
func SampleStringConcat() {
	println("Concat string: ", strings.Concat("hello", "world"))
}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringSubStr

```go
func SampleStringSubStr() {
	var str = "hello world, this is world"
	println("Sub string: ", strings.SubStr(str, 0, 5))
	println("Sub string: ", strings.SubStr(str, 0, 1))
	println("Sub string (only start): ", strings.SubStr(str, -5))
	println("Sub string (with negative index): ", strings.SubStr(str, -5, -1))
}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringSlice

```go
func SampleStringSlice() {
	var str = "hello world, this is world"
	println("Slice string: ", strings.Slice(str, 0, 5))
	println("Slice string: ", strings.Slice(str, 0, 1))
	println("Slice string (only start): ", strings.Slice(str, -5))
	println("Slice string (not parameters): ", strings.Slice(str))
	println("Slice string (with negative index): ", strings.Slice(str, -5, -1))
}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringSplice

```go
func SampleStringSplice() {
	var str = "helloworld, this is world"
	println("Splice string (default not replace): ", strings.Splice(str, 0, 5))
	println("Splice string (with replace): ", strings.Splice(str, 1, 5, "golang"))
	println("Splice string (with replace multiple): ", strings.Splice(str, 1, 5, "golang1", "golang2"))
}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringStartsWith

```go
func SampleStringStartsWith() {
	var str = "hello world, this is world"
	println("Starts with: ", strings.StartsWith(str, "hello"))
}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringEndsWith

```go
func SampleStringEndsWith() {
	var str = "hello world, this is world"
	println("Ends with: ", strings.EndsWith(str, "world"))
}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringEscape

```go
func SampleStringEscape() {
	var str = "This is a <b>bold</b> statement & \"quote\" – © ®"
	println("Escape string: ", strings.Escape(str))
}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringUnescape

```go
func SampleStringUnescape() {
	var str = "This is a &lt;b&gt;bold&lt;/b&gt; statement &amp; &quot;quote&quot; – © ®"
	println("Unescape string: ", strings.Unescape(str))
}```
## Imports

```go
import (
	"fmt"	"github.com/thuongtruong109/gouse/strings")
```
## Functions


### SampleStringPad

```go
func SampleStringPad() {
	var str = "hello world"
	println("Pad-left string: ", strings.PadStart(str, 20, '$'))
	println("Pad-right string: ", strings.PadEnd(str, 20, '@'))
}```
