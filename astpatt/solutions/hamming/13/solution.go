package hamming

import (
	"fmt"
)

// Distance returns the hamming distance of given strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("strings have different length")
	}

	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff, nil
}
