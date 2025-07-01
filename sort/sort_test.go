package sort

import (
	"reflect"
	"testing"
)

// testCases 定义了一组通用的测试用例，用于测试各种排序算法。
var testCases = []struct {
	name     string
	input    []int
	expected []int
}{
	{"empty", []int{}, []int{}},
	{"single", []int{1}, []int{1}},
	{"sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	{"reversed", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	{"duplicates", []int{3, 1, 4, 1, 5, 9, 2, 6}, []int{1, 1, 2, 3, 4, 5, 6, 9}},
	{"random", []int{4, 2, 1, 3, 5}, []int{1, 2, 3, 4, 5}},
	{"large_random", []int{9, 1, 8, 2, 7, 3, 6, 4, 5, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
}

// copySlice 辅助函数，用于复制切片，避免原地排序修改原始输入。
func copySlice(s []int) []int {
	dup := make([]int, len(s))
	copy(dup, s)
	return dup
}

func TestBubbleSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			arr := copySlice(tc.input)
			BubbleSort(arr)
			if !reflect.DeepEqual(arr, tc.expected) {
				t.Errorf("BubbleSort(%v) = %v, want %v", tc.input, arr, tc.expected)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			arr := copySlice(tc.input)
			SelectionSort(arr)
			if !reflect.DeepEqual(arr, tc.expected) {
				t.Errorf("SelectionSort(%v) = %v, want %v", tc.input, arr, tc.expected)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			arr := copySlice(tc.input)
			InsertionSort(arr)
			if !reflect.DeepEqual(arr, tc.expected) {
				t.Errorf("InsertionSort(%v) = %v, want %v", tc.input, arr, tc.expected)
			}
		})
	}
}

func TestShellSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			arr := copySlice(tc.input)
			ShellSort(arr)
			if !reflect.DeepEqual(arr, tc.expected) {
				t.Errorf("ShellSort(%v) = %v, want %v", tc.input, arr, tc.expected)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// MergeSort 返回一个新的切片，所以不需要复制输入
			result := MergeSort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("MergeSort(%v) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			arr := copySlice(tc.input)
			QuickSort(arr)
			if !reflect.DeepEqual(arr, tc.expected) {
				t.Errorf("QuickSort(%v) = %v, want %v", tc.input, arr, tc.expected)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			arr := copySlice(tc.input)
			HeapSort(arr)
			if !reflect.DeepEqual(arr, tc.expected) {
				t.Errorf("HeapSort(%v) = %v, want %v", tc.input, arr, tc.expected)
			}
		})
	}
}
