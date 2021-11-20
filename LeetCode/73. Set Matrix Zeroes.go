// https://leetcode.com/problems/set-matrix-zeroes

func setZeroes(matrix [][]int)  {
    // use existing column and row as flag is this row/column is 0, except first row(collide with first column). So we only init 1 variable
    // first for loop, fill column and row with zero if 0 exist
    // second for loop, fill 0 in data exept in first column and row
    // check in first column, if 0 exist than fill column with 0.
    // check in rowZero flag, if 0 exist than fill row with 0.
    
    rowZero := 0
    for r, v := range matrix {
        for c, _ := range v {
            if matrix[r][c] == 0 {
                matrix[0][c] = 0
                if r > 0 {
                    matrix[r][0] = 0
                } else {
                    rowZero = 1
                }
            }    
        }
    }
    
    for r:=1; r<len(matrix); r++ {
        for c:=1; c<len(matrix[r]); c++ {
            if matrix[0][c] == 0 || matrix[r][0] == 0 {
                matrix[r][c] = 0
            }
        }
    }
    
    if matrix[0][0] == 0 {
        for c, _ := range(matrix) {
            matrix[c][0] = 0
        }
    }
    
    if rowZero == 1{
        for r, _ := range matrix[0] {
            matrix[0][r] = 0
        }
    }
}
