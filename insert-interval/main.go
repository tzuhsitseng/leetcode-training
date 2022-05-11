package main

// https://leetcode.com/problems/insert-interval/

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)

	for i, v := range intervals {
		if newInterval[1] < v[0] {
			res = append(res, newInterval)
			res = append(res, intervals[i:]...)
			return res
		} else if newInterval[0] > v[1] {
			res = append(res, v)
		} else {
			newInterval = []int{min(newInterval[0], v[0]), max(newInterval[1], v[1])}
		}
	}

	res = append(res, newInterval)
	return res
}

// --
// The following solution is my v1 version, it can work but has a bit of luck
// --

//func insert(intervals [][]int, newInterval []int) [][]int {
//	if len(intervals) == 0 {
//		return [][]int{newInterval}
//	}
//
//	newStart, newEnd := newInterval[0], newInterval[1]
//	excludedStartIdx, excludedEndIdx := 0, len(intervals)-1
//	excludedStart, excludedEnd := newStart, newEnd
//
//	if newEnd < intervals[0][0] {
//		return append([][]int{newInterval}, intervals...)
//	}
//
//	if newStart > intervals[len(intervals)-1][1] {
//		return append(intervals, newInterval)
//	}
//
//	for i := 0; i < len(intervals); i++ {
//		start, end := intervals[i][0], intervals[i][1]
//
//		if newStart <= end {
//			excludedStartIdx = i
//			excludedStart = min(start, newStart)
//			break
//		}
//	}
//
//	for i := len(intervals) - 1; i >= 0; i-- {
//		start, end := intervals[i][0], intervals[i][1]
//
//		if newEnd >= start {
//			excludedEndIdx = i
//			excludedEnd = max(end, newEnd)
//			break
//		}
//	}
//
//	res := make([][]int, 0)
//	res = append(res, intervals[:excludedStartIdx]...)
//	res = append(res, []int{excludedStart, excludedEnd})
//	res = append(res, intervals[excludedEndIdx+1:]...)
//	return res
//}
