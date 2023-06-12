package LC26

func Test() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	index := m + n - 1
	for i, j := m-1, n-1; i >= 0 || j >= 0; {
		if i < 0 && j >= 0 {
			nums1[index] = nums2[j]
			j--
		} else if i >= 0 && j < 0 {
			nums1[index] = nums2[i]
			i--
		} else {
			if nums1[i] > nums2[j] {
				nums1[index] = nums1[i]
				i--
			} else {
				nums1[index] = nums2[j]
				j--
			}
		}
		index--
	}
}
