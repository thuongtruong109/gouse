package gouse

import (
	"os"
	"strconv"
)

func valueToString(value any) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'f', 2, 64)
	case bool:
		return strconv.FormatBool(v)
	default:
		return "unknown"
	}
}

func Sprint(args ...any) string {
	var result string
	for i, arg := range args {
		result += valueToString(arg)
		if i < len(args)-1 {
			result += " "
		}
	}
	return result
}

func Sprintln(args ...any) string {
	return Sprint(args...) + "\n"
}

func Sprintf(format string, args ...any) string {
	var result string
	index := 0

	for i := 0; i < len(format); i++ {
		if format[i] == '%' && i+1 < len(format) {
			i++
			if index >= len(args) {
				result += "%" + string(format[i])
				continue
			}
			switch format[i] {
			case 'd':
				if v, ok := args[index].(int); ok {
					result += strconv.Itoa(v)
				} else {
					result += "invalid"
				}
			case 's':
				if v, ok := args[index].(string); ok {
					result += v
				} else {
					result += "invalid"
				}
			case 'f':
				if v, ok := args[index].(float64); ok {
					result += strconv.FormatFloat(v, 'f', 2, 64)
				} else if v, ok := args[index].(float32); ok {
					result += strconv.FormatFloat(float64(v), 'f', 2, 64)
				} else {
					result += "invalid"
				}
			case 't':
				if v, ok := args[index].(bool); ok {
					result += strconv.FormatBool(v)
				} else {
					result += "invalid"
				}
			case '%':
				result += "%"
				index--
			case 'T':
				if v, ok := args[index].(bool); ok {
					if v {
						result += "true"
					} else {
						result += "false"
					}
				} else {
					result += "invalid"
				}
			case 'c':
				if v, ok := args[index].(int); ok {
					result += string(rune(v))
				} else {
					result += "invalid"
				}
			default:
				result += "%" + string(format[i])
			}
			index++
		} else {
			result += string(format[i])
		}
	}
	return result
}

func Println(args ...any) {
	result := Sprintln(args...)
	os.Stdout.WriteString(result)
}

func Print(args ...any) {
	result := Sprint(args...)
	os.Stdout.WriteString(result)
}

func Printf(format string, args ...any) {
	result := Sprintf(format, args...)
	os.Stdout.WriteString(result)
}
