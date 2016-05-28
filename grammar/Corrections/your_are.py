#!/usr/bin/env python

SET_ARE = {'are', "aren't"}


def do(self, cur):
    """Keyword: your
    Src: _ are(n't)?
    Dst: [you] ...
    Alt: your are[a], your ar[t]
    """
    if not self.sequence.next_has_continuous(1):
        return
    if self.sequence.next_word(1).word_lower not in SET_ARE:
        return
    self.matched('your-are')
    cur.replace_autocap('you')
    self.sequence.next_word(1).mark_common()
