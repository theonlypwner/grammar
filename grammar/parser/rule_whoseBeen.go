package parser

import (
	"victorz.ca/grammar/sequence"
)

/*
Keyword: whose
Src: _ been
Dst: [who's] been
*/

func (r *ruleMatcher) rule_whoseBeen(cur *sequence.Word) {
	if !r.HasNextCont(1) {
		return
	}
	next1 := r.NextWord(1)
	if next1.Lower != "been" {
		return
	}
	r.Matched("whose_has")
	cur.ReplaceCap("who's")
	next1.MarkCommon()
}
