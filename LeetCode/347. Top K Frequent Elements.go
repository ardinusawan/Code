// https://leetcode.com/problems/top-k-frequent-elements/

func topKFrequent(nums []int, k int) []int {
    // bucket sort algoritm

    // for loop from last until hit k
    // create map, key is frequent, value is list of num. size is len of nums
        // nums = [1,1,1,2,2,3], k = 2
        // map[3:[1], 2:[2], 1:[3]]
    m := make(map[int][]int, len(nums))
    setNums := make(map[int]int)
    for _, v := range nums {
        setNums[v] += 1
    }
    
    // set of num - frequent
    for i, v := range setNums {
        m[v] = append(m[v], i)
    }
    
    var result []int
    // from biggest to lowest, append result
    for i:=len(nums);i>=0;i-- {
        for _, v := range m[i] {
            if len(result) == k {
                break
            }
            result = append(result, v)    
        }
    }
    return result
}
