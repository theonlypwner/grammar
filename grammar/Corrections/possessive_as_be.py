#!/usr/bin/env python

SET_EFFECT = {'effect', 'effects'}
NEW_WORD = {
    'its': "it's",
    'your': "you're",
    'whose': "who's",
}


def do(self, cur):
    """Keyword: (its|your|whose)
    Src: _ (<article|possessive_determiner|possessive_pronoun|preposition>|here|not) <word>
    Dst: (it's|you're|who's) <word>
    
    NOTE: causes extension
    """
    if not self.sequence.next_has_continuous(2):
        return
    next_word_1 = self.sequence.next_word(1)
    # Exception 1: Some author, {whose THE} BOOK does, is
    # - book titles in ALLCAPS should be ignored
    if next_word_1.caps == 2:
        return
    # Exception 2: {its after} effects
    if next_word_1.word_lower == 'after':
        if self.sequence.next_word(2).word_lower in SET_EFFECT:
            return
    # Exception 3: {your all} but nothing system
    elif next_word_1.word_lower == 'all':
        if self.sequence.next_word(2).word_lower == 'but':
            return
    # 'not' removed: <possessive> not <participle_present> (gerund)
    elif next_word_1.word_lower != 'here':
        if (not next_word_1.is_preposition() and
                not next_word_1.is_determiner() and
                not next_word_1.is_possessive_pronoun()):
            return
    self.matched(cur.word_lower)  # in ['its', 'your', 'whose']
    cur.replace_autocap(NEW_WORD[cur.word_lower])
    next_word_1.mark_common()
