package sort

// selectsort 选择排序
func selectsort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minKey := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minKey] {
				minKey = j
			}
		}
		arr[i], arr[minKey] = arr[minKey], arr[i]
	}
	return arr
}
