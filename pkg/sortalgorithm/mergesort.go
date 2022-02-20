package sortalgorithm

func merge(arr []int, lowerBound int, middle int, upperBound int, aux []int) {

	idx := lowerBound
	jdx := middle + 1
	kdx := lowerBound

	for idx <= middle && jdx <= upperBound {
		if arr[idx] <= arr[jdx] {
			aux[kdx] = arr[idx]
			idx++

		} else {
			aux[kdx] = arr[jdx]
			jdx++
		}
		kdx++
	}

	if idx > middle {

		for jdx <= upperBound {
			aux[kdx] = arr[jdx]
			jdx++
			kdx++
		}

	} else {
		for idx <= middle {
			aux[kdx] = arr[idx]
			idx++
			kdx++
		}
	}

	for index := lowerBound; index <= upperBound; index++ {
		arr[index] = aux[index]
	}
}

func mergeSort(arr []int, lowerBound int, upperBound int, aux []int) {

	if lowerBound < upperBound {
		middle := (lowerBound + upperBound) / 2
		mergeSort(arr, lowerBound, middle, aux)
		mergeSort(arr, middle + 1, upperBound, aux)
		merge(arr, lowerBound, middle, upperBound, aux)
	}

}

func MergeSort(arr []int) {
	aux := make([]int, len(arr))
	mergeSort(arr, 0, len(arr) - 1, aux)
}

