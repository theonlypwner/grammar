// Package wording makes human-readable text for corrections.
package wording

import (
	"fmt"
	"math/rand"
	"unicode/utf8"
)

// MakeTweet randomly generates a tweet message to correct a user.
func MakeTweet(corrections, reasons []string, user string) string {
	// Build the sentence!
	clause := ""
	modal := ""
	verb := ""
	if p(.75) { // Add prefix (75%)
		prefix := msgPrefixes[rand.Intn(len(msgPrefixes))]
		clause = prefix.clause + " "
		if prefix.that && p(.50) { // use "that" (50%)
			clause += "that "
		}
	}
	secondPerson := clause == "" || p(.65)
	if secondPerson { // 2nd person instead of 3rd (65%)
		user = fmt.Sprintf("you, %v,", user)
	}
	randLoader()(secondPerson, &clause, &modal, &verb)

	// Build the entire sentence
	result := clause +
		user + " " +
		modal + " " +
		verb + " " +
		engJoin(corrections...) + " instead."
	result = firstCap(result)

	// Explain why, if we have space
	if len(reasons) != 0 {
		remain := 140 - utf8.RuneCountInString(result)
		// at least 3 characters have to be added
		if remain >= 3 {
			why := firstCap(engJoin(reasons...))
			if remain >= 2+utf8.RuneCountInString(why) {
				result += " " + why + "."
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
