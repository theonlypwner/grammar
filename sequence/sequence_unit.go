package sequence

type SequenceUnitType int

// Token is the "building block" of Sequence: a Word or a Space.
type Token struct {
	Text string
	New  bool
}

// NewToken makes a Token with the given original text.
func NewToken(text string) Token {
	return Token{
		Text: text,
		New:  false,
	}
}

// Replace changes the resulting text and marks it as modified.
func (s *Token) Replace(newText string) {
	s.Text = newText
	s.New = true
}

/*
// Final gets the text, with modifications if any.
func (s *Token) Final() string {
	return s.Text
}
*/
