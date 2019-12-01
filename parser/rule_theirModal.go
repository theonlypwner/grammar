package parser

import (
	"victorz.ca/grammar/sequence"
)

/*
Keyword: their
Src: _ <modal>
Dst: [there] <modal>
*/

func (r *ruleMatcher) rule_theirModal(cur *sequence.Word) {
	if !r.HasNextCont(1) {
		return
	}
	next1 := r.NextWord(1)
	if !(next1.IsModal() || next1.IsCopula()) {
		return
	}

	// Exception 1: their <noun abbreviation>
	// Exception 1b: their <title of article>
	if next1.Caps != sequence.WC_LOWER {
		return
	}

	if r.HasPrevInSentence(2) {
		// Exception 2: the difference between (there/they're), (they're/there)?, and {their is}
		// Exception 2b: (they're/their) / (they're/there) / {their is} sometimes confused
		// This exception set has three forms:
		// 1. T / T / _
		// 2. T and _
		// 3. T, T, and _ // covered by #2
		switch r.PrevWord(2).Lower {
		case "there", "they're":
			switch r.PrevWord(1).Lower {
			case "there", "they're", "and":
				return
			}
		}
	}

	if r.HasNextCont(2) {
		next2 := r.NextWord(2)
		switch next1.Lower {
		case "would":
			// Exception 3: their would be
			// actual: their would-be
			if next2.Lower == "be" {
				return
			}
		}
	}

	r.Matched("their_be", "‘their’ is possessive; ‘there’ is a pronoun or an adverb")
	cur.ReplaceCap("there")
	next1.MarkCommon()
}
