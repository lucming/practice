package mysort

//冒泡
func bible(arr []int) {
	lenArr := len(arr)
	for i := 0; i < lenArr; i++ {
		for j := 1; j < lenArr-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

//选择
func selectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minValueIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minValueIndex] > arr[j] {
				minValueIndex = j
			}
		}
		arr[i], arr[minValueIndex] = arr[minValueIndex], arr[i]
	}
}

//插入
//每次从未排序的数组中找一个数，插到有序的数组中，具体的点在于找到将要插入的数据在已排序数组中的位置，实现方式：
//ways1. 已排序数组从后往前，如果比将要插入的数据大，往后移一位，内层循环结束，就找到将要插入数据在有序数组中的位置
//ways2. 已排序数组从后往前，谁小谁往前走
func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i
		nextInsertVal := arr[i]
		for ; arr[j-1] > nextInsertVal; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
		arr[j] = nextInsertVal
	}
}

func insertSort1(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i
		nextInsertVal := arr[i]
		for ; j > 0; j-- {
			if arr[j-1] > nextInsertVal {
				arr[i], arr[j-1] = arr[j-1], arr[i]
			}
		}
	}
}

//希尔排序，带增量的插入排序
func shellSort(arr []int) {
	lenArr := len(arr)
	for step := lenArr / 2; step >= 1; step /= 2 {
		for i := step; i < lenArr; i += step {
			for j := i - step; j >= 0; j -= step {
				if arr[j+step] < arr[j] {
					arr[j], arr[j+step] = arr[j+step], arr[j]
					continue
				}
				break
			}
		}
	}
}

func merge(arr1, arr2 []int) []int {
	ret := []int{}
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			ret = append(ret, arr1[i])
			i++
		} else {
			ret = append(ret, arr2[j])
			j++
		}
	}
	for i < len(arr1) {
		ret = append(ret, arr1[i])
		i++
	}
	for j < len(arr2) {
		ret = append(ret, arr2[j])
		j++
	}

	return ret
}

func meregeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	midIdx := len(arr) / 2
	arr1 := meregeSort(arr[:midIdx])
	arr2 := meregeSort(arr[midIdx:])

	return merge(arr1, arr2)
}

func quickSort(arr []int, start int, end int) {
	if start < end {
		i, j := start, end
		tmp := arr[i]

		for i < j {
			for i < j && tmp < arr[j] {
				j--
			}
			arr[i] = arr[j]

			for i < j && tmp >= arr[i] {
				i++
			}
			arr[j] = arr[i]
		}
		arr[i] = tmp
		quickSort(arr, start, i-1)
		quickSort(arr, i+1, end)
	}
}
