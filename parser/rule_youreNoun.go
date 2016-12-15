package parser

import (
	"victorz.ca/grammar/sequence"
)

/*
Keyword: you're
Src: _ <possessed_noun:(day|life)|own>
Dst: [your] ...

Removed:
(you're man) enough
(you're life) changers
*/

func (r *ruleMatcher) rule_youreNoun(cur *sequence.Word, rerun *checkGroup) {
	if !r.HasNextCont(1) {
		return
	}
	next1 := r.NextWord(1)
	switch next1.Lower {
	default:
		return
	case "own":

	case "day", "life":
		if r.HasNextCont(2) {
			next2 := r.NextWord(2)
			// Exception 1: "[you're day/life] -ing"
			// Exception 2: "[you're day/life] -er(s?)"
			if next2.IsParticiplePresent() || next2.IsAgent() {
				return
			}
		} else if next1.Lower == "life" {
			// Exception 3: "you're life." [no following word]
			return
		}
	}

	r.Matched("your_po")
	cur.ReplaceCap("your")
	next1.MarkCommon()
	// Rerun: possessiveAsBe for "your you're own" -> "your [your] own" -> "[(you're) your] own"
	*rerun |= check_POSSESSIVE_AS_BE
}
