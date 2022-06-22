package main

// https://leetcode.com/problems/3sum/

import "sort"

func threeSum(nums []int) [][]int {
	l := len(nums)

	if l == 0 {
		return nil
	}

	res := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := l - 1
		target := 0 - nums[i]

		for j < k {
			numJ, numK := nums[j], nums[k]
			if target == numJ+numK {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] <= numJ {
					j++
				}
				for j < k && nums[k] >= numK {
					k--
				}
			} else if target < numJ+numK {
				k--
			} else {
				j++
			}
		}
	}

	return res
}
