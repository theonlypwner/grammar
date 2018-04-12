// Package wording makes human-readable text for corrections.
package wording

import (
	"fmt"
	"unicode/utf8"
)

// MakeTweet randomly generates a tweet message to correct a user.
func MakeTweet(corrections, reasons []string, user string) string {
	// Build the sentence!
	clause := ""
	if p(.75) { // Add prefix (75%)
		clause = randPrefix().String(p(.50)) // use "that" (50%)
	}
	secondPerson := clause == "" || p(.65)
	if secondPerson { // 2nd person instead of 3rd (65%)
		user = fmt.Sprintf("you, %v,", user)
	}
	prefix, suffix := randLoader()(secondPerson, clause)

	// Build the entire sentence
	result := prefix +
		user + " " +
		suffix + " " +
		engJoin(corrections) + " instead."
	result = firstCap(result, true)

	// Explain why, if we have space
	if len(reasons) != 0 {
		remain := 280 - utf8.RuneCountInString(result)
		// at least 3 characters have to be added per item
		for i := 0; remain >= 3 && i < len(reasons); i++ {
			why := firstCap(reasons[i], false)
			nextLen := 2 + utf8.RuneCountInString(why)
			if remain >= nextLen {
				result += " " + why + "."
				remain -= nextLen
			}
		}
	}

	return result
}

/*
// MakePost generates an extended description for the corrections.
// This method is currently unimplemented.
func MakePost(corrections, reasons []string, user string) string {
	return MakeTweet(corrections, reasons, user)
}
*/
