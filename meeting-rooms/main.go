package main

import (
	"fmt"
)

// https://www.lintcode.com/problem/920/

type Interval struct {
	Start, End int
}

func CanAttendMeetings(intervals []*Interval) bool {
	if len(intervals) == 0 {
		return true
	}

	doQuickSort(intervals, 0, len(intervals)-1)

	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start < intervals[i-1].End {
			return false
		}
	}

	return true
}

func doQuickSort(intervals []*Interval, left, right int) {
	if len(intervals) == 0 || len(intervals) == 1 {
		return
	}

	if right <= left {
		return
	}

	l, r := left, right
	m := (right-left)/2 + left
	pilot := intervals[m]

	for r >= l {
		for r >= l && intervals[l].Start < pilot.Start {
			l++
		}
		for r >= l && intervals[r].Start > pilot.Start {
			r--
		}
		if r >= l {
			intervals[l], intervals[r] = intervals[r], intervals[l]
			l++
			r--
		}
	}

	doQuickSort(intervals, left, r)
	doQuickSort(intervals, l, right)
}

func main() {
	intervals := []*Interval{
		{Start: 465, End: 497},
		{Start: 386, End: 462},
		{Start: 354, End: 380},
		{Start: 134, End: 189},
		{Start: 199, End: 282},
		{Start: 18, End: 104},
		{Start: 499, End: 562},
		{Start: 4, End: 14},
		{Start: 111, End: 129},
		{Start: 292, End: 345},
	}
	fmt.Println(CanAttendMeetings(intervals))
}
