package sortalgorithm

func swap(arr []int, start int, end int) {
	tmp := arr[start]
	arr[start] = arr[end]
	arr[end] = tmp
}

func partition(arr []int, lowerBound, upperBound int) int {
	pivot := arr[lowerBound]
	start := lowerBound
	end := upperBound

	for start < end {
		for start < len(arr) && arr[start] <= pivot {
			start++
		}

		for end > 0 && arr[end] > pivot {
			end--
		}

		if start < end {
			swap(arr, start, end)
		}
	}

	swap(arr, lowerBound, end)

	return end

}

func quickSort(arr []int, lowerBound int, upperBound int) {
	if lowerBound < upperBound {
		partitionDivisor := partition(arr, lowerBound, upperBound)
		quickSort(arr, lowerBound, partitionDivisor - 1)
		quickSort(arr, partitionDivisor + 1, upperBound)
	}
}

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr) - 1)
}