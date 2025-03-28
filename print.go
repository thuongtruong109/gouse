package gouse

import (
	"os"
	"strconv"
)

func Sprint(args ...any) string {
	var result string
	for i, arg := range args {
		result += ToStr(arg)
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
		// If the current character is '%', it's a format specifier
		if format[i] == '%' {
			// Check if it's a literal '%' (i.e. '%%')
			if i+1 < len(format) && format[i+1] == '%' {
				result += "%"
				i++ // Skip the next '%'
				continue
			}

			// Check if we've run out of arguments and append "invalid"
			if index >= len(args) {
				result += "invalid"
				i++ // Skip over the format specifier (d, s, etc.)
				continue
			}

			// Otherwise, process the format specifier
			i++ // Move to the format specifier (e.g., 'd', 's', etc.)
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
				// For unknown format specifiers, just append "invalid"
				result += "invalid"
			}
			index++ // Move to the next argument
		} else {
			// Regular character, just append it
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
