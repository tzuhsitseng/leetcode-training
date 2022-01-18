package main

import (
	"sort"
)

func sortArray(nums []int) []int {
	n := 5 // choose the case you want

	switch n {
	case 1:
		// 1. built-in sort method
		return builtInSort(nums)
	case 2:
		// 2. bubble sort
		return bubbleSort(nums)
	case 3:
		// 3. bucket sort
		return bucketSort(nums)
	case 4:
		// 4. merge sort
		return mergeSort(nums)
	case 5:
		// 5. quick sort
		return quickSort(nums)
	}

	return []int{}
}

func builtInSort(nums []int) []int {
	sort.Ints(nums)
	return nums
}

func bubbleSort(nums []int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := 0; j < l-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func bucketSort(nums []int) []int {
	bucket := make([]int, 100001)
	for _, v := range nums {
		bucket[v+50000] += 1
	}
	i := 0
	for j := range nums {
		for bucket[i] == 0 {
			i++
		}
		nums[j] = i - 50000
		bucket[i]--
	}
	return nums
}

func mergeSort(nums []int) []int {
	if len(nums) > 1 {
		mid := len(nums) / 2
		l := nums[:mid]
		r := nums[mid:]
		l = mergeSort(l)
		r = mergeSort(r)
		i, j := 0, 0
		res := make([]int, 0)

		for i < len(l) && j < len(r) {
			if l[i] < r[j] {
				res = append(res, l[i])
				i++
			} else {
				res = append(res, r[j])
				j++
			}
		}

		for i < len(l) {
			res = append(res, l[i])
			i++
		}

		for j < len(r) {
			res = append(res, r[j])
			j++
		}
		return res
	}
	return nums
}

func quickSort(nums []int) []int {
	doQuickSort(nums, 0, len(nums)-1)
	return nums
}

func doQuickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	if left+1 == right && nums[right] >= nums[left] {
		return
	}
	pivot := nums[left] // choose the left node as pivot
	l, r := left, right

	// partitioning
	for l <= r {
		for l <= r && nums[l] < pivot {
			l++
		}
		for l <= r && nums[r] > pivot {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	doQuickSort(nums, left, r)
	doQuickSort(nums, l, right)
	return
}
