// https://leetcode.com/problems/word-search/

var (
    rLen int
    cLen int
    b [][]byte
    w string
    lastPosition = make(map[string]bool)
)

func exist(board [][]byte, word string) bool {
    // loop from [0][0] to [len(n)][len(m)]
    // find first char to match
    // if found, check up-right-down-left using dfs
    // save last state, follow the line
    // if not found continue the searching from beginning
    
    rLen = len(board)
    cLen = len(board[0])
    b = board
    w = word
    lastPosition = make(map[string]bool)
    
    for row:=0;row<len(board);row++ {
        for col:=0;col<len(board[row]);col++ {
            // find where first char and call dfs
            // start from that position
            if dfs(row, col, 0) {
                return true
            }
        }
    }
    
    return false
}

func dfs(r, c, i int) bool {
    if i == len(w) {
        return true
    }
    
    pos := fmt.Sprintf("%d:%d", r, c)
    if (r<0 || c<0 ||
        r>=rLen || c>=cLen ||
        b[r][c] != w[i] ||
        lastPosition[pos]) {
        return false
    }
    
    lastPosition[pos] = true
    res := (dfs(r-1, c, i+1) || //up
            dfs(r+1, c, i+1) || // down
            dfs(r, c+1, i+1) || // right
            dfs(r, c-1, i+1))   // left
    lastPosition[pos] = false
    return res
}
