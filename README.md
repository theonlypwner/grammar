# grammar
[![Build Status](https://travis-ci.org/theonlypwner/grammar.svg?branch=master)](https://travis-ci.org/theonlypwner/grammar)
[![Coverage Status](https://coveralls.io/repos/github/theonlypwner/grammar/badge.svg?branch=master)](https://coveralls.io/github/theonlypwner/grammar?branch=master)
[![codecov](https://codecov.io/gh/theonlypwner/grammar/branch/master/graph/badge.svg)](https://codecov.io/gh/theonlypwner/grammar)
[![GoCover](http://gocover.io/_badge/victorz.ca/grammar)](https://gocover.io/victorz.ca/grammar)
[![Go Report Card](https://goreportcard.com/badge/victorz.ca/grammar)](https://goreportcard.com/report/victorz.ca/grammar)
[![GoDoc](https://godoc.org/victorz.ca/grammar?status.svg)](https://godoc.org/victorz.ca/grammar)

@\_grammar\_ might correct your grammar on Twitter!

This is a parser that corrects some specific common grammar errors.

# License
The code is licensed under a **modified** version of the AGPL. See LICENSE.txt and agpl-3.0.txt for more details.

# Example
```go
package main

import (
	"victorz.ca/grammar"

	"fmt"
)

func runTest(s string) (undetected bool) {
	r := grammar.MakeTweetReply(s, "@")
	if r == "" {
		r = "(No errors detected!)"
		undetected = true
	}
	fmt.Println(r)
	return
}

func main() {
	for _, test := range [...]string{
		"Nothing's wrong with this sentence.",
		"But it's true that their is a problem with this sentence.",
	} {
		for i := 0; i < 20; i++ {
			if runTest(test) {
				break
			}
		}
	}
}
```
