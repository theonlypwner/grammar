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

// firstCap makes the first character title-case.
func firstCap(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	if r != utf8.RuneError && !unicode.IsTitle(r) {
		return string(unicode.ToTitle(r)) + s[size:]
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
	opt := msgAlterOpt{
		p(.65), // secondPerson
	}

	// Build the sentence!
	clause := ""
	modals := modalsPerfect[:]
	verbs := saidPast[:]
	if p(.75) { // Add prefix (75%)
		prefix := msgPrefixes[rand.Intn(len(msgPrefixes))]
		clause = prefix.clause + " "
		if prefix.that && p(.50) { // use "that" (50%)
			clause += "that "
		}
	}
	if p(.40) { // Alter it! (40%)
		msgAlters[rand.Intn(len(msgAlters))](&opt, &clause, &modals, &verbs)
	}

	predicate := ""
	// choose
	if modal := choice(modals...); modal != "" {
		predicate += modal + " "
	}
	if verb := choice(verbs...); verb != "" {
		predicate += verb + " "
	}
	predicate += engJoin(corrections...) + " instead."

	// Build the entire sentence
	result := ""
	if opt.secondPerson { // 2nd person instead of 3rd (65%)
		if p(.85) {
			// Invert the subject so that we address one personally (85%)
			clause = user + ", " + clause
			user = "you"
		} else {
			user = fmt.Sprintf("you, %v,", user)
		}
	}
	if p(.10) { // Cleft sentence (10%)
		clause += "it is "
		user += choice(" who", " that")
		// TODO fix with "it is possible"
	}
	result = firstCap(clause) + user + " " + predicate

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

func choice(s ...string) string {
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

type msgAlterOpt struct{ secondPerson bool }

type msgAlter func(o *msgAlterOpt, clause *string, modals, verbs *[]string)

var msgAlters_HAD = [...]string{"had"}
var msgAlters_TO = [...]string{"to"}
var msgAlters_TO_HAVE = [...]string{"to have"}
var msgAlters = [...]msgAlter{
	func(_ *msgAlterOpt, _ *string, m, v *[]string) {
		*m = modalsInfinitive[:]
		*v = saidInfinitive[:]
	},
	func(_ *msgAlterOpt, c *string, m, v *[]string) {
		clauseAppend(c, fmt.Sprintf("it %v %v better if ",
			choice("could", "might", "would"),
			choice("have been", "be"),
		))
		*m = msgAlters_HAD[:]
		*v = saidPast[:]
	},
	func(_ *msgAlterOpt, c *string, m, v *[]string) {
		if p(.50) {
			clauseAppend(c, "it is possible for ")
			*m = msgAlters_TO[:]
			*v = saidInfinitive[:]
		} else {
			clauseAppend(c, "it was possible for ")
			*m = msgAlters_TO_HAVE[:]
			*v = saidPast[:]
		}
	},
	func(o *msgAlterOpt, _ *string, m, v *[]string) {
		if p(.50) { // infinitive rather than perfect (50%)
			*m = modalsInfinitive[:]
			*v = saidInfinitive[:]
		} else {
			*m = modalsPerfect[:]
			*v = saidPast[:]
		}
		if p(.50) { // perfect (have) instead of simple past (50%)
			h := "have"
			if !o.secondPerson {
				h = "has"
			}
			*m = []string{fmt.Sprintf("%v %v a%v and %v",
				h,
				choice("made", "created", "tweeted", "posted", "typed", "written"),
				choice("n error", " mistake", " solecism", " typo"),
				choice(*m...),
			)}
		} else {
			*m = []string{fmt.Sprintf("%v a%v and %v",
				choice("made", "created", "tweeted", "posted", "typed", "wrote"),
				choice("n error", " mistake", " solecism", " typo"),
				choice(*m...),
			)}
		}
	},
	func(o *msgAlterOpt, _ *string, m, v *[]string) {
		if p(.50) { // infinitive rather than perfect (50%)
			*m = modalsInfinitive[:]
			*v = saidInfinitive[:]
		} else {
			*m = modalsPerfect[:]
			*v = saidPast[:]
		}
		if p(.50) { // perfect (have) instead of simple past (50%)
			h := "have"
			if !o.secondPerson {
				h = "has"
			}
			*m = []string{fmt.Sprintf("%v %v %v and %v",
				h,
				choice("botched", "blundered", "messed up", "malformed", "screwed up", "mistyped", "miswritten"),
				choice(tweetNoun[:]...),
				choice(*m...),
			)}
		} else {
			*m = []string{fmt.Sprintf("%v %v and %v",
				choice("botched", "blundered", "messed up", "malformed", "screwed up", "mistyped", "miswrote"),
				choice(tweetNoun[:]...),
				choice(*m...),
			)}
		}
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
