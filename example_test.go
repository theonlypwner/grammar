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
	for i := 0; i <= 6; i++ {
		b, err := ioutil.ReadFile(fmt.Sprintf("in%v.txt", i))
		if err != nil {
			fmt.Println("[ERROR] ", err)
		}
		fmt.Println(grammar.Load(string(b)))
	}
}
