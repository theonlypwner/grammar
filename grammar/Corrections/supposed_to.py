#!/usr/bin/env python

SET_3 = {'he', 'she', 'it'}
SET_2 = {'we', 'you', 'they'}


def do(self, cur):
    """Keyword: supposed
    Src: (<word> doesn't|I don't|<word> (?<!I )don't) _ to
    Dst: [isn't|I'm not|aren't] _ to

    NOTE: causes extension - don't -> [aren't], didn't -> [wasn't/weren't]
    """
    if not self.sequence.next_has_continuous(1):
        return
    if self.sequence.next_word(1).word_lower != 'to':
        return
    if not self.sequence.prev_has_continuous(2):
        return
    person = 2  # Note that 2nd person = plural
    prev_word_1 = self.sequence.prev_word(1)
    prev_word_2 = self.sequence.prev_word(2)
    if prev_word_1.word_lower == "don't":
        if prev_word_2.word_lower == "i":
            person = 1
        elif prev_word_2.word_lower in SET_3:
            person = 3
        # else: person = 2
    elif prev_word_1.word_lower == "doesn't":
        if prev_word_2.word_lower == "i":
            person = 1
        elif not prev_word_2.word_lower in SET_2:
            person = 3
        # else: person = 2
    elif prev_word_1.word_lower == "didn't":
        if prev_word_2.word_lower == "i" or prev_word_2.word_lower in SET_2:
            person = 4
        elif prev_word_2.word_lower in SET_3:
            person = 5
        else:
            return  # unknown conjugation
    else:
        return
    self.matched('supposed-to')
    cur.mark_common()
    self.sequence.next_word(1).mark_common()
    if person == 1:
        # special: [I'm not] supposed to
        prev_word_2.replace("I'm")
        prev_word_1.replace('not')
    elif person == 2:
        prev_word_1.replace("aren't")
        self.rerun.add(11)  # Rerun: your_are for "you aren't"
    elif person == 3:
        prev_word_1.replace("isn't")
    elif person == 4:
        prev_word_1.replace("weren't")
    else:  # if person == 5:
        prev_word_1.replace("wasn't")
