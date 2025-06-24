package gouse

import (
	"fmt"
	"html"
)

func Capitalize(sentence string) string {
	sentenceBytes := []byte(sentence)

	for i := range sentenceBytes {
		if (i == 0 || sentenceBytes[i-1] == ' ') && IsLetter(sentenceBytes[i]) {
			sentenceBytes[i] = Upper(sentenceBytes[i])
		}
	}

	return string(sentenceBytes)
}

func CamelCase(s string) string {
	var result string
	upperNext := false

	for _, char := range s {
		if IsLetter(char) || IsDigit(char) {
			if upperNext {
				result += string(Upper(char))
				upperNext = false
			} else {
				result += string(Lower(char))
			}
		} else {
			upperNext = true
		}
	}

	return result
}

func concatCase(s string, sep string) string {
	spliStr := Split(s, " ")
	var result string

	if len(spliStr) == 1 {
		for i, char := range s {
			if i > 0 && (IsUpper(char) || IsDigit(char)) && !IsDigit(s[i-1]) {
				result += sep
			} else if i > 0 && !IsLetter(char) && !IsDigit(char) && string(char) != sep {
				result += sep
				continue
			}
			result += string(Lower(char))
		}

		return result
	} else {
		for i, str := range spliStr {
			if i > 0 {
				result += sep
			}
			result += str
		}

		return result
	}
}

func SnakeCase(s string) string {
	return concatCase(s, "_")
}

func KebabCase(s string) string {
	return concatCase(s, "-")
}

func SpaceCase(s string) string {
	return concatCase(s, " ")
}

func CustomCase(s, sep string) string {
	return concatCase(s, sep)
}

func IsLetter[T byte | rune](char T) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func IsDigit[T byte | rune](char T) bool {
	return char >= '0' && char <= '9'
}

func Includes(s string, substr string) bool {
	return FirstIdx(s, substr) != -1
}

func StartsWith(s string, substr string) bool {
	return FirstIdx(s, substr) == 0
}

func EndsWith(s string, substr string) bool {
	if len(s) < len(substr) {
		return false
	}
	return LastIdx(s, substr) == len(s)-len(substr)
}

func IsLower[T byte | rune](char T) bool {
	return char >= 'a' && char <= 'z'
}

func IsUpper[T byte | rune](char T) bool {
	return char >= 'A' && char <= 'Z'
}

func Split(s string, separator string) []string {
	var result []string
	var temp string

	if len(separator) == 0 {
		for _, v := range s {
			result = append(result, string(v))
		}
	} else {
		for _, v := range s {
			if string(v) == separator {
				result = append(result, temp)
				temp = ""
			} else {
				temp += string(v)
			}
		}
		result = append(result, temp)
	}

	return result
}

func Words(s string) []string {
	var result []string

	for _, v := range s {
		if v == ' ' || v == '\n' || v == '\t' {
			continue
		} else {
			result = append(result, string(v))
		}
	}

	return result
}

func Reverse(s string) string {
	var result string

	for _, v := range s {
		result = string(v) + result
	}

	return result
}

func Lowers(s string) string {
	var result string

	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			result += string(v + 32)
		} else {
			result += string(v)
		}
	}

	return result
}

func Lower[T byte | rune](char T) T {
	if char >= 'A' && char <= 'Z' {
		return char - 'A' + 'a'
	}
	return char
}

func LowerFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	return string(Lower(s[0])) + s[1:]
}

func Uppers(s string) string {
	var result string

	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			result += string(v - 32)
		} else {
			result += string(v)
		}
	}

	return result
}

func Upper[T byte | rune](char T) T {
	if char >= 'a' && char <= 'z' {
		return char - 'a' + 'A'
	}
	return char
}

func UpperFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	return string(Upper(s[0])) + s[1:]
}

func Repeat(s string, count int) string {
	var result string

	for range count {
		result += s
	}

	return result

}

func Truncate(s string, length int, omission ...string) string {
	if len(s) <= length {
		return s
	}

	if len(omission) == 0 {
		return s[:length] + "..."
	}

	if len(omission) > length {
		return s
	}

	return s[:length] + omission[0]
}

