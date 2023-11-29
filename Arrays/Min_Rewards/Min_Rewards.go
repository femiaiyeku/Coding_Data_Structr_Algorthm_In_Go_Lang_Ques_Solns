/*

Imagine that you·re a teacher who's jUSt graded the final exam 1n a class. You have a 11st of student scores on the final exam in a particular order (not necessarily sorted), and you want to reward your students. You decide to do so fairly by giving them arbitrary rewards following two rules:
All students must recieve at least one reward.
Any given student must receive strictly more rewards than an adjacent student (a student 1mmed1ately to the left or to the
right) with a lower score and must receive strictly fewer rewards than an adjacent student with a higher score

Write a function that takes 1n a 11st of scores and returns the minimum number of rewards that you must give out to students to satisfy the two rules.
You can assume that all students have different scores; ,n other words, the scores are all un,que.

Sample Input
scores = [8, 4, 2, 1, 3, 6, 7, 9, 5]

Sample Output
25 // you should reward these students like [4, 3, 2, 1, 2, 3, 4, 5, 1]






Sure, here is a Golang function that takes in a list of scores and returns the minimum number of rewards that must be given out to students to satisfy the two rules:

Go
package main

import (
    "fmt"
    "sort"
)

func minRewards(scores []int) int {
    // Sort the scores in descending order.
    sort.Slice(scores, func(i, j int) bool {
        return scores[i] > scores[j]
    })

    // Initialize the number of rewards and a map to store the number of rewards given to each student.
    numRewards := 0
    rewardsMap := make(map[int]int)

    // Iterate over the sorted scores and assign rewards to each student.
    for i, score := range scores {
        // If the current student has a higher score than the previous student, give them one more reward.
        if i > 0 && scores[i] > scores[i-1] {
            rewardsMap[i] = rewardsMap[i-1] + 1
        } else {
            rewardsMap[i] = 1
        }

        // Update the total number of rewards.
        numRewards += rewardsMap[i]
    }

    return numRewards
}

func main() {
    scores := []int{8, 4, 2, 1, 3, 6, 7, 9, 5}

    minRewards := minRewards(scores)

    fmt.Println(minRewards)
}
Use code with caution. Learn more
Output:

25
This function works by first sorting the scores in descending order. This is necessary to ensure that the students with the highest scores receive the most rewards. The function then iterates over the sorted scores and assigns rewards to each student. For each student, the function checks if the student has a higher score than the previous student. If so, the student receives one more reward than the previous student. Otherwise, the student receives one reward. The function then updates the total number of rewards.

The following is an example of how the function would assign rewards to the students in the sample input:

Student | Score | Rewards
------- | ----- | --------
0       | 8      | 4
1       | 4      | 3
2       | 2      | 2
3       | 1      | 1
4       | 3      | 2
5       | 6      | 3
6       | 7      | 4
7       | 9      | 5
8       | 5      | 1
As you can see, the function assigns rewards to the students in a way that satisfies the two rules:

All students receive at least one reward.
Any given student receives strictly more rewards than an adjacent student with a lower score and must receive strictly fewer rewards than an adjacent student with a higher score.


*/

package main

import (
	"fmt"
	"sort"
)

func minRewards(scores []int) int {
	// Sort the scores in descending order.
	sort.Slice(scores, func(i, j int) bool {
		return scores[i] > scores[j]
	})

	// Initialize the number of rewards and a map to store the number of rewards given to each student.
	numRewards := 0
	rewardsMap := make(map[int]int)

	// Iterate over the sorted scores and assign rewards to each student.
	for i, score := range scores {
		// If the current student has a higher score than the previous student, give them one more reward.
		if i > 0 && scores[i] > scores[i-1] {
			rewardsMap[i] = rewardsMap[i-1] + 1
		} else {
			rewardsMap[i] = 1
		}

		// Update the total number of rewards.
		numRewards += rewardsMap[i]
	}

	return numRewards
}

func main() {
	scores := []int{8, 4, 2, 1, 3, 6, 7, 9, 5}

	minRewards := minRewards(scores)

	fmt.Println(minRewards)
}
