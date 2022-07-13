func breakingRecords(scores []int32) []int32 {
    var max, min int32
    max, min = -1, 100000001
    result := []int32{-1, -1}
    for _, v := range scores {
        if max < v {
            max = v
            result[0] += 1 
        }
        if min > v {
            min = v
            result[1] += 1
        }
    }
    return result
}
