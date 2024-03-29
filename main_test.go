package main

import (
	"testing"
)

func FuzzSumInts(f *testing.F) {
    f.Add(int64(34),  int64(12))

    f.Fuzz(func(t *testing.T, i int64, j int64) {
		testscenario := map[string]int64{
			"first": i,
			"second": j,
		}

        SumInts(testscenario)
    })
}