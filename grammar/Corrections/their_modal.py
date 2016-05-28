#!/usr/bin/env python

# NOTE: no need for 'am' except for "there am I"
from .theyre_be import SET_MODAL
SET_THERE_THEYRE = {'there', "they're"}
SET_THERE_THEYRE_AND = {'there', "they're", 'and'}


def do(self, cur):
    """Keyword: their
    Src: _ <modal>
    Dst: [there] <modal>
    """
    if not self.sequence.next_has_continuous(1):
        return
    if self.sequence.next_word(1).word_lower not in SET_MODAL:
        return
    if self.sequence.prev_has_continuous(1, 0, True):
        # Exception 1: the difference between (there/they're), (they're/there)?, and {their is}
        # Exception 1b: (they're/their) / (they're/there) / {their is} sometimes confused
        # This exception set has three forms:
        # 1. T / T / _
        # 2. T and _
        # 3. T, T, and _ # covered by #2
        if (self.sequence.prev_has_continuous(2, 1, True) and
                self.sequence.prev_word(2).word_lower in SET_THERE_THEYRE and
                self.sequence.prev_word(1).word_lower in SET_THERE_THEYRE_AND):
            return
    self.matched('their_be')
    cur.replace_autocap("there")
    self.sequence.next_word(1).mark_common()
