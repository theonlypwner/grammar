package sequence

import (
	"strings"
)

// ConsiderNear returns whether this word may be included as a "near" word.
func (w *Word) ConsiderNear() bool {
	switch w.Lower {
	case "", "…":
		// can't detect "http://" and "https://" prefixes
		return false
	}
	return true
}

// ConsiderCommon returns whether this word may be considered a "common" word
func (w *Word) ConsiderCommon() bool {
	if w.common ||
		w.IsDeterminer() ||
		w.IsPreposition() ||
		w.IsConjunctionCoordinating() ||
		w.IsConjunctionSubordinating() ||
		w.IsPronounPersonal() ||
		w.IsModal() {
		return true
	}
	switch w.Lower {
	case
		"over", "to", // more prepositions
		"cannot", "ought", // more modals
		// interrogative
		"how", "what", "who", "whom", "why", "when", "where", "which",
		// to "be"
		"be", "am", "are", "is", "was", "were", "been", "being",
		// to "do"
		"do", "does", "did", "done", "doing",
		// to "have"
		"have", "has", "had", "having",
		// contractions
		"aren't", "can't", "couldn't", "didn't", "doesn't", "don't", "hadn't", "hasn't", "haven't", "he'd", "he'll", "he's", "here's", "how's", "i'd", "i'll", "i'm", "i've", "isn't", "it's", "let's", "mustn't", "shan't", "she'd", "she'll", "she's", "shouldn't", "that's", "there's", "they'd", "they'll", "they're", "they've", "wasn't", "we", "we'd", "we'll", "we're", "we've", "weren't", "what's", "when's", "where's", "who's", "why's", "won't", "wouldn't", "you'd", "you'll", "you're", "you've",
		// special
		"again",   // adverb
		"further", // adverb, adjective
		"not",     // adverb, noun
		"other",   // adjective
		"own",     // adjective, verb
		"same",    // adjective, pronoun, adverb
		"such",    // adjective
		"that",    // pronoun, adverb, adjective, conjunction
		"then",    // adverb
		"there",   // adverb
		"these",   // pronoun, adjective, adverb
		"this",    // ^
		"those",   // ^
		"too",     // adverb
		"very":    // adjective, adverb?
		return true
	}
	return false
}

func (w *Word) IsAgent() bool {
	// possible issues
	//   - false negatives: actor, actors
	//   - false positives: alter, deter, enter, prefer, ...
	return strings.HasSuffix(w.Lower, "er") ||
		strings.HasSuffix(w.Lower, "ers")
}

func (w *Word) IsArticle() bool {
	switch w.Lower {
	default:
		return false
	case "a", "an", "the":
		return true
	}
}

func (w *Word) IsComparative() bool {
	switch w.Lower {
	default:
		return false
	case "better", "worse", "more", "less", "fewer":
		return true
	}
}

func (w *Word) IsCopula() bool {
	switch w.Lower {
	default:
		return false
	case
		"be", "is", "are", "isn't", "aren't",
		"was", "wasn't":
		return true
	}
}

func (w *Word) IsDeterminer() bool {
	switch w.Lower {
	case
		"each", // "every",
		"either", "neither",
		"some", "any", "no", "none",
		// "much", "many",
		"most", "more",
		// "little", "less", "least",
		"few",         // "fewer", "fewest",
		"all", "both", // "half",
		// "several",
		"enough":
		return true
	}
	return w.IsArticle() || w.IsPossessiveDeterminer()
}

func (w *Word) IsConjunctionCoordinating() bool {
	switch w.Lower {
	default:
		return false
	case "for", "and", "nor", "but", "or", "yet", "so":
		return true
	}
}

func (w *Word) IsConjunctionSubordinating() bool {
	switch w.Lower {
	default:
		return false
	case
		"after",
		"although",
		"as", // also as if
		"because",
		"before",
		// "even though",
		"if",   // noun
		"once", // adverb
		"only", // adverb, adjective
		"than",
		// "that",
		"though",
		"unless",
		"until",
		"whether",
		"when",
		"where",
		"while": // noun, adverb, verb
		return true
	}
}

func (w *Word) IsModal() bool {
	switch w.Lower {
	default:
		return false
	// exclude "be" and its conjugations
	case "could", "should", "would", "must",
		"couldn't", "shouldn't", "wouldn't", "mustn't":
		return true
	}
}

