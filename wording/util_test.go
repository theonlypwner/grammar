package wording

import "testing"

func TestFirstCap(t *testing.T) {
	t.Parallel()

	for _, c := range []struct {
		in, want, want2 string
	}{
		{"", "", ""},
		{"abcd", "Abcd", "Abcd"},
		{"Already capitalized", "Already capitalized", "Already capitalized"},
		{"--Already OK--", "--Already OK--", "--Already OK--"},
		{"-e", "-e", "-E"},
		{"---", "---", "---"},
		{"F", "F", "F"},
		{"g", "G", "G"},
		{`123-{:[[;;""]:]hi`, `123-{:[[;;""]:]hi`, `123-{:[[;;""]:]Hi`},
		{"j-", "J-", "J-"},
	} {
		got := firstCap(c.in, true)
		if got != c.want {
			t.Errorf("firstCap(%q, true) == %q, want %q", c.in, got, c.want)
		}
		got = firstCap(c.in, false)
		if got != c.want2 {
			t.Errorf("firstCap(%q, false) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestEngJoin(t *testing.T) {
	t.Parallel()

	for _, c := range []struct {
		in   []string
		want string
	}{
		{[]string{}, ""},
		{[]string{"a"}, "a"},
		{[]string{"a", "b"}, "a and b"},
		{[]string{"a", "b", "c"}, "a, b and c"},
		{[]string{"a", "b", "c", "d"}, "a, b, c and d"},
	} {
		got := engJoin(c.in...)
		if got != c.want {
			t.Errorf("engJoin(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
