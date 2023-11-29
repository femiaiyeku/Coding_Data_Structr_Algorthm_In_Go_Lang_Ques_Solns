/*


The Fibonacci sequence is defined 35 follows: the first number of the sequence is 0 , the second number is 1 ,
and the nth number is the sum of the (n - 1 )th and {n - 2)th numbers.
Write a function that takes in an integer n and returns the nth Fibonacci number.
Important note: the Fibonacci sequence is often defined with its first two numbers as F0 = 0 and Fl = 1 .
For the purpose of this question,
the first Fibonacci number is F0; therefore, getNthFib(l) is equal to F0, getNthFib(2) is equal to Fl, etc..

Sample input #1: 2
Sample output #1: 1 (0, 1)

Sample input #2: 6
Sample output #2: 5 (0, 1, 1, 2, 3, 5)

Sample input #3: 8

Sample output #3: 13 (0, 1, 1, 2, 3, 5, 8, 13)





Hints
Hint 1
The formula to generate the nth Fibonacci number can be written as follows: F(n) = F(n - 1) + F(n - 2). Think of the case(s) for which this formula doesn't apply (the base case(s)) and try to implement a simple recursive algorithm to find the nth Fibonacci number with this formula.
Hint 2
What are the runtime implications of solving this problem as is described in Hint #1? Can you use memoization (caching) to improve the performance of your algorithm?
Hint 3
Realize that to calculate any single Fibonacci number you only need to have the two previous Fibonacci numbers. Knowing this, can you implement an iterative algorithm to solve this question, storing only the last two Fibonacci numbers at any given time?


This code first defines a function getNthFib that takes an integer n and returns the nth Fibonacci number. The function first checks if n is less than or equal to 1, in which case it returns n. Otherwise, it creates an array fib of size n+1 to store the Fibonacci numbers. It then sets the first two elements of fib to 0 and 1, respectively.

Next, it iterates from 2 to n, calculating the ith Fibonacci number as the sum of the (i-1)th and (i-2)th Fibonacci numbers. Finally, it returns the nth Fibonacci number.

The main function calls the getNthFib function for three different values of n and prints the results.




*/

package main

import (
	"fmt"
)

func getNthFib(n int) int {
	if n <= 1 {
		return n
	}

	fib := make([]int, n+1)
	fib[0] = 0
	fib[1] = 1

	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[n]
}

func main() {
	// Sample input #1
	n := 2
	fib := getNthFib(n)
	fmt.Println("Sample input #1:", n, "(0, 1)")
	fmt.Println("Sample output #1:", fib)

	// Sample input #2
	n = 6
	fib = getNthFib(n)
	fmt.Println("Sample input #2:", n, "(0, 1, 1, 2, 3, 5)")
	fmt.Println("Sample output #2:", fib)

	// Sample input #3
	n = 8
	fib = getNthFib(n)
	fmt.Println("Sample input #3:", n, "(0, 1, 1, 2, 3, 5, 8, 13)")
	fmt.Println("Sample output #3:", fib)
}
