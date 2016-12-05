// Package grammar generates grammar corrections.
package grammar

import (
	"victorz.ca/grammar/parser"
	"victorz.ca/grammar/sequence"
	"victorz.ca/grammar/transform"
	"victorz.ca/grammar/wording"
)

// Load runs some text through the transformer, lexer, parser, and formatter.
// If nothing is detected, corrections will be nil.
// Reasons for the corrections should be stored in why, but it may be nil.
func Load(s string) (corrections, why []string) {
	s = transform.DoAll(s)
	seq := sequence.New(s)
	why = parser.DoAll(&seq)
	if len(why) != 0 {
		corrections = seq.CorrectionsQuoted()
	}
	return
}

// MakeTweetReply generates a response for a tweet.
// If nothing was detected, it returns an empty string.
func MakeTweetReply(tweet, author string) string {
	corrections, why := Load(tweet)
	if len(corrections) == 0 {
		return ""
	}
	return wording.MakeTweet(corrections, why, author)
}
