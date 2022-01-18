package main

func peakIndexInMountainArray(arr []int) int {
	length := len(arr)
	left, right := 0, length-1

	for left <= right {
		mid := (left + right) / 2
		if mid == 0 {
			return 1
		}
		if mid == length-1 {
			return length - 2
		}
		if arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] {
			return mid
		}
		if arr[mid-1] > arr[mid] {
			right = mid - 1
		} else if arr[mid] < arr[mid+1] {
			left = mid + 1
		}

	}

	return -1
}
