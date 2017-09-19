package mergesort

import (
	"testing"
)

type Number struct {
	n int
}

func (a Number) Quantify() int {
	return a.n
}

func ToValues(ns []Number) []Value {
	values := make([]Value, len(ns))
	for i, v := range ns {
		values[i] = Value(v)
	}
	return values
}

func TestMergeSortNumber(t *testing.T) {
	tests := []struct {
		input []Number
		want  []Number
	}{
		{[]Number{{12}, {2}, {6}, {34}, {87}, {3}, {23}}, []Number{{2}, {3}, {6}, {12}, {23}, {34}, {87}}},
		{[]Number{{23}, {54}, {6}, {54}, {1}, {8}}, []Number{{1}, {6}, {8}, {23}, {54}, {54}}},
	}
	for i, test := range tests {
		got := MergeSort(ToValues(test.input))
		if len(got) != len(test.want) {
			t.Fatalf("got = %d numbers, want = %d", len(got), len(test.want))
		}
		for j := range got {
			if got[j].Quantify() != test.want[j].Quantify() {
				t.Errorf("Test %d: got[%d] = %v, want[%d] = %v", i, j, got[j].Quantify(), j, test.want[j].Quantify())
			}
		}
	}
}
