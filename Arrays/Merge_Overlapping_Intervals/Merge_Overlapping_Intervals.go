/*

Write a function that takes in a non-empty array of arbitrary Intervals, merges any overlapping intervals, and
returns the new Intervals in no particular order.Each interval interval is an array of two integers,
with interval [0] as the start of the interval and interval[1] as theend of the interval.Note that back-to-back intervals aren't considered to be overlapping.
 For example, [1, 5] and [6, 7] aren't overlapping,however, [1, 6] and [6, 7] are indeed overlapping.
Also note that the start of any particular interval will always be less than or equal to the end of that interval.


Sample Input
intervals = [
    [1, 2],
    [3, 5],
    [4, 7],
    [6, 8],
    [9, 10]
]

Sample Output
[
    [1, 2],
    [3, 8],
    [9, 10]
]
// Merge the intervals [3, 5], [4, 7], and [6, 8].
// The intervals could be ordered differently.




*/

package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func mergeIntervals(intervals []Interval) []Interval {
	// Sort the intervals by start time.
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	// Create a new slice to store the merged intervals.
	mergedIntervals := []Interval{}

	// Iterate over the sorted intervals and merge any overlapping intervals.
	for i := 0; i < len(intervals); i++ {
		// If the current interval overlaps with the previous interval, merge the two intervals.
		if i > 0 && intervals[i].Start <= intervals[i-1].End {
			newInterval := Interval{intervals[i-1].Start, intervals[i].End}
			mergedIntervals = append(mergedIntervals, newInterval)
			continue
		}

		// Add the current interval to the merged intervals slice.
		mergedIntervals = append(mergedIntervals, intervals[i])
	}

	return mergedIntervals
}

func main() {
	intervals := []Interval{
		{1, 2},
		{3, 5},
		{4, 7},
		{6, 8},
		{9, 10},
	}

	mergedIntervals := mergeIntervals(intervals)

	fmt.Println(mergedIntervals)
}
