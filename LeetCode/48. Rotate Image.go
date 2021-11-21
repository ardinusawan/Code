// https://leetcode.com/problems/rotate-image/

func rotate(matrix [][]int)  {
    // instead rotate to right, rotate to left (reverse). It only need to store top-left to temp variable. If rotate to right, we need to store 4 value of edge matrix.
    // create 4 pointer: left, right, up, down. left = 0 = top, right = len(matrix)-1 = buttom. We can combine it like matrix[top][left] = matrix[0][0], or matrix[buttom][right] = matrix[4][4]
    
    
    left := 0
    right := len(matrix) - 1
    
    for left < right {
        top := left
        buttom := right
        
        for i := 0; i < right-left; i++ { // distance between right and left will shrink because we have sub problem (matrix inside matrix)
            tempTopLeft := matrix[top][left+i]
            
            matrix[top][left+i] = matrix[buttom-i][left]
            
            matrix[buttom-i][left] = matrix[buttom][right-i]
            
            matrix[buttom][right-i] = matrix[top+i][right]
            
            matrix[top+i][right] = tempTopLeft
        }
        
        // shrink matrix to sub-problem
        left++
        right--
    }
}
