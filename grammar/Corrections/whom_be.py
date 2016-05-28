#!/usr/bin/env python

SET_BE_PRESENT = {'be', 'am', 'are', 'is'}
SET_BE_PAST = {'was', 'were'}
SET_1 = {'i', 'me', 'myself'}
SET_2 = {'people', 'persons', 'we', 'us', 'you', 'they', 'them'}
SET_3 = {'person', 'guy', 'he', 'him', 'she', 'her', 'it'}


def do(self, cur):
    """Keyword: whom/whomever
    Src: _ <be>
    Dst: [who/whoever] ...
    """
    person = 0
    if cur.word_lower == 'whomever':
        person = -1
    # elif cur.word_lower != 'whom':
    #     return
    if not self.sequence.next_has_continuous(1):
        return
    next_word_1 = self.sequence.next_word(1)
    if next_word_1.word_lower in SET_BE_PRESENT:
        present = True
    elif next_word_1.word_lower in SET_BE_PAST:
        present = False
    else:
        return
    if not person:
        if not self.sequence.prev_has_continuous(1, sentences=True):
            return
        prev_word_1_ = self.sequence.prev_word(1).word_lower
        if prev_word_1_ in SET_1:
            # Potential issue: the person (sitting across from {me) who is}
            person = 1
        elif prev_word_1_ in SET_2:
            person = 2
        elif prev_word_1_ in SET_3:
            person = 3
        else:
            return
    self.matched('whom')
    if person == -1:
        cur.replace_autocap('whoever')
        person = 3
    else:
        cur.replace('who')
    next_word_1.mark_common()
    next_word_real = (('am', 'are', 'is')
                      if present else ('was', 'were', 'was'))[person - 1]
    if next_word_1.word_lower != next_word_real:
        next_word_1.replace(next_word_real)
