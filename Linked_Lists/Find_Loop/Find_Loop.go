/*

Write a function that takes in the head of a Singly Linked List that contains a loop {in other words, the list's tail node points to some node in the list instead of None / null ). The function should return the node {the actual node--not just its value) from which the loop originates in consta1t space. 
Each Linked list node has an irteger value as well as a next node pointing to the next node in the list. 

Sample Input:
head = 0 -> 1 -> 2 -> 3 -> 4 -> 5 -> 6 // the head node with value 2
// 6 is connected to 3
Sample Output:
4 -> 5 -> 6 // the node with value 3
// 6 is connected to 3, so 3 is the node that the input Linked List loops from
// Note: that the input Linked List may not always have a loop
// and may terminate in a None / null value
// Also note: that the target node may be the first node in the list
// Note: Do not edit the linked list



Hints
Hint 1

Try traversing the linked list with two pointers, one iterating through every single node in the list and another iterating through every other node in the list (skipping a node every time). Eventually, both pointers will point to the same node since there is a loop in the list and since one pointer is moving faster than the other. Stop once the pointers overlap each other. How can you find the origin of the loop from here?

Hint 2

Can you come up with a mathematical relation between the respective distances traveled by each pointer? How far will the first pointer have traveled when the pointers overlap? What about the second pointer? How can this relation then help you find the actual origin of the loop in the list?

Hint 3

Let D be the distance between the start of the linked list and the origin of the loop in the list. Let P be distance between the origin of the loop and the node N where the first and second pointers overlap (going in the primary direction of the list). By the time the pointers reach N, the first pointer will have traveled a distance of length D + P, and the second pointer will have traveled a distance of length 2D + 2P, since it will have traveled twice as much as the first pointer. Thus, the distance between N and the origin of the loop (going in the primary direction of the list) can be arithmetically deduced to be 2D + 2P - D - 2P = D. With both pointers D length away from the origin of the loop, how can you find the origin?




*/


package main

type Node struct {
    value int
    next  *Node
}

func detectLoopAndFindStart(head *Node) *Node {
    slow := head
    fast := head

    for {
        if fast == nil || fast.next == nil {
            return nil
        }

        slow = slow.next
        fast = fast.next.next

        if slow == fast {
            break
        }
    }

    // Slow and fast are now pointing to the same node in the loop

    slow = head
    for slow != fast {
        slow = slow.next
        fast = fast.next
    }

    return slow
}

func main() {
    head := &Node{value: 0, next: &Node{value: 1, next: &Node{value: 2, next: &Node{value: 3, next: &Node{value: 4, next: &Node{value: 5, next: &Node{value: 6, next: &Node{value: 3}}}}}}}

    loopStart := detectLoopAndFindStart(head)
    if loopStart != nil {
        fmt.Println("Loop starts at node:", loopStart.value)
    } else {
        fmt.Println("No loop detected")
    }
}
