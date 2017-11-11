package grammar

import "testing"

func positive(t *testing.T, in, out string) {
	out = "“" + out + "”"
	corrections, why := Load(in)
	switch {
	case len(corrections) != 1:
		t.Errorf("corrections == %q, want %q", corrections, out)
	case corrections[0] != out:
		t.Errorf("corrections[0] == %q, want %q", corrections[0], out)
	case len(why) == 0:
		t.Errorf("len(why) == 0")
	}
}

func negative(t *testing.T, s string) {
	corrections, why := Load(s)
	switch {
	case corrections != nil:
		t.Errorf("corrections == %q, want nil", corrections)
	case why != nil:
		t.Errorf("why == %q, want nil", why)
	}
}

func TestRule_possessiveAsBe(t *testing.T) {
	t.Parallel()

	positive(t, "Its here or there", "[It's] here or there")
	positive(t, "Your the best", "[You're] the best")
	positive(t, "Whose your friend?", "[Who's] your friend?")
	// Exception 1
	negative(t, "Some author, whose THE BOOK does, is")
	// Exception 2
	negative(t, "Look at its after effects!")
	// Exception 3
	negative(t, "See your all but nothing system")
	// Boundary check: _ (2)
	negative(t, "It's its")
}

func TestRule_youreNoun(t *testing.T) {
	t.Parallel()

	positive(t, "of you're own!", "of [your] own!")
	// Exception 1
	negative(t, "You're day dreaming")
	negative(t, "You're day drinking")
	negative(t, "You're day dreamers")
	// Exception 2
	negative(t, "you're life changing!")
	negative(t, "You're life savers, you're life wasters, you're life changers!")
	// Exception 3
	negative(t, "You're life.")
	negative(t, "You're life!")
	negative(t, "You're life")
	// Boundary check: _ (1)
	negative(t, "See, you're")
	negative(t, "Note that you're.")
	negative(t, "you're.")
	negative(t, "you're")
}

func TestRule_itsOwn(t *testing.T) {
	t.Parallel()

	positive(t, "sees it's own", "sees [its] own")
	// Boundary check: (1) _
	negative(t, "It's those people")
}

func TestRule_thereOwn(t *testing.T) {
	t.Parallel()

	positive(t, "To each there own", "To each [their] own")
	// Exception 1
	negative(t, "Do any people there own something?")
	negative(t, "Does anyone there own this item?")
	negative(t, "Does someone there own this?")
	negative(t, "Does no one there own that?")
	negative(t, "Do any of you there own these items?")
	// Exception 1 with fused word
	negative(t, "anyone there own it?")
	// Exception 2
	negative(t, "People out there own it.")
	negative(t, "People over there own it.")
	negative(t, "People from there own it.")
	negative(t, "People close to there own it.")
	negative(t, "People near there own it.")
	// Boundary check: (there) _
	negative(t, "Own this item now")
}

func TestRule_whoseBeen(t *testing.T) {
	t.Parallel()

	positive(t, "Whose been?", "[Who's] been?")
	positive(t, "Whose been there?", "[Who's] been there?")
	positive(t, "See that person, whose been there.", "person, [who's] been there")
	// Exception
	negative(t, "whose BEEN")
	// Boundary check: _ (been)
	negative(t, "Whose?")
}

func TestRule_theyreBe(t *testing.T) {
	t.Parallel()

	positive(t, "They're is a cow", "[There's] a cow")
	positive(t, "They're is and they're are", "[There's] and [they] are")
	positive(t, "they're aren't any of those.", "[they] aren't any of")
	// Exception 1
	negative(t, "the difference between their, there, and they're is")
	negative(t, "the difference between there, their, and they're is")
	// Exception 1b
	negative(t, "their / there / they're is sometimes confused")
	negative(t, "there / their / they're is sometimes confused")
	// Exception 2
	negative(t, "what they're is good")
	// Exception 3
	negative(t, "they're is they are")
	negative(t, "they're, aren't they?") // 3
	// Boundary check: _ <be>
	negative(t, "See, they're")
}

func TestRule_theirModal(t *testing.T) {
	t.Parallel()

	positive(t, "Their is", "[There] is")
	positive(t, "Their must be something!", "[There] must be something!")
	// Exception 1
	negative(t, "their IS")
	negative(t, "their BE")
	negative(t, "their ARE")
	negative(t, "their MUST")
	// Exception 2
	negative(t, "the difference between there/they're/their is")
	negative(t, "Those who know the difference between there/they're/their are.")
	negative(t, "the difference between they're/there/their is")
	negative(t, "Those who know the difference between they're/there/their are.")
	// Boundary check: _ (1)
	negative(t, "their")
	// Restriction: _ <modal>
	negative(t, "their item")
}

func TestRule_beNoun(t *testing.T) {
	t.Parallel()

	positive(t, "I am hear", "I am [here]")
	positive(t, "I am board with", "I am [bored] with")
	positive(t, "I am hear to win", "I am [here] to win")
	positive(t, "They are hear", "They are [here]")
	positive(t, "Those people are hear", "people are [here]")
	positive(t, "He is hear", "He is [here]")
	// Tolerate misuse of verbs
	positive(t, "He are hear", "He are [here]")
	positive(t, "They is hear", "They is [here]")
	// Ignore misuse of 'am'
	negative(t, "He am hear")
	negative(t, "They am hear")
	// Boundary check: <be> _hear_
	negative(t, "Hear the silence")
	negative(t, "They can hear you")
	// Boundary check: <be> _board_ (with)
	negative(t, "Look at the board")
	negative(t, "Board the train")
	negative(t, "Board with them")
}

