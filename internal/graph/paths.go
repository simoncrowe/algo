package graph

type IPaths interface {
	HasPathTo(int) bool
	PathTo(int) []int
}
