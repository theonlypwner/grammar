package parser

import (
	"sequence"
)

/*
Keyword: than
Src: <word>(,) <conjunction:(and|but|yet)> _ <word>
Dst: ... [then] ...
*/

func (r *ruleMatcher) rule_than(cur *sequence.Word) {
	if !r.HasNextCont(1) {
		return
	}
	p := r.NumPrevInSentence()
	if p < 2 {
		return
	}
	prev1 := r.PrevWord(1)
	switch prev1.Lower {
	default:
		return
	case "and":
		// Exception 1: the difference between 'then' {and 'than'}
		if r.PrevWord(2).Lower == "then" {
			return
		}
		// Exception 2:
		// <comparative> than N<NP>+ {(and|or) than} <NP>+
		if p >= 4 {
			prev := r.PrevWord(3)
			for i := 4; i <= p; i++ { // do not check most recent 3: [better than <NP>+ and] than
				cur := r.PrevWord(i)
				if prev.Lower == "than" && cur.IsComparative() {
					return
				}
				prev = cur
			}
		}
	case "yet":
		// Exception 3:
		// <comparative> for <NP> not <adj> yet than <NP>
		// Exclude if <comparative> is found at least 6 words before 'than' in the same sentence
		for i := 6; i <= p; i++ {
			if r.PrevWord(i).IsComparative() {
				return
			}
		}
	case "but":
	}

	r.Matched("than")
	prev1.MarkCommon()
	cur.Replace("then")
}
