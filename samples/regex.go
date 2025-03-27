package samples

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
)

/*
Description: Check if the input string is match with regex pattern
Input params: (pattern, input string)
--> Note: regex pattern is not include one of (^, $, /g)
*/
func RegexIsMatch() {
	fmt.Println("Match string with regex: ", gouse.IsMatchReg(`[a-z]+\s[a-z]+`, "hello world"))
}

/*
Description: Find the first index and value of the match regex pattern
Input params: (pattern, input string)
--> Note: regex pattern is not include one of (^, $, /g)
*/
func RegexMatchIndex() {
	paragraph := "I think Ruth's dog is cuter than your dog!"
	matchIdx := gouse.MatchIdxReg(`[^\w\s']`, paragraph)
	if matchIdx != -1 {
		fmt.Printf("Match with regex (index: %d, value: %s)\n", matchIdx, string(paragraph[matchIdx]))
	} else {
		println("Not found index match regex")
	}
}

/*
Description: Find all index and value of the match regex pattern
Input params: (pattern, input string)
--> Note: regex pattern is not include one of (^, $, /g)
*/
func RegexMatch() {
	fmt.Println("Match string with regex: ", gouse.MatchReg(`[A-Z]`, "Hello World 123"))
}
