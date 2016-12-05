package wording

import "testing"

func TestFirstCap(t *testing.T) {
	t.Parallel()

	for _, c := range []struct {
		in, want string
	}{
		{"", ""},
		{"abcd", "Abcd"},
		{"-e", "-e"},
		{"---", "---"},
		{"F", "F"},
		{"g", "G"},
		{`123-{:[[;;""]:]hi`, `123-{:[[;;""]:]hi`},
		{"j-", "J-"},
	} {
		got := firstCap(c.in)
		if got != c.want {
			t.Errorf("firstCap(%q) == %q, want %q", c.in, got, c.want)
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
