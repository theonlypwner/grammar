#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""Unit Tests for grammar"""

import grammar

from io import StringIO
import logging
from logging import StreamHandler
from logging.handlers import MemoryHandler
import sys
import unittest


class ParserFunctions(unittest.TestCase):
    options = {
        'do_decode_html': False,
        'do_ellipsis': False,
        'do_quotations': True,  # uses regex
        'do_askfm': False,  # uses regex
        'do_fixcaps': False,  # uses regex
        'do_fixi': False,  # uses regex
        'do_fixnewline': True,  # uses regex
    }

    def setUp(self):
        self.parser = grammar.CorrectionManager()

    def load(self, text, **kwargs):
        options = self.options
        options.update(kwargs)
        if self.parser.load_text(text, **options):
            return self.parser.corrections[0]
        return False

    def positive(self, text, result=None):
        if result:
            self.assertEqual(self.load(text), result)
        else:
            self.assertTrue(self.load(text))

    def negative(self, text):
        self.assertFalse(self.load(text))

    # Misc Tests
    def test_wording_first_cap(self):
        self.assertEqual(grammar.Wording.first_cap('abc'), 'Abc')
        self.assertEqual(grammar.Wording.first_cap(
            '23-234234{:[[;;""]:]abc'), '23-234234{:[[;;""]:]Abc')

    def test_wording_english_join(self):
        self.assertEqual(grammar.Wording.english_join(['a']), 'a')
        self.assertEqual(grammar.Wording.english_join(['a', 'b']), 'a and b')
        self.assertEqual(
            grammar.Wording.english_join(['a', 'b', 'c']), 'a, b and c')

    def test_transform(self):
        """Test the text transformations."""
        # Create copy with all options set to True
        options = self.options
        for k in options:
            options[k] = True
        # Check helper function

        def check_transform(text, expected):
            self.assertEqual(
                grammar.Transformers.transform(text, **options), expected)
        # do_decode_html
        check_transform('&#39;&quot;&gt;&lt;&amp;', '\'"><&')
        # do_ellipsis
        check_transform('...', u'…')
        # do_quotations
        check_transform('He said, "hi everybody!" They heard.',
                        u'He said, … They heard.')
        # do_askfm
        check_transform(
            u'Who are you? — I am he. https://askfm/blah', 'I am he.')
        # do_fixcaps
        check_transform('CONTENT MUST NOT BE WRITTEN ENTIRELY IN CAPITALS',
                        'Content must not be written entirely in capitals')
        check_transform('ALLCAPS @lower http://is/ignored',
                        'Allcaps @lower http://is/ignored')
        check_transform('Titlecase Is Annoying Too For Some People',
                        'Titlecase is annoying too for some people')
        # do_fixi
        check_transform("i don't know how to capitalize i",
                        "I don't know how to capitalize I")
        # do_fixnewline
        check_transform('1\n2 \n 3', '1 / 2 / 3')

    # Parser Tests
    def test_load_regular(self):
        """Try to properly load a regular sentence."""
        self.negative("This sentence is fine.")

    def test_possessive_as_be(self):
        self.positive("Its here or there", "[It's] here or there")
        self.positive("Your the best", "[You're] the best")
        self.positive("Whose your friend?", "[Who's] your friend?")
        # Exception 1
        self.negative("Some author, whose THE BOOK does, is")
        # Exception 2
        self.negative("Look at its after effects!")
        # Exception 3
        self.negative("See your all but nothing system")
        # Boundary check: _ (2)
        self.negative("It's its")

    def test_youre_noun(self):
        self.positive("of you're own!", "of [your] own!")
        # Exception 1
        self.negative("You're day dreamers")
        self.negative("You're day dreaming")
        # Exception 2
        self.negative(
            "You're life savers, you're life wasters, you're life changers!")
        self.negative("you're life changing!")
        # Exception 3
        self.negative("You're life.")
        self.negative("You're life!")
        self.negative("You're life")
        # Boundary check: _ (1)
        self.negative("See, you're")

    def test_its_own(self):
        self.positive("sees it's own", "sees [its] own")
        # Boundary check: (1) _
        self.negative("It's those people")

    def test_there_own(self):
        self.positive("To each there own", "To each [their] own")
        # Exception
        self.negative("Do any people out there own something?")
        self.negative("Does anyone out there own this item?")
        self.negative("Does someone out there own this?")
        self.negative("Does no one out there own that?")
        self.negative("Do any of you out there own these items?")
        # Exception with fused word
        self.negative("anyone out there own it?")
        # Boundary check: (there) _
        self.negative("Own this item now")

    def test_whose_been(self):
        self.positive("Whose been there?", "[Who's] been there?")
        self.positive("See that person, whose been there.",
                      "person, [who's] been there")
        # Boundary check: _ (been)
        self.negative("Whose?")

    def test_theyre_be(self):
        self.positive("They're is a cow", "[There's] a cow")
        self.positive("They're is and they're are", "[There's] and [they] are")
        self.positive("they're aren't any of those.", "[they] aren't any of")
        # Exception 1
        self.negative("the difference between their, there, and they're is")
        self.negative("the difference between there, their, and they're is")
        # Exception 2
        self.negative("they're is they are")
        self.negative("they're, aren't they?")  # 2b
        # Boundary check: _ <be>
        self.negative("See, they're")

    def test_their_modal(self):
        self.positive("Their is", "[There] is")
        self.positive("Their must be something!", "[There] must be something!")
        # Boundary check: _ (1)
        self.negative("their")
        # Restriction: _ <modal>
        self.negative("their item")

    def test_be_noun(self):
        self.positive("I am hear", "I am [here]")
        self.positive("I am board with", "I am [bored] with")
        self.positive("I am hear to win", "I am [here] to win")
        self.positive("They are hear", "They are [here]")
        self.positive("Those people are hear", "people are [here]")
        self.positive("He is hear", "He is [here]")
        # Tolerate misuse of verbs
        self.positive("He are hear", "He are [here]")
        self.positive("They is hear", "They is [here]")
        # Ignore misuse of 'am'
        self.negative("He am hear")
        self.negative("They am hear")
        # Boundary check: <be> _hear_
        self.negative("Hear the silence")
        self.negative("They can hear you")
        # Boundary check: <be> _board_ (with)
        self.negative("Look at the board")
        self.negative("Board the train")
        self.negative("Board with them")

    def test_then(self):
        self.positive("this is better then that", "this is better [than] that")
        self.positive("I have more then you do", "I have more [than] you do")
        # Exception 1
        self.negative("wait until it's better then do it")
        # Exception 2
        self.negative("if it is better then you use it")
        self.negative("when it is better then get it")
        # Boundary check: (1) <comparative> _ (1)
        self.negative("Then they went somewhere")
        self.negative("Do it better then")

    def test_than(self):
        self.positive("I did this and than I did that",
                      "did this and [then] I did that")
        # Exception 1
        self.negative("the difference between then and than is")
        # Exception: 2
        self.negative("better than something and than something else")
        self.negative("Is it more than they do or than I do?")
        # Boundary check: (1) <and/but/yet> _
        self.negative("than")
        self.negative("this and than")

    def test_of(self):
        self.positive("I should of done it", "I should['ve] done it")
        self.positive("I should not of done it", "I should not['ve] done it")
        self.positive("I shouldn't of done it", "I shouldn't['ve] done it")
        self.positive("I should of went there", "I should['ve gone] there")
        # Exception 1
        self.negative("He could of course do")
        # Exception 2
        self.negative("would of themselves justify")
        # Exception 3a
        self.negative("Face the full might of our army!")
        self.negative("the might of some guy")
        self.negative("the full might of them")
        self.negative("no might of him")
        # Exception 3b
        self.negative("might of the people")
        # Exception: 4
        self.negative("more of this than they would of that")
        # Boundary check: (1) (not)? _ (1)
        self.negative("of")
        self.negative("not of this but that")

    def test_your_are(self):
        self.positive("your are", "[you] are")
        # Boundary check: _ (are)
        self.negative("your")

    def test_supposed_to(self):
        self.positive("I don't supposed to", "[I'm not] supposed to")
        self.positive("you don't supposed to", "you [aren't] supposed to")
        self.positive("he doesn't supposed to", "he [isn't] supposed to")
        # Correct mismatched verbs
        self.positive("I doesn't supposed to", "[I'm not] supposed to")
        self.positive("you doesn't supposed to", "you [aren't] supposed to")
        self.positive("he don't supposed to", "he [isn't] supposed to")
        # Past tense (detect person)
        self.positive("you didn't supposed to", "you [weren't] supposed to")
        self.positive("he didn't supposed to", "he [wasn't] supposed to")
        self.negative("this guy didn't supposed to")  # unknown person
        # Boundary check: (2) _ (to)
        self.negative("They supposed")
        self.negative("They supposed not")
        self.negative("become supposed to")
        # Restriction: <be>, not <modal> _
        self.negative("I can't supposed to")

    def test_whom_be(self):
        self.positive("he whom is", "he [who] is")
        self.positive("a person whom is", "person [who] is")
        self.positive("a person whom was", "person [who] was")
        self.positive("I whom was", "I [who] was")
        self.positive("people whom are", "people [who] are")
        self.positive("people whom were", "people [who] were")
        self.positive("see whomever is", "see [whoever] is")
        self.positive("Whomever is", "[Whoever] is")
        # "Whoever" is always singular
        self.positive("Whomever am", "[Whoever is]")
        self.positive("Whomever be", "[Whoever is]")
        self.positive("Whomever are", "[Whoever is]")
        # Correct improper verbs
        self.positive("I whom are", "I [who am]")
        self.positive("I whom were", "I [who was]")
        self.positive("a person whom are", "person [who is]")
        self.positive("a person whom were", "person [who was]")
        self.positive("people whom is", "people [who are]")
        self.positive("people whom was", "people [who were]")
        self.positive("You see me, whom is your friend.",
                      "see me, [who am] your friend")
        # A preceding word is required for verb conjugations
        self.negative("Whom is it?")
        # Other verbs might be valid for "whom"
        self.negative("Whom do they see?")
        # Boundary check: (1 if not whomever) _ <be>
        self.negative("for whom?")
        self.negative("whom they say")
        # Restrictions: unrecognized nouns
        self.negative("blah whom is")

    def test_unicode(self):
        self.positive(u"… → their is …", u"→ [there] is")

    def test_case(self):
        # lowercase detection
        self.positive("their is", "[there] is")
        # Titlecase Detection
        self.positive("Their is", "[There] is")
        # UPPERCASE DETECTION
        self.positive("THEIR is", "[THERE] is")

    def test_overlap(self):
        self.positive("Their is you're own", "[There] is [your] own")
        self.positive("Your you're own", "[You're your] own")
        self.positive("Its there own item", "[It's their] own item")
        self.positive("Your don't supposed to!", "[You aren't] supposed to!")

    def test_punctuation(self):
        # Keep question marks and exclamation marks
        self.positive("their is?", "[there] is?")
        self.positive("their is!", "[there] is!")
        # Original punctuation should be out of the brackets
        self.positive("I am hear!", "I am [here]!")
        # Drop leading and trailing spaces
        self.positive(" their is ", "[there] is")
        # Retain double spaces
        self.positive("their  is", "[there]  is")
        # Drop periods, commas, and semicolons
        self.positive("their is.", "[there] is")
        self.positive("their is,", "[there] is")
        self.positive("their is;", "[there] is")

    def test_wording(self):
        """Verify that wording can be generated without failing"""
        self.positive(
            "Their is and your don't supposed to! (blah) They think their is.")
        logging.debug(self.parser.generate_wording('@@')
                      .encode('ascii', 'replace'))
        # not implemented yet
        self.assertRaises(NotImplementedError,
                          self.parser.generate_wording_long, '')

if __name__ == '__main__':
    stream_null = StringIO()
    logging.basicConfig(stream=stream_null, level=logging.DEBUG)
    handler_stream = StreamHandler(stream=sys.stderr)
    handler_mem = MemoryHandler(1024, target=handler_stream)
    handler_mem.setLevel(logging.DEBUG)
    handler_mem.setFormatter(logging.Formatter())
    logging.getLogger().addHandler(handler_mem)
    unittest.main()
