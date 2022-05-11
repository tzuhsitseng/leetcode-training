package main

import "sort"

// https://www.lintcode.com/problem/919/

type Interval struct {
	Start, End int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinMeetingRooms(intervals []*Interval) int {
	if len(intervals) == 0 {
		return 0
	}

	start := make([]int, 0, len(intervals))
	end := make([]int, 0, len(intervals))
	for _, v := range intervals {
		start = append(start, v.Start)
		end = append(end, v.End)
	}
	sort.Ints(start)
	sort.Ints(end)

	room := 0
	res := 0
	i, j := 0, 0

	for i < len(intervals) && j < len(intervals) {
		if end[j] <= start[i] {
			room--
			j++
		} else {
			room++
			res = max(res, room)
			i++
		}
	}

	return res
}
