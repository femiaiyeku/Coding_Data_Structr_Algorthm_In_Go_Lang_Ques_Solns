/*



You recently started freelance software development and have been offered a variety of job opportunities. Each job has
a deadline, meaning there is no value in completing the work after the deadline. Additionally, each job has an associated
payment representing the profit for completing that job. Given this information, write a function that returns the
maximum profit that can be obtained in a 7-day period.
Each job will take 1 full day to complete, and the deadline will be given as the number of days left to complete the job.
For example, if a job has a deadline of 1, then it can only be completed if it is the first job worked on. If a job has a
deadline of 2, then it could be started on the first or second day.
Note: There is no requirement to complete all of the jobs. Only one job can be worked on at a time, meaning that in
some scenarios it will be impossible to complete them all.


Sample Input:

jobs = [
    {"deadline": 2, "profit": 100},
    {"deadline": 1, "profit": 19},
    {"deadline": 2, "profit": 27},
    {"deadline": 1, "profit": 25},
    {"deadline": 3, "profit": 15}

]

Sample Output:
3 // Job 2, then Job 1, then Job 3







*/

package main

import (
	"fmt"
	"sort"
)

type Job struct {
	deadline int
	profit   int
}

func maximumProfit(jobs []Job) int {
	// Sort the jobs by their deadlines in descending order
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].deadline > jobs[j].deadline
	})

	// Initialize a map to store the maximum profit for each deadline
	maxProfitByDeadline := make(map[int]int)

	// Calculate the maximum profit for each deadline
	for _, job := range jobs {
		currentProfit := job.profit

		// Check if there is a previous job with a deadline that allows us to fit this job in
		for previousDeadline := job.deadline - 1; previousDeadline >= 0; previousDeadline-- {
			previousProfit := maxProfitByDeadline[previousDeadline]
			if previousProfit > 0 {
				currentProfit = job.profit + previousProfit
				break
			}
		}

		// Update the maximum profit for the current deadline
		maxProfitByDeadline[job.deadline] = currentProfit
	}

	// Find the maximum profit across all deadlines
	maximumProfit := 0
	for _, profit := range maxProfitByDeadline {
		if profit > maximumProfit {
			maximumProfit = profit
		}
	}

	return maximumProfit
}

func main() {
	jobs := []Job{
		{deadline: 2, profit: 100},
		{deadline: 1, profit: 19},
		{deadline: 2, profit: 27},
		{deadline: 1, profit: 25},
		{deadline: 3, profit: 15},
	}

	maxProfit := maximumProfit(jobs)
	fmt.Println(maxProfit)
}
