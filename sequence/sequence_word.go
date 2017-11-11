package sequence

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Type of word capitalization
type WordCapType int

// Word flags
const (
	// first character is lowercase
	WC_LOWER WordCapType = iota
	// Title Case
	WC_TITLE
	// COMPLETELY UPPERCASE
	WC_UPPER
)

// Word is self-explanatory.
type Word struct {
	Token
	Lower  string // current text, in lower-case
	Caps   WordCapType
	common bool
}

// NewWord makes a Word for the given original text.
func NewWord(w string) Word {
	caps := WC_UPPER

	for i, r := range w {
		if unicode.IsLower(r) {
			if i == 0 {
				caps = WC_LOWER
			} else {
				caps = WC_TITLE
			}
			break
		}
	}

	if len(w) == 0 {
		caps = WC_LOWER
	}

	return Word{
		NewToken(w),
		strings.ToLower(w),
		caps,
		false,
	}
}

// Replace replaces the word. It does not check for capitalization.
func (w *Word) Replace(newText string) {
	w.Lower = strings.ToLower(newText)
	w.Token.Replace(newText)
}

// ReplaceCap replaces the word after fixing capitalization.
func (w *Word) ReplaceCap(newText string) {
	switch w.Caps {
	case WC_TITLE:
		r, size := utf8.DecodeRuneInString(newText)
		if r == utf8.RuneError || unicode.IsTitle(r) {
			break
		}
		newText = string(unicode.ToTitle(r)) + newText[size:]

	case WC_UPPER:
		newText = strings.ToUpper(newText)
	}
	w.Replace(newText)
}

// MarkCommon marks this word as common
func (w *Word) MarkCommon() {
	w.common = true
}
