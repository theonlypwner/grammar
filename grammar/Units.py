#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""Sets for the parser"""


class Sets:
    SET_ARTICLE = {'a', 'an', 'the'}

    SET_CONJUNCTION_COORDINATING = {
        'for', 'and', 'nor', 'but', 'or', 'yet', 'so'}

    SET_CONJUNCTION_SUBORDINATING = {'as'}

    SET_CONJUNCTION = SET_CONJUNCTION_COORDINATING | SET_CONJUNCTION_SUBORDINATING

    SET_PREPOSITION_SHORT = {
        'aboard',
        'about',
        'above',
        # 'according to',
        'across',
        'after',  # your after(-)effects [checked]
        'against',
        # 'ahead of',
        'along',
        'alongside',
        'amid',
        'amidst',
        'among',
        'amongst',
        'around',
        # 'as', # also conjunction
        #    'as far as',
        #    'as of',
        #    'as per',
        #    'as regards',
        #    'as well as',
        'aside',
        #    'aside from',
        # 'astride', # not common
        'at',
        # 'athwart', # not common
        'atop',
        # 'barring', # not common, also participle
        # 'because of',
        'before',
        # 'behind', # also noun
        'below',  # your below(-)average <nounp>
        'beneath',
        # 'beside', 'besides', # also noun
        'beyond',  # your beyond(-)<adj> <nounp>
        'between',
        'by',
        #    'by means of',
        # 'circa',  # not common
        # 'close to',
        # 'concerning', # also particple
        # 'despite', # not useful
        'down',  # your down(-)stairs computer
        # 'due to',
        'during',
        'except',
        # 'except for',
        'excluding',  # also particple
        'failing',  # also particple
        # 'far from',
        # 'following', # also particple
        'for',
        'from',
        # 'given', # also past particple
        'in',
        #    'in accordance with',
        #    'in addition to',
        #    'in case of',
        #    'in front of',
        #    'in lieu of',
        #    'in place of',
        #    'in point of',
        #    'in spite of',
        # 'including', # also particple
        # 'inside', # also noun
        #    'inside of',
        # 'instead of',
        'into',
        'like',  # also verb
        # 'mid', # also adj
        # 'minus', # also noun
        'near',
        # 'next', # also adj
        #    'next to',
        # 'notwithstanding of',
        'of',
        'off',
        # 'on', # also adj
        #    'on account of',
        #    'on behalf of',
        #    'on top of',
        'out',  # -of/from
        'outside',  # -of also noun
        # 'over', # your over(-)excited <nounp>
        # 'owing to',
        # 'pace', # outdated
        # 'past', # also noun
        # 'per', # not useful
        # 'plus', # also noun
        # 'prior to',
        # 'pursuant to',
        # 'qua', # formal only
        'regarding',  # also participle
        # 'regardless of',
        # 'round', # also adj
        'sans',
        'save',  # also verb
        # 'since', # also conjunction
        # 'subsequent to',
        # 'than', # also conjunction
        # 'thanks to',
        # 'that of',
        'through',
        'throughout',
        'till',  # informal until, also conjunction
        # 'times', # also noun
        # 'to', # your too-<adj> <nounp>
        'toward',
        'towards',
        'under',  # your under(-)prepared <nounp>
        'underneath',
        'unlike',  # also adj
        'until',
        # 'unto', # not common
        'up',
        'upon',
        'versus',
        'via',
        'with',
        #    'with regard to',
        #    'with respect to',
        'within',
        'without',
        # 'worth', # also noun
    }

    SET_PREPOSITION_EXTENSION = {'over', 'to'}

    SET_PRONOUN_PERSONAL = {
        'i', 'we', 'you', 'he', 'she', 'it', 'they',  # subjective
        'me', 'us', 'you', 'him', 'her', 'it', 'them',  # objective
        # reflexive
        'myself', 'ourselves', 'yourself', 'yourselves', 'himself', 'herself', 'itself', 'themselves',
        'myselves', 'ourself', 'themself',  # *reflexive
    }

    SET_DETERMINER_POSSESSIVE = {
        'my', 'our', 'thy', 'your', 'his', 'her', 'its', 'their'}
    SET_PRONOUN_POSSESSIVE = {
        'mine', 'ours', 'thine', 'yours', 'his', 'hers', 'its', 'theirs'}
    SET_POSSESSIVE = SET_DETERMINER_POSSESSIVE | SET_PRONOUN_POSSESSIVE

    SET_DETERMINER_YOUR = {
        'each',  # 'every',
        'either', 'neither',
        'some', 'any', 'no',
        # 'much', 'many', 'most', 'more',
        # 'little', 'less', 'least',
        # 'few', 'fewer', 'fewest',
        'all', 'both',  # 'half',
        # 'several',
        'enough',
    }
    # | set([])
    SET_DETERMINER = SET_ARTICLE | SET_DETERMINER_POSSESSIVE | SET_DETERMINER_YOUR

    # SET_STOP_ prefix: potentially useful for other purposes, but currently
    # only for stop words
    SET_STOP_DETERMINER_INC = {'each', 'none', 'few', 'most', 'more', }
    SET_STOP_CONJUNCITON_SUBORDINATING_INC = {
        'after',
        'although',
        'as',  # also as if
        'because',
        'before',
        # 'even though',
        'if',  # noun
        'once',  # adverb
        'only',  # adverb, adjective
        'than',
        # 'that',
        'though',
        'unless',
        'until',
        'whether',
        'when',
        'where',
        'while',  # noun, adverb, verb
    }
    SET_STOP_MODAL_INC = {'cannot', 'could', 'ought', 'should', 'would'}
    SET_STOP_INTERROGATIVE = {
        'how', 'what', 'who', 'whom', 'why', 'when', 'where', 'which'}

    # based on http://meta.wikimedia.org/wiki/Stop_word_list/google_stop_word_list
    # which uses http://www.ranks.nl/resources/stopwords.html
    SET_COMMON = SET_DETERMINER | SET_PREPOSITION_SHORT | SET_PREPOSITION_EXTENSION | SET_CONJUNCTION | SET_PRONOUN_PERSONAL | SET_STOP_DETERMINER_INC | SET_STOP_CONJUNCITON_SUBORDINATING_INC | SET_STOP_MODAL_INC | SET_STOP_INTERROGATIVE | {
        # to "be"
        'be', 'am', 'are', 'is', 'was', 'were', 'been', 'being',
        # to "do"
        'do', 'does', 'did', 'done', 'doing',
        # to "have"
        'have', 'has', 'had', 'having',
        # contractions
        "aren't", "can't", "couldn't", "didn't", "doesn't", "don't", "hadn't", "hasn't", "haven't", "he'd", "he'll", "he's", "here's", "how's", "i'd", "i'll", "i'm", "i've", "isn't", "it's", "let's", "mustn't", "shan't", "she'd", "she'll", "she's", "shouldn't", "that's", "there's", "they'd", "they'll", "they're", "they've", "wasn't", "we", "we'd", "we'll", "we're", "we've", "weren't", "what's", "when's", "where's", "who's", "why's", "won't", "wouldn't", "you'd", "you'll", "you're", "you've",
        # special
        'again',  # adverb
        'further',  # adverb, adjective
        'not',  # adverb, noun
        'other',  # adjective
        'own',  # adjective, verb
        'same',  # adjective, pronoun, adverb
        'such',  # adjective
        'that',  # pronoun, adverb, adjective, conjunction
        'then',  # adverb
        'there',  # adverb
        'these',  # pronoun, adjective, adverb
        'this',  # ^
        'those',  # ^
        'too',  # adverb
        'very',  # adjective, adverb?
    }

