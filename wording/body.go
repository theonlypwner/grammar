package wording

import (
	"fmt"
	"math/rand"
)

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
var tweetNounBase = [...]string{
	// singular
	"tweet", "post", "status", "message",
	"status update", "update",
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
	// I found that, I notice (that)
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

var msgLoaders_HaveBeen_Be = [...]string{"have been", "be"}
var msgLoadersWrote = [...]string{"wrote", "made", "created", "tweeted", "posted", "typed", "written"}
var msgLoadersWritePerfect = msgLoadersWrote[1:]
var msgLoadersWritePast = msgLoadersWrote[:len(msgLoadersWrote)-1]
var msgLoaders_Mistake = [...]string{"an error", "a mistake", "a solecism", "a typo"}
var msgLoaders_MistakeVerb = [...]string{"miswrote", "botched", "blundered", "messed up", "malformed", "screwed up", "mistyped", "miswritten"}
var msgLoaders_MistakeVerbPerfect = msgLoaders_MistakeVerb[1:]
var msgLoaders_MistakeVerbPast = msgLoaders_MistakeVerb[:len(msgLoaders_MistakeVerb)-1]
var msgLoaders = [...]msgLoader{
	func(secondPerson bool, c, m, v *string) {
		if secondPerson {
			yourReplace(c)
		}
		*m = choice(modalsPerfect[:]...)
		*v = choice(saidPast[:]...)
		cleft(c, m)
	},
	func(secondPerson bool, c, m, v *string) {
		if secondPerson {
			yourReplace(c)
		}
		*m = choice(modalsInfinitive[:]...)
		*v = choice(saidInfinitive[:]...)
		cleft(c, m)
	},
	func(secondPerson bool, c, m, v *string) {
		if secondPerson {
			yourReplace(c)
		}
		clauseAppend(c, fmt.Sprintf("it %v %v better if ",
			choice("could", "might", "would"),
			choice(msgLoaders_HaveBeen_Be[:]...),
		))
		*m = "had"
		*v = choice(saidPast[:]...)
	},
	func(secondPerson bool, c, m, v *string) {
		if secondPerson {
			yourReplace(c)
		}
		if p(.50) {
			clauseAppend(c, "it is possible for ")
			*m = "to"
			*v = choice(saidInfinitive[:]...)
		} else {
			clauseAppend(c, "it was possible for ")
			*m = "to have"
			*v = choice(saidPast[:]...)
		}
	},
	func(secondPerson bool, c, m, v *string) {
		if p(.50) { // infinitive rather than perfect (50%)
			*m = choice(modalsInfinitive[:]...)
			*v = choice(saidInfinitive[:]...)
		} else {
			*m = choice(modalsPerfect[:]...)
			*v = choice(saidPast[:]...)
		}
		if p(.50) { // perfect (have) instead of simple past (50%)
			h := "have"
			if !secondPerson {
				h = "has"
			}
			*m = fmt.Sprintf("%v %v %v and %v",
				h,
				choice(msgLoadersWritePerfect...),
				choice(msgLoaders_Mistake[:]...),
				*m,
			)
		} else {
			*m = fmt.Sprintf("%v %v and %v",
				choice(msgLoadersWritePast...),
				choice(msgLoaders_Mistake[:]...),
				*m,
			)
		}
		cleft(c, m)
	},
	func(secondPerson bool, c, m, v *string) {
		if p(.50) { // infinitive rather than perfect (50%)
			*m = choice(modalsInfinitive[:]...)
			*v = choice(saidInfinitive[:]...)
		} else {
			*m = choice(modalsPerfect[:]...)
			*v = choice(saidPast[:]...)
		}
		if p(.50) { // perfect (have) instead of simple past (50%)
			h := "have"
			if !secondPerson {
				h = "has"
			}
			*m = fmt.Sprintf("%v %v %v and %v",
				h,
				choice(msgLoaders_MistakeVerbPerfect...),
				choice(tweetNoun[:]...),
				*m,
			)
		} else {
			*m = fmt.Sprintf("%v %v and %v",
				choice(msgLoaders_MistakeVerbPast...),
				choice(tweetNoun[:]...),
				*m,
			)
		}
		cleft(c, m)
	},
	func(secondPerson bool, c, m, v *string) {
		*c = fmt.Sprintf("I %v the %v %v ",
			choice("consider", "deem", "declare"),
			choice(tweetNounBase[:]...),
			choice("of", "by"),
		)
		*m = choice("invalid", "incorrect", "wrong", "erroneous", "unacceptable", "unsuitable") + ";"
		*v = "it should " + choice("be", "say", "read")
	},
}

func randLoader() msgLoader {
	return msgLoaders[rand.Intn(len(msgLoaders))]
}

func clauseAppend(c *string, repl string) {
	if p(.50) {
		// 50% chance to append to the first clause
		*c += repl
	} else {
		*c = repl
	}
}

func yourReplace(c *string) {
	if p(.30) {
		*c = fmt.Sprintf("in your %v, ", choice(tweetNounBase[:]...))
	}
}

func cleft(clause, modal *string) {
	if p(.10) { // Cleft sentence (10%)
		*clause += "it is "
		*modal = choice("who ", "that ") + *modal
	}
}
