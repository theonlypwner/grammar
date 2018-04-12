package parser

import (
	"sequence"
)

/*
Keyword: own
Src: it's _
Dst: [its] _
*/

func (r *ruleMatcher) rule_itsOwn(cur *sequence.Word, rerun *checkGroup) {
	if !r.HasPrevCont(1) {
		return
	}
	prev := r.PrevWord(1)
	if prev.Lower != "it's" {
		return
	}

	r.Matched("it's", "‘it's’ means ‘it is’ or ‘it has’, but ‘its’ is possessive")
	prev.ReplaceCap("its")
	cur.MarkCommon()
	// Rerun: possessiveAsBe for "its it's own" -> "its [its] own" -> "[(it's) its] own"
	*rerun |= check_POSSESSIVE_AS_BE
}
