package binarytree

// TreeNode 二叉树节点结构
// 包含节点值和左右子树指针
type TreeNode struct {
	Value string    // 节点存储的字符串值
	Left  *TreeNode // 左子树指针
	Right *TreeNode // 右子树指针
}

// Travel 二叉树递归遍历示例
// 该函数展示了二叉树的三种经典递归遍历方式：前序、中序和后序。
// 实际使用时，通常只会启用其中一种遍历逻辑。
//
// 参数:
//   - root: 二叉树的根节点指针。
func Travel(root *TreeNode) {
	if root == nil {
		return
	}

	// 前序遍历 (Preorder Traversal): 根 -> 左 -> 右
	// fmt.Println(root.Value) // 在访问左子树和右子树之前访问根节点

	Travel(root.Left) // 递归遍历左子树

	// 中序遍历 (Inorder Traversal): 左 -> 根 -> 右
	// fmt.Println(root.Value) // 在访问左子树之后、右子树之前访问根节点

	Travel(root.Right) // 递归遍历右子树

	// 后序遍历 (Postorder Traversal): 左 -> 右 -> 根
	// fmt.Println(root.Value) // 在访问左子树和右子树之后访问根节点
}

// PreorderTravel 非递归前序遍历 (根-左-右)
//
// 核心思想:
//  1. 使用一个栈来辅助遍历。
//  2. 从根节点开始，只要当前节点不为空，就执行以下操作：
//     a. 访问当前节点（将其值加入结果列表）。
//     b. 将当前节点压入栈中（为了之后能找到右子树）。
//     c. 将当前节点更新为其左子节点，继续向左深入。
//  3. 当左子树走到尽头（当前节点为空），从栈中弹出一个节点，将其更新为该节点的右子节点，重复步骤 2。
//  4. 当栈为空且当前节点也为空时，遍历结束。
//
// 参数:
//   - root: 二叉树的根节点指针。
//
// 返回值:
//   - []string: 前序遍历的结果切片。
func PreorderTravel(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	result := make([]string, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) > 0 {
		// 持续向左遍历，沿途访问节点并入栈
		for root != nil {
			result = append(result, root.Value) // 访问根节点
			stack = append(stack, root)         // 根节点入栈
			root = root.Left                    // 移至左子树
		}

		// 左子树已遍历完，从栈中弹出节点
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 转向右子树，开始新一轮的“根-左-右”遍历
		root = node.Right
	}

	return result
}

// InorderTravel 非递归中序遍历 (左-根-右)
//
// 核心思想:
// 1. 使用一个栈来辅助遍历。
// 2. 从根节点开始，只要当前节点不为空，就将其压入栈中，并继续深入其左子树。
// 3. 当左子树走到尽头（当前节点为空），从栈中弹出一个节点。
// 4. 访问这个弹出的节点（将其值加入结果列表）。
// 5. 将当前节点更新为弹出节点的右子节点，重复步骤 2。
// 6. 当栈为空且当前节点也为空时，遍历结束。
//
// 参数:
//   - root: 二叉树的根节点指针。
//
// 返回值:
//   - []string: 中序遍历的结果切片。
func InorderTravel(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	result := make([]string, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) > 0 {
		// 持续将所有左子节点压入栈中
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 左子树已遍历完，从栈中弹出节点并访问
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Value) // 访问节点

		// 转向右子树，开始新一轮的中序遍历
		root = node.Right
	}

	return result
}

// PostorderTravel 非递归后序遍历 (左-右-根)
//
// 核心思想:
// 后序遍历是最复杂的，因为它需要在访问根节点之前确保其左右子树都已被访问。
//  1. 使用一个栈和一个 `lastVisit` 指针，`lastVisit` 记录上一个被访问的节点。
//  2. 从根节点开始，将所有左子节点压入栈中，直到最左边的节点。
//  3. 查看栈顶节点 `node`：
//     a. 如果 `node` 的右子树为空，或者其右子树已经被访问过（`node.Right == lastVisit`），
//     那么可以访问 `node`。将其值加入结果列表，弹出栈，并更新 `lastVisit`。
//     b. 否则，`node` 的右子树还未被访问。将当前节点 `root` 指向 `node` 的右子树，重复步骤 2。
//  4. 当栈为空且当前节点也为空时，遍历结束。
//
// 参数:
//   - root: 二叉树的根节点指针。
//
// 返回值:
//   - []string: 后序遍历的结果切片。
func PostorderTravel(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	var lastVisit *TreeNode // 记录上一个访问的节点
	result := make([]string, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) > 0 {
		// 持续将所有左子节点压入栈中
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 查看栈顶节点，但不弹出
		node := stack[len(stack)-1]

		// 如果右子树为空或已访问过，则可以访问当前节点
		if node.Right == nil || node.Right == lastVisit {
			result = append(result, node.Value) // 访问节点
			lastVisit = node
			stack = stack[:len(stack)-1] // 弹出节点
			// root 保持为 nil，下一轮循环将继续处理栈中节点
		} else {
			// 转向右子树
			root = node.Right
		}
	}

	return result
}