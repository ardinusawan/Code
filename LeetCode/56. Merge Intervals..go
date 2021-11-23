// https://leetcode.com/problems/merge-intervals/

// sort intervals based on first value i.e [[2, 3], [1, 5]] = [[1, 5], [2, 3]]
// create result [][]int, insert first intervals to result
// for loop from index one, compare last data on result with current intervals
// [left, right]. if right of last result >= left data of current interval
    // merge
// else append

import (
    "math"
)

func merge(intervals [][]int) [][]int {
    // sort intervals
    sort.Slice(intervals[:], func(i, j int) bool {
	    for x := range intervals[i] {
	        if intervals[i][x] == intervals[j][x] {
	            continue
	        }
	        return intervals[i][x] < intervals[j][x]
	    }
		return false
	})

    var result [][]int
    i := 1
    result = append(result, intervals[0])
    for i < len(intervals) {
        if result[len(result)-1][1] >= intervals[i][0] {
            // get min and max of 2 list to merge
            minLeft := int(math.Min(float64(result[len(result)-1][0]), float64(intervals[i][0])))
            maxRight := int(math.Max(float64(result[len(result)-1][1]), float64(intervals[i][1])))
            merge := []int{minLeft, maxRight}
            result[len(result)-1] = merge
        } else {
            result = append(result, intervals[i]) 
        }
        
        i++
    }
    return result
}
