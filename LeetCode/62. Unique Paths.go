// https://leetcode.com/problems/unique-paths/

func uniquePaths(m int, n int) int {
    // buttom up
    
    // make matrix 2d
    matrix := make([][]int, m)
    for c:=0;c<m;c++{
        matrix[c] = make([]int, n)
    }
    
    for row:=m-1;row>=0;row--{
        for column:=n-1;column>=0;column--{
            // init target value as 1
            if row==m-1 && column==n-1{
                matrix[row][column] = 1
            } else {
                // use modules as trick if out of boundary value will become 0
                // initial value
                right := matrix[row][(column+1)%n]
                down := matrix[(row+1)%m][column]
                // from current position, posibility of reaching target is 
                // by goes down or right. So every move we have sub-problem.
                // if we only have 2x2 box, the value will be like this
                // [2][1]
                // [1][1]
                matrix[row][column] = right + down
            }
        }
    }
    return matrix[0][0]
}
