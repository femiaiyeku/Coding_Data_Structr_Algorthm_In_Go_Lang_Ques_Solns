/*
Imagine that you want to schedule a meeting of a certain duration with a co-worker. You have access to your calendar and your co-worker's calendar (both of which contain your respective meetings for the day, in the form of [startTime, endTime] ), as well as both of your daily bounds (i.e., the earliest and latest times at which you're available for meetings every day, in the form of [earl iestTime, latestTime] ).
Write a function that takes in your calendar, your daily bounds, your co-worker's calendar, your co-worker's daily bounds, and the duration of the meeting that you want to schedule, and that returns a list of all the time blocks (in the form of [startT;me, endT;me] ) during which you could schedule the meeting, ordered from earliest time block to latest.
Note that times will be given and should be returned in military time. For example: 8: 30 , 9: 01 , and 23: 56 .
Also note that the given calendar times will be sorted by start time in ascending order, as you would expect them to appear in a calendar application like Google Calendar.


Sample Input
calendar1 = [['9:00', '10:30'], ['12:00', '13:00'J, ['16:00', '18:00']]
dai1yBoundsl = [ '9:00', '20:00' J
ca1endar2 = [('10:00', '11:30'), ['12:30', '14:30'), ['14:30', '15:00'J, ['16:00', '17:00')) dai1yBounds2 = ['10:00', '18:30']
meetingDuration = 30


Sample Output
[['11:30', '12:00'], ['15:00', '16:00'], ['18:00', '18:30']]


Hints
Hint 1

In order to find blocks of time during which you and your coworker can have a meeting, you have to first find all of the unavailabilities between the two of you. An unavailability is any block of time during which at least one of you is busy.

Hint 2

You'll have to start by taking into account the daily bounds; the daily bounds can be represented by two additional meetings in each of your calendars.

Hint 3

Once you've taken the daily bounds into account, you'll want to merge the two calendars into a single calendar of meetings and then flatten that calendar in order to eliminate overlapping and back-to-back meetings. This will give you a calendar of unavailabilities, from which you can pretty easily find the list of matching availabilities.







*/

package main

import (
	"sort"
	"strconv"
	"strings"
)

func findAvailableTimeBlocks(calendar1 [][]string, dailyBounds1 []string, calendar2 [][]string, dailyBounds2 []string, meetingDuration int) [][]string {

	// Merge the two calendars.
	mergedCalendar := mergeCalendars(calendar1, calendar2)

	// Find all the time blocks during which you are both available.
	availableTimeBlocks := findAvailableTimeBlocksInMergedCalendar(mergedCalendar, dailyBounds1, dailyBounds2, meetingDuration)

	// Return the available time blocks, ordered from earliest time block to latest.
	return sortAvailableTimeBlocks(availableTimeBlocks)
}

// Merge two calendars.
func mergeCalendars(calendar1 [][]string, calendar2 [][]string) [][]string {

	// Create a new slice to store the merged calendar.
	mergedCalendar := make([][]string, 0)

	// Iterate over the two calendars and add each time block to the merged calendar.
	i := 0
	j := 0
	for i < len(calendar1) && j < len(calendar2) {

		// If the start time of the current time block in calendar1 is earlier than the start time of the current time block in calendar2, add the current time block in calendar1 to the merged calendar.
		if compareTimes(calendar1[i][0], calendar2[j][0]) == -1 {
			mergedCalendar = append(mergedCalendar, calendar1[i])
			i++
		} else if compareTimes(calendar1[i][0], calendar2[j][0]) == 1 {
			mergedCalendar = append(mergedCalendar, calendar2[j])
			j++
		} else {
			// The start times of the current time blocks in the two calendars are the same.
			mergedCalendar = append(mergedCalendar, calendar1[i])
			i++
			j++
		}
	}

	// Add any remaining time blocks from calendar1 to the merged calendar.
	for i < len(calendar1) {
		mergedCalendar = append(mergedCalendar, calendar1[i])
		i++
	}

	// Add any remaining time blocks from calendar2 to the merged calendar.
	for j < len(calendar2) {
		mergedCalendar = append(mergedCalendar, calendar2[j])
		j++
	}

	// Return the merged calendar.
	return mergedCalendar
}

