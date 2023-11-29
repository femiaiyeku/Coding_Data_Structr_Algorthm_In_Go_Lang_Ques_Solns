/*

There's an algorithms tournament taking place in which teams of programmers compete against each other to solve algorithmic problems as fast as possible. Teams compete in a round robin, where each team faces off against all other teams. Only two teams compete against each other at a time, and for each competition, one team is designated the home team, while the other team is the away team. In each competition there's always one winner and one loser; there are no ties. A team receives 3 points if it wins and O points if it loses. The winner of the tournament is the team that receives the most amount of points.
Given an array of pairs representing the teams that have competed against each other and an array containing the results of each competition, write a function that returns the winner of the tournament. The input arrays are named competitions and results , respectively. The competitions array has elements in the form of [homeTeam, awayTeam] , where each team is a string of at most 30 characters representing the name of the team. The results array contains information about the winner of each corresponding competition in the competitions array. Specifically, results [i] denotes the winner of competitions[i] , where a 1 in the results array means that the home team in the corresponding competition won and a 0 means that the away team won.
It's guaranteed that exactly one team will win the tournament and that each team will compete against all other teams exactly once. It's also guaranteed that the tournament will always have at least two teams.


Sample Input
competitions = [
    ["HTML", "C#"],
    ["C#", "Python"],
    ["Python", "HTML"],
]
results = [0, 0, 1]


Sample Output
"Python"
// C# beats HTML, Python Beats C#, and Python Beats HTML.
// HTML - 0 points
// C# -  3 points
// Python -  6 points



This function works by first initializing a map to store the points of each team. The function then iterates over the competitions array. For each competition, the function gets the home team, the away team, and the result of the competition. The function then updates the points of the winning team.

Once the function has finished iterating over the competitions array, it finds the team with the most points and returns that team as the winner of the tournament.

*/

package main

import (
	"fmt"
)

func findTournamentWinner(competitions [][]string, results []int) string {
	// Initialize a map to store the points of each team.
	teamPoints := make(map[string]int)

	// Iterate over the competitions array.
	for i, competition := range competitions {
		// Get the home team and the away team.
		homeTeam := competition[0]
		awayTeam := competition[1]

		// Get the result of the competition.
		result := results[i]

		// Update the points of the winning team.
		if result == 1 {
			teamPoints[homeTeam] += 3
		} else {
			teamPoints[awayTeam] += 3
		}
	}

	// Find the team with the most points.
	var winner string
	var maxPoints int
	for team, points := range teamPoints {
		if points > maxPoints {
			maxPoints = points
			winner = team
		}
	}

	// Return the winner of the tournament.
	return winner
}

func main() {
	competitions := [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	}
	results := []int{0, 0, 1}

	winner := findTournamentWinner(competitions, results)

	fmt.Println(winner)
}
