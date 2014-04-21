#!/usr/bin/env python

# NOTE: no need for 'am' except for "there am I"
from .theyre_be import SET_MODAL


def do(self, cur):
    """Keyword: their
    Src: _ <modal>
    Dst: [there] <modal>
    """
    if not self.sequence.next_has_continuous(1):
        return
    if self.sequence.next_word(1).word_lower not in SET_MODAL:
        return
    self.matched('their_be')
    cur.replace_autocap("there")
    self.sequence.next_word(1).mark_common()
