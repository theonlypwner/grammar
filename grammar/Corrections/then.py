#!/usr/bin/env python

SET_COMPARATIVE = set(['better', 'worse', 'more', 'less'])
SET_NOFOLLOW = set(['lol', 'be', 'do', 'did', 'get', 'got'])


def do(self, cur):
    """Keyword: then
    Src: <word> (better|more|less) _ <word>
    Dst: ... [than] <word>
    """
    prev_words = self.sequence.prev_get_words_continuous()
    if len(prev_words) < 2:
        return
    if prev_words[0].word_lower not in SET_COMPARATIVE:
        return
    if not self.sequence.next_has_continuous(1):
        return
    # Exception 1: then <verb:(be|do|did|get|got)|lol> <noun>
    if self.sequence.next_word(1).word_lower in SET_NOFOLLOW:
        return
    # Exception 2: if/when ... {better(,) then}
    for p in prev_words:
        if p.word_lower == 'if' or p.word_lower == 'when':
            return
    self.matched('then')
    prev_words[0].mark_common()
    cur.replace('than')