func TestRule_then(t *testing.T) {
	t.Parallel()

	positive(t, "this is better then that", "this is better [than] that")
	positive(t, "I have more then you do", "I have more [than] you do")
	// Exception 1
	negative(t, "wait until it's better then do it")
	// Exception 2
	negative(t, "if it is better then you use it")
	negative(t, "when it is better then get it")
	// Boundary check: no <comparative> before
	negative(t, "Then they went somewhere")
	// Boundary check: <comparative> _ |
	negative(t, "Do it better then")
	// Boundary check: <comparative> | _
	negative(t, "Do you want more? Then you should do this.")
	negative(t, "I need one more; then I have all.")
}

func TestRule_than(t *testing.T) {
	t.Parallel()

	positive(t, "I did this and than I did that", "did this and [then] I did that")
	positive(t, "I did this but than I did that", "did this but [then] I did that")
	positive(t, "I did this yet than I did that", "did this yet [then] I did that")
	positive(t, "I did this, and than I did that", "did this, and [then] I did that")
	// Exception 1
	negative(t, "the difference between then and than is")
	// Exception 2
	negative(t, "better than something and than something else")
	negative(t, "Is it more than they do or than I do?")
	// Exception 3
	negative(t, "They take more action for someone who isn't important yet than they ever did someone who is.")
	negative(t, "They take more for those not important yet than those that are.")
	// Boundary check: (1) <and/but/yet> _
	negative(t, "than")
	negative(t, "this and than")
}

func TestRule_of(t *testing.T) {
	t.Parallel()

	positive(t, "I could of done it", "I could['ve] done it")
	positive(t, "I could not of done it", "I could not['ve] done it")
	positive(t, "I couldn't of done it", "I couldn't['ve] done it")
	positive(t, "I could of went there", "I could['ve gone] there")
	positive(t, "I could not of went there", "I could not['ve gone] there")
	positive(t, "I could of not went there", "I could['ve] not [gone] there")
	positive(t, "I would of done it", "I would['ve] done it")
	positive(t, "I should of done it", "I should['ve] done it")
	positive(t, "I must of done it", "I must['ve] done it")
	// Exception 1
	negative(t, "He could of course do")
	// Exception 2
	negative(t, "would of themselves justify")
	// Exception 3a
	negative(t, "Face the full might of our army!")
	negative(t, "the might of some guy")
	negative(t, "the full might of those people")
	negative(t, "the full might of people")
	negative(t, "no might of him")
	negative(t, "the must of those")
	// Exception 3b
	negative(t, "might of the people")
	negative(t, "might of our people")
	negative(t, "might of them")
	negative(t, "must of them")
	negative(t, "must of the item")
	// Exception 4
	negative(t, "more of this than they would of that")
	// Boundary check: (1) (not)? _ (1)
	negative(t, "of")
	negative(t, "not of this but that")
}

func TestRule_yourAre(t *testing.T) {
	t.Parallel()

	positive(t, "your are", "[you] are")
	// Exception
	negative(t, "your ARE")
	// Boundary check: _ (are)
	negative(t, "your")
}

func TestRule_supposedTo(t *testing.T) {
	t.Parallel()

	positive(t, "I don't supposed to", "[I'm not] supposed to")
	positive(t, "you don't supposed to", "you [aren't] supposed to")
	positive(t, "he doesn't supposed to", "he [isn't] supposed to")
	// Correct mismatched verbs
	positive(t, "I doesn't supposed to", "[I'm not] supposed to")
	positive(t, "you doesn't supposed to", "you [aren't] supposed to")
	positive(t, "he don't supposed to", "he [isn't] supposed to")
	// Past tense (detect person)
	positive(t, "you didn't supposed to", "you [weren't] supposed to")
	positive(t, "he didn't supposed to", "he [wasn't] supposed to")
	negative(t, "this guy didn't supposed to") // unknown person
	// Boundary check: (2) _ (to)
	negative(t, "They supposed")
	negative(t, "They supposed not")
	negative(t, "become supposed to")
	// Restriction: <be>, not <modal> _
	negative(t, "I can't supposed to")
}

func TestRule_whomBe(t *testing.T) {
	t.Parallel()

	positive(t, "he whom is", "he [who] is")
	positive(t, "a person whom is", "person [who] is")
	positive(t, "a person whom was", "person [who] was")
	positive(t, "I whom was", "I [who] was")
	positive(t, "people whom are", "people [who] are")
	positive(t, "people whom were", "people [who] were")
	positive(t, "those whom are", "those [who] are")
	positive(t, "those whom were", "those [who] were")
	positive(t, "see whomever is", "see [whoever] is")
	positive(t, "Whomever is", "[Whoever] is")
	// "Whoever" is always singular
	positive(t, "Whomever am", "[Whoever is]")
	positive(t, "Whomever be", "[Whoever is]")
	positive(t, "Whomever are", "[Whoever is]")
	// Correct improper verbs
	positive(t, "I whom are", "I [who am]")
	positive(t, "I whom were", "I [who was]")
	positive(t, "a person whom are", "person [who is]")
	positive(t, "a person whom were", "person [who was]")
	positive(t, "people whom is", "people [who are]")
	positive(t, "people whom was", "people [who were]")
	positive(t, "You see me, whom is your friend.", "see me, [who am] your friend")
	// REMOVED: Guess third-person singular if nothing precedes it
	// positive(t, "Whom is it?", "[Who] is it?")
	// A preceding word is required for verb conjugations
	negative(t, "UnknownEntity whom is")
	// Other verbs might be valid for "whom"
	negative(t, "Whom do they see?")
	// Boundary check: (1 if not whomever) _ <be>
	negative(t, "for whom?")
	negative(t, "whom they say")
	// Restrictions: unrecognized nouns
	negative(t, "blah whom is")
}
