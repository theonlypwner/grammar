#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""Functions for the parser, mainly to transform the string"""

from .Units import Space
from .Units import Word

SPACECHARS = r' ,;:()[]{}.!?/'  # no '-' dashes; they form compound words


def partitionize(text):
    """Yield a list consisting of [Word, Space] * N for a given string."""
    def partitionize_(text):
        word_start = 0
        space_start = None
        end = len(text)
        for i in range(0, end + 1):  # i: current character
            if i != end and (text[i] in SPACECHARS):
                # space
                if space_start is None:
                    space_start = i
                    if word_start is None:  # when i == 0
                        word_start = i  # preceding word will be blanked
            else:
                if i == end and space_start is None:
                    space_start = i
                if space_start is not None:
                    # start of new word -- process last block first
                    yield Word(text[word_start:space_start])
                    yield Space(text[space_start:i])
                    word_start = i
                    space_start = None
    return list(partitionize_(text))

import re

# Capture any two words within the smallest possible quotation marks
REGEX_QUOTATION = re.compile(u'["“].+? .+?["”]', re.DOTALL)
REGEX_QUOTATION_REPL = u"…"

# Specific to ask.fm


def REGEX_QUOTATION_ASKFM_REPL(mo):
    return mo.group(1)
REGEX_QUOTATION_ASKFM = re.compile(u'[^—]+— (.*) https?://.*')


def REGEX_FIX_ALLCAPS_REPL(mo):
    return mo.group(0).capitalize()
# Match continuous sentences
REGEX_FIX_ALLCAPS = re.compile(r'[^ .!?"][^.!?"]*', re.DOTALL)


def REGEX_FIXI_REPL(mo):
    return "I" + mo.group(1)  # lambda mo: "I" + mo.group(1).lower(),
# "[I]%s", word must follow if %s is empty
# DO NOT USE re.I or the expression will waste time
REGEX_FIXI = re.compile(
    r"(?:(?<=^)|(?<=[ ,;.]))i(|'(?:[dD]|[lL][lL])(?:'[vV][eE])?|'[mM]|'[vV][eE])(?=$|[ ,;.])")


REGEX_FIXNL = re.compile(r" *[\r\n]+ *")
REGEX_FIXNL_REPL = ' / '


REGEX_LINKS = re.compile(r"https?://[\w.-]+(?:/[\w.-]*(?:\?.*)?)?")
REGEX_LINKS_REPL = u"…"


def transform(text, do_decode_html=False,
              do_ellipsis=False,
              do_quotations=False,
              do_askfm=False,
              do_links=False,
              do_fixcaps=False,
              do_fixi=False,
              do_fixnewline=False,
              **kwargs):
    """Perform transformations on some text"""
    # Decode &[...]; -> [char]
    if do_decode_html:
        text = text.replace('&#39;', "'")
        text = text.replace('&quot;', '"')
        text = text.replace('&gt;', '>')
        text = text.replace('&lt;', '<')
        text = text.replace('&amp;', '&')
    # Compress "..." -> '…'
    if do_ellipsis:
        text = text.replace('...', u'…')
    # Remove quotations
    if do_quotations:
        text = REGEX_QUOTATION.sub(REGEX_QUOTATION_REPL, text)
    # Remove ask.fm question quotations
    if do_askfm:
        text = REGEX_QUOTATION_ASKFM.sub(REGEX_QUOTATION_ASKFM_REPL, text)
    # Trim links
    if do_links:
        text = REGEX_LINKS.sub(REGEX_LINKS_REPL, text)
    # Detect content entirely written in ALLCAPS or in Title Case Text
    if do_fixcaps:
        words_total = 0
        words_upper = 0
        words_title = 0
        for word in text.split():
            if word:
                # people tend to use auto-@reply, which leaves
                # @<lowercase_name>, and C&P'd links
                if (word[0] == '@' or
                        word.startswith('http://') or
                        word.startswith('https://')):
                    continue
                words_total += 1
                if word.isupper():
                    words_upper += 1
                if word.istitle():
                    words_title += 1
        # 65% all-caps or 80% title-case
        if (words_upper >= words_total * 13 / 20 or
                words_title >= words_total * 4 / 5):
            # Now, only the first letter of every sentence is now capitalized.
            text = REGEX_FIX_ALLCAPS.sub(REGEX_FIX_ALLCAPS_REPL, text)
            do_fixi = True
    # Fix i* -> I*
    if do_fixi:
        text = REGEX_FIXI.sub(REGEX_FIXI_REPL, text)
    # fix new-lines
    if do_fixnewline:
        text = REGEX_FIXNL.sub(REGEX_FIXNL_REPL, text)
    return text
