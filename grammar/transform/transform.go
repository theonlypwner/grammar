// Package transform changes a string to another string, to make it more suitable for grammar checking.
package transform

import (
	"regexp"
	"strings"
	"unicode"
)

var replacerHTML = strings.NewReplacer(
	"&#39;", "'",
	"&quot;", `"`,
	"&gt;", ">",
	"&lt;", "<",
	"&amp;", "&",
)

// DecodeHTML decodes common HTML entities.
func DecodeHTML(s string) string { return replacerHTML.Replace(s) }

// Ellipsis replaces "..." -> '…'
func Ellipsis(s string) string { return strings.Replace(s, "...", "…", -1) }

var rQuotation = regexp.MustCompile(`(?m)["“].+? .+?["”]`).ReplaceAllLiteralString

// Quotation removes quotations with at least two words.
func Quotation(s string) string { return rQuotation(s, "…") }

var rASKfm = regexp.MustCompile(`[^—]+— (.*) https?://.*`).ReplaceAllString

// ASKfm removes ask.fm question quotations
func ASKfm(s string) string { return rASKfm(s, "$1") }

var rLinks = regexp.MustCompile(`(?i)https?://[\w.-]+(?:/[\w.-]*(?:\?.*)?)?`).ReplaceAllLiteralString

// Links removes HTTP and HTTPS links.
func Links(s string) string { return rLinks(s, "…") }

// ShouldFixCaps determines whether FixCaps would be useful.
func ShouldFixCaps(s string) bool {
	total, upper, title := WordStats(s)
	return upper*20 > total*13 || // 65% all-caps
		title*5 > total*4 // 80% title-case
}

// WordStats detects content entirely written in "ALL CAPS" or in "Title Case Text"
func WordStats(s string) (total, upper, title int) {
	const (
		TOK_SPACE = iota
		TOK_NOLOWER
		TOK_HASLOWER
	)
	prev := TOK_SPACE
	for _, r := range s {
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			prev = TOK_SPACE
		} else {
			switch prev {
			case TOK_SPACE:
				total++
				if !unicode.IsUpper(r) && !unicode.IsTitle(r) {
					prev = TOK_HASLOWER
				} else {
					upper++
					prev = TOK_NOLOWER
				}

			case TOK_NOLOWER:
				if !unicode.IsUpper(r) {
					upper--
					title++
					prev = TOK_HASLOWER
				}
			case TOK_HASLOWER:
			}
		}
	}
	return
}

var rFixALLCAPS = regexp.MustCompile(`(?m)[^ .!?"][^.!?"]*`).ReplaceAllStringFunc

func rFixALLCAPSRepl(s string) string {
	f := unicode.ToTitle
	return strings.Map(func(r rune) rune {
		r = f(r)
		f = unicode.ToLower
		return r
	}, s)
}

// FixCaps makes the first letter of each sentence title-case, and everything else lower-case.
func FixCaps(s string) string { return rFixALLCAPS(s, rFixALLCAPSRepl) }

var rFixI = regexp.MustCompile(`(^|[ ,;.])i((?i)|'(?:d|ll)(?:'ve)?|'m|'ve)($|[ ,;.])`).ReplaceAllString

// FixI fixes the capitalization on the word "I"
func FixI(s string) string { return rFixI(s, "${1}I$2$3") }

var rNewLines = regexp.MustCompile(" *[\r\n]+ *").ReplaceAllLiteralString

// NewLines replaces new lines with "/" virgules.
func NewLines(s string) string { return rNewLines(s, " / ") }

// DoAll performs all transformations on a string.
func DoAll(s string) string {
	s = DecodeHTML(s)
	s = Ellipsis(s)
	s = Quotation(s)
	s = ASKfm(s)
	s = Links(s)
	if ShouldFixCaps(s) {
		s = FixCaps(s)
	}
	s = FixI(s)
	s = NewLines(s)
	return s
}
