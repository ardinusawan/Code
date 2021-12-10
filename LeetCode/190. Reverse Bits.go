// https://leetcode.com/problems/reverse-bits/
import (
    "strconv"
    "strings"
)

func reverseBits(num uint32) uint32 {
    // uint32 to str
    numBit := strconv.FormatUint(uint64(num), 2)
    var resultBit []string
    for i:=len(numBit)-1;i>=0;i--{
        resultBit = append(resultBit, string(numBit[i]))
    }

    // add more 0 to fill the space
    for i:=len(resultBit);i<32;i++{
        resultBit = append(resultBit, "0")
    }
    
    // convert to uint32 back
    strBit := strings.Join(resultBit, "")
    result, _ := strconv.ParseUint(strBit, 2, 32)
    return uint32(result)
}
