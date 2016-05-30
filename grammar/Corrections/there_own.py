#!/usr/bin/env python

SET_do_check_there_own_FUSED = {
    'anyone', 'anybody', 'someone', 'somebody', 'no-one', 'no-body', 'noone', 'nobody'}
SET_do_check_there_own_UNFUSED1 = {'any', 'some', 'no'}
SET_do_check_there_own_UNFUSED2 = {'one', 'person', 'people', 'body', 'of'}


def do(self, cur):
    """Keyword: own
    Src: there _
    Dst: [their] _

    Removed: Is {there a} <noun>?
    """
    prev_words = self.sequence.prev_get_words_continuous()
    if (not prev_words) or prev_words[0].word_lower != 'there':
        return
    # Exception: Do/es (any/some/no-/no )one out {there own} something?"
    for i in range(1, len(prev_words)):
        if prev_words[i].word_lower in SET_do_check_there_own_FUSED:
            return
        elif i + 1 < len(prev_words) and prev_words[i].word_lower in SET_do_check_there_own_UNFUSED2:
            if prev_words[i + 1].word_lower in SET_do_check_there_own_UNFUSED1:
                return
        # no need to find ['do', 'does'] since people sometimes skip it
    self.matched('there_their')
    prev_words[0].replace_autocap("their")
    cur.mark_common()
    self.rerun.add(10)  # Rerun: possessive_as_be for their
