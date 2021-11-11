// https://leetcode.com/problems/longest-common-subsequence/

import (
    "math"
)
func longestCommonSubsequence(text1 string, text2 string) int {
    // create matrix t1 * t2
    // fill right and down with 0 as bounday
    matrix := make([][]int, len(text1)+1)
    for i:=len(text1); i>=0;i-- {
        matrix[i] = make([]int, len(text2)+1)
    }
    for i:=len(text1)-1; i>=0;i-- {
        for j:=len(text2)-1; j>=0;j-- {
            if text1[i] == text2[j] {
                // if char match, set cell as 1.
                // But we can optimize calculation add check of diagonal cell.
                // If diagonal cell value is 1, then add with 1 on existing value.
                matrix[i][j] = int(math.Max(float64(1), float64(1 + matrix[i+1][j+1])))
            } else {
                // Do check right and below cell.
                // If value is 1, make existing cell as 1.
                // So value is move from buttom to higher up.
                matrix[i][j] = int(math.Max(float64(matrix[i+1][j]), float64(matrix[i][j+1])))
            }
        }
    }
    return matrix[0][0]
}
