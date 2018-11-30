package sort

// mergeSort 归并排序，分治法，递归
func mergeSort(arr []int) []int {
	switch {
	case len(arr) <= 1:
		return arr
	case len(arr) == 2:
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}
		return arr
	case len(arr) > 2:
		middle := len(arr) / 2
		return mergeArr(mergeSort(arr[:middle]), mergeSort(arr[middle:]))
	}
	return arr
}

// mergeArr 合并两个有序数据，分别从a,b中查找最小的元素，放入新的数组中，直至放完
func mergeArr(a, b []int) (arr []int) {
	i, j := 0, 0
	for {
		if i < len(a) && j < len(b) {
			if a[i] < b[j] {
				arr = append(arr, a[i])
				i++
			} else {
				arr = append(arr, b[j])
				j++
			}
		} else if i < len(a) {
			arr = append(arr, a[i])
			i++
		} else if j < len(b) {
			arr = append(arr, b[j])
			j++
		} else {
			break
		}
	}
	return arr
}
