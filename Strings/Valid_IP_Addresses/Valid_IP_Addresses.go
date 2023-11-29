/*


You're given a string of length 12 or smaller, containing only digits. Write a function that returns all the possible IP addresses that can be created by inserting three . sin the string.
An IP address is a sequence of four positive integers that are separated by . s, where each individual integer is within the range 0 - 255 , inclusive.
An IP address isn't valid if any of the individual integers contains leading e s. For example, "192.168.0.1 " is a valid IP
address, but "192. 168.00.1" and " 192.168. 0. 01" are not valid IP addresses, because they contain "0" as the first integer in a group.
example of a valid IP address is "99 .1.1.10" ; conversely, "991.1.1. 0" isn't valid, because "991" is greater than 255.
Your function should return the IP addresses in string format and in no particular order. If no valid IP addresses can be created from the string, your function should return an empty list.
Note: check out our Systems Design Fundamentals on Systems Expert to learn more about IP addresses!
For this problem, you can assume that input string will always be valid.



Sample Input:
string = "1921680"


Sample Output:
[
  "
    1. 9. 216. 80",
    "
    1. 92. 16. 80",
    "
    1. 92. 168. 0",
    "
    19. 2. 16. 80",
    "
    19. 2. 168. 0",
    "
    19. 21. 6. 80",
    "
    19. 21. 68. 0",
    "
    19. 216. 8. 0",
    "
    192. 1. 6. 80",
    "
    192. 1. 68. 0",
    "
    192. 16. 8. 0"
    ]
    // The IP addresses could be ordered differently.




Hints
Hint 1

How can you split this problem into subproblems to make it easier?

Hint 2

Each IP address is comprised of four parts; try finding one valid part at a time and then combining sets of four valid parts to create one valid IP address.

Hint 3

Go through all possible combinations of valid IP-address parts. You'll do this by generating a valid first part, then generating all valid second parts given the first part, then finally all valid third and fourth parts given first and second parts. If you find a set of four valid parts, then simply combine them together and add that IP address to some final array. You can start by creating all the possible first parts of an IP address; these will be substrings of the main string that start at the first character and that have lengths 1, 2 and 3. Then you can repeat this process for the second part, where the substrings in this part will start where the first part ended. The same thing applies for the third and fourth parts. After going through all possible parts and storing valid IP addresses, you'll have found all of the IP addresses that can be formed from the input string.





This code first generates all possible combinations of three dot placements within the given string. Then, for each combination, it checks if the resulting string is a valid IP address using the isValidIP function. If it is, it appends it to the list of valid IP addresses. Finally, it returns the list of valid IP addresses.



*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func generateIPAddresses(str string) []string {
	validIPs := []string{}

	for i := 1; i < len(str)-3; i++ {
		for j := i + 1; j < len(str)-2; j++ {
			for k := j + 1; k < len(str)-1; k++ {
				ip := str[:i] + "." + str[i:j] + "." + str[j:k] + "." + str[k:]

				if isValidIP(ip) {
					validIPs = append(validIPs, ip)
				}
			}
		}
	}

	return validIPs
}

func isValidIP(ipAddress string) bool {
	parts := strings.Split(ipAddress, ".")

	if len(parts) != 4 {
		return false
	}

	for _, part := range parts {
		if num, err := strconv.Atoi(part); err != nil || num < 0 || num > 255 {
			return false
		}

		if len(part) > 1 && part[0] == '0' {
			return false
		}
	}

	return true
}

func main() {
	inputString := "1921680"
	validIPs := generateIPAddresses(inputString)

	for _, ip := range validIPs {
		fmt.Println(ip)
	}
}
