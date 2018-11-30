package sort

// quickSort 快速排序
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	p := arr[0]
	var before, after []int
	for i := 1; i < len(arr); i++ {
		if arr[i] < p {
			before = append(before, arr[i])
		} else {
			after = append(after, arr[i])
		}
	}
	return append(append(quickSort(before), p), quickSort(after)...)

}

// quick2Sort 优化版的快拍，空间复杂度O(1)，原地排序
func quick2Sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	p := arr[0]
	i := 1
	for j := 1; j < len(arr); j++ {
		if arr[j] < p {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[0], arr[i-1] = arr[i-1], arr[0]
	return append(append(quickSort(arr[:i-1]), p), quickSort(arr[i:])...)
}
