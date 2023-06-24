package symbol_table

import (
	"algo/internal/queue"
	"errors"
	"golang.org/x/exp/constraints"
)

const Red = true
const Black = false

type RedBlackTreeNode[K constraints.Ordered, V any] struct {
	key    K
	val    V
	left   *RedBlackTreeNode[K, V]
	right  *RedBlackTreeNode[K, V]
	size   int
	colour bool
}

func isRed[K constraints.Ordered, V any](n *RedBlackTreeNode[K, V]) bool {
	if n == nil {
		return false
	}
	return n.colour == Red
}

func rotateLeft[K constraints.Ordered, V any](n *RedBlackTreeNode[K, V]) *RedBlackTreeNode[K, V] {
	x := n.right
	n.right = x.left
	x.left = n
	x.colour = n.colour
	n.colour = Red
	x.size = n.size
	n.size = 1 + sizeRedBlack(n.left) + sizeRedBlack(n.right)
	return x
}

func rotateRight[K constraints.Ordered, V any](n *RedBlackTreeNode[K, V]) *RedBlackTreeNode[K, V] {
	x := n.left
	n.left = x.right
	x.right = n
	x.colour = n.colour
	n.colour = Red
	x.size = n.size
	n.size = 1 + sizeRedBlack(n.left) + sizeRedBlack(n.right)
	return x
}

func flipColours[K constraints.Ordered, V any](n *RedBlackTreeNode[K, V]) {
	n.colour = !n.colour
	n.left.colour = !n.left.colour
	n.right.colour = !n.right.colour
}

func moveRedLeft[K constraints.Ordered, V any](n *RedBlackTreeNode[K, V]) *RedBlackTreeNode[K, V] {
	flipColours(n)
	if isRed(n.right.left) {
		n.right = rotateRight(n.right)
		n = rotateLeft(n)
		flipColours(n)
	}
	return n
}

func moveRedRight[K constraints.Ordered, V any](n *RedBlackTreeNode[K, V]) *RedBlackTreeNode[K, V] {
	flipColours(n)
	if isRed(n.left.left) {
		n = rotateRight(n)
		flipColours(n)
	}
	return n
}

func balance[K constraints.Ordered, V any](n *RedBlackTreeNode[K, V]) *RedBlackTreeNode[K, V] {
	if isRed(n.right) && !isRed(n.left) {
		n = rotateLeft(n)
	}
	if isRed(n.left) && isRed(n.left.left) {
		n = rotateRight(n)
	}
	if isRed(n.left) && isRed(n.right) {
		flipColours(n)
	}

	n.size = 1 + sizeRedBlack(n.left) + sizeRedBlack(n.right)
	return n
}

func newRedBlackTreeNode[K constraints.Ordered, V any](key K, val V, size int, colour bool) *RedBlackTreeNode[K, V] {
	return &RedBlackTreeNode[K, V]{
		key:    key,
		val:    val,
		left:   nil,
		right:  nil,
		size:   size,
		colour: colour,
	}
}

type RedBlackBST[K constraints.Ordered, V any] struct {
	root *RedBlackTreeNode[K, V]
}

func NewRedBlackBST[K constraints.Ordered, V any]() *RedBlackBST[K, V] {
	return &RedBlackBST[K, V]{root: nil}
}

func (st RedBlackBST[K, V]) Size() int {
	return sizeRedBlack(st.root)
}

func sizeRedBlack[K constraints.Ordered, V any](x *RedBlackTreeNode[K, V]) int {
	if x == nil {
		return 0
	}
	return x.size
}

func (st RedBlackBST[K, V]) IsEmpty() bool {
	return st.root == nil
}

func (st RedBlackBST[K, V]) Keys() []K {
	q := queue.NewQueue[K]()
	if st.root == nil {
		return q.Data()
	}
	lo := minRedBlack(st.root).key
	hi := maxRedBlack(st.root).key
	keysRedBlack(st.root, q, lo, hi)
	return q.Data()
}

