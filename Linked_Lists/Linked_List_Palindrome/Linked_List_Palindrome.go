/*

Write a function that takes ,n the head of a Singly Linked List and returns a boolean representing whether the linked list's nodes form a palindrome.
 Your function shouldn't make use of any aux1l1ary data structure 
A palindrome ,s usually defined as a string that's written the same forward and backward. For a linked list's nodes to form a palindrome, 
their values must be the same when read from left to right and from right to left. Note that single-character strings are palindromes, 
which means that single-node linked lists form palindromes. 
Each Linked  list node has an integer value as well as a next node po,nt,ng to the next node ,n the list or to None / null if it's the tail of the list. 
You can assume that the input linked list will always have at least one node; ,n other words. the head will never be None / null. 

Sample Input:
head = 0 -> 1 -> 2 -> 2 -> 1 -> 0 // the head node with value 0

Sample Output:
true // it's a palindrome






*/


package main

type Node struct {
    value int
    next  *Node
}

func isPalindrome(head *Node) bool {
    slow := head
    fast := head

    for fast != nil && fast.next != nil {
        slow = slow.next
        fast = fast.next.next
    }

    mid := slow

    slow = head
    var prev *Node

    for mid != nil {
        next := slow.next
        slow.next = prev
        prev = slow
        slow = next

        if mid == slow {
            break
        }
    }

    for slow != nil && prev != nil {
        if slow.value != prev.value {
            return false
        }

        slow = slow.next
        prev = prev.next
    }

    return true
}

func main() {
    head := &Node{value: 0, next: &Node{value: 1, next: &Node{value
