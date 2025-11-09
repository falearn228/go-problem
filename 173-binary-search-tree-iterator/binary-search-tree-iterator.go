/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    st []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    it := BSTIterator{st: make([]*TreeNode, 0)}
    it.pushLeft(root)
    return it
}

func (this *BSTIterator) pushLeft(node *TreeNode) {
    for node != nil {
        this.st = append(this.st, node)
        node = node.Left
    }
}


func (this *BSTIterator) Next() int {
    n := this.st[len(this.st)-1]
    this.st = this.st[:len(this.st)-1]
    if n.Right != nil {
        this.pushLeft(n.Right)
    }
    return n.Val
}


func (this *BSTIterator) HasNext() bool {
    return len(this.st) > 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */