package sort

// insert 插入排序
func insert(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := i; j >= 0; j-- {
			// fmt.Println(arr[i], arr[j], arr[i+1])
			if (j > 0 && arr[i+1] <= arr[j] && arr[i+1] >= arr[j-1]) || (j == 0 && arr[i+1] <= arr[j]) {
				tmp := arr[i+1]
				for l := i; l >= j; l-- {
					arr[l+1] = arr[l]
				}
				arr[j] = tmp
				// fmt.Println(arr)
			}
			// fmt.Println(arr)
		}
	}
	return arr
}
