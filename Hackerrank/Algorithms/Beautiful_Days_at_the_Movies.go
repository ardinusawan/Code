// https://www.hackerrank.com/challenges/beautiful-days-at-the-movies/problem
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func reverseNumber(number int) int {
	reverse := 0

	for number > 0 {
		lastDigit := number % 10
		reverse = (reverse * 10) + lastDigit
		number = number / 10
	}
	return reverse
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var a, b, c int

	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &a, &b, &c)
	total := 0
	for i := a; i <= b; i++ {
		temp := math.Abs(float64(i - reverseNumber(i)))
		var y int = int(temp)
		if y%c == 0 {
			total++
		}
		// fmt.Println(math.Abs(float64(i-reverseNumber(i))))
	}
	fmt.Print(total)
}
