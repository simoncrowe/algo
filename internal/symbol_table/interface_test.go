package symbol_table

import (
	"reflect"
	"testing"
)

func TestSize(t *testing.T) {
	testCases := []struct {
		name string
		impl Interface[string, int]
	}{
		{"SequentialSearch", NewSequentialSearch[string, int]()},
		{"BinarySearch", NewBinarySearch[string, int]()},
		{"BinarySearchTree", NewBinarySearchTree[string, int]()},
		{"RedBlackBST", NewRedBlackBST[string, int]()},
		{"SeparateChainingHashTable", NewSeparateChainingHashTable[string, int]()},
		{"LinearProbingHashTable", NewLinearProbingHashTable[string, int]()},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var words = []string{"able", "bake", "card", "dusk", "gold", "hero"}

			for i, word := range words {
				tc.impl.Put(word, i)
				expected_size := i + 1
				size := tc.impl.Size()
				if size != expected_size {
					t.Errorf("%s: expected size %d after adding item %d; expected %d", tc.name, expected_size, i, size)
				}
			}

		})
	}
}
func TestIsEmpty(t *testing.T) {
	testCases := []struct {
		name string
		impl Interface[string, int]
	}{
		{"SequentialSearch", NewSequentialSearch[string, int]()},
		{"BinarySearch", NewBinarySearch[string, int]()},
		{"BinarySearchTree", NewBinarySearchTree[string, int]()},
		{"RedBlackBST", NewRedBlackBST[string, int]()},
		{"SeparateChainingHashTable", NewSeparateChainingHashTable[string, int]()},
		{"LinearProbingHashTable", NewLinearProbingHashTable[string, int]()},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if !tc.impl.IsEmpty() {
				t.Errorf("%s: newly created symbol table not empty", tc.name)
			}
		})
	}
}

func TestPutAndGet(t *testing.T) {
	testCases := []struct {
		name string
		impl Interface[string, int]
	}{
		{"SequentialSearch", NewSequentialSearch[string, int]()},
		{"BinarySearch", NewBinarySearch[string, int]()},
		{"BinarySearchTree", NewBinarySearchTree[string, int]()},
		{"RedBlackBST", NewRedBlackBST[string, int]()},
		{"SeparateChainingHashTable", NewSeparateChainingHashTable[string, int]()},
		{"LinearProbingHashTable", NewLinearProbingHashTable[string, int]()},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			key, value := "volts", 9001

			tc.impl.Put(key, value)
			result, err := tc.impl.Get(key)
			if err != nil {
				t.Errorf("%s: error getting key '%s'", tc.name, key)
			}

			if result != value {
				t.Errorf("%s: got value %d for key '%s'; expected %d", tc.name, result, key, value)
			}
		})
	}
}

func TestDeleteFirst(t *testing.T) {
	testCases := []struct {
		name string
		impl Interface[string, int]
	}{
		{"SequentialSearch", NewSequentialSearch[string, int]()},
		{"BinarySearch", NewBinarySearch[string, int]()},
		{"BinarySearchTree", NewBinarySearchTree[string, int]()},
		{"RedBlackBST", NewRedBlackBST[string, int]()},
		{"SeparateChainingHashTable", NewSeparateChainingHashTable[string, int]()},
		{"LinearProbingHashTable", NewLinearProbingHashTable[string, int]()},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var wordsAfter = []string{"able", "bake", "card", "dusk", "gold", "hero"}
			key, value := "volts", 9001

			tc.impl.Put(key, value)
			for i, word := range wordsAfter {
				tc.impl.Put(word, i)
			}

			tc.impl.Delete(key)

			result, err := tc.impl.Get(key)
			if err == nil {
				t.Errorf("%s: expected error whem getting deleted key '%s', got value %d", tc.name, key, result)
			}

		})
	}
}

