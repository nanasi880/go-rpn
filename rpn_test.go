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
		{
			expr: "0xFFFFFFFFFFFFFFFF",
			want: "18446744073709551615",
		},
		{
			expr: "0x1fe $ -",
			want: "378",
		},
	}

	for _, tt := range tests {

		r := RPN{
			Tokens: strings.Split(tt.expr, " "),
		}
		v, err := r.Eval(func(name string) (decimal.Decimal, error) {
			return decimal.New(132, 0), nil
		})
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
