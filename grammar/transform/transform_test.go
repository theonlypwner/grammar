package transform

import "testing"

func testTransform(t *testing.T, transform func(string) string, tests ...string) {
	//t.Parallel()

	n := len(tests)
	if n&1 != 0 {
		t.Errorf("Invalid len(tests) == %v", n)
	}

	for i := 0; i != len(tests); i += 2 {
		in := tests[i]
		want := tests[i+1]
		got := transform(in)
		if got != want {
			t.Errorf("T(%q) == %q, want %q", in, got, want)
		}
	}
}

func TestDecodeHTML(t *testing.T) {
	testTransform(t, DecodeHTML,
		"&#39;&quot;&gt;&lt;&amp;", `'"><&`,
	)
}

func TestEllipsis(t *testing.T) {
	testTransform(t, Ellipsis,
		"...", "…",
	)
}

func TestQuotation(t *testing.T) {
	testTransform(t, Quotation,
		`He said, "hi everybody!" They heard.`,
		"He said, … They heard.",
	)
}
func TestASKfm(t *testing.T) {
	testTransform(t, ASKfm,
		"Who are you? — I am he. https://askfm/blah",
		"I am he.",
	)
}

func TestLinks(t *testing.T) {
	testTransform(t, Links,
		"http://example.invalid-domain", "…",

		"Visit http://example.invalid-domain when you can",
		"Visit … when you can",
	)
}
func TestFixCaps(t *testing.T) {
	testTransform(t, func(s string) string {
		if ShouldFixCaps(s) {
			s = FixCaps(s)
		}
		return s
	},
		"CONTENT MUST NOT BE WRITTEN ENTIRELY IN CAPITALS",
		"Content must not be written entirely in capitals",

		"ALLCAPS @lower http://is/ignored !THIS PART IS WRITTEN ALL IN ALLCAPS",
		"Allcaps @lower http://is/ignored !This part is written all in allcaps",

		"Titlecase Is Annoying Too For Some People",
		"Titlecase is annoying too for some people",
	)
}
func TestFixI(t *testing.T) {
	testTransform(t, FixI,
		"i don't know how to capitalize i",
		"I don't know how to capitalize I",

		"i'd write a test case.",
		"I'd write a test case.",

		"i'll write a test case.",
		"I'll write a test case.",

		"i'd've write a test case.",
		"I'd've write a test case.",

		"i'll've write a test case.",
		"I'll've write a test case.",

		"i'm writing a test case.",
		"I'm writing a test case.",

		"i've written a test case.",
		"I've written a test case.",
	)
}
func TestNewLines(t *testing.T) {
	testTransform(t, NewLines,
		"1\n2 \n 3",
		"1 / 2 / 3",
	)
}
func TestTransformAll(t *testing.T) {
	testTransform(t, DoAll,
		"Visit http://example.invalid-domain\n or not \"this is a quotation\"",
		"Visit … / or not …",
	)
}
