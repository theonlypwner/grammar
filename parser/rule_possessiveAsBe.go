package parser

import (
	"victorz.ca/grammar/sequence"
)

/*
Keyword: (its|your|whose)
Src: _ (<article|possessive_determiner|possessive_pronoun|preposition>|here|not) <word>
Dst: (it's|you're|who's) <word>

NOTE: causes extension
*/

func (r *ruleMatcher) rule_possessiveAsBe(cur *sequence.Word) {
	if !r.HasNextCont(2) {
		return
	}

	repl := "[ERROR]"
	switch cur.Lower {
	case "its":
		repl = "it's"
	case "your":
		repl = "you're"
	case "whose":
		repl = "who's"
	}

	next1 := r.NextWord(1)
	// Exception 1: Some author, {whose THE} BOOK does, is
	// - book titles in ALLCAPS should be ignored
	if next1.Caps == sequence.WC_UPPER {
		return
	}
	switch next1.Lower {
	case "after":
		// Exception 2: {its after} effect(s?)
		switch r.NextWord(2).Lower {
		case "effect", "effects":
			return
		}

	case "all":
		// Exception 3: {your all} but nothing system
		if r.NextWord(2).Lower == "but" {
			return
		}

	case "here":
	// case "not": // removed: <possessive> not <participle_present(gerund)> <noun>

	default:
		if !(next1.IsPreposition() ||
			next1.IsDeterminer() ||
			next1.IsPossessivePronoun()) {
			return
		}
	}

	r.Matched(cur.Lower)
	cur.ReplaceCap(repl)
	next1.MarkCommon()
}
