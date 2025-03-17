package binarytree

// TreeNode 二叉树节点结构
// 包含节点值和左右子树指针
type TreeNode struct {
	Value string    // 节点存储的字符串值
	Left  *TreeNode // 左子树指针
	Right *TreeNode // 右子树指针
}

// Travel 前、中、后序遍历示例（递归实现）
// 展示了三种遍历方式的代码位置
// 参数:
//   - root: 二叉树根节点指针
func Travel(root *TreeNode) {
	if root == nil {
		return
	}

	// fmt.Println(root.Value) 前序遍历（根-左-右）
	Travel(root.Left)
	// fmt.Println(root.Value) 中序遍历（左-根-右）
	Travel(root.Right)
	// fmt.Println(root.Value) 后序遍历（左-右-根）
}

// PreorderTravel 非递归前序遍历
// 使用栈实现二叉树的前序遍历（根-左-右）
// 参数:
//   - root: 二叉树根节点指针
//
// 返回值:
//   - []string: 前序遍历结果数组
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
// 使用栈实现二叉树的中序遍历（左-根-右）
// 参数:
//   - root: 二叉树根节点指针
//
// 返回值:
//   - []string: 中序遍历结果数组
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

		// pop栈顶节点并访问
		node := stack[len(stack)-1]
		result = append(result, node.Value)
		stack = stack[:len(stack)-1]
		root = node.Right // 处理右子树
	}

	return result
}

// PostorderTravel 后续非递归遍历
// 使用栈实现二叉树的后序遍历（左-右-根）
// 参数:
//   - root: 二叉树根节点指针
//
// 返回值:
//   - []string: 后序遍历结果数组
func PostorderTravel(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	var lastVisit *TreeNode // 记录上一个访问的节点
	result := make([]string, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVisit {
			// 向右探到底或从右边底部往上回时，访问节点
			result = append(result, node.Value)
			lastVisit = node
			stack = stack[:len(stack)-1]
		} else {
			root = node.Right // 处理右子树
		}
	}

	return result
}
