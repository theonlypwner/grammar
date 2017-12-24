package wording

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// firstCap makes the first character or letter title-case.
// If onlyFirst is true, only the first character will be checked;
// otherwise, the string s will be searched for a letter.
func firstCap(s string, onlyFirst bool) string {
	result := []rune(s)
	for i, r := range result {
		if unicode.IsLetter(r) {
			if !unicode.IsTitle(r) {
				result[i] = unicode.ToTitle(r)
				return string(result)
			}
			break
		} else if onlyFirst {
			break
		}
	}
	return s
}

// englishJoin joins a slice of strings using an English conjunction.
func englishJoin(conjunction string, a []string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return a[0]
	case 2:
		return a[0] + " " + conjunction + " " + a[1]
	default:
		commaJoined := strings.Join(a[:len(a)-1], ", ")
		// omit serial comma because of Twitter's character limit
		return commaJoined + " " + conjunction + " " + a[len(a)-1]
	}
}

// engJoin calls englishJoin with "and"
func engJoin(a []string) string {
	return englishJoin("and", a)
}

func p(x float64) bool { return rand.Float64() < x }

func choice(s ...string) string {
	if len(s) == 0 {
		return ""
	}
	return s[rand.Intn(len(s))]
}
