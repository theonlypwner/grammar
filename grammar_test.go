package grammar

import "testing"

func TestLoadRegular(t *testing.T) {
	t.Parallel()

	// Try to properly load a regular sentence.
	negative(t, "This sentence is fine.")
}

func TestUnicode(t *testing.T) {
	t.Parallel()

	positive(t, "… → their is …", "→ [there] is")
}

func TestWordCase(t *testing.T) {
	t.Parallel()

	// lowercase detection
	positive(t, "their is", "[there] is")
	// Titlecase Detection
	positive(t, "Their is", "[There] is")
	// UPPERCASE DETECTION
	positive(t, "THEIR is", "[THERE] is")
}

func TestOverlap(t *testing.T) {
	t.Parallel()

	positive(t, "Their is you're own", "[There] is [your] own")
	positive(t, "Your you're own", "[You're your] own")
	positive(t, "Its there own item", "[It's their] own item")
	positive(t, "Its it's own item", "[It's its] own item")
	positive(t, "Your don't supposed to!", "[You aren't] supposed to!")
}

func TestPunctuation(t *testing.T) {
	t.Parallel()

	// Keep question marks and exclamation marks
	positive(t, "their is?", "[there] is?")
	positive(t, "their is!", "[there] is!")
	// Original punctuation should be out of the brackets
	positive(t, "I am hear!", "I am [here]!")
	// Drop leading and trailing spaces
	positive(t, " their is ", "[there] is")
	// Retain double spaces
	positive(t, "their  is", "[there]  is")
	// Drop periods, commas, and semicolons
	positive(t, "their is.", "[there] is")
	positive(t, "their is,", "[there] is")
	positive(t, "their is;", "[there] is")
}

func TestWording(t *testing.T) {
	// t.Parallel()

	// Verify that wording can be generated without failing
	r := MakeTweetReply("Their is and your don't supposed to! (blah) They think their is.", "@@")
	if r != "" {
		t.Logf("Reply: %q", r)
	} else {
		t.Errorf("Expected non-empty reply")
	}
}

var corrections, why []string

func BenchmarkShortOK(b *testing.B) {
	var c, w []string
	for n := 0; n < b.N; n++ {
		c, w = Load("Nothing's wrong with this sentence.")
	}
	corrections = c
	why = w
}

func BenchmarkShortDetect(b *testing.B) {
	var c, w []string
	for n := 0; n < b.N; n++ {
		c, w = Load("But it's true that their is a problem with this sentence.")
	}
	corrections = c
	why = w
}

func BenchmarkLongOK(b *testing.B) {
	s := "Nothing's wrong with this sentence. "
	for i := 0; i < 10; i++ {
		s += s
	}

	var c, w []string
	for n := 0; n < b.N; n++ {
		c, w = Load(s)
	}
	corrections = c
	why = w
}

func BenchmarkLongDetect(b *testing.B) {
	s := "But it's true that their is a problem with this sentence. "
	for i := 0; i < 10; i++ {
		s += s
	}

	var c, w []string
	for n := 0; n < b.N; n++ {
		c, w = Load(s)
	}
	corrections = c
	why = w
}

func BenchmarkVeryLongOK(b *testing.B) {
	s := "Nothing's wrong with this sentence. "
	for i := 0; i < 15; i++ {
		s += s
	}

	var c, w []string
	for n := 0; n < b.N; n++ {
		c, w = Load(s)
	}
	corrections = c
	why = w
}
