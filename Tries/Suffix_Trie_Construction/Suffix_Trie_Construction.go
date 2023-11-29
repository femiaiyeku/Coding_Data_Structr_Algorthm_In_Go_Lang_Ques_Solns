/*


Write a Suffix Trie class for a Suffix-Trie-like data structure. The class should have a root property set to be
the root node of the trie and should support:
• Creating the trie from a string; this will be done by calling the populateSuffixTrieFrom method upon
class instantiation, which should populate the root of the class.
• Searching for strings in the trie.
Note that every string added to the trie should end with the special endSymbol character: "*"
If you're unfamiliar with Suffix Tries, we recommend watching the Conceptual Overview section of this question's

video explanation before starting to code.
Sample Input (for creation):
string = "babc"

Sample Output (for creation) [The strings below are in no particular order]:

{
"c": {"*": True},
"b": {
"c": {"*": True},
"a": {"b": {"c": {"*": True}}},
"*": True
},

"a": {"b": {"c": {"*": True}}},
"*": True
}

Sample Input (for searching in the suffix trie above):
string = "abc"

Sample Output (for searching in the suffix trie above):
true





This code defines a SuffixTrieNode struct to represent each node in the trie and a SuffixTrie struct to manage the overall trie structure. The populateSuffixTrieFrom method constructs the trie from a given string by inserting each substring of the string into the trie. The search method searches for a given string in the trie by traversing the trie and checking if the end node for the string exists.



*/

package main

import (
	"fmt"
)

type SuffixTrieNode struct {
	children map[string]*SuffixTrieNode
	isEnd    bool
}

type SuffixTrie struct {
	root *SuffixTrieNode
}

func NewSuffixTrie() *SuffixTrie {
	return &SuffixTrie{root: &SuffixTrieNode{children: make(map[string]*SuffixTrieNode)}}
}

func (trie *SuffixTrie) populateSuffixTrieFrom(str string) {
	for i := len(str) - 1; i >= 0; i-- {
		currentChar := string(str[i])
		currentNode := trie.root

		for j := i; j >= 0; j-- {
			currentSubstr := string(str[j:])
			if _, exists := currentNode.children[currentChar]; !exists {
				currentNode.children[currentChar] = &SuffixTrieNode{children: make(map[string]*SuffixTrieNode)}
			}

			currentNode = currentNode.children[currentChar]
			if len(currentSubstr) == 1 {
				currentNode.isEnd = true
			}
		}
	}
}

func (trie *SuffixTrie) search(str string) bool {
	currentNode := trie.root

	for i := 0; i < len(str); i++ {
		currentChar := string(str[i])
		if _, exists := currentNode.children[currentChar]; !exists {
			return false
		}

		currentNode = currentNode.children[currentChar]
	}

	return currentNode.isEnd
}

func main() {
	trie := NewSuffixTrie()
	trie.populateSuffixTrieFrom("babc")

	str := "abc"
	found := trie.search(str)
	fmt.Println("Found:", found)
}
