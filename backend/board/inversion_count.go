package board

// Merge function that counts inversions
func invCountMerge(arr []int, temp []int, left, mid, right int) int {
	i := left    // Starting index for left subarray
	j := mid + 1 // Starting index for right subarray
	k := left    // Starting index to be sorted
	inv_count := 0

	// Conditions are checked to ensure that i doesn't exceed mid and j doesn't exceed right
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			inv_count += (mid - i + 1) // All remaining elements in left subarray are inversions
			j++
		}
		k++
	}

	// Copy the remaining elements of left subarray, if any
	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}

	// Copy the remaining elements of right subarray, if any
	for j <= right {
		temp[k] = arr[j]
		j++
		k++
	}

	// Copy the sorted subarray into Original array
	for i = left; i <= right; i++ {
		arr[i] = temp[i]
	}

	return inv_count
}

// Function to use invCountMergeSort to count inversion
func invCountMergeSort(arr []int, temp []int, left, right int) int {
	inv_count := 0
	if left < right {
		// Divide the array into two halves
		mid := (left + right) / 2

		inv_count += invCountMergeSort(arr, temp, left, mid)
		inv_count += invCountMergeSort(arr, temp, mid+1, right)

		// Conquer step: Merge the sorted arrays and count inversions
		inv_count += invCountMerge(arr, temp, left, mid, right)
	}
	return inv_count
}

func countInversions(arr []int) int {
	n := len(arr)
	temp := make([]int, n)
	return invCountMergeSort(arr, temp, 0, n-1)
}
