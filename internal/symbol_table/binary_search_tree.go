package symbol_table

import (
	"algo/internal/queue"
	"errors"
	"golang.org/x/exp/constraints"
)

type TreeNode[K constraints.Ordered, V any] struct {
	key   K
	val   V
	left  *TreeNode[K, V]
	right *TreeNode[K, V]
	size  int
}

func newTreeNode[K constraints.Ordered, V any](key K, val V, size int) *TreeNode[K, V] {
	return &TreeNode[K, V]{
		key:   key,
		val:   val,
		left:  nil,
		right: nil,
		size:  size,
	}
}

type BinarySearchTree[K constraints.Ordered, V any] struct {
	root *TreeNode[K, V]
}

func NewBinarySearchTree[K constraints.Ordered, V any]() *BinarySearchTree[K, V] {
	return &BinarySearchTree[K, V]{root: nil}
}

func (st BinarySearchTree[K, V]) Size() int {
	return size(st.root)
}

func size[K constraints.Ordered, V any](x *TreeNode[K, V]) int {
	if x == nil {
		return 0
	}
	return x.size
}

func (st BinarySearchTree[K, V]) IsEmpty() bool {
	return st.root == nil
}

func (st BinarySearchTree[K, V]) Keys() []K {
	q := queue.NewQueue[K]()
	if st.root == nil {
		return q.Data()
	}
	lo := min(st.root).key
	hi := max(st.root).key
	keys(st.root, q, lo, hi)
	return q.Data()
}

func keys[K constraints.Ordered, V any](x *TreeNode[K, V], q *queue.Queue[K], lo K, hi K) {
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

func (st BinarySearchTree[K, V]) Get(key K) (V, error) {
	val, err := get(st.root, key)
	if err != nil {
		var nothing V
		return nothing, errors.New("Key not found")
	}
	return val, nil
}

func get[K constraints.Ordered, V any](x *TreeNode[K, V], key K) (V, error) {
	if x == nil {
		var nothing V
		return nothing, errors.New("Null node")
	}
	if key < x.key {
		return get(x.left, key)
	}
	if key > x.key {
		return get(x.right, key)
	}
	return x.val, nil
}

func (st *BinarySearchTree[K, V]) Put(key K, val V) {
	st.root = put(st.root, key, val)
}

func put[K constraints.Ordered, V any](x *TreeNode[K, V], key K, val V) *TreeNode[K, V] {
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

func (st *BinarySearchTree[K, V]) Contains(key K) bool {
	_, err := st.Get(key)
	return err == nil
}

func (st *BinarySearchTree[K, V]) Delete(key K) {
	st.root = delete(st.root, key)
}

func delete[K constraints.Ordered, V any](x *TreeNode[K, V], key K) *TreeNode[K, V] {
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

func min[K constraints.Ordered, V any](x *TreeNode[K, V]) *TreeNode[K, V] {
	if x.left == nil {
		return x
	} else {
		return min(x.left)
	}
}

func max[K constraints.Ordered, V any](x *TreeNode[K, V]) *TreeNode[K, V] {
	if x.right == nil {
		return x
	} else {
		return max(x.right)
	}
}

func deleteMin[K constraints.Ordered, V any](x *TreeNode[K, V]) *TreeNode[K, V] {
	if x.left == nil {
		return x.right
	}
	x.left = deleteMin(x.left)
	x.size = 1 + size(x.left) + size(x.right)
	return x
}
