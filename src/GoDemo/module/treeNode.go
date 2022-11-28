package module

type TreeNodes struct {
	Val   int
	Left  *TreeNodes
	Right *TreeNodes
}

func NewTreeNode(val int) *TreeNodes {
	return &TreeNodes{
		Val: val,
	}
}
