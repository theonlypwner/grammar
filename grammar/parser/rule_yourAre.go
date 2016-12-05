package parser

import (
	"victorz.ca/grammar/sequence"
)

/*
Keyword: your
Src: _ are(n't)?
Dst: [you] ...
Alt: your are[a], your ar[t]
*/

func (r *ruleMatcher) rule_yourAre(cur *sequence.Word) {
	if !r.HasNextCont(1) {
		return
	}
	next1 := r.NextWord(1)
	switch next1.Lower {
	default:
		return
	case "are", "aren't":
	}

	r.Matched("your-are")
	cur.ReplaceCap("you")
	next1.MarkCommon()
}
