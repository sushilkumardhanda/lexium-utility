package datahandler

type TreeNode struct {
	Parent   *TreeNode   `json:"-"`
	NodeData Element     `json:"nodeData"`
	Children []*TreeNode `json:"children"`
}
