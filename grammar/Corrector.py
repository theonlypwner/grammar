#!/usr/bin/env python

"""Parser and correction manager for Your_Parser"""

from . import Corrections
from . import Wording
from .Transformers import partitionize
from .Transformers import transform
from .Units import SequenceManager
from .Units import UnitToString


class CorrectionManager(object):

    """This manager takes text as input and creates a list of corrections."""

    def __init__(self):
        """Constructor for the manager, which resets the state"""
        self.reset()

    def reset(self):
        """Reset the state"""
        self.corrections = []
        self.sequence = None
        self.corrected = {}
        self.rerun = set([1])  # 0 = done, 1 = initial run, 2+ = rerun groups

    def load_text(self, text, **options):
        """Load some text into the state and return whether there are detections"""
        text = transform(text, **options)
        self.sequence = SequenceManager(partitionize(text))
        # Do checks
        self.do_checks_all()
        # Were there any corrections?
        if not self.corrected:
            return False
        # Generate texts and return
        return self.generate_texts()  # False should never occur

    def matched(self, kind):
        """Flag a type of correction"""
        self.corrected[kind] = True

    def do_checks_all(self):
        """Repeat checks until nothing is detected"""
        while self.rerun:
            current_rerun = self.rerun
            self.rerun = set()
            for cur in self.sequence.iter_words():
                self.do_checks(
                    current_rerun=current_rerun,
                    cur=cur,
                )

    SET_possessive_as_be = set(['its', 'your', 'whose'])

    def do_checks(self, current_rerun, cur):
        if 1 in current_rerun:
            # First pass
            if cur.word_lower in self.SET_possessive_as_be:
                # conflicts with 'whose'
                Corrections.possessive_as_be.do(self, cur)
            if cur.word_lower == "you're":
                Corrections.youre_noun.do(self, cur)
            elif cur.word_lower == 'own':
                Corrections.its_own.do(self, cur)
                Corrections.there_own.do(self, cur)
            # elif cur.word_lower == 'going':
            # Corrections.im_going.do(self, cur) # DISABLED
            elif cur.word_lower == 'whose':
                Corrections.whose_been.do(self, cur)
            elif cur.word_lower == "they're":
                Corrections.theyre_be.do(self, cur)
            elif cur.word_lower == "their":
                Corrections.their_modal.do(self, cur)
            elif cur.word_lower == "hear" or cur.word_lower == "board":
                Corrections.be_noun.do(self, cur)
            elif cur.word_lower == "then":
                Corrections.then.do(self, cur)
            elif cur.word_lower == "than":
                Corrections.than.do(self, cur)
            elif cur.word_lower == "of":
                Corrections.of.do(self, cur)
            elif cur.word_lower == "your":
                Corrections.your_are.do(self, cur)
            elif cur.word_lower == "supposed":
                Corrections.supposed_to.do(self, cur)
            elif cur.word_lower == "whom" or cur.word_lower == "whomever":
                Corrections.whom_be.do(self, cur)
        else:
            if 10 in current_rerun:
                # Caused by: youre_noun, its_own, there_own
                if cur.word_lower in self.SET_possessive_as_be:
                    Corrections.possessive_as_be.do(self, cur)
            if 11 in current_rerun:
                # Caused by: supposed_to
                if cur.word_lower == 'your':
                    Corrections.your_are.do(self, cur)

    def generate_texts(self):
        """Generate a list of corrections"""
        self.corrections = []
        sequence = list(self.sequence)  # convert the SequenceManager to a list
        sequence_len = len(sequence)

        current_correction = []
        current_end = None

        # Include: [corrected],common{0,2},near (on both sides)
        # Allow one gap: [included] [[word]] [included]
        # Add special punctuation: [final word] [?!]
        for i in range(0, sequence_len + 1, 2):
            if i == sequence_len or sequence[i].flags == 1:
                this_start = i
                this_end = i
                if i != sequence_len:
                    # left check
                    while this_start:  # at least one word before
                        if sequence[this_start - 1].sentenceBreaker:
                            # wall on left
                            break
                        if current_end is not None and this_start <= current_end + 4:
                            # overlap
                            break
                        this_start -= 2
                        if not sequence[this_start].is_common():
                            if not sequence[this_start].is_near():
                                this_start += 2
                            break
                        if this_start + 6 == i:
                            # reached limit (3 words)
                            break
                    # right check
                    while this_end + 2 < sequence_len:
                        # at least one word after
                        if sequence[this_end + 1].sentenceBreaker:
                            # wall on right
                            break
                        this_end += 2
                        if sequence[this_end].flags == 1:
                            # don't parse the corrected word yet
                            this_end -= 2
                            break
                        if not sequence[this_end].is_common():
                            if not sequence[this_end].is_near():
                                this_end -= 2
                            break
                        if this_end == i + 6:
                            # reached limit (3 words)
                            break
                # overlap
                if i != sequence_len and current_end is not None and this_start <= current_end + 4:
                    # merge ranges
                    current_correction.extend(
                        map(UnitToString, sequence[current_end + 1: i]))
                    if sequence[current_end].flags != 1:
                        current_correction.append('[')
                    current_correction.append(UnitToString(sequence[i]))
                    if i + 2 == sequence_len or sequence[i + 2].flags != 1:
                        current_correction.append(']')
                    if this_end > current_end:
                        current_correction.extend(
                            map(UnitToString, sequence[i + 1:this_end + 1]))
                        current_end = this_end
                else:
                    # previous range
                    if current_end is not None:
                        last_punctuation = sequence[
                            current_end + 1].get_final()[:1]
                        if last_punctuation in '?!':
                            current_correction.append(last_punctuation)
                        self.corrections.append(''.join(current_correction))
                    else:
                        current_end = this_start
                    if i != sequence_len:
                        # set to current range
                        current_correction = list(
                            map(UnitToString, sequence[this_start: i])) + ['[', UnitToString(sequence[i])]
                        if i + 2 == sequence_len or sequence[i + 2].flags != 1:
                            current_correction.append(']')
                        current_correction.extend(
                            list(map(UnitToString, sequence[i + 1:this_end + 1])))
                        current_end = this_end
        return bool(self.corrections)  # len(self.corrections) > 0

    def generate(self, user, callback):
        """ Output a random message for the corrections"""
        return callback(self.corrections, self.corrected.keys(), user)

    def generate_wording(self, user):
        """ Output a random tweet for the corrections"""
        return self.generate(user, Wording.generate)

    def generate_wording_long(self, user):
        """ Output a random post for the corrections"""
        return self.generate(user, Wording.generate_long)
