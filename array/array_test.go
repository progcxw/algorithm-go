package array

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "Example 1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "Example 2",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "Example 3",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "No Solution",
			nums:   []int{1, 2, 3, 4},
			target: 10,
			want:   nil,
		},
		{
			name:   "Negative Numbers",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{2, 4},
		},
		{
			name:   "Zero Target",
			nums:   []int{-3, 0, 3},
			target: 0,
			want:   []int{0, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			// Sort the result to handle different orderings if any
			if got != nil && len(got) == 2 {
				if got[0] > got[1] {
					got[0], got[1] = got[1], got[0]
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			name:  "Example 1",
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.0,
		},
		{
			name:  "Example 2",
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},
		{
			name:  "Empty nums1",
			nums1: []int{},
			nums2: []int{1, 2, 3, 4, 5},
			want:  3.0,
		},
		{
			name:  "Empty nums2",
			nums1: []int{1, 2, 3, 4, 5},
			nums2: []int{},
			want:  3.0,
		},
		{
			name:  "Both Empty",
			nums1: []int{},
			nums2: []int{},
			want:  0.0, // Or handle error, depends on spec. For now, 0.0 is fine.
		},
		{
			name:  "Single element",
			nums1: []int{1},
			nums2: []int{},
			want:  1.0,
		},
		{
			name:  "Single element 2",
			nums1: []int{},
			nums2: []int{2},
			want:  2.0,
		},
		{
			name:  "Odd total length",
			nums1: []int{1, 2, 3},
			nums2: []int{4, 5},
			want:  3.0,
		},
		{
			name:  "Even total length",
			nums1: []int{1, 2, 3},
			nums2: []int{4, 5, 6},
			want:  3.5,
		},
		{
			name:  "Overlapping ranges",
			nums1: []int{1, 5, 7, 9},
			nums2: []int{2, 3, 4, 6, 8, 10},
			want:  5.5,
		},
		{
			name:  "Large numbers",
			nums1: []int{100000},
			nums2: []int{100001},
			want:  100000.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMedianSortedArrays(tt.nums1, tt.nums2)
			if got != tt.want {
				t.Errorf("FindMedianSortedArrays(%v, %v) = %f, want %f", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}

func TestFindMedianSortedArraysByBinary(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			name:  "Example 1",
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.0,
		},
		{
			name:  "Example 2",
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},
		{
			name:  "Empty nums1",
			nums1: []int{},
			nums2: []int{1, 2, 3, 4, 5},
			want:  3.0,
		},
		{
			name:  "Empty nums2",
			nums1: []int{1, 2, 3, 4, 5},
			nums2: []int{},
			want:  3.0,
		},
		{
			name:  "Both Empty",
			nums1: []int{},
			nums2: []int{},
			want:  0.0, // Or handle error, depends on spec. For now, 0.0 is fine.
		},
		{
			name:  "Single element",
			nums1: []int{1},
			nums2: []int{},
			want:  1.0,
		},
		{
			name:  "Single element 2",
			nums1: []int{},
			nums2: []int{2},
			want:  2.0,
		},
		{
			name:  "Odd total length",
			nums1: []int{1, 2, 3},
			nums2: []int{4, 5},
			want:  3.0,
		},
		{
			name:  "Even total length",
			nums1: []int{1, 2, 3},
			nums2: []int{4, 5, 6},
			want:  3.5,
		},
		{
			name:  "Overlapping ranges",
			nums1: []int{1, 5, 7, 9},
			nums2: []int{2, 3, 4, 6, 8, 10},
			want:  5.5,
		},
		{
			name:  "Large numbers",
			nums1: []int{100000},
			nums2: []int{100001},
			want:  100000.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMedianSortedArraysByBinary(tt.nums1, tt.nums2)
			if got != tt.want {
				t.Errorf("FindMedianSortedArraysByBinary(%v, %v) = %f, want %f", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Example 1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "Example 2",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "Example 3",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "No Solution",
			nums: []int{1, 2, 3, 4, 5},
			want: [][]int{},
		},
		{
			name: "Contains Duplicates",
			nums: []int{-2, 0, 0, 2, 2},
			want: [][]int{{-2, 0, 2}},
		},
		{
			name: "All Zeros",
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "Mixed Numbers",
			nums: []int{-4, -2, -2, -1, 0, 1, 2, 2, 3, 3, 4, 6},
			want: [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ThreeSum(tt.nums)
			// Sort inner slices and then the outer slice for consistent comparison
			for _, triplet := range got {
				sort.Ints(triplet)
			}
			sort.Slice(got, func(i, j int) bool {
				for k := 0; k < len(got[i]); k++ {
					if got[i][k] != got[j][k] {
						return got[i][k] < got[j][k]
					}
				}
				return false
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThreeSum(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
		{
			name: "Example 2",
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
		{
			name: "Smallest k",
			nums: []int{1, 2, 3, 4, 5},
			k:    1,
			want: 5,
		},
		{
			name: "Largest k",
			nums: []int{1, 2, 3, 4, 5},
			k:    5,
			want: 1,
		},
		{
			name: "Duplicates",
			nums: []int{3, 3, 3, 3, 3},
			k:    3,
			want: 3,
		},
		{
			name: "Negative numbers",
			nums: []int{-1, -5, -2, -8, -3},
			k:    2,
			want: -2,
		},
		{
			name: "Mixed numbers",
			nums: []int{-1, 0, 1, 2, -2},
			k:    3,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of nums because quickSelect modifies the slice in place
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			got := FindKthLargest(numsCopy, tt.k)
			if got != tt.want {
				t.Errorf("FindKthLargest(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func TestFindKthLargestByHeap(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
		{
			name: "Example 2",
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
		{
			name: "Smallest k",
			nums: []int{1, 2, 3, 4, 5},
			k:    1,
			want: 5,
		},
		{
			name: "Largest k",
			nums: []int{1, 2, 3, 4, 5},
			k:    5,
			want: 1,
		},
		{
			name: "Duplicates",
			nums: []int{3, 3, 3, 3, 3},
			k:    3,
			want: 3,
		},
		{
			name: "Negative numbers",
			nums: []int{-1, -5, -2, -8, -3},
			k:    2,
			want: -2,
		},
		{
			name: "Mixed numbers",
			nums: []int{-1, 0, 1, 2, -2},
			k:    3,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of nums because heap sort modifies the slice in place
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			got := FindKthLargestByHeap(numsCopy, tt.k)
			if got != tt.want {
				t.Errorf("FindKthLargestByHeap(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func TestFindKthLargestByMinHeap(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
		{
			name: "Example 2",
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
		{
			name: "Smallest k",
			nums: []int{1, 2, 3, 4, 5},
			k:    1,
			want: 5,
		},
		{
			name: "Largest k",
			nums: []int{1, 2, 3, 4, 5},
			k:    5,
			want: 1,
		},
		{
			name: "Duplicates",
			nums: []int{3, 3, 3, 3, 3},
			k:    3,
			want: 3,
		},
		{
			name: "Negative numbers",
			nums: []int{-1, -5, -2, -8, -3},
			k:    2,
			want: -2,
		},
		{
			name: "Mixed numbers",
			nums: []int{-1, 0, 1, 2, -2},
			k:    3,
			want: 0,
		},
		{
			name: "k out of bounds (k=0)",
			nums: []int{1, 2, 3},
			k:    0,
			want: -1,
		},
		{
			name: "k out of bounds (k > len)",
			nums: []int{1, 2, 3},
			k:    4,
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of nums because heap operations modify the slice in place
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			got := FindKthLargestByMinHeap(numsCopy, tt.k)
			if got != tt.want {
				t.Errorf("FindKthLargestByMinHeap(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func TestFindMinDifference(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "Example 1",
			arr:  []int{1, 5, 2, 8},
			want: 1, // (2-1)
		},
		{
			name: "Example 2",
			arr:  []int{10, 100, 300, 200},
			want: 90, // (100-10)
		},
		{
			name: "Negative numbers",
			arr:  []int{-5, -1, -10, -3},
			want: 2, // (-3 - -5)
		},
		{
			name: "Duplicates",
			arr:  []int{1, 2, 2, 3},
			want: 0, // (2-2)
		},
		{
			name: "Two numbers",
			arr:  []int{5, 10},
			want: 5,
		},
		{
			name: "Less than two numbers (empty)",
			arr:  []int{},
			want: 0,
		},
		{
			name: "Less than two numbers (one element)",
			arr:  []int{1},
			want: 0,
		},
		{
			name: "Large range",
			arr:  []int{1, 1000000},
			want: 999999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of arr because FindMinDifference sorts the slice in place
			arrCopy := make([]int, len(tt.arr))
			copy(arrCopy, tt.arr)
			got := FindMinDifference(arrCopy)
			if got != tt.want {
				t.Errorf("FindMinDifference(%v) = %d, want %d", tt.arr, got, tt.want)
			}
		})
	}
}
