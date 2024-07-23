package main

import "testing"

func TestSort(t *testing.T) {

	tests := [][][]int{
		{{3, 7, 5, 6, 9}, {1, 3}},
		{{1, 2, 3, 4, 2, 7, 9}, {2, 4}},
	}

	for _, v := range tests {
		out := sortWin(v[0])
		expectedResult := v[1]

		if out[0] != expectedResult[0] || out[1] != expectedResult[1] {
			t.Log("out", out, "expected", expectedResult)
			t.Fail()
		}
	}
}
