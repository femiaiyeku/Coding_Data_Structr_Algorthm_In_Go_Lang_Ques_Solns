/*

You're given a list of arbitrary jobs that need to be completed; these jobs are represented by distinct integers.
You're also given a list of dependencies. A dependency is represented as a pair of jobs where the first job is a prerequisite of the second one.
In other words, the second job depends on the first one; it can only be completed once the first job is completed.
Write a function that takes in a list of jobs and a list of dependencies and returns a list containing a valid order in which the given jobs can be completed.
If no such order exists, the function should return an empty array.


Sample Input:
jobs = [1, 2, 3, 4]
deps = [
    [1, 2],
    [1, 3],
    [3, 2],
    [4, 2],
    [4, 3],
]

Sample Output:
[1, 4, 3, 2] or [4, 1, 3, 2]


Hints
Hint 1

Try representing the jobs and dependencies as a graph, where each vertex is a job and each edge is a dependency. How can you traverse this graph to topologically sort the list of jobs?

Hint 2

One approach to solving this problem is to traverse the graph mentioned in Hint #1 using Depth-first Search. Starting at a random job, traverse its prerequisite jobs in Depth-first Search fashion until you reach a job with no prerequisites; such a job can safely be appended to the final order. Once you've traversed and added all prerequisites of a job to the final order, you can append the job in question to the order. This approach will have to track whether nodes have been traversed already, whether they're in the process of being traversed (which would indicate a cycle in the graph and therefore no valid topological order), or whether they're ready to be traversed.

Hint 3

Another approach to solving this problem is to traverse the graph mentioned in Hint #1 starting specifically with jobs that have no prerequisites. Keep track of all the jobs that have no prerequisites, traverse them one by one, and append them to the final order. For all of these jobs, remove their dependencies from the graph and update the number of prerequisites for each of these dependencies accordingly (these dependencies should now have one prerequisite less since one of their prerequisite job has just been added to the final order). As you update the number of prerequisites for these other jobs, keep track of the ones that no longer have prerequisites and that are ready to be traversed. You'll eventually go through all of the jobs if there are no cycles in the graph. If there is a cycle in the graph, there will still be jobs with prerequisites and you'll know that there is no valid topological order. This approach will involve keeping track of the number of prerequisites per job as well as all the actual dependencies of each job.


*/

package main

import (
	"fmt"
)

type Job struct {
	id            int
	prerequisites []int
	inDegree      int
}

func findValidOrder(jobs []Job, deps [][]int) []int {
	inDegreeMap := make(map[int]int)
	prerequisitesMap := make(map[int][]int)

	for _, job := range jobs {
		inDegreeMap[job.id] = 0
		prerequisitesMap[job.id] = []int{}
	}

	for _, dep := range deps {
		prerequisite, dependent := dep[0], dep[1]
		inDegreeMap[dependent]++
		prerequisitesMap[dependent] = append(prerequisitesMap[dependent], prerequisite)
	}

	queue := []int{}
	for jobID, inDegree := range inDegreeMap {
		if inDegree == 0 {
			queue = append(queue, jobID)
		}
	}

	orderedJobs := []int{}
	for len(queue) > 0 {
		jobID := queue[0]
		queue = queue[1:]

		orderedJobs = append(orderedJobs, jobID)

		for _, prerequisite := range prerequisitesMap[jobID] {
			inDegreeMap[prerequisite]--
			if inDegreeMap[prerequisite] == 0 {
				queue = append(queue, prerequisite)
			}
		}
	}

	if len(orderedJobs) != len(jobs) {
		return []int{}
	}

	return orderedJobs
}

func main() {
	jobs := []Job{
		{id: 1, prerequisites: []int{2, 3}},
		{id: 2, prerequisites: []int{1, 3}},
		{id: 3, prerequisites: []int{1}},
		{id: 4},
	}

	deps := [][]int{
		{1, 2},
		{1, 3},
		{3, 2},
		{4, 2},
		{4, 3},
	}

	order := findValidOrder(jobs, deps)
	fmt.Println(order)
}
