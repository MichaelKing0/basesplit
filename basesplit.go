package basesplit

import (
	"fmt"
	"strconv"
	"strings"
)

// Number struct
type Number struct {
	Base      int
	AsBase    string
	AsDecimal int64
}

// Split a number in any base
func Split(number string, base int) []Number {
	characters := []rune(number)

	sections := []Number{}

	for i, value := range characters {

		paddedNumber := fmt.Sprintf("%-*s", len(characters)-i, string(value))
		paddedNumber = strings.ReplaceAll(paddedNumber, " ", "0")

		realNumber, _ := strconv.ParseInt(paddedNumber, base, 0)

		sections = append(sections, Number{Base: base, AsBase: paddedNumber, AsDecimal: realNumber})
	}

	return sections
}
