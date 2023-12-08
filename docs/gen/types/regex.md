
### IsInt
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
func IsInt(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "int")
}
```

### IsFloat
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
func IsFloat(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "float")
}
```

### IsString
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
func IsString(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "string")
}
```

### IsBool
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
func IsBool(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "bool")
}
```

### IsSlice
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
func IsSlice(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "[]")
}
```

### IsMap
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
func IsMap(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "map")
}
```

### IsStruct
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
func IsStruct(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "struct")
}
```

### IsArray
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
func IsArray(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "[")
}
```

### IsPointer
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
func IsPointer(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "*")
}
```

### IsFunc
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
func IsFunc(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "func")
}
```

### IsNil
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
func IsNil(v interface{}) bool {
	return v == nil
}
```

### IsEmpty
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
func IsEmpty(v interface{}) bool {
	if v == nil {
		return true
	}
	switch v.(type) {
	case string:
		return v.(string) == ""
	case int:
		return v.(int) == 0
	case int8:
		return v.(int8) == 0
	case int16:
		return v.(int16) == 0
	case int32:
		return v.(int32) == 0
	case int64:
		return v.(int64) == 0
	case uint:
		return v.(uint) == 0
	case uint8:
		return v.(uint8) == 0
	case uint16:
		return v.(uint16) == 0
	case uint32:
		return v.(uint32) == 0
	case uint64:
		return v.(uint64) == 0
	case float32:
		return v.(float32) == 0
	case float64:
		return v.(float64) == 0
	case bool:
		return !v.(bool)
	case []interface{}:
		return len(v.([]interface{})) == 0
	case map[string]interface{}:
		return len(v.(map[string]interface{})) == 0
	default:
		return false
	}
}
```

### IsUndefined
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
func IsUndefined(v interface{}) bool {
	return v == nil || IsEmpty(v)
}
```

### IsZero
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
func IsZero(v interface{}) bool {
	return v == nil || IsEmpty(v)
}
```

### IsUUID
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
func IsUUID(input string) (bool, error) {
	if input == "" {
		return false, constants.ErrRequiredUUID
	}

	_, err := uuid.Parse(input)
	if err != nil {
		return false, constants.ErrInvalidUUID
	}

	return true, nil
}
```

### StructToString
```go
import (
	"fmt"

	"reflect"
)
```

```go
func StructToString(data interface{}) string {
	v := reflect.ValueOf(data)
	t := v.Type()

	var fields []string

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i).Interface()
		fields = append(fields, fmt.Sprintf("%s: %v", fieldName, fieldValue))
	}

	// replace with array.Join()
	var result string
	for i, v := range fields {
		result += v
		if i != len(fields)-1 {
			result += ", "
		}
	}

	return fmt.Sprintf("%s{%s}", t.Name(), result)
}
```

### StructToMap
```go
import (
	"fmt"

	"reflect"
)
```

```go
func StructToMap(data interface{}) map[string]interface{} {
	v := reflect.ValueOf(data)
	t := v.Type()

	result := make(map[string]interface{})

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i).Interface()
		result[fieldName] = fieldValue
	}

	return result
}
```

### StringToInt
```go
import (
	"fmt"

	"reflect"
)
```

```go
func StringToInt(data string) int {
	var result int
	fmt.Sscanf(data, "%d", &result)
	return result
}
```

### StringToFloat
```go
import (
	"fmt"

	"reflect"
)
```

```go
func StringToFloat(data string) float64 {
	var result float64
	fmt.Sscanf(data, "%f", &result)
	return result
}
```

### StringToBool
```go
import (
	"fmt"

	"reflect"
)
```

```go
func StringToBool(data string) bool {
	var result bool
	fmt.Sscanf(data, "%t", &result)
	return result
}
```

### IntToString
```go
import (
	"fmt"

	"reflect"
)
```

```go
func IntToString(data int) string {
	return fmt.Sprintf("%d", data)
}
```

### FloatToString
```go
import (
	"fmt"

	"reflect"
)
```

```go
func FloatToString(data float64) string {
	return fmt.Sprintf("%f", data)
}
```

### BoolToString
```go
import (
	"fmt"

	"reflect"
)
```

```go
func BoolToString(data bool) string {
	return fmt.Sprintf("%t", data)
}
```

### InterfaceToString
```go
import (
	"fmt"

	"reflect"
)
```

```go
func InterfaceToString(data interface{}) string {
	return fmt.Sprintf("%v", data)
}
```

### BytesToString
```go
import (
	"fmt"

	"reflect"
)
```

```go
func BytesToString(data []byte) string {
	return string(data)
}
```

### StringToBytes
```go
import (
	"fmt"

	"reflect"
)
```

```go
func StringToBytes(data string) []byte {
	return []byte(data)
}
```

### StringsToBytes
```go
import (
	"fmt"

	"reflect"
)
```

```go
func StringsToBytes(data []string) []byte {
	var result []byte
	for _, v := range data {
		result = append(result, StringToBytes(v)...)
	}
	return result
}
```

### MatchRegex
```go
import (
	"regexp"
)
```

```go
func MatchRegex(regex, chain string) bool {
	return regexp.MustCompile(regex).MatchString(chain)
}
```