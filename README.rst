========
grammar
========

.. image:: https://travis-ci.org/theonlypwner/grammar.svg?branch=master
	:target: https://travis-ci.org/theonlypwner/grammar
	:alt: Build status

.. image:: https://coveralls.io/repos/theonlypwner/grammar/badge.svg?branch=master
	:target: https://coveralls.io/r/theonlypwner/grammar?branch=master
	:alt: Coverage Status

.. image:: https://www.codacy.com/project/badge/1a50eeb1918e4d76912d05bf468c93e0
	:target: https://www.codacy.com/app/victor/grammar
	:alt: Codacy Badge

@_grammar_ might correct your grammar on Twitter!

This is a Python parser that corrects some common grammar errors.

============
License
============

The code is licensed under a **modified** version of the AGPL. See LICENSE.txt and agpl-3.0.txt for more details.

============
Usage
============

.. code-block:: python

	# Create parser
	parser = grammar.CorrectionManager()

	# For each text sample,
	if self.parser.load_text(tweet['text'], **options):
		tweet = self.parser.generate_wording('@' + tweet['username'])
		# ... publish the tweet
	else:
		# no errors detected
		pass
