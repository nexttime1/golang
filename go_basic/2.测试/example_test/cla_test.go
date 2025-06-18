package cla

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type test struct {
		inputA int
		inputB int
		output int
	}

	tests := []test{
		{inputA: 1, inputB: 2, output: 3},
		{inputA: 3, inputB: 3, output: 6},
		{inputA: 1, inputB: 1, output: 4},
	}
	for _, x := range tests {
		got := Add(x.inputA, x.inputB)
		if got != x.output {
			t.Errorf("Add(%d,%d) = %d; want 1", x.inputA, x.inputB, got)
		}
	}
}

func TestSub(t *testing.T) {
	type test struct {
		inputA int
		inputB int
		output int
	}
	tests := map[string]test{
		"NO1": {inputA: 1, inputB: 1, output: 0},
		"NO2": {inputA: 3, inputB: 2, output: 12},
		"NO3": {inputA: 5, inputB: 1, output: 4},
	}
	for i, j := range tests {
		t.Run(i, func(t *testing.T) {
			got := Sub(j.inputA, j.inputB)
			if got != j.output {
				t.Errorf("Sub(%d,%d) = %d; want %d", j.inputA, j.inputB, got, j.output)
			}
		})
	}
}
