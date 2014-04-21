#!/usr/bin/env python


def do(self, cur):
    """Keyword: own
    Src: it's _
    Dst: [its] _
    """
    if not self.sequence.prev_has_continuous(1):
        return
    if self.sequence.prev_word(1).word_lower != "it's":
        return
    self.matched('its_po')
    self.sequence.prev_word(1).replace_autocap("its")
    self.rerun.add(10)  # Rerun: possesive_as_be for its
    cur.mark_common()
