package union_find

type UnionFind interface {
	Union(int, int)
	Find(int, int) bool
}
