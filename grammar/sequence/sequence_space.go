package sequence

type SpaceLevel int

// Space break levels
const (
	// Words around Space are in same clause
	SL_SPACE SpaceLevel = iota
	// Space breaks the clause; Words around Space are in same sentence
	SL_CLAUSE
	// Space breaks the sentence
	SL_SENTENCE
)

// Space is a word separator, possibly including punctuation.
type Space struct {
	Token
	Level SpaceLevel
}

// NewSpace makes a new Space for the given original text.
func NewSpace(s string, level SpaceLevel) Space {
	return Space{
		NewToken(s),
		level,
	}
}
