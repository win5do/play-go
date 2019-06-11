package sortAlgorithm

import (
	"math/rand"
	"testing"
	"time"
)

const (
	COUNT = 10 // 数据长度
)

// 生成测试数据
func mockData() []int {
	arr := make([]int, COUNT)
	for i := 0; i < COUNT; i++ {
		arr[i] = i
	}

	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, i := range arr {
		x := rd.Intn(COUNT)
		if x != i {
			arr[i], arr[x] = arr[x], arr[i]
		}
	}
	return arr
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

// 生产随机数据函数
func TestMock(t *testing.T) {
	for i := 0; i < COUNT; i++ {
		r := mockData()
		t.Log(r)
	}
}

// 冒泡排序
func TestBubbleSort(t *testing.T) {
	r := bubbleSort(mockData())
	t.Log(r)
	check(t, r)
}

// 选择排序
func TestChooseSort(t *testing.T) {
	r := chooseSort(mockData())
	t.Log(r)
	check(t, r)
}

// 插入排序
func TestInsertSort(t *testing.T) {
	r := insertSort(mockData())
	t.Log(r)
	check(t, r)
}

// 快速排序
func TestQuickSort(t *testing.T) {
	r := quickSort(mockData())
	t.Log(r)
	check(t, r)
}

func TestQuickSort2(t *testing.T) {
	r := quickSort2(mockData())
	t.Log(r)
	check(t, r)
}

// 希尔排序
func TestShellSort(t *testing.T) {
	r := shellSort(mockData())
	t.Log(r)
	check(t, r)
}

// 归并排序
func TestMergeSort(t *testing.T) {
	r := mergeSort(mockData())
	t.Log(r)
	check(t, r)
}

func TestMergeSort2(t *testing.T) {
	r := mergeSort2(mockData())
	t.Log(r)
	check(t, r)
}

// 堆排序
func TestHeapSort(t *testing.T) {
	r := heapSort(mockData())
	t.Log(r)
	check(t, r)
}
