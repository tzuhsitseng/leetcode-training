package main

import "sort"

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	l := len(nums)

	// 先排序
	sort.Ints(nums)

	// 開始使用 two pointers 去查訪
	for i := 0; i < l-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			j := i + 1
			k := l - 1
			target := 0 - nums[i]

			for j < k {
				nj := nums[j]
				nk := nums[k]
				if nj+nk == target {
					result = append(result, []int{nums[i], nums[j], nums[k]})

					for j < k && nums[j] <= nj {
						j++
					}

					for j < k && nums[k] >= nk {
						k--
					}
				} else if nj+nk < target {
					j++
				} else if nj+nk > target {
					k--
				}
			}
		}
	}

	return result
}
