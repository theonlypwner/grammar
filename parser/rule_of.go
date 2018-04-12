package parser

import (
	"victorz.ca/grammar/sequence"
)

/*
Keyword: of
Src: <modal:((could|should|would|might|must)(n't)?)> (not)? _ <verb_past>
Dst: ... _['ve] to_past_participle(<verb_past>)

Removed:
Bad Modals:
	*can
	first {May of} 2000
	will
	shall
	ought to (low traffic)
	in {need not of} better days
	had better (awkward 've)
*/

func (r *ruleMatcher) rule_of(cur *sequence.Word) {
	p := r.NumPrevCont()
	if p < 1 {
		return
	}

	notShift := 0
	prev1 := r.PrevWord(1)
	if prev1.Lower == "not" {
		if p < 2 {
			return
		}
		prev1.MarkCommon()
		notShift = 1
		prev1 = r.PrevWord(2)
	}

	next1 := (*sequence.Word)(nil)
	if r.HasNextCont(1) {
		next1 = r.NextWord(1)
		if next1.Lower == "course" {
			// Exception 1: He {could(,) of} course(,) ...
			return
		} else if next1.IsPronounPersonal() {
			// Exception 2: this {would of} themselves justify
			return
		}
	}

	switch {
	default:
		return
	case prev1.Lower == "might" ||
		prev1.Lower == "must":
		// Exception 3a: <determiner> <adjP>? {(might|must) of}
		for i := 2 + notShift; i <= p; i++ {
			if r.PrevWord(i).IsDeterminer() {
				return
			}
		}
		// Exception 3b: {(might|must) of} <determiner|NP>
		if next1 != nil &&
			(next1.IsDeterminer() ||
				next1.IsPronounPersonal()) {
			return
		}
	case prev1.IsModal():
	}

	// Exception 4: (more|less) (of) <NP>+ than <NP>+ {<modal> of} <NP>+
	// [more of <word>+ than <word>+ (prev) not?] of
	if p >= 6+notShift {
		const (
			TOKEN_THAN = iota
			//TOKEN_WORD
			TOKEN_OF
		)
		finding := TOKEN_THAN
		// skip <modal> and <word>, and remove 1 at end
		for i := 3 + notShift; i < p; i++ {
			switch finding {
			case TOKEN_THAN:
				if r.PrevWord(i).Lower == "than" {
					// finding = TOKEN_WORD
					i++ // skip <word>
					finding = TOKEN_OF
				}
			/*
				case TOKEN_WORD:
					// skip <word>
					finding = TOKEN_OF
			*/
			case TOKEN_OF:
				if r.PrevWord(i).Lower == "of" {
					if r.PrevWord(i + 1).IsComparative() {
						return
					}
				}
			}
		}
	}

	r.Matched("of", "‘of’ is not a verb like ‘have’ is")
	prev1.MarkCommon()
	r.PrevSpace(1).Replace("") // collapse space
	cur.Replace("'ve")

	// Fix: 've <verb_past_simple> -> <verb_past_perfect>
	baseVerb := (*sequence.Word)(nil)
	if next1 != nil {
		if next1.Lower != "not" {
			baseVerb = next1
		} else if r.HasNextCont(2) {
			baseVerb = r.NextWord(2)
		}
	}
	if baseVerb != nil {
		if newVerb, ok := fixPastToParticiple[baseVerb.Lower]; ok {
			baseVerb.Replace(newVerb)
		}
	}
}

var fixPastToParticiple = map[string]string{
	// "was": "been", // be
	"went": "gone", // go
	// "laid": "lain", // lay
	// "lay":  "lain", // lie
	// "lied": "lied", // lie
	"showed": "shown", // show
	"slew":   "slain", // slay
	"have":   "had",   // have (typo)

	"began":  "begun",  // begin
	"drank":  "drunk",  // drink
	"rang":   "rung",   // ring
	"sang":   "sung",   // sing
	"sank":   "sunk",   // sink
	"sprang": "sprung", // spring
	"swam":   "swum",   // swim

	"arose": "arisen",  // arise
	"drove": "driven",  // drive
	"rode":  "ridden",  // ride
	"rose":  "risen",   // rise
	"wrote": "written", // write

	"broke": "broken", // break
	"chose": "chosen", // choose
	"spoke": "spoken", // speak
	"stole": "stolen", // steal
	"woke":  "woken",  // wake

	"fell":      "fallen",     // fall
	"saw":       "seen",       // see
	"see":       "seen",       // see (typo)
	"shook":     "shaken",     // shake
	"took":      "taken",      // take
	"undertook": "undertaken", // undertake

	"became":   "become",   // _
	"came":     "come",     // _
	"overcame": "overcome", // _
	"ran":      "run",      // _

	"ate":     "eaten",     // eat
	"forbade": "forbidden", // forbid
	"forgot":  "forgotten", // forget
	"forgave": "forgiven",  // forgive
	"froze":   "frozen",    // freeze
	// "got": "gotten", // get (have got is sometimes a false positive)
	"gave": "given",  // give
	"hid":  "hidden", // hide

	"bore":     "borne",     // bear
	"beat":     "beaten",    // _
	"blew":     "blown",     // blow
	"did":      "done",      // do
	"drew":     "drawn",     // draw
	"flew":     "flown",     // fly
	"grew":     "grown",     // grow
	"knew":     "known",     // know
	"tore":     "torn",      // tear
	"threw":    "thrown",    // throw
	"wore":     "worn",      // wear
	"withdrew": "withdrawn", // withdraw
}