// Find all the time blocks during which you are both available in the merged calendar.
func findAvailableTimeBlocksInMergedCalendar(mergedCalendar [][]string, dailyBounds1 []string, dailyBounds2 []string, meetingDuration int) [][]string {

	// Create a new slice to store the available time blocks.
	availableTimeBlocks := make([][]string, 0)

	// Iterate over the merged calendar.
	for i := 0; i < len(mergedCalendar)-1; i++ {

		// If the end time of the current time block is less than the start time of the next time block, and the duration of the current time block is greater than or equal to the meeting duration, and the current time block is within the daily bounds of both you and your co-worker, add the current time block to the available time blocks.
		if compareTimes(mergedCalendar[i][1], mergedCalendar[i+1][0]) == -1 && compareTimes(mergedCalendar[i][1], mergedCalendar[i][0]) >= meetingDuration && isWithinDailyBounds(mergedCalendar[i], dailyBounds1, dailyBounds2) {
			availableTimeBlocks = append(availableTimeBlocks, mergedCalendar[i])
		}
	}

	// Return the available time blocks.
	return availableTimeBlocks
}

// Sort the available time blocks in ascending order.
func sortAvailableTimeBlocks(availableTimeBlocks [][]string) [][]string {

	// Sort the available time blocks by start time.
	sort.Slice(availableTimeBlocks, func(i, j int) bool {
		return compareTimes(availableTimeBlocks[i][0], availableTimeBlocks[j][0]) == -1
	})

	// Return the sorted available time blocks.
	return availableTimeBlocks
}

// Compare two times.
func compareTimes(time1 string, time2 string) int {

	// Convert the times to minutes.
	minutes1 := convertTimeToMinutes(time1)
	minutes2 := convertTimeToMinutes(time2)

	// Compare the times.
	if minutes1 < minutes2 {
		return -1
	} else if minutes1 > minutes2 {
		return 1
	} else {
		return 0
	}
}

// Convert a time to minutes.
func convertTimeToMinutes(time string) int {

	// Split the time into hours and minutes.
	timeSplit := strings.Split(time, ":")
	hours, _ := strconv.Atoi(timeSplit[0])
	minutes, _ := strconv.Atoi(timeSplit[1])

	// Convert the time to minutes.
	return hours*60 + minutes
}

// Check if a time block is within the daily bounds of both you and your co-worker.
func isWithinDailyBounds(timeBlock []string, dailyBounds1 []string, dailyBounds2 []string) bool {

	// Convert the start time of the time block to minutes.
	startTimeMinutes := convertTimeToMinutes(timeBlock[0])

	// Convert the daily bounds to minutes.
	dailyBounds1Minutes := convertDailyBoundsToMinutes(dailyBounds1)
	dailyBounds2Minutes := convertDailyBoundsToMinutes(dailyBounds2)

	// Check if the start time of the time block is within the daily bounds of both you and your co-worker.
	return startTimeMinutes >= dailyBounds1Minutes[0] && startTimeMinutes >= dailyBounds2Minutes[0] && startTimeMinutes < dailyBounds1Minutes[1] && startTimeMinutes < dailyBounds2Minutes[1]
}

// Convert daily bounds to minutes.
func convertDailyBoundsToMinutes(dailyBounds []string) []int {

	// Create a new slice to store the daily bounds in minutes.
	dailyBoundsMinutes := make([]int, 0)

	// Convert the daily bounds to minutes.
	for _, bound := range dailyBounds {
		dailyBoundsMinutes = append(dailyBoundsMinutes, convertTimeToMinutes(bound))
	}

	// Return the daily bounds in minutes.
	return dailyBoundsMinutes
}

// O(c1 + c2) time | O(c1 + c2) space - where c1 and c2 are the numbers of meetings in calendar1 and calendar2, respectively
