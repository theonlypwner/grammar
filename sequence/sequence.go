// Package sequence represents a sequence of tokens, which can be replaced with corrections.
package sequence

// S represents a sequence of [Word, Space] * N.
// The first Word and last Space may be empty.
type S struct {
	Words  []Word
	Spaces []Space
	pos    int
}

// Has returns whether the current word exists.
func (s *S) Has() bool {
	return s.pos < len(s.Words)
}

// Cur gets the current word.
func (s *S) Cur() *Word {
	return &s.Words[s.pos]
}

// Advance moves the iterator forward by one position.
// It does nothing if the end is reached.
func (s *S) Advance() {
	if s.Has() {
		s.pos++
	}
}

// Reset returns the iterator position to the beginning.
func (s *S) Reset() {
	s.pos = 0
}

// PrevWord gets the nth previous word.
// n = 0 gets the current word.
func (s *S) PrevWord(n int) *Word {
	return &s.Words[s.pos-n]
}

// NextWord gets the nth next word.
// n = 0 gets the current word.
func (s *S) NextWord(n int) *Word {
	return &s.Words[s.pos+n]
}

// PrevSpace gets the nth previous space, n>0.
func (s *S) PrevSpace(n int) *Space {
	return &s.Spaces[s.pos-n]
}

// NextSpace gets the nth next space, n>0.
func (s *S) NextSpace(n int) *Space {
	return &s.Spaces[s.pos+n-1]
}

// HasPrev checks whether there are at least n words earlier.
func (s *S) HasPrev(n int) bool {
	return s.pos >= n
}

// HasNext checks whether there are at least n words later.
func (s *S) HasNext(n int) bool {
	return s.pos+n < len(s.Words)
}

func (s *S) hasLoop(a, b int, maxLevel SpaceLevel) bool {
	if a < 0 || b >= len(s.Spaces) {
		return false
	}
	for i := a; i < b; i++ {
		if s.Spaces[i].Level > maxLevel {
			return false
		}
	}
	return true
}

// HasPrevCont checks whether there are at least n words earlier in the same clause.
func (s *S) HasPrevCont(n int) bool {
	return s.hasLoop(s.pos-n, s.pos, SL_CLAUSE)
}

// HasNextCont checks whether there are at least n words later in the same clause.
func (s *S) HasNextCont(n int) bool {
	return s.hasLoop(s.pos, s.pos+n, SL_CLAUSE)
}

// HasPrevInSentence checks whether there are at least n words earlier in the same sentence.
func (s *S) HasPrevInSentence(n int) bool {
	return s.hasLoop(s.pos-n, s.pos, SL_SENTENCE)
}

// HasNextInSentence checks whether there are at least n words later in the same sentence.
func (s *S) HasNextInSentence(n int) bool {
	return s.hasLoop(s.pos, s.pos+n, SL_SENTENCE)
}

func (s *S) numLoopPrev(maxLevel SpaceLevel) int {
	n := 0
	for i := s.pos - 1; i >= 0; i-- {
		if s.Spaces[i].Level >= maxLevel {
			break
		}
		n++
	}
	return n
}

// NumPrevCont counts how many previous words are in the same clause.
func (s *S) NumPrevCont() int {
	return s.numLoopPrev(SL_CLAUSE)
}

// NumPrevInSentence counts how many previous words are in the same sentence.
func (s *S) NumPrevInSentence() int {
	return s.numLoopPrev(SL_SENTENCE)
}
