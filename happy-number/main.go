package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		d := n % 10
		sum += d * d
		n /= 10
	}
	return sum
}

func isHappy(n int) bool {
	seen := make(map[int]bool)
	for n != 1 && !seen[n] {
		seen[n] = true
		n = sumOfSquares(n)
	}
	return n == 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		if isHappy(num) {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}
