// https://www.hackerrank.com/challenges/mini-max-sum/problem

func miniMaxSum(arr []int32) {    
    if len(arr) == 0 {
        fmt.Println(0, 0)
    }
    
    var total int64
    min, max := int64(arr[0]), int64(arr[0])
    for _, v := range arr {
        v := int64(v)
        total += v
        if min > v {
            min = v
        }
        if max < v {
            max = v
        }
    }
    fmt.Println(total - max, total - min)
}
