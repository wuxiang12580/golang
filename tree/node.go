package main

import "fmt"

type TreeNode struct {
	value int
	left,right *TreeNode
}

func (node TreeNode)print()  {
	fmt.Println(node.value)
}
func (node *TreeNode)setValue(value int)  {
	node.value = value
}
func createNodeTree(value int) *TreeNode {
	return &TreeNode{value:value}
}
func main()  {
	root := TreeNode{}
	//fmt.Println(root)
	//root.left = &TreeNode{}
	//fmt.Println(root)
	//root.right = &TreeNode{5,nil,nil}
	//fmt.Println(root)
	//root.left.right = createNodeTree(2)
	root.print()
	//fmt.Println(root)
	root.setValue(33)
	//fmt.Println(root)
	root.print()

	pRoot := &root
	//pRoot.print()
	pRoot.setValue(200)
	pRoot.print()
}
