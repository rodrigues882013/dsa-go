// Package utils provides some utilities common to all data structure problems.
package utils

import "sort"

// MinFromIterable get the minimum element from a iterable at cost of O(n log (n))
func MinFromIterable(iterable... int) int {
	sort.Ints(iterable)
	return iterable[0]
}

// MaxFromIterable get the maximum element from a iterable at cost of O(n log (n))
func MaxFromIterable(iterable... int) int {
	MinFromIterable(iterable...)
	return iterable[len(iterable) - 1]
}

// Min get the minimum between two int at cost of O(1)
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max get the maximum between two int at cost of O(1)
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Fill fills an array with the provided value
func Fill(arr []int, size, value int) {
	for j := 0; j < size; j++ {
		arr[j] = value
	}
}

// Make2dArray create a 2d slice given their dimensions
func Make2dArray(m, n int) [][]int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}

	return matrix
}

