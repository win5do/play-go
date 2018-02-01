package sort_algorithm

import "testing"

var (
	data = []int{5, 0, 4, 2, 1, 3}
)

func init() {

}

func TestBubble(t *testing.T) {
	r := bubble(data)
	t.Log(r)
	for i := 0; i < 5; i++ {
		if r[i] != i {
			t.Fatal("排序不正确")
		}
	}
}

func TestChoose(t *testing.T) {
	r := choose(data)
	t.Log(r)
	for i := 0; i < 5; i++ {
		if r[i] != i {
			t.Fatal("排序不正确")
		}
	}
}

func TestInsert(t *testing.T) {
	r := insert(data)
	t.Log(r)
	for i := 0; i < 5; i++ {
		if r[i] != i {
			t.Fatal("排序不正确")
		}
	}
}
