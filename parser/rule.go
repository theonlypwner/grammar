// Package parser applies rules and makes a list of corrections
package parser

import (
	"victorz.ca/grammar/sequence"
)

type checkGroup int

const (
	check_INITIAL checkGroup = 1 << iota
	check_POSSESSIVE_AS_BE
	check_YOUR_ARE
)

type word *sequence.Word

type ruleMatcher struct {
	*sequence.S
	matched map[string]struct{}
	why     []string
}

func (r *ruleMatcher) Matched(reasonID, reasonFull string) {
	_, ok := r.matched[reasonID]
	if ok {
		// already matched
		return
	}
	r.matched[reasonID] = struct{}{}
	r.why = append(r.why, reasonFull)
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
		case "whom", "whomever", "whomsoever":
			r.rule_whomModal(cur)
		}
	}
	return
}
