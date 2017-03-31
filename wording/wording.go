// Package wording makes human-readable text for corrections.
package wording

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// firstCap makes the first letter title-case.
func firstCap(s string) string {
	result := []rune(s)
	for i, r := range result {
		if unicode.IsLetter(r) {
			if !unicode.IsTitle(r) {
				result[i] = unicode.ToTitle(r)
				return string(result)
			}
			break
		}
	}
	return s
}

// englishJoin joins a (slice of strings) using an English conjunction.
func englishJoin(conjunction string, a ...string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return a[0]
	case 2:
		return a[0] + " " + conjunction + " " + a[1]
	default:
		commaJoined := strings.Join(a[:len(a)-1], ", ")
		// omit serial comma because Twitter only allows 140 characters
		return englishJoin(conjunction, commaJoined, a[len(a)-1])
	}
}

// engJoin calls englishJoin with "and"
func engJoin(a ...string) string {
	return englishJoin("and", a...)
}

// MakeTweet randomly generates a tweet message to correct a user.
func MakeTweet(corrections, reasons []string, user string) string {
	secondPerson := p(.65)

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
	msgLoaders[rand.Intn(len(msgLoaders))](secondPerson, &clause, &modal, &verb)

	// Build the entire sentence
	result := ""
	if secondPerson { // 2nd person instead of 3rd (65%)
		if p(.85) {
			// Invert the subject so that we address one personally (85%)
			clause = "to " + user + ", " + clause
			user = "you"
		} else {
			user = fmt.Sprintf("you, %v,", user)
		}
	} else if clause == "" {
		clause = "to "
	}
	result = clause +
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

func p(x float64) bool { return rand.Float64() < x }

func choice(s []string) string {
	if len(s) == 0 {
		return ""
	}
	return s[rand.Intn(len(s))]
}

// Constants
var modalsInfinitive = [...]string{
	"should", "ought to", "could", "can", "meant to", "intended to",
}
var modalsPerfect = [...]string{
	"should have", "ought to have", "could have",
}
var saidPast = [...]string{
	// (simple [past) perfect]
	"used", "said", "tweeted", "posted", "typed",
}
var saidInfinitive = [...]string{
	// without "to"
	"use", "say", "tweet", "post", "type", "write",
}
var tweetNoun = [...]string{
	// singular
	"a tweet", "a post", "a status", "a message",
	"a status update", "an update",
}

type msgPrefix struct {
	clause string
	that   bool
}

var msgPrefixes = [...]msgPrefix{
	// confident
	{"it is the case that", false},
	{"it is true that", false},
	{"in this case,", false},
	// in ___'s/your tweet,
	// in the tweet of ___,
	{"I am confident", true},
	{"I am sure", true},
	{"I say", true},
	{"I claim", true},
	{"I aver that", false},
	{"I assert that", false},
	{"I contend that", false},
	{"I declare that", false},
	{"I insist that", false},
	{"I note that", false},
	{"I state that", false},
	{"I noticed", true},
	{"I discovered", true},
	{"I see", true},
	// weaker
	{"it seems", true},
	{"to me, it seems", true},
	{"it seems to me", true},
	{"it appears", true},
	{"to me, it appears", true},
	{"it appears to me", true},
	{"it seems like", false},
	{"it looks like", false},
	{"it seems to be the case that", false},
	{"it appears to be the case that", false},
	{"it seems to be true that", false},
	{"it appears to be true that", false},
	{"it is likely that", false},
	{"it is probable that", false},
	// weak
	{"I think", true},
	{"I believe", true},
	{"I reckon", true},
	{"I suppose", true},
	{"I suspect", true},
	{"I feel", true},
	{"it is my opinion that", false},
	// sort of uncertain
	{"I guess", true},
}

type msgLoader func(secondPerson bool, clause, modal, verb *string)

var msgLoadersCouldMightWould = [...]string{"could", "might", "would"}
var msgLoaders_HaveBeen_Be = [...]string{"have been", "be"}
var msgLoadersWrote = [...]string{"wrote", "made", "created", "tweeted", "posted", "typed", "written"}
var msgLoadersWritePerfect = msgLoadersWrote[1:]
var msgLoadersWritePast = msgLoadersWrote[:len(msgLoadersWrote)-1]
var msgLoaders_Mistake = [...]string{"an error", "a mistake", "a solecism", "a typo"}
var msgLoaders_MistakeVerb = [...]string{"miswrote", "botched", "blundered", "messed up", "malformed", "screwed up", "mistyped", "miswritten"}
var msgLoaders_MistakeVerbPerfect = msgLoaders_MistakeVerb[1:]
var msgLoaders_MistakeVerbPast = msgLoaders_MistakeVerb[:len(msgLoaders_MistakeVerb)-1]
var msgLoaders = [...]msgLoader{
	func(_ bool, c, m, v *string) {
		*m = choice(modalsPerfect[:])
		*v = choice(saidPast[:])
		cleft(c, m)
	},
	func(_ bool, c, m, v *string) {
		*m = choice(modalsInfinitive[:])
		*v = choice(saidInfinitive[:])
		cleft(c, m)
	},
	func(_ bool, c, m, v *string) {
		clauseAppend(c, fmt.Sprintf("it %v %v better if ",
			choice(msgLoadersCouldMightWould[:]),
			choice(msgLoaders_HaveBeen_Be[:]),
		))
		*m = "had"
		*v = choice(saidPast[:])
	},
	func(_ bool, c, m, v *string) {
		if p(.50) {
			clauseAppend(c, "it is possible for ")
			*m = "to"
			*v = choice(saidInfinitive[:])
		} else {
			clauseAppend(c, "it was possible for ")
			*m = "to have"
			*v = choice(saidPast[:])
		}
	},
	func(secondPerson bool, c, m, v *string) {
		if p(.50) { // infinitive rather than perfect (50%)
			*m = choice(modalsInfinitive[:])
			*v = choice(saidInfinitive[:])
		} else {
			*m = choice(modalsPerfect[:])
			*v = choice(saidPast[:])
		}
		if p(.50) { // perfect (have) instead of simple past (50%)
			h := "have"
			if !secondPerson {
				h = "has"
			}
			*m = fmt.Sprintf("%v %v %v and %v",
				h,
				choice(msgLoadersWritePerfect),
				choice(msgLoaders_Mistake[:]),
				*m,
			)
		} else {
			*m = fmt.Sprintf("%v %v and %v",
				choice(msgLoadersWritePast),
				choice(msgLoaders_Mistake[:]),
				*m,
			)
		}
		cleft(c, m)
	},
	func(secondPerson bool, c, m, v *string) {
		if p(.50) { // infinitive rather than perfect (50%)
			*m = choice(modalsInfinitive[:])
			*v = choice(saidInfinitive[:])
		} else {
			*m = choice(modalsPerfect[:])
			*v = choice(saidPast[:])
		}
		if p(.50) { // perfect (have) instead of simple past (50%)
			h := "have"
			if !secondPerson {
				h = "has"
			}
			*m = fmt.Sprintf("%v %v %v and %v",
				h,
				choice(msgLoaders_MistakeVerbPerfect),
				choice(tweetNoun[:]),
				*m,
			)
		} else {
			*m = fmt.Sprintf("%v %v and %v",
				choice(msgLoaders_MistakeVerbPast),
				choice(tweetNoun[:]),
				*m,
			)
		}
		cleft(c, m)
	},
}

func clauseAppend(c *string, repl string) {
	if p(.50) {
		// 50% chance to append to the first clause
		*c += repl
	} else {
		*c = repl
	}
}

var cleftModals = [...]string{"who ", "that "}

func cleft(clause, modal *string) {
	if p(.10) { // Cleft sentence (10%)
		*clause += "it is "
		*modal = choice(cleftModals[:]) + *modal
	}
}
