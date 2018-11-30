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
