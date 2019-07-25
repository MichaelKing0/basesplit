package basesplit

import "testing"

func TestSplit(t *testing.T) {
	got := Split("1010", 2)

	expectation := []Number{
		Number{base: 2, asDecimal: 8, asBase: "1000"},
		Number{base: 2, asDecimal: 0, asBase: "000"},
		Number{base: 2, asDecimal: 2, asBase: "10"},
		Number{base: 2, asDecimal: 0, asBase: "0"},
	}

	for i, want := range expectation {
		if got[i] != want {
			t.Errorf("Objects don't match!")
		}
	}
}
