#!/usr/bin/env python

SET_BUTYET = set(['but', 'yet'])
from .then import SET_COMPARATIVE    # replicated


def do(self, cur):
    """Keyword: than
    Src: <word>(,) <conjunction:(and|but|yet)> _ <word>
    Dst: ... [then] ...
    """
    if not self.sequence.next_has_continuous(1):
        return
    prev_words = self.sequence.prev_get_words_continuous()
    if len(prev_words) < 2:
        return
    prev_word = prev_words[0]
    if prev_word.word_lower == 'and':
        # Exception 1: the difference between 'then' {and 'than'}
        if prev_words[1].word_lower == 'then':
            return
    elif prev_word.word_lower not in SET_BUTYET:
        return
    # Exception 2:
    # <comparative:(better|worse|more|less)> than N<NP>+ {and/or than} <NP>+
    if prev_word.word_lower == 'and':  # in ['and', 'or']
        if len(prev_words) >= 4:  # [better than <NP>+ (prev)] than
            # do not check most recent 3
            for i in range(3, len(prev_words)):
                # inverted with below
                if prev_words[i - 1].word_lower == 'than':
                    if prev_words[i].word_lower in SET_COMPARATIVE:
                        # one scenario [1] is <word>+ and [0] is prev_word
                        return
    self.matched('than')
    prev_word.mark_common()
    cur.replace('then')
