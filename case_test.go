package deltacase_test

import (
	"testing"

	"github.com/crhntr/deltacase"
)

var (
	christopherHunter = []string{
		"christopherRobertHunter",
		"ChristopherRobertHunter",
		"christopher-robert-hunter",
		"christopher_robert_hunter",
		"christopher robert hunter",
		"CHRISTOPHER_ROBERT_HUNTER",
	}

	hello = []string{
		"hello",
		"Hello",
		"hello",
		"hello",
		"hello",
		"HELLO",
	}
)

func Test(t *testing.T) {

	fns := []func(string) string{
		deltacase.Camel,
		deltacase.Pascal,
		deltacase.Kebab,
		deltacase.Snake,
		deltacase.Spaced,
		deltacase.UpperSnake,
	}

	for i, fn := range fns {
		for _, name := range christopherHunter {
			if fn(name) != christopherHunter[i] {
				t.Fail()
			}
		}
	}
	for i, fn := range fns {
		for j, greeting := range hello {
			if val := fn(greeting); val != hello[i] {
				t.Errorf("%d %d fn(%q) != %q", i, j, val, hello[i])
			}
		}
	}
}