func keysRedBlack[K constraints.Ordered, V any](x *RedBlackTreeNode[K, V], q *queue.Queue[K], lo K, hi K) {
	if x == nil {
		return
	}
	if x.key > lo {
		keysRedBlack(x.left, q, lo, hi)
	}
	if x.key >= lo && x.key <= hi {
		q.Enqueue(x.key)
	}
	if x.key < hi {
		keysRedBlack(x.right, q, lo, hi)
	}
}

func (st RedBlackBST[K, V]) Get(key K) (V, error) {
	val, err := getRedBlack(st.root, key)
	if err != nil {
		var nothing V
		return nothing, errors.New("Key not found")
	}
	return val, nil
}

func getRedBlack[K constraints.Ordered, V any](x *RedBlackTreeNode[K, V], key K) (V, error) {
	if x == nil {
		var nothing V
		return nothing, errors.New("Null node")
	}
	if key < x.key {
		return getRedBlack(x.left, key)
	}
	if key > x.key {
		return getRedBlack(x.right, key)
	}
	return x.val, nil
}

func (st *RedBlackBST[K, V]) Put(key K, val V) {
	st.root = putRedBlack(st.root, key, val)
}

func putRedBlack[K constraints.Ordered, V any](x *RedBlackTreeNode[K, V], key K, val V) *RedBlackTreeNode[K, V] {
	if x == nil {
		return newRedBlackTreeNode(key, val, 1, Red)
	}

	if key < x.key {
		x.left = putRedBlack(x.left, key, val)
	} else if key > x.key {
		x.right = putRedBlack(x.right, key, val)
	} else {
		x.val = val
	}

	if isRed(x.right) && !isRed(x.left) {
		x = rotateLeft(x)
	}
	if isRed(x.left) && isRed(x.left.left) {
		x = rotateRight(x)
	}
	if isRed(x.left) && isRed(x.right) {
		flipColours(x)
	}

	x.size = 1 + sizeRedBlack(x.left) + sizeRedBlack(x.right)

	return x
}

func (st *RedBlackBST[K, V]) Contains(key K) bool {
	_, err := st.Get(key)
	return err == nil
}

func (st *RedBlackBST[K, V]) Delete(key K) {
	st.root = deleteRedBlack(st.root, key)

}

func deleteRedBlack[K constraints.Ordered, V any](x *RedBlackTreeNode[K, V], key K) *RedBlackTreeNode[K, V] {
	if x == nil {
		return nil
	}
	if key < x.key {
		if !isRed(x.left) && !isRed(x.left.left) {
			x = moveRedLeft(x)
		}
		x.left = deleteRedBlack(x.left, key)
	} else {
		if isRed(x.left) {
			x = rotateRight(x)
		}
		if key == x.key && x.right == nil {
			return nil
		}
		if !isRed(x.right) && !isRed(x.right.left) {
			x = moveRedRight(x)
		}
		if key == x.key {
			y := minRedBlack(x.right)
			x.key = y.key
			x.val = y.val
			x.right = deleteMinRedBlack(x.right)
		} else {
			x.right = deleteRedBlack(x.right, key)
		}
	}
	return balance(x)
}

func minRedBlack[K constraints.Ordered, V any](x *RedBlackTreeNode[K, V]) *RedBlackTreeNode[K, V] {
	if x.left == nil {
		return x
	} else {
		return minRedBlack(x.left)
	}
}

func maxRedBlack[K constraints.Ordered, V any](x *RedBlackTreeNode[K, V]) *RedBlackTreeNode[K, V] {
	if x.right == nil {
		return x
	} else {
		return maxRedBlack(x.right)
	}
}

func deleteMinRedBlack[K constraints.Ordered, V any](x *RedBlackTreeNode[K, V]) *RedBlackTreeNode[K, V] {
	if x.left == nil {
		return x.right
	}
	x.left = deleteMinRedBlack(x.left)
	x.size = 1 + sizeRedBlack(x.left) + sizeRedBlack(x.right)
	return x
}
