// https://leetcode.com/problems/spiral-matrix/

var up, down, left, right bool

func spiralOrder(matrix [][]int) []int {
    // create new matrix with size +1 around input matrix. Value is -101 and is act like boundary.
    // just like snake, and because it circular we can do:
    // if right is -101, it goes down. If snake already goes to that direction, fill with -101 (new boundary)
    // if down is -101, it goes left.
    // if left is -101, it goes up.
    // if up is -101, it goes right.
    
    boundary := -101
    newMatrix := make([][]int, len(matrix)+2)
    var result []int
    for i, _ := range newMatrix {
        newMatrix[i] = make([]int, len(matrix[0])+2)
        for ii, _ := range newMatrix[i] {
            newMatrix[i][ii] = boundary
        }
    }
    
    for row:=0;row<len(matrix);row++ {
        for col:=0;col<len(matrix[0]);col++ {
            newMatrix[row+1][col+1] = matrix[row][col]
        }
    }
    
    row := 1
    col := 1
    step := len(matrix)*len(matrix[0])
    resetDirection()
    for step>0 { 
        if newMatrix[row][col] != boundary {
            result = append(result, newMatrix[row][col])
            newMatrix[row][col] = boundary
            step--
            if up {
               row--
               up = true 
            } else if down {
               row++ 
               down = true
            } else if left {
               col--
               left = true
            } else {
               col++
               right = true
            } 
        } else {
            // check previous direction
            if right { // go down
                col--
                row++
                resetDirection()
                down = true
            } else if left { // go up 
                col++
                row--
                resetDirection()
                up = true
            } else if up { // go right
                col++
                row++
                resetDirection()
                right = true
            } else if down { // go left
                row--
                col--
                resetDirection()
                left = true
            }  
        }
    }
    
    return result
}

func resetDirection() {
    up = false
    down = false
    left = false
    right = false
}
