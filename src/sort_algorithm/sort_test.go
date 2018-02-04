package sort_algorithm

import (
	"math/rand"
	"testing"
	"time"
)

const (
	COUNT = 6 // 数据长度
)

// 生成测试数据
func mockData() []int {
	var arr []int
	src := make([]int, COUNT)
	for i := 0; i < COUNT; i++ {
		src[i] = i
	}
	random(&src, &arr)
	return arr
}

// 乱序
func random(srcp *[]int, arrp *[]int) {
	src := *srcp
	arr := *arrp
	ls := len(src)
	if ls <= 0 {
		return
	}
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := rd.Intn(ls)
	*arrp = append(arr, src[index])
	*srcp = append(src[:index], src[index+1:]...)
	random(srcp, arrp)
}

// 检查排序结果
func check(t *testing.T, r []int) {
	if len(r) != COUNT {
		t.Fatal("排序不正确")
		return
	}

	for i := 0; i < COUNT; i++ {
		if r[i] != i {
			t.Fatal("排序不正确")
			break
		}
	}
}

// 测试生产随机数据函数
func TestMock(t *testing.T) {
	for i := 0; i < COUNT; i++ {
		r := mockData()
		t.Log(r)
	}
}

// 测试冒泡排序
func TestBubbleSort(t *testing.T) {
	r := bubbleSort(mockData())
	t.Log(r)
	check(t, r)
}

// 测试选择排序
func TestChooseSort(t *testing.T) {
	r := chooseSort(mockData())
	t.Log(r)
	check(t, r)
}

// 插入排序排序
func TestInsertSort(t *testing.T) {
	r := insertSort(mockData())
	t.Log(r)
	check(t, r)
}

func TestInsertSort2(t *testing.T) {
	r := insertSort2(mockData())
	t.Log(r)
	check(t, r)
}

// 快速排序排序
func TestIQuickSort(t *testing.T) {
	r := quickSort(mockData())
	t.Log(r)
	check(t, r)
}

func TestIQuickSort2(t *testing.T) {
	r := quickSort2(mockData())
	t.Log(r)
	check(t, r)
}
