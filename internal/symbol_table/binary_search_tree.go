package symbol_table

import (
	"errors"
	"algo/internal/queue"
)

type TreeNode struct {
	key   string
	val   int
	left  *TreeNode
	right *TreeNode
	size  int
}

func newTreeNode(key string, val int, size int) *TreeNode {
	return &TreeNode{
		key:   key,
		val:   val,
		left:  nil,
		right: nil,
		size:  size,
	}
}

type BinarySearchTree struct {
	root *TreeNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{root: nil}
}

func (st BinarySearchTree) Size() int {
	return size(st.root)
}

func size(x *TreeNode) int {
	if x == nil {
		return 0
	}
	return x.size
}

func (st BinarySearchTree) IsEmpty() bool {
	return st.root == nil
}

func (st BinarySearchTree) Keys() []string {
	q := queue.NewStringQueue()
	if st.root == nil {
		return q.Data()
	}
	lo := min(st.root).key
	hi := max(st.root).key
	keys(st.root, q, lo, hi)
	return q.Data()
}

func keys(x *TreeNode, q *queue.StringQueue, lo string, hi string) {
	if x == nil {
		return
	}
	if x.key > lo {
		keys(x.left, q, lo, hi)
	}
	if x.key >= lo && x.key <= hi {
		q.Enqueue(x.key)
	}
	if x.key < hi {
		keys(x.right, q, lo, hi)
	}
}

func (st BinarySearchTree) Get(key string) (int, error) {
	val, err := get(st.root, key)
	if err != nil {
		return 0, errors.New("Key not found")
	}
	return val, nil
}

func get(x *TreeNode, key string) (int, error) {
	if x == nil {
		return 0, errors.New("Null node")
	}
	if key < x.key {
		return get(x.left, key)
	}
	if key > x.key {
		return get(x.right, key)
	}
	return x.val, nil
}

func (st *BinarySearchTree) Put(key string, val int) {
	st.root = put(st.root, key, val)
}

func put(x *TreeNode, key string, val int) *TreeNode {
	if x == nil {
		return newTreeNode(key, val, 1)
	}
	if key < x.key {
		x.left = put(x.left, key, val)
	} else if key > x.key {
		x.right = put(x.right, key, val)
	} else {
		x.val = val
	}
	x.size = 1 + size(x.left) + size(x.right)
	return x
}

func (st *BinarySearchTree) Contains(key string) bool {
	_, err := st.Get(key)
	return err == nil
}

func (st *BinarySearchTree) Delete(key string) {
	st.root = delete(st.root, key)
}

func delete(x *TreeNode, key string) *TreeNode {
	if x == nil {
		return nil
	}
	if key < x.key {
		x.left = delete(x.left, key)
	} else if key > x.key {
		x.right = delete(x.right, key)
	} else {
		if x.right == nil {
			return x.left
		}
		if x.left == nil {
			return x.right
		}
		t := x
		x = min(t.right)
		x.right = deleteMin(t.right)
		x.left = t.left
	}
	x.size = 1 + size(x.left) + size(x.right)
	return x
}

func min(x *TreeNode) *TreeNode {
	if x.left == nil {
		return x
	} else {
		return min(x.left)
	}
}

func max(x *TreeNode) *TreeNode {
	if x.right == nil {
		return x
	} else {
		return max(x.right)
	}
}

func deleteMin(x *TreeNode) *TreeNode {
	if x.left == nil {
		return x.right
	}
	x.left = deleteMin(x.left)
	x.size = 1 + size(x.left) + size(x.right)
	return x
}
