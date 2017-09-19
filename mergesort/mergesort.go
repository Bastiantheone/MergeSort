package mergesort

// Value is an interface of what is going to be sorted.
type Value interface {
	// Quantify returns the an integer associated with the value.
	Quantify() int
}

// MergeSort sorts the input based on the Compare function.
//
// The time complexity is O(n*log(n)).
func MergeSort(input []Value) []Value {
	if len(input) <= 1 {
		return input
	}
	middle := len(input) / 2
	left := MergeSort(input[:middle])
	right := MergeSort(input[middle:])
	return merge(left, right)
}

// merge merges the slices left and right and returns the resulting slice.
func merge(left, right []Value) []Value {
	result := make([]Value, len(left)+len(right))
	for i := 0; len(left) > 0 || len(right) > 0; i++ {
		if len(left) > 0 && len(right) > 0 {
			if left[0].Quantify() < right[0].Quantify() {
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
		} else if len(left) > 0 {
			result[i] = left[0]
			left = left[1:]
		} else if len(right) > 0 {
			result[i] = right[0]
			right = right[1:]
		}
	}
	return result
}
