package main

// https://space.bilibili.com/206214
func getCommon(nums1, nums2 []int) int {
	i, n := 0, len(nums1)
	j, m := 0, len(nums2)
	for i < n && j < m {
		if nums1[i] == nums2[j] {
			return nums1[i]
		}
		if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return -1
}
