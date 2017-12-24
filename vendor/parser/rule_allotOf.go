package parser

import (
	"sequence"
)

/*
Keyword: allot
Src: _ of
Dst: [a lot] of
*/

func (r *ruleMatcher) rule_allotOf(cur *sequence.Word) {
	if !r.HasNextCont(1) {
		return
	}
	next := r.NextWord(1)
	if next.Lower != "of" {
		return
	}

	r.Matched("allot-of")
	cur.ReplaceCap("a lot")
	next.MarkCommon()
}
