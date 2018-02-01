package sort_algorithm

import (
	"testing"
)

func mockData() []int {
	data := []int{5, 0, 4, 2, 1, 3}
	dst := make([]int, len(data))
	copy(dst, data)
	return dst
}

func TestBubble(t *testing.T) {
	r := bubble(mockData())
	t.Log(r)
	for i := 0; i < 5; i++ {
		if r[i] != i {
			t.Fatal("排序不正确")
		}
	}
}

func TestChoose(t *testing.T) {
	r := choose(mockData())
	t.Log(r)
	for i := 0; i < 5; i++ {
		if r[i] != i {
			t.Fatal("排序不正确")
		}
	}
}

func TestInsert(t *testing.T) {
	r := insert(mockData())
	t.Log(r)
	for i := 0; i < 5; i++ {
		if r[i] != i {
			t.Fatal("排序不正确")
		}
	}
}

func TestInsert2(t *testing.T) {
	r := insert2(mockData())
	t.Log(r)
	for i := 0; i < 5; i++ {
		if r[i] != i {
			t.Fatal("排序不正确")
		}
	}
}