func TestDeleteMiddle(t *testing.T) {
	testCases := []struct {
		name string
		impl Interface[string, int]
	}{
		{"SequentialSearch", NewSequentialSearch[string, int]()},
		{"BinarySearch", NewBinarySearch[string, int]()},
		{"BinarySearchTree", NewBinarySearchTree[string, int]()},
		{"RedBlackBST", NewRedBlackBST[string, int]()},
		{"SeparateChainingHashTable", NewSeparateChainingHashTable[string, int]()},
		{"LinearProbingHashTable", NewLinearProbingHashTable[string, int]()},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			wordsBefore := []string{"balm", "dart", "echo", "fizz", "gush"}
			wordsAfter := []string{"able", "bake", "card", "dusk", "gold", "hero"}
			key, value := "volts", 9001

			for i, word := range wordsBefore {
				tc.impl.Put(word, i)
			}
			tc.impl.Put(key, value)
			for i, word := range wordsAfter {
				tc.impl.Put(word, i)
			}

			tc.impl.Delete(key)

			result, err := tc.impl.Get(key)
			if err == nil {
				t.Errorf("%s: expected error whem getting deleted key '%s', got value %d", tc.name, key, result)
			}

		})
	}
}

func TestDeleteEnd(t *testing.T) {
	testCases := []struct {
		name string
		impl Interface[string, int]
	}{
		{"SequentialSearch", NewSequentialSearch[string, int]()},
		{"BinarySearch", NewBinarySearch[string, int]()},
		{"BinarySearchTree", NewBinarySearchTree[string, int]()},
		{"RedBlackBST", NewRedBlackBST[string, int]()},
		{"SeparateChainingHashTable", NewSeparateChainingHashTable[string, int]()},
		{"LinearProbingHashTable", NewLinearProbingHashTable[string, int]()},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			wordsBefore := []string{"balm", "dart", "echo", "fizz", "gush"}
			key, value := "volts", 9001

			for i, word := range wordsBefore {
				tc.impl.Put(word, i)
			}
			tc.impl.Put(key, value)

			tc.impl.Delete(key)

			result, err := tc.impl.Get(key)
			if err == nil {
				t.Errorf("%s: expected error whem getting deleted key '%s', got value %d", tc.name, key, result)
			}

		})
	}
}

func TestContains(t *testing.T) {
	testCases := []struct {
		name string
		impl Interface[string, int]
	}{
		{"SequentialSearch", NewSequentialSearch[string, int]()},
		{"BinarySearch", NewBinarySearch[string, int]()},
		{"BinarySearchTree", NewBinarySearchTree[string, int]()},
		{"RedBlackBST", NewRedBlackBST[string, int]()},
		{"SeparateChainingHashTable", NewSeparateChainingHashTable[string, int]()},
		{"LinearProbingHashTable", NewLinearProbingHashTable[string, int]()},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			unexpectedKey := "Spanish Inquisition"
			expectedKey, value := "volts", 9001

			tc.impl.Put(expectedKey, value)

			if !tc.impl.Contains(expectedKey) {
				t.Errorf("%s: expected to contain %s", tc.name, expectedKey)
			}
			if tc.impl.Contains(unexpectedKey) {
				t.Errorf("%s: expected not to contain %s", tc.name, unexpectedKey)
			}
		})
	}
}

func TestKeys(t *testing.T) {
	testCases := []struct {
		name string
		impl Interface[string, int]
	}{
		{"SequentialSearch", NewSequentialSearch[string, int]()},
		{"BinarySearch", NewBinarySearch[string, int]()},
		{"BinarySearchTree", NewBinarySearchTree[string, int]()},
		{"RedBlackBST", NewRedBlackBST[string, int]()},
		{"SeparateChainingHashTable", NewSeparateChainingHashTable[string, int]()},
		{"LinearProbingHashTable", NewLinearProbingHashTable[string, int]()},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			words := [...]string{"pyx", "fey", "kep", "gym", "haw"}
			wordSet := map[string]bool{}
			for i, word := range words {
				wordSet[word] = true
				tc.impl.Put(word, i)
			}

			keys := tc.impl.Keys()

			keySet := map[string]bool{}
			for i := 0; i < len(keys); i++ {
				keySet[keys[i]] = true
			}
			if !reflect.DeepEqual(keySet, wordSet) {
				t.Errorf("%s: The result of Keys (%v) is not equal to the keys pushed (%v)", tc.name, keySet, wordSet)
			}

		})
	}
}
