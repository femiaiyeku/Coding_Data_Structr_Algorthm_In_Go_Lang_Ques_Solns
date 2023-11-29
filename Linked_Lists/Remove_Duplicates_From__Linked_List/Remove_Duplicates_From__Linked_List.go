/*



Remove Duplicates From Linked List
You're given the head of a Singly Linked List whose nodes are in sorted order with respect to their values. Write a
function that returns a modified version of the Linked List that doesn't contain any nodes with duplicate values.
The Linked List should be modified in place (i.e., you shouldn't create a brand new list), and the modified Linked
List should still have its nodes sorted with respect to their values.
Each Linked List node has an integer value as well as a next node pointing to the next node in the list or
to None/null if it's the tail of the list.

Sample Input:
linkedList = 1 -> 1 -> 3 -> 4 -> 4 -> 4 -> 5 -> 6 -> 6

Sample Output:
1 -> 3 -> 4 -> 5 -> 6




*/




package main

type Node struct {
    value int
    next  *Node
}

func removeDuplicates(head *Node) {
    if head == nil || head.next == nil {
        return
    }

    current := head
    prev := head

    for current.next != nil {
        if current.value == current.next.value {
            prev.next = current.next.next
        } else {
            prev = current
        }

        current = current.next
    }
}

func main() {
    head := &Node{value: 1, next: &Node{value: 1, next: &Node{value: 3, next: &Node{value: 4, next: &Node{value: 4, next: &Node{value: 4, next: &Node{value: 5, next: &Node{value: 6, next: &Node{value: 6}}}}}}}

    removeDuplicates(head)

    for current := head; current != nil; current = current.next {
        fmt.Print(current.value, " -> ")
    }
}
