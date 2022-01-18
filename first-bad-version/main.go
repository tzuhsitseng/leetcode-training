package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// It's a mock function
func isBadVersion(version int) bool {
	return false
}

func firstBadVersion(n int) int {
	// 題目為 1 ~ N
	low := 1
	high := n

	// 黃金交叉後代表查找結束, 迴圈退出
	for low <= high {
		mid := (low + high) / 2

		// 先確認 mid 是否為 bad version
		// 若是的話直接檢查本身是否是首位或是 mid -1 是否不為 bad version，若符合條件就代表找到答案了
		// 若不是的話就再往前或往後找
		if isBadVersion(mid) {
			if mid == 1 || !isBadVersion(mid-1) {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	// 代表沒找到半個 bad version
	return -1
}
