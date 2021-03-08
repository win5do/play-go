package sortAlgorithm

import (
	"math/rand"
	"testing"
	"time"
)

//go:generate go test -v ./...

const (
	COUNT = 100 // 数据长度
)

// 生成测试数据
func mockData() []int {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, COUNT)
	for i := 0; i < COUNT; i++ {
		arr[i] = rd.Intn(COUNT) - 5
	}

	return arr
}

// 检查排序结果
func check(t *testing.T, r []int) {
	if len(r) != COUNT {
		t.Fatal("bug")
	}

	for i := 0; i < len(r)-1; i++ {
		if r[i] > r[i+1] {
			t.Fatal("bug")
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
	t.Run("Lomuto", func(t *testing.T) {
		arr := mockData()
		QuickSort_Lomuto(arr, 0, len(arr)-1)
		t.Log(arr)
		check(t, arr)
	})

	t.Run("Hoare", func(t *testing.T) {
		arr := mockData()
		QuickSort_Hoare(arr, 0, len(arr)-1)
		t.Log(arr)
		check(t, arr)
	})
}

// 希尔排序
func TestShellSort(t *testing.T) {
	r := shellSort(mockData())
	t.Log(r)
	check(t, r)
}

// 归并排序
func TestMergeSort(t *testing.T) {

	t.Run("recurse", func(t *testing.T) {
		r := mergeSort(mockData())
		t.Log(r)
		check(t, r)
	})

	t.Run("loop", func(t *testing.T) {
		r := mergeSortLoop(mockData())
		t.Log(r)
		check(t, r)
	})
}

// 堆排序
func TestHeapSort(t *testing.T) {
	r := heapSort(mockData())
	t.Log(r)
	check(t, r)
}
