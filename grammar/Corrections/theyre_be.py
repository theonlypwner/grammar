#!/usr/bin/env python

from .be_noun import SET_BE
from .of import SET_MODAL as SET_OF_MODAL
SET_THERETHEIR = {'there', 'their'}
SET_THERETHEIRAND = {'there', 'their', 'and'}
SET_MODAL_SINGULAR = {'is', "isn't"}
# 'is' is already checked
SET_MODAL = SET_BE | SET_OF_MODAL


def do(self, cur):
    """Keyword: they're
    Src: _ is
    Dst: [there's]
    Src: _ are
    Dst: [there] are
    Alt: [they] are
     - they are being (GOOD)
     - there are being (BAD)
    # Removed: {they're day} dreaming
    # Removed: Be {there day} and night
    # Removed: "they're be" is full of improper usage
    """

    if self.sequence.prev_has_continuous(1, 0, True):
        # Exception 1: the difference between (their/there), (there/their)?, and {they're is}
        # Exception 1b: (their/there) / (there/their) / {they're is} sometimes confused
        # This exception set has three forms:
        # 1. T / T / _
        # 2. T and _
        # 3. T, T, and _ # covered by #2
        if (self.sequence.prev_has_continuous(2, 1, True) and
                self.sequence.prev_word(2).word_lower in SET_THERETHEIR and
                self.sequence.prev_word(1).word_lower in SET_THERETHEIRAND):
            return

        # Exception 2: what {they're is}
        if self.sequence.prev_word(1).word_lower == 'what':
            return
    # Exception 3b: {'they're' is} 'they are'
    # Exception 3: {they're, aren't} they?
    if self.sequence.next_has_continuous(2):
        if self.sequence.next_word(2).word_lower == 'they':
            return
    elif not self.sequence.next_has_continuous(1):
        return
    next_word = self.sequence.next_word(1)
    if next_word.word_lower in SET_MODAL_SINGULAR:
        self.matched('theyre_be')
        cur.replace_autocap("there's")
        self.sequence.next_space(1).replace('')  # collapse space
        next_word.replace('')  # delete next word (is)
    elif next_word.word_lower in SET_MODAL:
        self.matched('theyre_are')
        # cur.replace_autocap("they/there")
        # they are being is OK, but there are being is NOT, but they can
        # always replace there
        cur.replace_autocap("they")
        next_word.mark_common()
