/*


One of the most efficient ways to run a factory is to use an assembly line, with multiple stations
performing different assembling steps simultaneously in order to save time. But an assembly line
only as fast as its slowest station/step; for example, if an assembly line has 100 different steps
performed by 100 different stations, with 99 steps taking 1 minute each to complete and 1 step ta
1 hour to complete, then the entire assembly line is dramatically slowed down by the 1-hour-long
step.
Write a function that takes in a non-empty array of positive Integers stepDurations and a positive
integer numStations. The Input array of Integers represents the times that the various steps in
assembly process take, and the Input Integer represents the number of stations that this assembl
process has access to. For this particular assembly process, a single station can perform multiple
steps, so long as these steps are performed in order, meaning that a single station can perform
multiple steps whose times appear consecutively in the stepDurations array. Your function
should return the longest duration of a single station in the assembly line after optimizing the
assembly line (1.e., minimizing its slowest station/step).
You can assume that there will never be more stations than steps.


Sample Input:
stepDurations = [15,15 30, 30, 45]
numStations = 3

Sample Output:

60



This code first initializes two-dimensional arrays dpStationTimes and dpStationChoices to store the optimal station times and the corresponding step choices for each station and step. The dpStationTimes array stores the minimum time it takes to complete all steps up to a given step using a given number of stations. The dpStationChoices array stores the index of the previous step that led to the current minimum time.

The code then fills the base cases of the dpStationTimes and dpStationChoices arrays. The base cases represent the scenarios with no stations or no steps.

Next, the code fills the remaining elements of the dpStationTimes and dpStationChoices arrays. For each station, it iterates over all possible previous steps and calculates the minimum time it takes to complete all steps up to the current step using the current station and the previous step. It then updates the dpStationTimes and dpStationChoices arrays accordingly.

Finally, the code finds the maximum station time by iterating over the last row of the dpStationTimes array. The maximum station time represents the longest duration of a single station in the optimized assembly line.


*/

package main

import (
	"fmt"
)

func optimizeAssemblyLine(stepDurations []int, numStations int) int {
	if numStations > len(stepDurations) {
		return -1 // Invalid input: more stations than steps
	}

	// Initialize the dp arrays
	dpStationTimes := make([][]int, numStations+1)
	dpStationChoices := make([][]int, numStations+1)

	for i := range dpStationTimes {
		dpStationTimes[i] = make([]int, len(stepDurations)+1)
		dpStationChoices[i] = make([]int, len(stepDurations)+1)
	}

	// Fill the base cases
	for station := 1; station <= numStations; station++ {
		dpStationTimes[station][0] = 0
		for step := 1; step <= len(stepDurations); step++ {
			dpStationTimes[station][step] = dpStationTimes[station-1][step-1] + stepDurations[step-1]
			dpStationChoices[station][step] = step - 1
		}
	}

	// Fill the remaining dp arrays
	for station := 1; station <= numStations; station++ {
		for step := len(stepDurations); step > 0; step-- {
			for prevStep := step - 1; prevStep >= 0; prevStep-- {
				currentStationTime := dpStationTimes[station-1][prevStep] + stepDurations[step-1]

				if currentStationTime < dpStationTimes[station][step] {
					dpStationTimes[station][step] = currentStationTime
					dpStationChoices[station][step] = prevStep
				}
			}
		}
	}

	// Find the maximum station time
	maxStationTime := 0
	for stationTime := range dpStationTimes[numStations] {
		if stationTime > maxStationTime {
			maxStationTime = stationTime
		}
	}

	return maxStationTime
}

func main() {
	stepDurations := []int{15, 15, 30, 30, 45}
	numStations := 3

	optimizedStationTime := optimizeAssemblyLine(stepDurations, numStations)
	fmt.Println("Optimized assembly line station time:", optimizedStationTime)
}
