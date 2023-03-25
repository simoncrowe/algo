package union_find

import (
	"testing"
)

func TestUnionAndFind(t *testing.T) {
	testCases := []struct {
		name string
		impl UnionFind
	}{
		{"Quick Find", NewQuickFind(10)},
		{"Quick Union", NewQuickUnion(10)},
		{"Weighted Quick Union", NewWeightedQuickUnion(10)},
		{"Weighted Quick Union With Path Compression", NewWeightedQuickUnionPathComp(10)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.impl.Union(1, 2)
			tc.impl.Union(2, 3)
			oneRoot := tc.impl.Find(1)
			threeRoot := tc.impl.Find(3)
			if oneRoot != threeRoot {
				t.Errorf("Expected the root of 1 and 3 to be the same, got %d and %d", oneRoot, threeRoot)
			}
			fourRoot := tc.impl.Find(4)
			if oneRoot == fourRoot {
				t.Errorf("Expected the root of 1 and 4 differ, they are both %d", oneRoot)
			}

		})
	}
}
