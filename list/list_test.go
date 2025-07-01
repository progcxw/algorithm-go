package list

import (
	"reflect"
	"testing"
)

// Helper function to create a slice from a ListNode for easy comparison
func listToSlice(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Value)
		head = head.Next
	}
	return result
}

// Helper function to create a ListNode from a slice
func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Value: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Value: nums[i]}
		current = current.Next
	}
	return head
}

func TestDoubleLinkedList(t *testing.T) {
	l := NewDoubleLinkedList()

	if l.Len() != 0 {
		t.Errorf("Expected new list to have length 0, got %d", l.Len())
	}
	if l.Head() != nil || l.Tail() != nil {
		t.Errorf("Expected new list to have nil Head and Tail")
	}

	n1 := &DLLNode{key: 1, value: 10}
	n2 := &DLLNode{key: 2, value: 20}
	n3 := &DLLNode{key: 3, value: 30}

	l.InsertHead(n1)
	if l.Len() != 1 || l.Head() != n1 || l.Tail() != n1 {
		t.Errorf("InsertHead failed for the first node")
	}

	l.InsertHead(n2)
	if l.Len() != 2 || l.Head() != n2 || l.Tail() != n1 {
		t.Errorf("InsertHead failed for the second node")
	}
	if n2.Next() != n1 || n1.Prev() != n2 {
		t.Errorf("Pointers are not set correctly after InsertHead")
	}

	l.InsertHead(n3)
	if l.Len() != 3 || l.Head() != n3 || l.Tail() != n1 {
		t.Errorf("InsertHead failed for the third node")
	}

	l.Remove(n2) // Remove from middle
	if l.Len() != 2 || l.Head() != n3 || l.Tail() != n1 {
		t.Errorf("Remove from middle failed")
	}
	if n3.Next() != n1 || n1.Prev() != n3 {
		t.Errorf("Pointers are not set correctly after removing from middle")
	}

	l.Remove(n3) // Remove from head
	if l.Len() != 1 || l.Head() != n1 || l.Tail() != n1 {
		t.Errorf("Remove from head failed")
	}

	l.Remove(n1) // Remove the last node
	if l.Len() != 0 || l.Head() != nil || l.Tail() != nil {
		t.Errorf("Remove the last node failed")
	}
}

func TestLRUCache(t *testing.T) {
	cache := New(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	if cache.Get(1) != 1 {
		t.Errorf("Expected Get(1) to return 1")
	}

	cache.Put(3, 3) // This should evict key 2

	if cache.Get(2) != -1 {
		t.Errorf("Expected Get(2) to return -1 after eviction")
	}

	if cache.Get(3) != 3 {
		t.Errorf("Expected Get(3) to return 3")
	}

	cache.Put(1, 4) // Update value of key 1
	if cache.Get(1) != 4 {
		t.Errorf("Expected Get(1) to return updated value 4")
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"normal case", []int{3, 2, 1, 5, 4}, []int{1, 2, 3, 4, 5}},
		{"sorted case", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reversed case", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"duplicates", []int{3, 2, 1, 2, 3}, []int{1, 2, 2, 3, 3}},
		{"empty list", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.input)
			sortedHead := QuickSort(head)
			result := listToSlice(sortedHead)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("QuickSort failed. Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		k        int
		expected []int
	}{
		{"k=2", []int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{"k=3", []int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
		{"k=1", []int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5}},
		{"k > len", []int{1, 2, 3}, 5, []int{1, 2, 3}},
		{"k = len", []int{1, 2, 3, 4}, 4, []int{4, 3, 2, 1}},
		{"empty list", []int{}, 2, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.input)
			reversedHead := ReverseKGroup(head, tt.k)
			result := listToSlice(reversedHead)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ReverseKGroup failed. Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
