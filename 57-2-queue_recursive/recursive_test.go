package recursive

import (
	"testing"
)

func TestRecursive(t *testing.T) {
	m := make(map[int]int, 5)
	m[1] = 1
	m[2] = 2
	m[3] = 6
	m[5] = 120
	m[9] = 362880
	for k, v := range m {
		if factorial(k) != v {
			t.Fail()
		}
	}
}
