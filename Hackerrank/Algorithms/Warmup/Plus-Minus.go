https://www.hackerrank.com/challenges/plus-minus/problem

func plusMinus(arr []int32) {
    // Write your code here
    dict := make(map[int]float32, 3)
    for _, v := range arr {
        if v < 0 {
            dict[-1] += 1
        } else if v == 0 {
            dict[0] += 1
        } else {
            dict[1] += 1
        }
    }
    l := float32(len(arr))
    fmt.Printf("%.6f\n%.6f\n%.6f", dict[1]/l, dict[-1]/l, dict[0]/l)
}
