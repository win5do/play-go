package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	l := m - 1
	r := m + n - 1
	for i := n - 1; i >= 0 && l < r; i-- {
		for l >= 0 && nums1[l] > nums2[i] {
			swap(nums1, l, r)
			r--
			l--
		}

		nums1[r] = nums2[i]
		r--
	}
}

func swap(arr []int, l, r int) {
	arr[r], arr[l] = arr[l], arr[r]
}
