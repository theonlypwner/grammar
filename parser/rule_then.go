package parser

import (
	"victorz.ca/grammar/sequence"
)

/*
Keyword: then
Src: * <comparative> _ *
Dst: ... [than] *
*/

func (r *ruleMatcher) rule_then(cur *sequence.Word) {
	p := r.NumPrevInSentence()
	if p < 2 || !r.HasNextCont(1) {
		return
	}

	prev1 := r.PrevWord(1)
	if !prev1.IsComparative() {
		return
	}

	// Exception 1: then <verb:(be|do|did|get|got)|lol> <noun>
	switch r.NextWord(1).Lower {
	case "lol", "be", "do", "did", "get", "got":
		return
	}

	// Exception 2: if/when ... {better(,) then}
	for i := 2; i <= p; i++ {
		switch r.PrevWord(i).Lower {
		case "if", "when":
			return
		}
	}

	r.Matched("then")
	prev1.MarkCommon()
	cur.Replace("than")
}
