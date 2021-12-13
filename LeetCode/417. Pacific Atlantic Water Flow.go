// https://leetcode.com/problems/pacific-atlantic-water-flow/

var (
    ROWS int
    COLS int
)

func pacificAtlantic(heights [][]int) [][]int {
    ROWS, COLS = len(heights), len(heights[0])
    pac, atl := make([][]int, len(heights)), make([][]int, len(heights))
    for r, _ := range heights {
        // pac and atl is marked of water from their side
        pac[r], atl[r] = make([]int, len(heights[0])), make([]int, len(heights[0])) 
    }
    
    for c:=0;c<COLS;c++ {
        // left side (pacific ocean side) 
        dfs(0, c, pac, heights[0][c], heights)
        
        // right side (atlantic ocean side)
        dfs(ROWS - 1, c, atl, heights[ROWS-1][c], heights)
    }
    
    for r:=0;r<ROWS;r++ {
        // up side (pacific ocean side) 
        dfs(r, 0, pac, heights[r][0], heights)
        
        // down side (atlantic ocean side)
        dfs(r, COLS-1, atl, heights[r][COLS-1], heights)
    }
    
    var result [][]int
    for r, _ := range heights {
        for c, _ := range heights[0] {
            // if match in those two place, return it
            if pac[r][c] * atl[r][c] == 1 {
                result = append(result, []int{r, c})
            }
        }
    }
    
    return result
}

func dfs(r int, c int, visit [][]int, prevHeight int, h [][]int) {
    // if out of boundaries or water cannot be come from prev position
    if r < 0 || c < 0 || r == ROWS || c == COLS || 
    h[r][c] < prevHeight {
        return
    }
    
    if visited := visit[r][c]; visited == 1 {
        return
    }
    
    visit[r][c] = 1
    
    // do dfs on 4 side
    dfs(r+1, c, visit, h[r][c], h)
    dfs(r-1, c, visit, h[r][c], h)
    dfs(r, c+1, visit, h[r][c], h)
    dfs(r, c-1, visit, h[r][c], h)
}
