package parser

import (
	"victorz.ca/grammar/sequence"

	"fmt"
)

/*
Keyword: whom/whomever/whomsoever
Src: _ <be>
Dst: [who/whoever/whosoever] ...
*/

func (r *ruleMatcher) rule_whomModal(cur *sequence.Word) {
	if !r.HasNextCont(1) {
		return
	}

	const (
		WHOM_MODAL = iota
		WHOM_PRESENT
		WHOM_PAST
		WHOM_PRESENT_HAS
		WHOM_PAST_HAS
	)

	next1 := r.NextWord(1)
	whom := WHOM_MODAL
	possibleBaseVerb := (*sequence.Word)(nil)
	switch next1.Lower {
	case "be", "am", "are", "is":
		whom = WHOM_PRESENT
	case "was", "were":
		whom = WHOM_PAST
	case "have", "has":
		whom = WHOM_PRESENT_HAS
		if r.HasNextCont(2) {
			possibleBaseVerb = r.NextWord(2)
		}
	case "had":
		whom = WHOM_PAST_HAS
		if r.HasNextCont(2) {
			possibleBaseVerb = r.NextWord(2)
		}
	default:
		if !next1.IsModal() {
			return
		}
	}

	next1New := "was"
	if whom == WHOM_PRESENT {
		next1New = "is"
	}

	rule := cur.Lower
	repl := ""

	switch rule {
	// default: return
	case "whomever":
		repl = "whoever"
	case "whomsoever":
		repl = "whosoever"
	case "whom":
		if r.HasPrevInSentence(1) {
			switch r.PrevWord(1).Lower {
			case "i", "me", "myself":
				// Potential issue: the person (sitting across from {me) who is}
				if whom == WHOM_PRESENT {
					next1New = "am"
				}

			case "people", "persons", "we", "us", "you", "they", "them", "those":
				switch whom {
				case WHOM_PRESENT:
					next1New = "are"
				case WHOM_PAST:
					next1New = "were"
				case WHOM_PRESENT_HAS:
					next1New = "have"
				}

			case "person", "guy", "he", "him", "she", "her", "it":
				switch whom {
				case WHOM_PRESENT:
					// next1New = "is" // already set
				case WHOM_PAST:
					// next1New = "was" // already set
				case WHOM_PRESENT_HAS:
					next1New = "has"
				}

			default:
				// unknown person
				return
			}
			repl = "who"
		} else {
			// Exception: "{Whom was} it from?"
			// Exception: "{Whom were} they from?"
			return
		}
	}

	cur.ReplaceCap(repl)
	next1.MarkCommon()
	if whom != WHOM_MODAL && next1.Lower != next1New {
		next1.Replace(next1New)
	}
	if possibleBaseVerb != nil {
		// Fix: who have <verb_past_simple> -> <verb_past_perfect>
		possibleBaseVerb.FixParticiplePast()
	}
	r.Matched("whom", fmt.Sprintf("unlike ‘%v’, ‘%v’ is the subject of ‘%v’", rule, repl, next1New))
}
