// https://leetcode.com/problems/non-overlapping-intervals/

import (
    "sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
    result := 0
    
    // sort based on first value if same sort based on second value
    sort.Slice(intervals, func(a, b int) bool {
        for x := range intervals[a] {
            if intervals[a][x] == intervals[b][x] {
                continue
            }
            return intervals[a][x] < intervals[b][x]
        }
        return false
    })
    
    fmt.Println(intervals)
    for i:=1;i<len(intervals);i++ {
        // if overlap
        if intervals[i][0] >= intervals[i-1][0] && intervals[i][0] < intervals[i-1][1] {
            // we want to minimize overlap
            // so find the lowest second pair
            if intervals[i][1] > intervals[i-1][1] {
                // overwrite current value with prev value
                intervals[i] = intervals[i-1]    
            }
            
            result+=1
        }
    }

    return result
}
