package rpn

import (
	"strings"
	"testing"

	"github.com/shopspring/decimal"
)

func TestEval(t *testing.T) {

	tests := []struct {
		expr string
		want string
	}{
		{
			expr: "1 1 +",
			want: "2",
		},
		{
			expr: "10 20 30 * +",
			want: "610",
		},
	}

	for _, tt := range tests {

		r := RPN{
			Tokens: strings.Split(tt.expr, " "),
		}
		v, err := r.Eval(nil)
		if err != nil {
			t.Fatal(err, " ", tt.expr)
		}

		want, err := decimal.NewFromString(tt.want)
		if err != nil {
			t.Fatal(err)
		}

		if !v.Equal(want) {
			t.Fatal(v.String())
		}
	}
}
