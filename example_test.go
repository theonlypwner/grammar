package grammar_test

import (
	"victorz.ca/grammar"

	"fmt"
	"io/ioutil"
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

func ExampleMakeTweetReply() {
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

func ExampleLoad() {
	// Test 7 files
	for i := 0; i <= 6; i++ {
		f := fmt.Sprintf("in%v.txt", i)
		b, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Println("[ERROR] ", err)
			continue
		}
		fmt.Print(f, " ")
		fmt.Println(grammar.Load(string(b)))
	}
}
