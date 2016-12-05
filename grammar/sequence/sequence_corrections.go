package sequence

import (
	"bytes"
)

// String gets a string representing the sequence after modifications.
func (s *S) String() string {
	var buf bytes.Buffer
	for i := 0; i < len(s.Words); i++ {
		buf.WriteString(s.Words[i].Text)
		buf.WriteString(s.Spaces[i].Text)
	}
	return buf.String()
}

// Corrections returns a slice of strings showing differences in the new sequence.
func (s *S) Corrections() []string {
	return s.corrections(false)
}

// CorrectionsQuoted is like CorrectionsQuoted, but it adds quotation marks around each string.
func (s *S) CorrectionsQuoted() []string {
	return s.corrections(true)
}

func (s *S) corrections(quoted bool) []string {
	var result []string
	n := len(s.Words)

	const END_NONE = -3

	var last bytes.Buffer
	lastEnd := END_NONE

	if quoted {
		last.WriteRune('“')
	}

	// Include: [<corrected>] <common>{0,2} <near> (on both sides)
	// Allow one gap: <right boundary> <word> <left boundary>
	// Add special punctuation: <final word> [?!]

	addLast := func() {
		lastPunctuation := s.Spaces[lastEnd].Text
		if len(lastPunctuation) != 0 {
			c := lastPunctuation[0]
			if c == '?' || c == '!' {
				last.WriteByte(c)
			}
		}
		if quoted {
			last.WriteRune('”')
		}
		result = append(result, last.String())
	}

	for i := 0; i < n; i++ {
		if !s.Words[i].New {
			continue
		}
		L := i
		R := i
		// left check
		for L != 0 && // at least one word before
			L != lastEnd+1 && // no overlap
			s.Spaces[L-1].Level < SL_SENTENCE && // no wall on left
			L+3 != i { // limit (3 words)

			L--
			if !s.Words[L].ConsiderCommon() {
				if !s.Words[L].ConsiderNear() {
					L++
				}
				break
			}
		}
		// right check
		for R+1 != n && // at least one word after
			!s.Words[R+1].New && // don't parse next corrected word yet
			s.Spaces[R].Level < SL_SENTENCE && // no wall on right
			R != i+3 { // limit (3 words)

			R++
			if !s.Words[R].ConsiderCommon() {
				if !s.Words[R].ConsiderNear() {
					R--
				}
				break
			}
		}
		switch {
		case L == lastEnd+2:
			// one word in between
			L--
			fallthrough
		case L == lastEnd+1:
			// merge overlap or touching boundaries
			last.WriteString(s.Spaces[lastEnd].Text)

		case lastEnd != END_NONE:
			// no overlap, but not first
			addLast()
			last.Reset()
			if quoted {
				last.WriteRune('“')
			}
		}
		for j := L; j != i; j++ {
			last.WriteString(s.Words[j].Text)
			last.WriteString(s.Spaces[j].Text)
		}
		if i == 0 || !s.Words[i-1].New {
			last.WriteByte('[')
		}
		last.WriteString(s.Words[i].Text)
		if i+1 == n || !s.Words[i+1].New {
			last.WriteByte(']')
		}
		for j := i; j != R; j++ {
			last.WriteString(s.Spaces[j].Text)
			last.WriteString(s.Words[j+1].Text)
		}
		lastEnd = R
	}
	if lastEnd != END_NONE {
		addLast()
	}

	return result
}
