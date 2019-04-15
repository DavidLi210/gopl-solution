package exe7_13_14_15_16

import (
	"log"
	"net/http"
	"testing"
)

func TestString(t *testing.T) {
	tests := []string{
		"-1 + -x",
		"-1 - x",
		"5 / 9 * (F - 32)",
		"pow(x, 3) + pow(y, 3)",
		"sqrt(A / pi)",
	}

	for _, test := range tests {
		expr, err := Parse(test)
		if err != nil {
			t.Error(err)
			continue
		}

		expr2, err2 := Parse(expr.String())
		if err2 != nil {
			t.Error(err2)
		}

		if expr.String() != expr2.String() {
			t.Errorf("%s != %s", expr2.String(), expr.String())
		}
	}
}

func TestExe7_13(t *testing.T) {
	/*tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
		//!-Eval
		// additional tests that don't appear in the book
		{"-1 + -x", Env{"x": 1}, "-2"},
		{"-1 - x", Env{"x": 1}, "-2"},
		//!+Eval
	}
	var prevExpr string
	for _, test := range tests {
		// Print expr only when it changes.
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}*/
	http.HandleFunc("/plot", plot)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
