package tokenizer

import (
	"testing"
)

func TestParse(t *testing.T) {

	const expr = `1+2-3*(4/var)`
	tokens, err := Parse(expr)
	if err != nil {
		t.Fatal(err)
	}

	want := []string{
		"1", "+", "2", "-", "3", "*", "(", "4", "/", "var", ")",
	}

	if len(tokens) != len(want) {
		t.Fatal(tokens)
	}

	for i := range tokens {
		if tokens[i] != want[i] {
			t.Fatal(tokens)
		}
	}
}

func TestParse2(t *testing.T) {

	const expr = `0x1234abcd + 1`
	tokens, err := Parse(expr)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		"0x1234abcd", "+", "1",
	}

	if len(tokens) != len(want) {
		t.Fatal(tokens)
	}

	for i := range tokens {
		if tokens[i] != want[i] {
			t.Fatal(tokens)
		}
	}
}
