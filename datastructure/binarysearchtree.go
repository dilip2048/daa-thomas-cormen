/*
Invariants:
The key in the left subtree is smaller than root's key
The key in the right subtree is greater than the root's key
*/

package main

import (
	"fmt"
)

// TreeNode is a placeholder that holds data and pointers of its children
type TreeNode struct {
	left   *TreeNode
	right  *TreeNode
	parent *TreeNode
	data   int
}

// This method will traverse in in-order fashion
func inorderTreeWalk(node *TreeNode) {
	if node != nil {
		inorderTreeWalk(node.left)
		fmt.Printf("%v ", node.data)
		inorderTreeWalk(node.right)
	}
}

// This method will traverse in pre-order fashion
func preorderTreeWalk(node *TreeNode) {
	if node != nil {
		fmt.Printf("%v ", node.data)
		inorderTreeWalk(node.left)
		inorderTreeWalk(node.right)
	}
}

// This method will traverse in post-order fashion
func postorderTreeWalk(node *TreeNode) {
	if node != nil {
		inorderTreeWalk(node.left)
		inorderTreeWalk(node.right)
		fmt.Printf("%v ", node.data)
	}
}

// This method will search for the element in the tree
func search(t *TreeNode, k int) {
	if t != nil {
		if k == t.data {
			fmt.Printf("\nFound")
			return
		} else if k < t.data {
			search(t.left, k)
		} else {
			search(t.right, k)
		}
	}
}

// This method will search for the key iteratively
func searchIterative(t *TreeNode, k int) {
	temp := t
	for temp != nil && temp.data != k {
		if k < t.data {
			temp = temp.left
		} else if k > t.data {
			temp = temp.right
		}
	}
	if temp.data == k {
		fmt.Printf("\nfound")
	} else {
		fmt.Printf("\nnot found")
	}
}

// This method will find the successor of the element
func findSuccessor(t *TreeNode) *TreeNode {
	if t.right != nil {
		return treeMinimum(t.right)
	}
	y := t.parent
	for y != nil && t == y.right {
		t = y
		y = y.parent
	}
	return y
}

// This method will find the predecessor of the element
func findPredecessor(t *TreeNode) *TreeNode {
	if t.left != nil {
		return treeMaximum(t.left)
	}
	y := t.parent
	for y != nil && t == y.left {
		t = y
		y = y.parent
	}
	return y
}

// This method will find the minimum in a subtree
func treeMinimum(t *TreeNode) *TreeNode {
	temp := t
	for temp.left != nil {
		temp = temp.left
	}
	return temp
}

// This method will find the maximum in a subtree
func treeMaximum(t *TreeNode) *TreeNode {
	temp := t
	for temp.right != nil {
		temp = temp.right
	}
	return temp
}

func main() {
	var root TreeNode
	root = TreeNode{data: 6}
	root.left = &TreeNode{data: 5}
	root.left.parent = &root
	root.left.left = &TreeNode{data: 2}
	root.left.left.parent = root.left
	root.left.right = &TreeNode{data: 5}
	root.left.right.parent = root.left
	root.right = &TreeNode{data: 7}
	root.right.right = &TreeNode{data: 8}
	root.right.right.parent = root.right
	// all types of traversals
	fmt.Println("\nThe Inorder traversal is")
	inorderTreeWalk(&root)
	fmt.Println("\nThe preorder traversal is")
	preorderTreeWalk(&root)
	fmt.Println("\nThe postorder traversal is")
	postorderTreeWalk(&root)

	// searches recursively and iteratively
	search(&root, 2)
	searchIterative(&root, 2)

	// todo: write delete method

	// find the successor of the element
	successor := findSuccessor(root.left.left)
	fmt.Printf("\nThe successor is %v\n", successor.data)

	// find the predecessor of the element
	predecessor := findPredecessor(&root)
	fmt.Printf("\nThe predecessor is %v\n", predecessor.data)

}
