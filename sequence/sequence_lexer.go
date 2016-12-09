package sequence

import (
	"strings"
)

// New makes a Sequence from an input string.
func New(s string) S {
	seq := S{}

	if s == "" {
		return seq
	}

	sLower := strings.ToLower(s)

	word := 0
	space := -1
	level := SL_SPACE

	appendWS := func(i int) {
		seq.Words = append(seq.Words, NewWord(s[word:space], sLower[word:space]))
		seq.Spaces = append(seq.Spaces, NewSpace(s[space:i], level))
	}

	for i, r := range s {
		switch r {
		case '.', '!', '?':
			level = SL_SENTENCE
			fallthrough
		case ',', ';', ':', '(', ')', '[', ']', '{', '}':
			if level < SL_CLAUSE {
				level = SL_CLAUSE
			}
			fallthrough
		case ' ', '/': // hyphens and dashes might form compound words
			if space == -1 {
				space = i
			}

		default:
			if space != -1 {
				appendWS(i)

				word = i
				space = -1
				level = SL_SPACE
			}
		}
	}

	// Append last part
	if space == -1 {
		space = len(s)
	}
	appendWS(len(s))

	// Shrink to fit
	seq.Words = append(([]Word)(nil), seq.Words...)
	seq.Spaces = append(([]Space)(nil), seq.Spaces...)

	return seq
}
