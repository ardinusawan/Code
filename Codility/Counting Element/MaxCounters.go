// Source : https://app.codility.com/programmers/lessons/4-counting_elements/max_counters/
// Inspired by : https://www.martinkysel.com/codility-maxcounters-solution/

package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
func CompareMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Solution(N int, A []int) []int {
	max_now := 0
	last_increment := 0

	arr := make([]int, N)

	for _, val := range A {
		if val == N+1 {
			last_increment = max_now
		} else {
			arr[val-1] = CompareMax(arr[val-1], last_increment)
			arr[val-1] += 1
			max_now = CompareMax(max_now, arr[val-1])
		}
	}
	for i, _ := range arr {
		arr[i] = CompareMax(last_increment, arr[i])
	}
	return arr
}
