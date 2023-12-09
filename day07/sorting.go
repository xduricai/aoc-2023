package day07

func qs(slice []Hand, low int, high int) {
	if low >= high {
		return
	}

	pivotIdx := partition(slice, low, high)
	qs(slice, low, pivotIdx-1)
	qs(slice, pivotIdx+1, high)
}

func partition(slice []Hand, low int, high int) int {
	pivot := slice[high]
	idx := low - 1

	for i := low; i < high; i++ {
		if slice[i].Compare(&pivot) < 0 {
			idx++
			slice[i], slice[idx] = slice[idx], slice[i]
		}
	}

	idx++
	slice[high], slice[idx] = slice[idx], slice[high]

	return idx
}

func quickSort(slice []Hand) {
	qs(slice, 0, len(slice)-1)
}
