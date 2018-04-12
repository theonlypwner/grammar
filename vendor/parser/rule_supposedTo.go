package parser

import (
	"sequence"
)

/*
Keyword: supposed
Src: (<word> doesn't|I don't|<word> (?<!I )don't) _ to
Dst: [isn't|I'm not|aren't] _ to

NOTE: causes extension
	don't -> [aren't]
	didn't -> [wasn't/weren't]
*/

func (r *ruleMatcher) rule_supposedTo(cur *sequence.Word, rerun *checkGroup) {
	if !r.HasNextCont(1) ||
		!r.HasPrevCont(2) ||
		r.NextWord(1).Lower != "to" {
		return
	}

	const (
		PERSON_FIRST = iota
		PERSON_PLURAL
		PERSON_THIRD
		PERSON_PAST_PLURAL
		PERSON_PAST_SINGULAR
	)

	person := PERSON_PLURAL
	prev1 := r.PrevWord(1)
	prev2 := r.PrevWord(2)
	switch prev1.Lower {
	default:
		return

	case "don't":
		switch prev2.Lower {
		case "i":
			person = PERSON_FIRST
		case "he", "she", "it":
			person = PERSON_THIRD
			// default: person = PERSON_PLURAL
		}

	case "doesn't":
		switch prev2.Lower {
		case "i":
			person = PERSON_FIRST

		case "we", "you", "they":
			// person = PERSON_PLURAL

		default:
			person = PERSON_THIRD
		}

	case "didn't":
		switch prev2.Lower {
		case "i", "we", "you", "they":
			person = PERSON_PAST_PLURAL

		case "he", "she", "it":
			person = PERSON_PAST_SINGULAR

		default:
			// unknown conjugation, don't try to guess
			return
		}
	}

	r.Matched("supposed-to", "‘supposed’ is a participle, not a bare infinitive")
	cur.MarkCommon()
	r.NextWord(1).MarkCommon()
	switch person {
	case PERSON_FIRST:
		// special: [I'm not] supposed to
		prev2.Replace("I'm")
		prev1.Replace("not")

	case PERSON_PLURAL:
		prev1.Replace("aren't")
		// Rerun: yourAre for "your don't" -> "your [aren't]" -> "[(you) aren't]"
		*rerun |= check_YOUR_ARE

	case PERSON_THIRD:
		prev1.Replace("isn't")

	case PERSON_PAST_PLURAL:
		prev1.Replace("weren't")

	case PERSON_PAST_SINGULAR:
		prev1.Replace("wasn't")
	}
}
