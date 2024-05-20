package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		if root.Left == nil {
			newNode := &TreeNode{Data: data, Parent: root}
			root.Left = newNode
			return newNode
		} else {
			return BTreeInsertData(root.Left, data)
		}
	} else if data > root.Data {
		if root.Right == nil {
			newNode := &TreeNode{Data: data, Parent: root}
			root.Right = newNode
			return newNode
		} else {
			return BTreeInsertData(root.Right, data)
		}
	}

	// If the data already exists in the tree, return the existing node
	return root
}
