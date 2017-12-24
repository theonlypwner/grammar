package parser

import (
	"sequence"
)

/*
Keyword: whom/whomever
Src: _ <be>
Dst: [who/whoever] ...
*/

func (r *ruleMatcher) rule_whomBe(cur *sequence.Word) {
	if !r.HasNextCont(1) {
		return
	}
	next1 := r.NextWord(1)
	present := false
	switch next1.Lower {
	case "be", "am", "are", "is":
		present = true
	case "was", "were":
		// present = false
	default:
		return
	}

	next1New := "was"
	if present {
		next1New = "is"
	}

	if cur.Lower == "whomever" {
		cur.ReplaceCap("whoever")
	} else {
		if r.HasPrevInSentence(1) {
			switch r.PrevWord(1).Lower {
			case "i", "me", "myself":
				// Potential issue: the person (sitting across from {me) who is}
				if present {
					next1New = "am"
				}

			case "people", "persons", "we", "us", "you", "they", "them", "those":
				if present {
					next1New = "are"
				} else {
					next1New = "were"
				}

			case "person", "guy", "he", "him", "she", "her", "it":
				// already set

			default:
				// unknown person
				return
			}
			cur.Replace("who")
		} else {
			// Exception: "{Whom was} it from?"
			// Exception: "{Whom were} they from?"
			return
		}
	}

	next1.MarkCommon()
	if next1.Lower != next1New {
		next1.Replace(next1New)
	}
	r.Matched("whom")
}
