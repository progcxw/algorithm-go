package binarytree

type TreeNode struct {
	Value string
	Left  *TreeNode // 左子树
	Right *TreeNode // 右子树
}

// Travel 前、中、后序遍历示例
func Travel(root *TreeNode) {
	if root == nil {
		return
	}

	// fmt.Println(root.Value) 前序遍历
	Travel(root.Left)
	// fmt.Println(root.Value) 中序遍历
	Travel(root.Right)
	// fmt.Println(root.Value) 后序遍历
}

// PreorderTravel 非递归前序遍历
func PreorderTravel(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	result := make([]string, 0)
	stack := make([]*TreeNode, 0)
	// 向左探到底后需在右子树继续向左探底，所以需要两个root不等于nil的for循环
	for root != nil || len(stack) > 0 {
		for root != nil {
			// 前序遍历，先将value加入到result
			result = append(result, root.Value)
			stack = append(stack, root)
			root = root.Left
		}

		// 左子树探到底，移到右子树，继续向左探底
		node := stack[len(stack)-1]
		root = node.Right
		// pop栈顶
		stack = stack[:len(stack)-1]
	}

	return result
}

// InorderTravel 中序非递归遍历
func InorderTravel(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	result := make([]string, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// pop
		node := stack[len(stack)-1]
		result = append(result, node.Value)
		stack = stack[:len(stack)-1]
		root = node.Right
	}

	return result
}

// PostorderTravel 后续非递归遍历
func PostorderTravel(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	var lastVisit *TreeNode
	result := make([]string, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVisit {
			// 向右探到底或从右边底部往上回时，pop
			result = append(result, node.Value)
			lastVisit = node
			stack = stack[:len(stack)-1]
		} else {
			root = node.Right
		}
	}

	return result
}