# The "building blocks" of sequences!


class SequenceUnit(object):

    def __init__(self, original):
        self.original = original
        # Word flags:
        # -1 = hidden, 0 = verbatim, 1 = modified
        # 2 = common word (from correction)
        self.flags = 0
        self.new_text = None

    def replace(self, new_text):
        self.flags = 1
        self.new_text = new_text

    def get_final(self):
        if self.new_text is not None:  # and self.flags == 1
            return self.new_text
        return self.original

UnitToString = SequenceUnit.get_final


class Space(SequenceUnit):  # word separators - can also have punctuation

    def __init__(self, spacer):
        # (space) and '/' do not break
        self.sentenceBreaker = ('.' in spacer) or (
            '!' in spacer) or ('?' in spacer)
        self.anyBreaker = self.sentenceBreaker or (',' in spacer) or (';' in spacer) or (':' in spacer) or (
            '(' in spacer) or (')' in spacer) or ('[' in spacer) or (']' in spacer) or ('{' in spacer) or ('}' in spacer)
        super(Space, self).__init__(spacer)


class Word(SequenceUnit):

    def __init__(self, word):
        super(Word, self).__init__(word)
        if word:
            # use lower-case word for comparing (current state)
            self.word_lower = word.lower()
            if self.original.isupper():
                self.caps = 2  # all caps
            elif self.original[0].isupper():
                self.caps = 1  # title caps
            else:
                self.caps = 0  # lowercase
        else:
            # empty word
            self.flags = -1
            self.word_lower = ''
            self.caps = 0  # as lowercase

    def replace(self, new_text):
        self.word_lower = new_text.lower()
        return super(Word, self).replace(new_text)

    def replace_autocap(self, new_text):
        if self.caps == 1:
            new_text = new_text[0].upper() + new_text[1:]
        elif self.caps == 2:
            new_text = new_text.upper()
        return self.replace(new_text)

    def mark_common(self):
        """Mark this as a common word from a correction"""
        if self.flags != 1:
            self.flags = 2

    NEAR_EXCEPTIONS = {'', u"…"}

    def is_near(self):
        """Returns whether this word may be included as a "near" word"""
        return not(self.word_lower in self.NEAR_EXCEPTIONS or self.word_lower.startswith('http://') or self.word_lower.startswith('https://'))

    def is_common(self):
        """Returns whether this word may be considered a "common" word"""
        return self.flags == 2 or self.word_lower in Sets.SET_COMMON

    def is_article(self):
        return self.word_lower in Sets.SET_ARTICLE

    def is_agent(self):
        # possible issues
        #  - false negatives: actor, actors
        #  - false positives: alter, deter, enter, prefer, ...
        return self.word_lower.endswith('er') or self.word_lower.endswith('ers')

    def is_determiner(self):
        return self.word_lower in Sets.SET_DETERMINER

    def is_conjunction_coordinating(self):
        return self.word_lower in Sets.SET_CONJUNCTION_COORDINATING

    def is_preposition(self):
        return self.word_lower in Sets.SET_PREPOSITION_SHORT

