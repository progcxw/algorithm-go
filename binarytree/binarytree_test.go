package binarytree

import (
	"reflect"
	"testing"
)

// buildTestTree 创建一个用于测试的二叉树
//      A
//     / \
//    B   C
//   /   / \
//  D   E   F
func buildTestTree() *TreeNode {
	root := &TreeNode{Value: "A"}
	root.Left = &TreeNode{Value: "B"}
	root.Right = &TreeNode{Value: "C"}
	root.Left.Left = &TreeNode{Value: "D"}
	root.Right.Left = &TreeNode{Value: "E"}
	root.Right.Right = &TreeNode{Value: "F"}
	return root
}

func TestPreorderTravel(t *testing.T) {
	root := buildTestTree()
	expected := []string{"A", "B", "D", "C", "E", "F"}
	result := PreorderTravel(root)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PreorderTravel failed. Expected %v, got %v", expected, result)
	}

	// 测试空树
	var nilRoot *TreeNode
	if PreorderTravel(nilRoot) != nil {
		t.Errorf("PreorderTravel with nil root should return nil")
	}
}

func TestInorderTravel(t *testing.T) {
	root := buildTestTree()
	expected := []string{"D", "B", "A", "E", "C", "F"}
	result := InorderTravel(root)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("InorderTravel failed. Expected %v, got %v", expected, result)
	}

	// 测试空树
	var nilRoot *TreeNode
	if InorderTravel(nilRoot) != nil {
		t.Errorf("InorderTravel with nil root should return nil")
	}
}

func TestPostorderTravel(t *testing.T) {
	root := buildTestTree()
	expected := []string{"D", "B", "E", "F", "C", "A"}
	result := PostorderTravel(root)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PostorderTravel failed. Expected %v, got %v", expected, result)
	}

	// 测试空树
	var nilRoot *TreeNode
	if PostorderTravel(nilRoot) != nil {
		t.Errorf("PostorderTravel with nil root should return nil")
	}
}
