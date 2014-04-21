#!/usr/bin/env python

SET_DAY_EXCEPT = set(['dreamers', 'dreaming'])
SET_LIFE_EXCEPT = set(
    ['saver', 'savers', 'waster', 'wasters', 'changer', 'changers'])


def do(self, cur):
    """Keyword: you're
    Src: _ <possessed_noun:(day|life)|own>
    Dst: [your] ...

    Removed:
    (you're man) enough
    (you're life) changers
    """
    if not self.sequence.next_has_continuous(1):
        return
    next_word_1 = self.sequence.next_word(1)
    if next_word_1.word_lower != 'own':
        if next_word_1.word_lower == 'day':
            if self.sequence.next_has_continuous(2, 1):
                # Exception 1: {you're day} dream(ers|ing)
                if self.sequence.next_word(2).word_lower in SET_DAY_EXCEPT:
                    return
        elif next_word_1.word_lower == 'life':
            if self.sequence.next_has_continuous(2, 1):
                # Exception 2a: "[you're life] -ing"
                if self.sequence.next_word(2).is_participle_present():
                    return
                # Exception 2b: "[you're life] (sav|wast|chang)ers?"
                elif self.sequence.next_word(2).word_lower in SET_LIFE_EXCEPT:
                    return
                # Allowed: you're life _
            else:
                # Exception 3: "you're life." [no following word]
                return
        else:
            return
    self.matched('your_po')
    cur.replace_autocap("your")
    self.rerun.add(10)  # Rerun: possessive_as_be for your
    next_word_1.mark_common()
