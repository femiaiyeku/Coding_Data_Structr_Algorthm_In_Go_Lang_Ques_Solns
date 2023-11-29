/*

From Wikipedia, the concept of six degrees of separation "is the idea that all people are six, or fewer, social connections away from each other."

For example, if Clement and Antoine are friends, they share one degree of separation. If Simon is Antoine's friend but not Clement's friend, then Clement and Simon share two degrees of separation, since they're connected by Antoine.

You're given a directory of friends lists (a map of people's names pointing to lists of their friends' names) as well as the names of two people (found in the directory).

Write a function that returns the name of the person (one of the two input names) that is more than six degrees of separation away from the fewest amount of people (in the directory).

If both input persons have the same number of people who are more than six degrees of separation away, your function should return the empty string "".

Note that if Antoine is Clement's friend, it follows that Clement is Antoine's friend. Thus, if person A's name is found in person B's friends list, then person B's name will be found in person A's friends list.




Sample Input
friendsList = {
  "Aaron": ["Paul"],
  "Akshay": [],
  "Alex": ["Antoine", "Clement", "Ryan", "Simon"],
  "Antoine": ["Alex", "Clement", "Simon"],
  "Ayushi": ["Lee"],
  "Changpeng": ["Kelly", "Sandeep"],
  "Clement": ["Alex", "Antoine", "Sandeep", "Simon"],
  "Hannah": ["Lexi", "Paul", "Sandeep"],
  "James": ["Paul"],
  "Kelly": ["Changpeng", "Molly"],
  "Lee": ["Ayushi", "Molly"],
  "Lexi": ["Hannah"],
  "Molly": ["Kelly", "Lee"],
  "Paul": ["Aaron", "James", "Hannah"],
  "Ryan": ["Alex"],
  "Sandeep": ["Changpeng", "Clement", "Hannah"],
  "Simon": ["Alex", "Antoine", "Clement"]
},
personOne = "Antoine"
personTwo = "Clement"
Sample Output
"Clement"
// Neither Antoine nor Clement have any connection to Akshay.
// Antoine is seven degrees of separation away from Ayushi.
// So Clement only has one person who is more than six degrees of
// separation away (Akshay), whereas Antoine has two (Akshay and Ayushi).










*/

package degrees_of_separation

// important question

import (
	"math"
)

// DegreesOfSeparation
// O(v + e) time | O(v) space - where v is the number of vertices (people) in the
// friends graph and e is the number of edges (total friends) in the friends graph
// https://stackoverflow.com/questions/18604803/why-is-the-complexity-of-bfs-ove-instead-of-ove
func DegreesOfSeparation(friendsLists FriendsGraph, personOne, personTwo string) string {
	degreesOne := getDegreesOfSeparation(friendsLists, personOne)
	degreesTwo := getDegreesOfSeparation(friendsLists, personTwo)
	numDegreesOverSixOne := getNumDegreesOverSix(friendsLists, degreesOne)
	numDegreesOverSixTwo := getNumDegreesOverSix(friendsLists, degreesTwo)
	if numDegreesOverSixOne == numDegreesOverSixTwo {
		return ""
	} else if numDegreesOverSixOne < numDegreesOverSixTwo {
		return personOne
	}
	return personTwo
}

type DegreePair struct {
	Person string
	Degree int
}

func getDegreesOfSeparation(friendsLists FriendsGraph, origin string) map[string]int {
	degrees := map[string]int{}
	visited := map[string]bool{}
	queue := []DegreePair{{Person: origin, Degree: 0}}
	for len(queue) > 0 {
		var pair DegreePair
		pair, queue = queue[0], queue[1:]

		person, degree := pair.Person, pair.Degree
		degrees[person] = degree

		for _, friend := range friendsLists[person] {
			if visited[friend] {
				continue
			}
			visited[friend] = true
			queue = append(queue, DegreePair{
				Person: friend, Degree: degree + 1,
			})
		}
	}

	for person := range friendsLists {
		if !visited[person] {
			degrees[person] = math.MaxInt32
		}
	}
	return degrees
}

func getNumDegreesOverSix(friendsLists FriendsGraph, degrees map[string]int) int {
	numDegreesOverSix := 0
	for person := range friendsLists {
		if degrees[person] > 6 {
			numDegreesOverSix++
		}
	}
	return numDegreesOverSix
}
