package sortalg

// time complexity: O(n^2) worst/average, O(n) best
// space complexity: O(1)

func BubbleSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
