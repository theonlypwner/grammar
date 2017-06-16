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

// firstCap makes the first character title-case.
func firstCap(s string) string {
	result := []rune(s)
	for _, r := range result {
		if !unicode.IsTitle(r) {
			result[0] = unicode.ToTitle(r)
			return string(result)
		}
		break
	}
	return s
}

// englishJoin joins a (slice of strings) using an English conjunction.
func englishJoin(conjunction string, a ...string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return a[0]
	case 2:
		return a[0] + " " + conjunction + " " + a[1]
	default:
		commaJoined := strings.Join(a[:len(a)-1], ", ")
		// omit serial comma because Twitter only allows 140 characters
		return englishJoin(conjunction, commaJoined, a[len(a)-1])
	}
}

// engJoin calls englishJoin with "and"
func engJoin(a ...string) string {
	return englishJoin("and", a...)
}

func p(x float64) bool { return rand.Float64() < x }

func choice(s ...string) string {
	if len(s) == 0 {
		return ""
	}
	return s[rand.Intn(len(s))]
}
