#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""Wording generator"""

import random


def first_cap(s):
    """Capitalize the first alphabetical character in the string"""
    s = list(s)
    for i in range(len(s)):
        if s[i].isalpha():
            s[i] = s[i].upper()
            break
    return ''.join(s)


def english_join(x, conjunction='and'):
    """English-join a list of strings"""
    if len(x) <= 2:
        return (' ' + conjunction + ' ').join(x)
    # omit serial comma because Twitter only allows 140 characters
    else:
        return '%s %s %s' % (', '.join(x[:-1]), conjunction, x[-1])

# Correction reasons
why_reasons = {
    'its': (u"‘its’ belongs to ‘it’", u"‘its’ belongs to ‘it’; ‘it's’ means ‘it is’",),
    'your': (u"‘your’ doesn't mean ‘you are’; ‘you're’ does", u"‘your’ belongs to ‘you’",),
    'its_po': (u"‘it's’ doesn't belong to ‘it’; ‘its’ does", u"‘it's’ means ‘it is’",),
    'your_po': (u"‘you're’ doesn't belong to ‘you’; ‘your’ does", u"‘you're’ means ‘you are’",),
    'there_their': (u"‘there’ doesn't belong to ‘them’; ‘their’ does",),
    'whose': (u"‘whose’ belongs to ‘whom’; ‘who's’ means ‘who is’",),
    'whose_has': (u"‘whose’ belongs to ‘whom’; ‘who's’ means ‘who has’",),
    'theyre_be': (u"‘they're’ means ‘they are’, not ‘there’",),
    'their_be': (u"‘their’ belongs to ‘them’", u"‘there’ is ‘their’ item",),
    'theyre_are': (u"‘they're’ means ‘they are’, not ‘they’ or ‘there’",),
    'hear': (u"I ‘hear’ but am ‘here’",),
    'board': (u"‘bored’ is a verb; ‘board’ is a noun", u"‘board’ is a noun; ‘bored’ is a verb",),
    'than': (u"‘than’ isn't the adverb ‘then’",),
    'then': (u"‘then’ doesn't compare like ‘than’", u"‘then’ doesn't compare; ‘than’ does",),
    'of': (u"‘of’ isn't a verb", u"‘have’ is a real verb"),
    'your-are': (u"‘you’ are; ‘your’ belongs to ‘you’", u"‘you’ are rather than ‘your’ are",),
    'supposed-to': (u"‘supposed’ isn't a bare infinitive", u"‘supposed’ is really a participle",),
    'whom': (u"‘whom’ is not nominative", u"‘whom’ isn't in subjective case"),
}
# More Constants
modals_infinitive = (
    'should', 'ought to', 'could', 'can', 'meant to', 'intended to')
modals_perfect = ('should have', 'ought to have', 'could have')
said_past = ('used', 'said', 'tweeted', 'posted')  # (simple [past) perfect]
said_infinitive = ('use', 'say', 'tweet', 'post')  # without 'to'
tweet_noun = ('a tweet', 'a post', 'a status', 'a message',
              'a status update', 'an update')  # singular
MESSAGE = [
    # (clause, add ' that'),
    # confident
    ('it is the case that', False),
    ('in this case,', False),
    # in ___'s/your tweet,
    ('I am confident', True),
    ('I say', True),
    ('I note that', False),
    ('I declare that', False),
    ('I noticed', True),
    ('I discovered', True),
    ('I see', True),
    # weaker
    ('it seems', True),
    ('to me, it seems', True),
    ('it seems to me', True),
    ('it appears', True),
    ('to me, it appears', True),
    ('it appears to me', True),
    ('it seems like', False),
    ('it looks like', False),
    # weak
    ('I think', True),
    ('I believe', True),
    ('I reckon', True),
    ('I suppose', True),
    ('I suspect', True),
    ('I feel', True),
    ('it seems to be the case that', False),
    ('it appears to be the case that', False),
]


