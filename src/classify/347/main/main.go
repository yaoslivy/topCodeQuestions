package main

func topKFrequent(nums []int, k int) []int {
	// 使用map记录每个元素出现的次数
	mm := make(map[int]int)
	res := make([]int, 0)
	for _, val := range nums {
		//集合元素唯一
		if _, ok := mm[val]; !ok {
			res = append(res, val)
		}
		mm[val]++
	}
	if len(res) <= k {
		return res
	}
	//对每个元素按照出现次数排序
	// sort.Slice(res, func (i, j int) bool {
	//     return mm[res[i]] > mm[res[j]]
	// })

	// 使用快排找到k大元素的位置
	// quickSort(res, 0, len(res)-1, k, mm)
	// return res[:k]

	//使用堆排序思路，先建立起大小为k的小顶堆，
	//从非叶子节点开始调整前k个为小顶堆
	for i := k/2 - 1; i >= 0; i-- {
		adjustHeap(res, i, k, mm)
	}
	//当下一个元素大于堆顶元素，则弹出堆顶元素，然后将下一个元素换到堆顶，同时调整堆为小顶堆，直到最后
	for i := k; i < len(res); i++ {
		if mm[res[i]] > mm[res[0]] {
			res[0], res[i] = res[i], res[0]
			adjustHeap(res, 0, k, mm)
		}
	}
	return res[:k]
}

//调整当前堆，i为需要调整的根节点，n为堆的个数
func adjustHeap(arr []int, i, n int, mm map[int]int) {
	temp := arr[i] //当前根元素
	for k := 2*i + 1; k < n; k = 2*k + 1 {
		// 当右子树小于左子树的个数时
		if k+1 < n && mm[arr[k+1]] < mm[arr[k]] {
			k++
		}
		//如果左右子节点右比根节点小的次数则换到根节点
		if mm[temp] > mm[arr[k]] {
			arr[i] = arr[k]
			i = k //下一个需要比较的根位置
		} else { //没有则退出循环，因为后面子孙节点都是次数比根大的
			break
		}
	}
	arr[i] = temp
}

//快排，找到中间元素为第k大则返回
func quickSort(arr []int, left, right int, k int, mm map[int]int) {
	if left > right {
		return
	}
	i, j := left, right
	temp := arr[i]
	for i < j {
		//从后往前，找到一个比temp位置次数大的元素
		for i < j && mm[arr[j]] <= mm[temp] {
			j--
		}
		arr[i] = arr[j]
		//从前往后，找到一个比temp小的元素
		for i < j && mm[arr[i]] > mm[temp] {
			i++
		}
		arr[j] = arr[i]
	}
	arr[i] = temp
	//说明前k个元素都是大于后面元素的
	if i == k {
		return
	}
	//k在i左边，才需要左递归找，在i右边才需要右递归找
	if i > left && k < i {
		quickSort(arr, left, i-1, k, mm)
	}
	if j < right && k > j {
		quickSort(arr, j+1, right, k, mm)
	}
}
