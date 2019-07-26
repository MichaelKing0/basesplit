package basesplit

import "testing"

func TestSplit(t *testing.T) {
	got := Split("1010", 2)

	expectation := []Number{
		Number{Base: 2, AsDecimal: 8, AsBase: "1000"},
		Number{Base: 2, AsDecimal: 0, AsBase: "000"},
		Number{Base: 2, AsDecimal: 2, AsBase: "10"},
		Number{Base: 2, AsDecimal: 0, AsBase: "0"},
	}

	for i, want := range expectation {
		if got[i] != want {
			t.Errorf("Objects don't match!")
		}
	}
}