func Replace(s string, old string, new string) string {
	if len(old) == 0 {
		return s
	}

	var result string

	for i := 0; i < len(s); {
		// Check if the remaining part of the input matches the old
		if len(s)-i >= len(old) && s[i:i+len(old)] == old {
			// If there's a match, append the newSubstr to the result and move the index past the old
			result += new
			i += len(old)
		} else {
			// If no match, append the current character to the result and move to the next character
			result += string(s[i])
			i++
		}
	}

	return result
}

func LTrim(s string) string {
	var result string
	var flag bool

	for _, v := range s {
		if v != ' ' {
			flag = true
			result += string(v)
		} else if flag {
			result += string(v)
		}
	}

	return result
}

func RTrim(s string) string {
	var result string
	var flag bool

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			flag = true
			result = string(s[i]) + result
		} else if flag {
			result = string(s[i]) + result
		}
	}

	return result
}

func Trim(s string) string {
	tmp := LTrim(s)
	return RTrim(tmp)
}

func TrimBlank(s string) string {
	start := 0
	end := len(s) - 1

	for start <= end && (s[start] == ' ' || s[start] == '\t' || s[start] == '\n') {
		start++
	}

	for end >= start && (s[end] == ' ' || s[end] == '\t' || s[end] == '\n') {
		end--
	}

	return s[start : end+1]
}

func TrimPre(s, prefix string) string {
	if len(s) < len(prefix) {
		return s
	}

	if s[:len(prefix)] == prefix {
		return s[len(prefix):]
	}

	return s
}

func TrimSuf(s, suffix string) string {
	if len(s) < len(suffix) {
		return s
	}

	if s[len(s)-len(suffix):] == suffix {
		return s[:len(s)-len(suffix)]
	}

	return s
}

func Join(s []string, separator string) string {
	var result string

	for i, v := range s {
		result += v
		if i != len(s)-1 {
			result += separator
		}
	}

	return result
}

func Concat(s ...string) string {
	var result string

	for _, v := range s {
		result += v
	}

	return result
}

func SubStr(s string, start int, length ...int) string {
	if len(length) > 0 {
		if start < 0 {
			start = len(s) + start
		}

		if length[0] < 0 || start == length[0] || start > len(s) {
			return ""
		}

		if start+length[0] > len(s) || start == length[0]-1 {
			if (start >= len(s)) || (length[0] >= len(s)) {
				if start >= len(s) && length[0] >= len(s) {
					return ""
				} else {
					if length[0] >= len(s) {
						return s[start:]
					} else if start >= len(s) {
						return ""
					}
				}
			} else {
				if start == length[0]-1 {
					return string(s[start])
				}
			}
		}

		if start+length[0] == len(s) {
			if start == 0 {
				return s
			} else {
				return s[start : start+length[0]-1]
			}
		}

		return s[start : start+length[0]]
	}
	return s[start:]
}

func Slice(s string, index ...int) string {
	var start, end int
	if len(index) == 0 {
		return s
	}

	if len(index) == 1 {
		start = index[0]
		end = len(s)
	} else {
		start = index[0]
		end = index[1]
	}

	if start < 0 {
		start = len(s) + start
	}

	if end < 0 {
		end = len(s) + end
	}

	if end > len(s) {
		end = len(s)
	}

	if start > end || start == end {
		return ""
	}

	if start == end-1 {
		return string(s[start])
	}

	return s[start:end]
}

func Splice(s string, start, replaceCount int, items ...string) string {
	if len(items) == 0 {
		return s[:start] + s[start+replaceCount:]
	}

	if start < 0 {
		start = len(s) + start
	}

	if replaceCount < 0 {
		replaceCount = 0
	}

	if start > len(s) {
		start = len(s)
	}

	if replaceCount > len(s)-start {
		replaceCount = len(s) - start
	}

	if replaceCount > len(items) {
		replaceCount = len(items) - 1
	}

	return s[:start] + Join(items, "") + s[start+replaceCount:]
}

