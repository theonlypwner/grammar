#!/usr/bin/env python


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
        if next_word_1.word_lower != 'day' and next_word_1.word_lower != 'life':
            return

        if self.sequence.next_has_continuous(2, 1):
            # Exception 1: "[you're day/life] -ing"
            if self.sequence.next_word(2).is_participle_present():
                return
            # Exception 2: "[you're day/life] -ers?"
            elif self.sequence.next_word(2).is_agent():
                return
        elif next_word_1.word_lower == 'life':
            # Exception 3: "you're life." [no following word]
            return

    self.matched('your_po')
    cur.replace_autocap("your")
    self.rerun.add(10)  # Rerun: possessive_as_be for your
    next_word_1.mark_common()