func (w *Word) IsPreposition() bool {
	switch w.Lower {
	default:
		return false
	case "aboard",
		"about",
		"above",
		// "according to",
		"across",
		"after", // your after(-)effects [checked]
		"against",
		// "ahead of",
		"along",
		"alongside",
		"amid",
		"amidst",
		"among",
		"amongst",
		"around",
		// "as", // also conjunction
		//    "as far as",
		//    "as of",
		//    "as per",
		//    "as regards",
		//    "as well as",
		"aside",
		//    "aside from",
		// "astride", // not common
		"at",
		// "athwart", // not common
		"atop",
		// "barring", // not common, also participle
		// "because of",
		"before",
		// "behind", // also noun
		"below", // your below(-)average <nounp>
		"beneath",
		// "beside", "besides", // also noun
		"beyond", // your beyond(-)<adj> <nounp>
		"between",
		"by",
		//    "by means of",
		// "circa",  // not common
		// "close to",
		// "concerning", // also participle
		// "despite", // not useful
		"down", // your down(-)stairs computer
		// "due to",
		"during",
		"except",
		// "except for",
		"excluding", // also participle
		"failing",   // also participle
		// "far from",
		// "following", // also participle
		"for",
		"from",
		// "given", // also past participle
		"in",
		//    "in accordance with",
		//    "in addition to",
		//    "in case of",
		//    "in front of",
		//    "in lieu of",
		//    "in place of",
		//    "in point of",
		//    "in spite of",
		// "including", // also participle
		// "inside", // also noun
		//    "inside of",
		// "instead of",
		"into",
		"like", // also verb
		// "mid", // also adj
		// "minus", // also noun
		"near",
		// "next", // also adj
		//    "next to",
		// "notwithstanding of",
		"of",
		"off",
		// "on", // also adj
		//    "on account of",
		//    "on behalf of",
		//    "on top of",
		"out",     // -of/from
		"outside", // -of also noun
		"over",    // your over(-)excited <nounp>
		// "owing to",
		// "pace", // outdated
		// "past", // also noun
		// "per", // not useful
		// "plus", // also noun
		// "prior to",
		// "pursuant to",
		// "qua", // formal only
		"regarding", // also participle
		// "regardless of",
		// "round", // also adj
		"sans",
		"save", // also verb
		// "since", // also conjunction
		// "subsequent to",
		// "than", // also conjunction
		// "thanks to",
		// "that of",
		"through",
		"throughout",
		"till", // informal until, also conjunction
		// "times", // also noun
		"to", // your too-<adj> <nounp>
		"toward",
		"towards",
		"under", // your under(-)prepared <nounp>
		"underneath",
		"unlike", // also adj
		"until",
		// "unto", // not common
		"up",
		"upon",
		"versus",
		"via",
		"with",
		//    "with regard to",
		//    "with respect to",
		"within",
		"without":
		// "worth", // also noun
		return true
	}
}

func (w *Word) IsParticiplePast() bool {
	switch w.Lower {
	case "been", "thrown":
		return true
	default:
		return strings.HasSuffix(w.Lower, "ed")
	}
}

func (w *Word) IsParticiplePresent() bool {
	return strings.HasSuffix(w.Lower, "ing")
}

func (w *Word) IsPronounPersonal() bool {
	switch w.Lower {
	default:
		return false
	case
		"i", "we", "you", "he", "she", "it", "they", // subjective
		"me", "us", "him", "her", "them", // objective
		// reflexive
		"myself", "ourselves", "yourself", "yourselves", "himself", "herself", "itself", "themselves",
		// *reflexive (nonstandard)
		"myselves", "ourself", "themself":
		return true
	}
}

func (w *Word) IsPossessive() bool {
	return w.IsPossessiveDeterminer() ||
		w.IsPossessivePronoun()
}

func (w *Word) IsPossessiveDeterminer() bool {
	switch w.Lower {
	default:
		return false
	case "my", "our", "thy", "your", "his", "her", "its", "their":
		return true
	}
}

func (w *Word) IsPossessivePronoun() bool {
	switch w.Lower {
	default:
		return false
	case "mine", "ours", "thine", "yours", "his", "hers", "its", "theirs":
		return true
	}
}
