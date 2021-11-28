// https://leetcode.com/problems/decode-ways/

var (
    memo map[int]int
    sInput string
)

func numDecodings(s string) int {
    memo = make(map[int]int) 
    sInput = s
    memo[len(s)] = 1
    return dfs(0)
}

func dfs(i int) int {
    var result int
    
    // create map with length of input
    if _, ok := memo[i]; ok {
        if memo[i] > 0 {
            fmt.Println("1", i)
            return memo[i]
        }
    }
    
    // input 0 means 0 ways
    if string(sInput[i]) == "0" {
        return result
    }
    
    // if input is 1 char, it will return 1 because we initialize on memo
    // if input is > 1 char, increment of(current result = 1 + memo[len(s)] = 1)
    result = dfs(i+1)
    // check if value is >= 10 and <= 26
    if (i+1 < len(sInput) && (string(sInput[i]) == "1" || string(sInput[i]) == "2" && string(sInput[i+1]) <= "6")) {
        // if it two value, combination will always 2
        result += dfs(i+2) 
    }
    
    // buttom-up
    memo[i] = result
    return result
}
