// Package parser applies rules and makes a list of corrections
package parser

import (
	"sequence"

	"math/rand"
)

type checkGroup int

const (
	check_INITIAL checkGroup = 1 << iota
	check_POSSESSIVE_AS_BE
	check_YOUR_ARE
)

var why_reasons = map[string][]string{
	"its":         {"‘its’ belongs to ‘it’", "‘its’ belongs to ‘it’; ‘it's’ means ‘it is’ or ‘it has’"},
	"your":        {"‘your’ doesn't mean ‘you are’; ‘you're’ does", "‘your’ belongs to ‘you’"},
	"its_po":      {"‘it's’ doesn't belong to ‘it’; ‘its’ does", "‘it's’ means ‘it is’ or ‘it has’"},
	"your_po":     {"‘you're’ doesn't belong to ‘you’; ‘your’ does", "‘you're’ means ‘you are’"},
	"there_their": {"‘there’ doesn't belong to ‘them’; ‘their’ does"},
	"whose":       {"‘whose’ belongs to ‘whom’; ‘who's’ means ‘who is’"},
	"whose_has":   {"‘whose’ belongs to ‘whom’; ‘who's’ means ‘who has’"},
	"theyre_be":   {"‘they're’ means ‘they are’, not ‘there’"},
	"their_be":    {"‘their’ belongs to ‘them’", "‘there’ is ‘their’ item"},
	"theyre_are":  {"‘they're’ means ‘they are’, not ‘they’ or ‘there’"},
	"hear":        {"I ‘hear’ but am ‘here’"},
	"board":       {"‘bored’ is a verb; ‘board’ is a noun", "‘board’ is a noun; ‘bored’ is a verb"},
	"than":        {"‘than’ isn't the adverb ‘then’"},
	"then":        {"‘then’ doesn't compare like ‘than’", "‘then’ doesn't compare; ‘than’ does"},
	"of":          {"‘of’ isn't a verb", "‘have’ is a real verb"},
	"your-are":    {"‘you’ are; ‘your’ belongs to ‘you’", "‘you’ are rather than ‘your’ are"},
	"supposed-to": {"‘supposed’ isn't a bare infinitive", "‘supposed’ is really a participle"},
	"whom":        {"‘whom’ is not nominative", "‘whom’ isn't in subjective case"},
	"allot-of":    {"‘allot’ is a verb"},
}

type word *sequence.Word

type ruleMatcher struct {
	*sequence.S
	matched map[string]struct{}
	why     []string
}

func (r *ruleMatcher) Matched(reason string) {
	_, ok := r.matched[reason]
	if ok {
		// already matched
		return
	}
	r.matched[reason] = struct{}{}
	why := "[ERROR]"
	if reasons, ok := why_reasons[reason]; ok {
		// randomly pick a reason
		why = reasons[rand.Intn(len(reasons))]
	}
	r.why = append(r.why, why)
}

// DoAll repeatedly checks until nothing is detected
func DoAll(s *sequence.S) (why []string) {
	r := ruleMatcher{
		s,
		make(map[string]struct{}),
		nil,
	}
	r.doChecksRecursively()
	return r.why
}

func (r *ruleMatcher) doChecksRecursively() {
	run := check_INITIAL
	for run != 0 {
		curRun := run
		run = 0
		r.Reset()
		for r.Has() {
			run |= r.doChecks(curRun)
			r.Advance()
		}
	}
}

func (r *ruleMatcher) doChecks(run checkGroup) (rerun checkGroup) {
	cur := r.Cur()
	if run&(check_INITIAL|check_POSSESSIVE_AS_BE) != 0 {
		switch cur.Lower {
		case "its", "your", "whose":
			r.rule_possessiveAsBe(cur)
		}
	}
	if run&(check_INITIAL|check_YOUR_ARE) != 0 {
		if cur.Lower == "your" {
			r.rule_yourAre(cur)
		}
	}
	if run&check_INITIAL != 0 {
		switch cur.Lower {
		case "you're":
			r.rule_youreNoun(cur, &rerun)
		case "own":
			r.rule_itsOwn(cur, &rerun)
			r.rule_thereOwn(cur, &rerun)

		//case "going":
		// r.rule_imGoing(cur) // DISABLED

		case "allot":
			r.rule_allotOf(cur)
		case "whose":
			r.rule_whoseBeen(cur)
		case "they're":
			r.rule_theyreBe(cur)
		case "their":
			r.rule_theirModal(cur)
		case "hear", "board":
			r.rule_beNoun(cur)
		case "then":
			r.rule_then(cur)
		case "than":
			r.rule_than(cur)
		case "of":
			r.rule_of(cur)
		case "supposed":
			r.rule_supposedTo(cur, &rerun)
		case "whom", "whomever":
			r.rule_whomBe(cur)
		}
	}
	return
}
