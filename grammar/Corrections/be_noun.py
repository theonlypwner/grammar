#!/usr/bin/env python

# no 'am'
SET_BE = {'be', 'is', 'are', "isn't", "aren't"}


def do(self, cur):
    """Keyword: hear/board
    Src: ((?<=I )am|are|be) _hear_
    Dst: ... here
    Src: ((?<=I )am|are|be) _board_ with
    Dst: ... bored with

    Removed: they {are board of} directors
    """
    if cur.word_lower != 'hear':
        if cur.word_lower == 'board':
            if not self.sequence.next_has_continuous(1):
                return
            if self.sequence.next_word(1).word_lower != 'with':
                return
        # else: return
    if not self.sequence.prev_has_continuous(2):
        return
    if self.sequence.prev_word(1).word_lower == 'am':
        if self.sequence.prev_word(2).word_lower == 'i':
            self.sequence.prev_word(2).mark_common()
        else:
            return
    elif self.sequence.prev_word(1).word_lower not in SET_BE:
        return
    self.matched(cur.word_lower)  # in ['hear', 'board']
    self.sequence.prev_word(1).mark_common()
    cur.replace('here' if cur.word_lower == 'hear' else 'bored')
