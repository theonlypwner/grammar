package wording

import (
	"fmt"
	"math/rand"
)

// Prefix

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
	{"I comment that", false},
	{"I argue that", false},
	{"I opine that", false},
	{"I maintain that", false},
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
	{"I am of the opinion that", false},
	{"it is in my opinion that", false},
	{"it is my opinion that", false},
	// sort of uncertain
	{"I guess", true},
}

func randPrefix() *msgPrefix {
	return &msgPrefixes[rand.Intn(len(msgPrefixes))]
}

func (p *msgPrefix) String(that bool) string {
	clause := p.clause + " "
	if that && p.that {
		clause += "that "
	}
	return clause
}

// Random verbs

func randInfinitive() (modal, verb string) {
	modal = choice("should", "ought to", "could", "can", "meant to", "intended to")
	// without "to"
	return modal, randInfinitiveVerb()
}
func randInfinitiveVerb() string {
	// without "to"
	return choice("use", "say", "tweet", "post", "type", "write")
}

func randPastPerfect() (modal, verb string) {
	modal = choice("should have", "ought to have", "could have")
	return modal, randPastVerb()
}
func randPastVerb() string {
	// (simple [past) perfect]
	return choice("used", "said", "tweeted", "posted", "typed")
}

// Random nouns

func randTweetNoun(includeArticle bool) string {
	if includeArticle {
		return choice("a tweet", "a post", "a status", "a message",
			"a status update", "an update")
	}
	return choice("tweet", "post", "status", "message",
		"status update", "update")
}

func randMistakeNoun() string {
	return choice("an error", "a mistake", "a solecism", "a typo")
}

type msgLoader func(secondPerson bool, clause string) (prefix, suffix string)

var msgLoaders = [...]msgLoader{
	func(secondPerson bool, c string) (string, string) {
		m, v := randPastPerfect()
		return cleft(yourPrepend(secondPerson, c), m+" "+v)
	},
	func(secondPerson bool, c string) (string, string) {
		m, v := randInfinitive()
		return cleft(yourPrepend(secondPerson, c), m+" "+v)
	},
	func(secondPerson bool, c string) (string, string) {
		return clauseAppend(yourPrepend(secondPerson, c),
				fmt.Sprintf("it %v %v better if ",
					choice("could", "might", "would"),
					choice("have been", "be"),
				)),
			"had " + randPastVerb()
	},
	func(secondPerson bool, c string) (string, string) {
		c = yourPrepend(secondPerson, c)
		s := ""
		if p(.50) {
			c = clauseAppend(c, "it is possible for ")
			s = "to " + randInfinitiveVerb()
		} else {
			c = clauseAppend(c, "it was possible for ")
			s = "to have " + randPastVerb()
		}
		return c, s
	},
	func(secondPerson bool, c string) (string, string) {
		var suffix, m, v string
		if p(.50) { // infinitive rather than perfect (50%)
			m, v = randInfinitive()
		} else {
			m, v = randPastPerfect()
		}
		if p(.50) { // perfect (have) instead of simple past (50%)
			h := "have"
			if !secondPerson {
				h = "has"
			}
			suffix = fmt.Sprintf("%v %v %v and %v %v",
				h,
				choice("written", "made", "created", "tweeted", "posted", "typed"),
				randMistakeNoun(),
				m,
				v,
			)
		} else {
			suffix = fmt.Sprintf("%v %v and %v %v",
				choice("wrote", "made", "created", "tweeted", "posted", "typed"),
				randMistakeNoun(),
				m,
				v,
			)
		}
		return cleft(c, suffix)
	},
	func(secondPerson bool, c string) (string, string) {
		var suffix, m, v string
		if p(.50) { // infinitive rather than perfect (50%)
			m, v = randInfinitive()
		} else {
			m, v = randPastPerfect()
		}
		if p(.50) { // perfect (have) instead of simple past (50%)
			h := "have"
			if !secondPerson {
				h = "has"
			}
			suffix = fmt.Sprintf("%v %v %v and %v %v",
				h,
				choice("miswritten", "botched", "blundered", "messed up", "malformed", "screwed up", "mistyped"),
				randTweetNoun(true),
				m,
				v,
			)
		} else {
			suffix = fmt.Sprintf("%v %v and %v %v",
				choice("miswrote", "botched", "blundered", "messed up", "malformed", "screwed up", "mistyped"),
				randTweetNoun(true),
				m,
				v,
			)
		}
		return cleft(c, suffix)
	},
	func(secondPerson bool, _ string) (string, string) {
		prefix := fmt.Sprintf("I %v %v %v %v ",
			choice("consider", "deem", "declare"),
			choice("the", "this"),
			randTweetNoun(false),
			choice("of", "by"),
		)
		suffix := choice("invalid", "incorrect", "wrong", "erroneous", "unacceptable", "unsuitable") +
			"; it should " + choice("be", "say", "read")
		return prefix, suffix
	},
	func(secondPerson bool, c string) (string, string) {
		prefix := fmt.Sprintf("I %v that ",
			choice("suggest", "recommend"),
		)
		suffix := randInfinitiveVerb()
		return prefix, suffix
	},
}

func randLoader() msgLoader {
	return msgLoaders[rand.Intn(len(msgLoaders))]
}

func clauseAppend(c, repl string) string {
	if p(.50) {
		// 50% chance to append to the first clause
		c += repl
	} else {
		c = repl
	}
	return c
}

func yourPrepend(ok bool, c string) string {
	if ok && p(.30) {
		return fmt.Sprintf("in %v %v, ",
			choice("your", "this"),
			randTweetNoun(false),
		)
	}
	return c
}

func cleft(prefix, suffix string) (string, string) {
	if p(.10) { // Cleft sentence (10%)
		prefix += "it is "
		suffix = choice("who ", "that ") + suffix
	}
	return prefix, suffix
}