#    def is_participle_past(self):
#        if self.word_lower in ['been', 'thrown']:
#            return True
#        return self.word_lower.endswith('ed')

    def is_participle_present(self):
        return self.word_lower.endswith('ing')

    def is_pronoun_personal(self):
        return self.word_lower in Sets.SET_PRONOUN_PERSONAL

    def is_possessive(self):
        return self.word_lower in Sets.SET_POSSESSIVE

    def is_possessive_determiner(self):
        return self.word_lower in Sets.SET_DETERMINER_POSSESSIVE

    def is_possessive_pronoun(self):
        return self.word_lower in Sets.SET_PRONOUN_POSSESSIVE


class SequenceManager(list):

    def __init__(self, *args, **kwargs):
        self.position = 0
        super(SequenceManager, self).__init__(*args, **kwargs)

    def iter_words(self, start=0):
        """Returns an iterator that loops through the words."""
        self.position = 0
        while self.position < len(self):
            yield self[self.position]
            self.position += 2

    def prev_has(self, words):
        """Determine whether there are at least n preceding words"""
        # @ 0: [ W ] S W S
        # @ 2: W S [ W ] S
        # change to > if SWS is required
        return self.position >= (2 * words)

    def prev_has_continuous(self, words, already_checked=0, sentences=False):
        """Check if the previous n words are within the same block/sentence."""
        if not self.prev_has(words):
            return False
        # check for lack of breakers
        if sentences:
            for i in range(
                    self.position - 1 - already_checked * 2,
                    self.position - 1 - words * 2,
                    -2):
                if self[i].sentenceBreaker:
                    return False
        else:
            for i in range(
                    self.position - 1 - already_checked * 2,
                    self.position - 1 - words * 2,
                    -2):
                if self[i].anyBreaker:
                    return False
        return True

    def next_has(self, words):
        """Determine whether there are at least n remaining words"""
        # @ -4 W S [ W ] S W S
        # @ -2 W S W S [ W ] S
        # there is always a space after
        return self.position + words * 2 < len(self)

    def next_has_continuous(self, words, already_checked=0, sentences=False):
        """Check if the next n words are within the same block/sentence."""
        if not self.next_has(words):
            return False
        # check for lack of breakers
        if sentences:  # pragma: no cover
            for i in range(
                    self.position + 1 + already_checked * 2,
                    self.position + 1 + words * 2,
                    2):
                if self[i].sentenceBreaker:
                    return False
        else:
            for i in range(
                    self.position + 1 + already_checked * 2,
                    self.position + 1 + words * 2,
                    2):
                if self[i].anyBreaker:
                    return False
        return True

    def prev_get_words_continuous(self, sentences=False):
        """Returns an list of the closest previous words until a breaker."""
        if not self.position:
            return []
        i = self.position - 1
        if sentences:  # pragma: no cover
            while i >= 0:  # 1
                if self[i].sentenceBreaker:
                    i -= 2
                    break
                i -= 2
        else:
            while i >= 0:  # 1
                if self[i].anyBreaker:
                    i -= 2
                    break
                i -= 2
        if i == -1:  # because -2 is needed to include 0, which is rerouted...
            ret = self[self.position - 2:i + 1:-2]
            ret.append(self[0])
            return ret
        return self[self.position - 2:i + 1:-2]

    def prev_word(self, n):
        """Get the nth previous word"""
        return self[self.position - n * 2]

    def next_word(self, n):
        """Get the nth next word"""
        return self[self.position + n * 2]

    def prev_space(self, n):
        """Get the nth previous space"""
        return self[self.position - n * 2 + 1]

    def next_space(self, n):
        """Get the nth next space"""
        return self[self.position + n * 2 - 1]
