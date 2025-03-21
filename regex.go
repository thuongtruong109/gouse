package gouse

import "regexp"

func IsMatchReg(regex, chain string) bool {
	return regexp.MustCompile(regex).MatchString(chain)
}

func MatchReg(regex, chain string) []string {
	var result []string
	for _, v := range chain {
		if IsMatchReg(regex, string(v)) {
			result = append(result, string(v))
		}
	}
	return result
}

func MatchIdxReg(regex, chain string) int {
	for i, v := range chain {
		if IsMatchReg(regex, string(v)) {
			return i
		}
	}

	return -1
}
