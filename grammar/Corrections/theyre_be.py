#!/usr/bin/env python

from .be_noun import SET_BE
from .of import SET_MODAL as SET_OF_MODAL
SET_THERETHEIR = set(['there', 'their'])
SET_MODAL_SINGULAR = set(['is', "isn't"])
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
    # Exception 1: the difference between their/there, <, and {they're is}
    # Exception 2: what {they're is}
    if (self.sequence.prev_has_continuous(1) and self.sequence.prev_has(2) and
        self.sequence.prev_word(1).word_lower == 'and' and
            self.sequence.prev_word(2).word_lower in SET_THERETHEIR):
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
