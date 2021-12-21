package search

func binary(arr []int, data int) int {
	start, end := 0, len(arr)-1
	for start < end {
		midIdx := (start + end) / 2
		if data == arr[midIdx] {
			return midIdx
		}
		if data > arr[midIdx] {
			start = midIdx + 1
		}
		if data < arr[midIdx] {
			end = midIdx - 1
		}
	}

	return -1
}
