package main

func main() {

}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	index := m + n - 1
	i, j := m-1, n-1
	for i >= 0 && j >= 0 { //逆序方法依次从后往前填
		if nums1[i] >= nums2[j] {
			nums1[index] = nums1[i]
			index--
			i--
		} else {
			nums1[index] = nums2[j]
			index--
			j--
		}
	}
	for i >= 0 {
		nums1[index] = nums1[i]
		index--
		i--
	}
	for j >= 0 {
		nums1[index] = nums2[j]
		index--
		j--
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	copyArr := make([]int, 0)
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			copyArr = append(copyArr, nums1[i])
			i++
		} else {
			copyArr = append(copyArr, nums2[j])
			j++
		}
	}
	for i < m {
		copyArr = append(copyArr, nums1[i])
		i++
	}
	for j < n {
		copyArr = append(copyArr, nums2[j])
		j++
	}
	copy(nums1, copyArr)
}
