#!/usr/bin/env python

SET_MODAL = {'could', 'should', 'would', 'must',
                 "couldn't", "shouldn't", "wouldn't", "mustn't"}
SET_CMP = {'more', 'less'}


def do(self, cur):
    """Keyword: of
    Src: <modal:((could|should|would|might|must)(n't)?)> (not)? _ <verb_past>
    Dst: ...['ve] past_participle(<verb_past>)

    Removed: Bad Modals:
    *can
    first {May of} 2000
    will
    shall
    ought to (low traffic)
    in {need not of} better days
    had better (awkward 've)
    """
    prev_words = self.sequence.prev_get_words_continuous()
    if len(prev_words) < 1:
        return
    not_shift = 0
    if prev_words[0].word_lower == 'not':
        if len(prev_words) < 2:
            return
        not_shift = 1
    next_word_1 = None
    if self.sequence.next_has_continuous(1):
        next_word_1 = self.sequence.next_word(1)
        # Exception 1: He {could(,) of} course(,) ...
        if next_word_1.word_lower == 'course':
            return
        # Exception 2: this {would of} themselves justify
        elif next_word_1.is_pronoun_personal():
            return
    if prev_words[not_shift].word_lower == 'might':
        # Exception 3a: <determiner> {might of}
        if len(prev_words) >= 2 + not_shift and prev_words[1 + not_shift].is_determiner():
            return
        # Exception 3b: {might of} <determiner|NP>
        if next_word_1 and (next_word_1.is_determiner() or next_word_1.is_pronoun_personal()):
            return
    elif prev_words[not_shift].word_lower not in SET_MODAL:
        return
    # Exception 4: (more|less) (of) <NP>+ than <NP>+ {<modal> of} <NP>+
    # [more of <word>+ than <word>+ (prev) not?] of
    if len(prev_words) >= 6 + not_shift:
        finding = 0
        # skip <modal> and <word>, and remove 1 at end
        for i in range(2 + not_shift, len(prev_words) - 1):
            if finding == 0:
                if prev_words[i].word_lower == 'than':
                    finding = 1
            elif finding == 1:  # skip <word>
                finding = 2
            elif finding == 2:
                if prev_words[i].word_lower == 'of':
                    # if i >= 1 and  # (updated range instead)
                    if prev_words[i + 1].word_lower in SET_CMP:
                        return
    self.matched('of')
    # for i in range(not_shift + 1):
    prev_words[0].mark_common()
    if not_shift:
        prev_words[1].mark_common()
    self.sequence.prev_space(1).replace('')  # collapse space
    cur.replace("'ve")
    # Fix: 've <verb_past_simple> -> <verb_past_perfect>
    base_verb = next_word_1
    if next_word_1.word_lower == 'not':
        # this check can be moved up, especially if needed by other exceptions
        if self.sequence.next_has_continuous(2):
            base_verb = self.sequence.next_word(2)
        else:
            base_verb = None
    if base_verb and base_verb.word_lower in FIX_past_to_participle:
        base_verb.replace(
            FIX_past_to_participle[base_verb.word_lower])

FIX_past_to_participle = {
    # 'was': 'been', # be
    'went': 'gone',  # go
    # 'laid': 'lain', # lay
    # 'lay': 'lain', # lie
    # 'lied': 'lied', # lie
    'showed': 'shown',  # show
    'slew': 'slain',  # slay
    'have': 'had',  # have (typo)

    'began': 'begun',  # begin
    'drank': 'drunk',  # drink
    'rang': 'rung',  # ring
    'sang': 'sung',  # sing
    'sank': 'sunk',  # sink
    'sprang': 'sprung',  # spring
    'swam': 'swum',  # swim

    'arose': 'arisen',  # arise
    'drove': 'driven',  # drive
    'rode': 'ridden',  # ride
    'rose': 'risen',  # rise
    'wrote': 'written',  # write

    'broke': 'broken',  # break
    'chose': 'chosen',  # choose
    'spoke': 'spoken',  # speak
    'stole': 'stolen',  # steal
    'woke': 'woken',  # wake

    'fell': 'fallen',  # fall
    'saw': 'seen',  # see
    'see': 'seen',  # see (typo)
    'shook': 'shaken',  # shake
    'took': 'taken',  # take
    'undertook': 'undertaken',  # undertake

    'became': 'become',  # _
    'came': 'come',  # _
    'overcame': 'overcome',  # _
    'ran': 'run',  # _

    'ate': 'eaten',  # eat
    'forbade': 'forbidden',  # forbid
    'forgot': 'forgotten',  # forget
    'forgave': 'forgiven',  # forgive
    'froze': 'frozen',  # freeze
    # 'got': 'gotten', # get (have got is sometimes a false positive)
    'gave': 'given',  # give
    'hid': 'hidden',  # hide

    'bore': 'borne',  # bear
    'beat': 'beaten',  # _
    'blew': 'blown',  # blow
    'did': 'done',  # do
    'drew': 'drawn',  # draw
    'flew': 'flown',  # fly
    'grew': 'grown',  # grow
    'knew': 'known',  # know
    'tore': 'torn',  # tear
    'threw': 'thrown',  # throw
    'wore': 'worn',  # wear
    'withdrew': 'withdrawn',  # withdraw
}
