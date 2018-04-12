package parser

import (
	"sequence"
)

/*
Keyword: own
Src: there _
Dst: [their] _

Removed: Is {there a} <noun>?
*/

func (r *ruleMatcher) rule_thereOwn(cur *sequence.Word, rerun *checkGroup) {
	p := r.NumPrevCont()
	if p == 0 {
		return
	}
	prev1 := r.PrevWord(1)
	if prev1.Lower != "there" {
		return
	}

	// Exception 1: "Do/es (any/some/no-/no )one {there own} something?"
	prev := prev1
	for i := 2; i <= p; i++ {
		cur := r.PrevWord(i)
		switch cur.Lower {
		case "anyone", "anybody", "someone", "somebody",
			"no-one", "no-body", "noone", "nobody":
			return
		case "any", "some", "no":
			switch prev.Lower {
			case "one", "person", "people", "body", "of":
				return
			}
		}
		// no need to find ['do', 'does'] since people sometimes skip it
		prev = cur
	}

	// Exception 2: "<NP>+ <preposition> there own"
	if p >= 3 && r.PrevWord(2).IsPreposition() {
		return
	}

	r.Matched("there_their", "‘there’ is not possessive, but ‘their’ is")
	prev1.ReplaceCap("their")
	cur.MarkCommon()
	// Rerun: possessiveAsBe for "Its there own" -> "Its [their] own" -> "[(It's) their] own"
	*rerun |= check_POSSESSIVE_AS_BE
}
