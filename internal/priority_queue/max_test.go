package priority_queue

import (
	"testing"
)

func TestMax(t *testing.T) {
	words := []string{"is", "it", "at", "up", "by", "do", "go", "my", "of", "on"}
	pq := NewMaxPriorityQueue(words)

	max := pq.Max()

	if max != "up" {
		t.Errorf("The resulting max is '%s' rather than 'up'", max)
	}
}

func TestDelMax(t *testing.T) {
	words := []string{"is", "it", "at", "up", "by", "do", "go", "my", "of", "on"}
	pq := NewMaxPriorityQueue(words)

	first_max := pq.DelMax()

	if first_max != "up" {
		t.Errorf("The first max is '%s' rather than 'up'", first_max)
	}

	second_max := pq.DelMax()

	if second_max != "on" {
		t.Errorf("The second max is '%s' rather than 'on'", second_max)
	}
}

func TestInsert(t *testing.T) {
	words := []string{"is", "it", "at", "up", "by", "do", "go", "my", "of", "on"}
	pq := NewMaxPriorityQueue(words)
	inserted := "zoo"

	pq.Insert(inserted)
	max := pq.Max()

	if max != inserted {
		t.Errorf("The resulting max is '%s' rather than '%s'", max, inserted)
	}
}