func Esc(s string) string {
	var result string

	for _, v := range s {
		if v == '\\' || v == '/' || v == '[' || v == ']' || v == '{' || v == '}' || v == '(' || v == ')' || v == '*' || v == '+' || v == '?' || v == '.' || v == '^' || v == '$' || v == '|' {
			result += "\\" + string(v)
		}

		switch v {
		case '&':
			result += "&amp;"
		case '<':
			result += "&lt;"
		case '>':
			result += "&gt;"
		case '"':
			result += "&quot;"
		case '\'':
			result += "&#39;"
		case '\u2013': // En dash
			result += "&ndash;"
		case '\u2014': // Em dash
			result += "&mdash;"
		case '\u2018', '\u2019': // Left and right single quotation mark
			result += "&lsquo;"
		case '\u201C', '\u201D': // Left and right double quotation mark
			result += "&ldquo;"
		case '\u00A9': // Copyright symbol
			result += "&copy;"
		case '\u00AE': // Registered trademark symbol
			result += "&reg;"
		default:
			result += string(v)
		}
	}

	return result
}

func Unesc(s string) string {
	s = html.UnescapeString(s)

	for i := 0; i < len(s); i++ {
		if s[i] == '\\' {
			if i+1 < len(s) {
				i++
			}
		}
	}

	return s
}

func pad(s string, addAmount int, padChar byte) string {
	padding := make([]byte, addAmount-len(s))
	for i := range padding {
		padding[i] = padChar
	}

	return string(padding)
}

func LPad(s string, addAmount int, padChar byte) string {
	if len(s) >= addAmount {
		return s
	}
	return pad(s, addAmount, padChar) + s
}

func RPad(s string, addAmount int, padChar byte) string {
	if len(s) >= addAmount {
		return s
	}

	return s + pad(s, addAmount, padChar)
}

func Count(s string, substr ...string) int {
	if len(substr) == 0 {
		return len(s)
	}

	var count int

	for i := 0; i < len(s); i++ {
		if s[i] == substr[0][0] && len(s)-i >= len(substr[0]) {
			if s[i:i+len(substr[0])] == substr[0] {
				count++
				i += len(substr[0]) - 1
			}
		}
	}

	return count
}

func Lines(s string) int {
	var result []string
	var temp string

	for _, v := range s {
		if v == '\n' {
			result = append(result, temp)
			temp = ""
		} else {
			temp += string(v)
		}
	}

	result = append(result, temp)

	return len(result)
}

func IdxSubStr(s, substr string) (int, int) {
	firstIndex := -1
	lastIndex := -1

	for i := range len(s) {
		if len(s)-i < len(substr) {
			break
		}

		if s[i:i+len(substr)] == substr {
			if firstIndex == -1 {
				firstIndex = i
			}
			lastIndex = i
		}
	}

	return firstIndex, lastIndex
}

func FirstIdx(s string, substr string) int {
	firstIndex, _ := IdxSubStr(s, substr)
	return firstIndex
}

func LastIdx(s string, substr string) int {
	_, lastIndex := IdxSubStr(s, substr)
	return lastIndex
}

func At(s string, index int) string {
	if index < 0 {
		index = len(s) + index
	}

	if index < 0 || index >= len(s) {
		return ""
	}

	return string(s[index])
}

func CodePointAt(s string, index int) int {
	if index < 0 {
		index = len(s) + index
	}

	if index < 0 || index >= len(s) {
		return -1
	}

	return int(s[index])
}

func CodePoint(input string) []int {
	asciiValues := make([]int, len(input))

	for i, char := range input {
		asciiValues[i] = int(char)
	}

	return asciiValues
}

func FromCodePointAt(codePoint int) string {
	return string(rune(codePoint))
}

func FromCodePoint(codePoint ...int) string {
	var result string

	for _, v := range codePoint {
		result += fmt.Sprintf("%c", v)
	}

	return result
}

func MatchStr(text string, pattern string) int {
	if len(pattern) == 0 {
		return 0
	}

	LPS := func(pattern string) []int {
		lps := make([]int, len(pattern))
		i := 1
		j := 0
		for i < len(pattern) {
			if pattern[i] == pattern[j] {
				j++
				lps[i] = j
				i++
			} else {
				if j != 0 {
					j = lps[j-1]
				} else {
					lps[i] = 0
					i++
				}
			}
		}
		return lps
	}

	lps := LPS(pattern)
	i := 0
	j := 0
	for i < len(text) {
		if text[i] == pattern[j] {
			i++
			j++
		}
		if j == len(pattern) {
			return i - j
		} else if i < len(text) && text[i] != pattern[j] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return -1
}
