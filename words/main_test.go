package main

import "testing"

func TestSolve(t *testing.T) {
	var test = []struct {
		s            string
		UniqExpected int
		RepExpected  string
	}{
		{"hola si is no no", 3, "no"},
		{"this is another for merlin test test test test", 5, "test"},
	}

	for _, tc := range test {
		u, w := solve(tc.s)

		if tc.UniqExpected != u {
			t.Errorf("Expected %v, but get %v", tc.UniqExpected, u)
		}

		if tc.RepExpected != w {
			t.Errorf("Expected %v, but get %v", tc.RepExpected, w)
		}
	}
}
