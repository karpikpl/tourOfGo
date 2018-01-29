package slices

import "errors"

// An array has a fixed size. A slice, on the other hand, is a dynamically-sized,
// flexible view into the elements of an array. In practice, slices are much more common than arrays.

func firstTwoSlice(array []int) ([]int, error) {

	if len(array) < 3 {
		return nil, errors.New("Array needs to have 3 of more elements")
	}

	// s is []int
	// A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	// a[low : high]
	// This selects a half-open range which includes the first element, but excludes the last one.
	s := array[1:3]
	return s, nil
}
