// Package parser applies rules and makes a list of corrections
package parser

import (
	"sequence"
)

type checkGroup int

const (
	check_INITIAL checkGroup = 1 << iota
	check_POSSESSIVE_AS_BE
	check_YOUR_ARE
)

var why_reason = map[string]string{
	"its":   "‘its’ is possessive; ‘it's’ means ‘it is’ or ‘it has’",
	"your":  "‘your’ is possessive; ‘you're’ means ‘you are’",
	"whose": "‘whose’ is possessive; ‘who's’ means ‘who is’",

	"its_po":  "‘it's’ means ‘it is’ or ‘it has’, but ‘its’ is possessive",
	"your_po": "‘you're’ means ‘you are’; ‘your’ is possessive",

	"there_their": "‘there’ is not possessive, but ‘their’ is",
	"whose_has":   "‘whose’ is possessive; ‘who's’ means ‘who has’",
	"theyre_be":   "‘they're’ means ‘they are’, not ‘there’",
	"theyre_are":  "‘they're’ means ‘they are’, not ‘they’ or ‘there’",
	"their_be":    "‘their’ is possessive; ‘there’ is a pronoun or an adverb",
	"hear":        "I am ‘here’ to ‘hear’",
	"board":       "‘board’ is a noun; ‘bored’ is a verb",
	"than":        "‘than’ compares, but ‘then’ is an adverb",
	"then":        "unlike the adverb ‘then’, ‘than’ compares",
	"of":          "‘of’ is not a verb like ‘have’ is",
	"your-are":    "‘your’ is a possessive determiner; ‘you’ is a pronoun",
	"supposed-to": "‘supposed’ is a participle, not a bare infinitive",
	"whom":        "unlike ‘whom’, ‘who’ is a subject",
	"allot-of":    "‘allot’ is a verb; ‘a lot’ is a noun or adverb",
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
	if reasonText, ok := why_reason[reason]; ok {
		why = reasonText
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
