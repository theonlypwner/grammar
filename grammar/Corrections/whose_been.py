#!/usr/bin/env python


def do(self, cur):
    """Keyword: whose
    Src: _ been
    Dst: [who's] been
    """
    if not self.sequence.next_has_continuous(1):
        return
    if self.sequence.next_word(1).word_lower != 'been':
        return
    self.matched('whose_has')
    cur.replace_autocap("who's")
    self.sequence.next_word(1).mark_common()
