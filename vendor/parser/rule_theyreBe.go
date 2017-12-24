package parser

import (
	"sequence"
)

/*
Keyword: they're
Src: _ is
Dst: [there's]
Src: _ are
Dst: [there] are
Alt: [they] are
 - they are being (GOOD)
 - there are being (BAD)

Removed:
	{they're day} dreaming
	Be {there day} and night
	"they're be" is full of improper usage
*/

func (r *ruleMatcher) rule_theyreBe(cur *sequence.Word) {
	if r.HasPrevInSentence(1) {
		prev1 := r.PrevWord(1)
		// Exception 1: the difference between (their/there), (there/their)?, and {they're is}
		// Exception 1b: (their/there) / (there/their) / {they're is} sometimes confused
		// This exception set has three forms:
		// 1. T / T / _
		// 2. T and _
		// 3. T, T, and _ // covered by #2
		if r.HasPrevInSentence(2) {
			prev2 := r.PrevWord(2)
			switch prev2.Lower {
			case "there", "their":
				switch prev1.Lower {
				case "there", "their", "and":
					return
				}
			}
		}

		// Exception 2: what {they're is}
		if prev1.Lower == "what" {
			return
		}
	}
	switch {
	case !r.HasNextCont(1):
		return
	// Exception 3b: {'they're' is} 'they are'
	// Exception 3: {they're, aren't} they?
	case r.HasNextCont(2) && r.NextWord(2).Lower == "they":
		return
	}
	next1 := r.NextWord(1)
	switch {
	case next1.Lower == "is" ||
		next1.Lower == "isn't":
		r.Matched("theyre_be")
		cur.ReplaceCap("there's")
		r.NextSpace(1).Replace("") // collapse space
		next1.Replace("")          // delete next word (is)
	case next1.IsCopula() ||
		next1.IsModal():
		r.Matched("theyre_are")
		// they are being is OK, but there are being is NOT, but they can
		// always replace there
		cur.ReplaceCap("they") // cur.ReplaceCap("they/there")
		next1.MarkCommon()
	}
}
