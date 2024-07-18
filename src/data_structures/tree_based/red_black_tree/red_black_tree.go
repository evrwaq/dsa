package data_structures

import (
	"errors"
)

type Color bool

const (
	Red   Color = false
	Black Color = true
)

type RBTreeNode struct {
	Value       int
	Color       Color
	Left, Right *RBTreeNode
	Parent      *RBTreeNode
}

type RBTree struct {
	Root *RBTreeNode
}

func NewRBTree() *RBTree {
	return &RBTree{}
}

func (t *RBTree) IsEmpty() bool {
	return t.Root == nil
}

func (t *RBTree) Insert(value int) {
	newNode := &RBTreeNode{Value: value, Color: Red}
	if t.Root == nil {
		newNode.Color = Black
		t.Root = newNode
	} else {
		insertRB(t.Root, newNode)
		fixInsert(t, newNode)
	}
}

func insertRB(root, newNode *RBTreeNode) {
	if newNode.Value < root.Value {
		if root.Left == nil {
			root.Left = newNode
			newNode.Parent = root
		} else {
			insertRB(root.Left, newNode)
		}
	} else {
		if root.Right == nil {
			root.Right = newNode
			newNode.Parent = root
		} else {
			insertRB(root.Right, newNode)
		}
	}
}

func (t *RBTree) Search(value int) (bool, error) {
	return searchRB(t.Root, value)
}

func searchRB(node *RBTreeNode, value int) (bool, error) {
	if node == nil {
		return false, nil
	}
	if value == node.Value {
		return true, nil
	} else if value < node.Value {
		return searchRB(node.Left, value)
	} else {
		return searchRB(node.Right, value)
	}
}

func (t *RBTree) Remove(value int) error {
	node := t.Root
	for node != nil {
		if value == node.Value {
			t.Root = deleteRB(t, node)
			return nil
		} else if value < node.Value {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return errors.New("value not found")
}

func deleteRB(t *RBTree, node *RBTreeNode) *RBTreeNode {
	var child, parent *RBTreeNode
	var color Color

	if node.Left != nil && node.Right != nil {
		replacement := node.Right
		for replacement.Left != nil {
			replacement = replacement.Left
		}
		if node.Parent != nil {
			if node.Parent.Left == node {
				node.Parent.Left = replacement
			} else {
				node.Parent.Right = replacement
			}
		} else {
			t.Root = replacement
		}

		child = replacement.Right
		parent = replacement.Parent
		color = replacement.Color

		if parent == node {
			parent = replacement
		} else {
			if child != nil {
				child.Parent = parent
			}
			parent.Left = child
			replacement.Right = node.Right
			node.Right.Parent = replacement
		}

		replacement.Parent = node.Parent
		replacement.Color = node.Color
		replacement.Left = node.Left
		node.Left.Parent = replacement

		if color == Black {
			fixDelete(t, child, parent)
		}

		return t.Root
	}

	if node.Left != nil {
		child = node.Left
	} else {
		child = node.Right
	}

	parent = node.Parent
	color = node.Color

	if child != nil {
		child.Parent = parent
	}

	if parent != nil {
		if parent.Left == node {
			parent.Left = child
		} else {
			parent.Right = child
		}
	} else {
		t.Root = child
	}

	if color == Black {
		fixDelete(t, child, parent)
	}

	return t.Root
}

func fixDelete(t *RBTree, node, parent *RBTreeNode) {
	for (node == nil || node.Color == Black) && node != t.Root {
		if parent.Left == node {
			sibling := parent.Right
			if sibling.Color == Red {
				sibling.Color = Black
				parent.Color = Red
				rotateLeft(t, parent)
				sibling = parent.Right
			}

			if (sibling.Left == nil || sibling.Left.Color == Black) &&
				(sibling.Right == nil || sibling.Right.Color == Black) {
				sibling.Color = Red
				node = parent
				parent = node.Parent
			} else {
				if sibling.Right == nil || sibling.Right.Color == Black {
					sibling.Left.Color = Black
					sibling.Color = Red
					rotateRight(t, sibling)
					sibling = parent.Right
				}

				sibling.Color = parent.Color
				parent.Color = Black
				if sibling.Right != nil {
					sibling.Right.Color = Black
				}
				rotateLeft(t, parent)
				node = t.Root
				break
			}
		} else {
			sibling := parent.Left
			if sibling.Color == Red {
				sibling.Color = Black
				parent.Color = Red
				rotateRight(t, parent)
				sibling = parent.Left
			}

			if (sibling.Left == nil || sibling.Left.Color == Black) &&
				(sibling.Right == nil || sibling.Right.Color == Black) {
				sibling.Color = Red
				node = parent
				parent = node.Parent
			} else {
				if sibling.Left == nil || sibling.Left.Color == Black {
					sibling.Right.Color = Black
					sibling.Color = Red
					rotateLeft(t, sibling)
					sibling = parent.Left
				}

				sibling.Color = parent.Color
				parent.Color = Black
				if sibling.Left != nil {
					sibling.Left.Color = Black
				}
				rotateRight(t, parent)
				node = t.Root
				break
			}
		}
	}

	if node != nil {
		node.Color = Black
	}
}

func fixInsert(t *RBTree, node *RBTreeNode) {
	for node != t.Root && node.Parent.Color == Red {
		if node.Parent == node.Parent.Parent.Left {
			uncle := node.Parent.Parent.Right
			if uncle != nil && uncle.Color == Red {
				node.Parent.Color = Black
				uncle.Color = Black
				node.Parent.Parent.Color = Red
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Right {
					node = node.Parent
					rotateLeft(t, node)
				}
				node.Parent.Color = Black
				node.Parent.Parent.Color = Red
				rotateRight(t, node.Parent.Parent)
			}
		} else {
			uncle := node.Parent.Parent.Left
			if uncle != nil && uncle.Color == Red {
				node.Parent.Color = Black
				uncle.Color = Black
				node.Parent.Parent.Color = Red
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Left {
					node = node.Parent
					rotateRight(t, node)
				}
				node.Parent.Color = Black
				node.Parent.Parent.Color = Red
				rotateLeft(t, node.Parent.Parent)
			}
		}
	}
	t.Root.Color = Black
}

func rotateLeft(t *RBTree, node *RBTreeNode) {
	right := node.Right
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Parent = node.Parent
	if node.Parent == nil {
		t.Root = right
	} else if node == node.Parent.Left {
		node.Parent.Left = right
	} else {
		node.Parent.Right = right
	}
	right.Left = node
	node.Parent = right
}

func rotateRight(t *RBTree, node *RBTreeNode) {
	left := node.Left
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Parent = node.Parent
	if node.Parent == nil {
		t.Root = left
	} else if node == node.Parent.Right {
		node.Parent.Right = left
	} else {
		node.Parent.Left = left
	}
	left.Right = node
	node.Parent = left
}