def generate(corrections, corrected, user):
    """Generate a tweet message for the purpose of correcting a user."""
    # Random wording
    second_person = 13 >= random.randrange(20)
    use_infinitive = 50 >= random.randrange(100)  # rather than perfect
    use_inflected_have = 50 >= random.randrange(100)
    inflected_have = 'have' if second_person else 'has'
    inflected_have_optional = (
        inflected_have + ' ') if use_inflected_have else ''
    message_alter = [
        # can_override, subclause, new_modals, new_verbs
        # NOTE: had = past participle, 0 = simple past
        # (False, None, modals_perfect, said_past), # (default)
        (False, None, modals_infinitive, said_infinitive),
        (True, 'it %s %s better if' %
         (random.choice([
             'could', 'might', 'would']), random.choice(['have been', 'be'])), ('had',), said_past),
        (True, 'it %s possible for' % ('is' if use_infinitive else 'was'),
         ('to' if use_infinitive else 'to have',), said_infinitive if use_infinitive else said_past),
        (False, None, ('%s%s %s and %s' % (
            inflected_have_optional,
            random.choice(['made', 'created', 'tweeted', 'posted',
                           'written' if use_inflected_have else 'wrote']),
            random.choice(
                ['an error', 'a mistake', 'a solecism', 'a typo']),
            random.choice(
                modals_infinitive if use_infinitive else modals_perfect)
        ),), said_infinitive if use_infinitive else said_past),
        (False, None, ('%s%s %s and %s' % (
            inflected_have_optional,
            random.choice(
                ['botched', 'blundered', 'messed up', 'malformed', 'screwed up',
                 'miswritten' if use_inflected_have else 'miswrote', 'mistyped']),
            random.choice(tweet_noun),
            random.choice(
                modals_infinitive if use_infinitive else modals_perfect)
        ),), said_infinitive if use_infinitive else said_past),
        # (Must?, 'it is', ('who should',), said_infinitive), # cleft
    ]
    # Build the sentence!
    message = random.choice(MESSAGE)
    clause = message[0]
    modals = modals_perfect
    verbs = said_past
    # 50% chance to use "that"
    if message[1] and 1 == random.randrange(2):  # pragma: no cover
        clause += ' that'
    # Alter it! (40%)
    if 2 >= random.randrange(5):  # pragma: no cover
        message_alter = random.choice(message_alter)
        if message_alter[1]:
            # 50% chance to bypass the first clause, if possible
            if message_alter[0] and 1 == random.randrange(2):
                clause = message_alter[1]
            else:
                clause += ' ' + message_alter[1]
        if message_alter[2]:
            modals = message_alter[2]
        if message_alter[3]:
            verbs = message_alter[3]
    else:  # pragma: no cover
        message_alter = None  # save memory
    predicate = ''
    # choose
    modals = random.choice(modals)
    if modals:  # pragma: no cover
        predicate += modals + ' '
    verbs = random.choice(verbs)
    if verbs:  # pragma: no cover
        predicate += verbs + ' '
    # jump that quote
    predicate += '%s instead.' % (
        english_join(tuple(map(u'“{0}”'.format, corrections))))
    # 2nd person instead of 3rd (65%)
    if second_person and 17 >= random.randrange(20):  # pragma: no cover
        # Invert the subject so that we address one personally (85%)
        result = '%s, %s you %s' % (user, clause, predicate)
    # No subject inversion, but we have to make the first letter uppercase
    else:  # pragma: no cover
        if second_person:
            user = 'you, %s,' % (user)
        result = '%s %s %s' % (
            clause[0].upper() + clause[1:], user, predicate)
    # Do we need to check? Any space for why?
    # at least 3 characters have to be added
    if corrected and len(result) <= 137:  # pragma: no cover
        why = []
        for w in corrected:
            why.append(random.choice(why_reasons[w]))
        why = ' %s.' % first_cap(english_join(why))
        if len(result) + len(why) <= 140:
            result += why
# Trim to 140 characters
#  if len(result) > 140:
#      result = result[:139] + u'…'
    return result


def generate_long(corrections, corrected, user):
    raise NotImplementedError
