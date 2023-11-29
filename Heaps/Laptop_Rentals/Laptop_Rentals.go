/*

You're given a list of time intervals during which students at a school need a laptop. These time intervals are represented by pairs of integers [start, end] ,
 where 0 <= start < end . However, start and end don't represent real times; therefore, they may be greater than 24 .
No two students can use a laptop at the same time, but immediately after a student is done using a laptop,
another student can use that same laptop. For example, if one student rents a laptop during the time interval [0, 2] ,
another student can rent the same laptop during any time interva starting with 2 .
Write a function that returns the minimum number of laptops that the school needs to rent such that all students will always have access to a laptop when they need one.


Sample Input:
times = [
    [0, 2],
    [1, 4],
    [4, 6],
    [0, 4],
    [7, 8],
    [9, 11],
    [3, 10]
]

Sample Output:
3



Hints
Hint 1

There are many different ways to solve this problem, but only a few of them run in the optimal time. Can you come up with an algorithm that solves this problem in O(nlog(n)) time?

Hint 2

Suppose that you're given two time intervals: [s1, e1] and [s2, e2], where s1 < s2. If e1 <= s2, then the second time interval can use the same laptop as the first time interval. One method to solve this problem with an optimal time complexity is to use a Min Heap. If you loop through time intervals that have been sorted by their start times and keep track of the smallest end time of time intervals for laptops that have already been rented out, you can determine how many laptops are required. Use the Min Heap to efficiency determine if any previous rental time intervals have ended as you loop through all the time intervals. If a rental time interval is done and another one starts after it, no extra laptop is required.

Hint 3

Another way to efficiently solve this problem is to realize that we don't need to know what start time corresponds with what end time. So long as we know all start times and all end times, we can determine the number of laptops required.

Hint 4

Start by creating two arrays—one for start times and one for end times—and sort them both in ascending order. We can simply loop through the start times and end times at the same time and compare the current start time to the current end time. If the current start time is greater than the current end time, then that means a laptop that was previously used is no longer being used and can be given to the student renting a laptop at this starting time. Thus, we can increment both our start-time and end-time pointers and continue without needing an additional laptop. If the current start time is smaller than the current end time, then another rental has started before a previous rental has ended, and we thus require another laptop, so we increment the start pointer and a variable keeping track of the number of laptops required. See the Conceptual Overview section of this question's video explanation for a more in-depth explanation.


*/

package main

import (
	"container/heap"
	"sort"
)

type TimeInterval struct {
	start int
	end   int
}

func minimumNumberOfLaptops(times []TimeInterval) int {
	// Sort the time intervals by their start times
	sort.Slice(times, func(i, j int) bool {
		return times[i].start < times[j].start
	})

	// Initialize a max heap to store the currently active time intervals
	maxActiveTimes := &maxHeap{}
	laptopsNeeded := 0

	// Iterate through the time intervals
	for _, timeInterval := range times {
		// Remove any time intervals that have ended from the max heap
		for timeInterval.start < maxActiveTimes.Peek().start {
			heap.Pop(maxActiveTimes)
		}

		// Add the current time interval to the max heap
		heap.Push(maxActiveTimes, timeInterval)

		// Update the number of laptops needed if the current time interval requires more laptops than the current maximum
		if maxActiveTimes.Len() > laptopsNeeded {
			laptopsNeeded = maxActiveTimes.Len()
		}
	}

	return laptopsNeeded
}

type maxHeap []TimeInterval

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i].end < h[j].end
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	*h = old[:n-1]
	return old[n-1]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(TimeInterval))
}
