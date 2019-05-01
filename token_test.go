package rpn

import "testing"

func TestTokenToDecimal(t *testing.T) {

	_, err := tokenToDecimal("0xFFFFFFFFFFFFFFFF")
	if err != nil {
		t.Fatal(err)
	}
}
