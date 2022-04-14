package main

// https://leetcode.com/problems/sort-an-array/

import (
	"fmt"
	"sort"
)

func main() {
	d := []int{5, 2, 3, 1}
	fmt.Println(sortArray(d))
}

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
	size := len(nums)
	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
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

	i, j := 0, 0
	for j = range nums {
		for bucket[i] == 0 {
			i++
		}

		nums[j] = i - 50000
		bucket[i]--
	}

	return nums
}

func mergeSort(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}

	result := make([]int, 0, len(nums))
	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid:]
	left = mergeSort(left)
	right = mergeSort(right)

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func quickSort(nums []int) []int {
	doQuickSort(nums, 0, len(nums)-1)
	return nums
}

func doQuickSort(nums []int, left, right int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	if right <= left {
		return
	}

	l, r := left, right
	m := (right-left)/2 + left
	pilot := nums[m]

	for r >= l {
		for r >= l && nums[l] < pilot {
			l++
		}
		for r >= l && nums[r] > pilot {
			r--
		}
		if r >= l {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	doQuickSort(nums, left, r)
	doQuickSort(nums, l, right)
}
