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
		case "is", "are", "was", "were", "isn't", "aren't", "wasn't", "weren't":
			// Fix copulas
			var copulaFixMap map[string]string
			switch next2.Lower {
			case "a", "an", "another", "anything", "something", "nothing":
				copulaFixMap = copulaToSingular
			case "people", "persons":
				copulaFixMap = copulaToPlural
			}
			if copulaFixMap != nil {
				if next1new, ok := copulaFixMap[next1.Lower]; ok {
					next1.Replace(next1new)
				}
			}
		}
	}

	r.Matched("their_be", "‘their’ is possessive; ‘there’ is a pronoun or an adverb")
	cur.ReplaceCap("there")
	next1.MarkCommon()
}

var copulaToSingular = make(map[string]string)
var copulaToPlural = make(map[string]string)

func init() {
	for _, pair := range [...][2]string{
		{"is", "are"},
		{"was", "were"},
		{"isn't", "aren't"},
		{"wasn't", "weren't"},
	} {
		copulaToSingular[pair[1]] = pair[0]
		copulaToPlural[pair[0]] = pair[1]
	}
}
