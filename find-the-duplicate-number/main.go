package main

func findDuplicate(nums []int) int {
	if len(nums) < 2 {
		return -1
	}

	// 龜兔賽跑 phase I
	slow, fast := nums[0], nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	// 龜兔賽跑 phase II
	fast = nums[0]

	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
