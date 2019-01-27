// Source : https://www.hackerrank.com/challenges/breaking-best-and-worst-records/problem

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the breakingRecords function below.
 */
func breakingRecords(score []int32) []int32 {
    /*
     * Write your code here.
     */
    
    min := score[0]
    max := score[0]
    var min_c, max_c int32
    for _ , v := range score {
        if(v<min) {
            min_c+=1
            min = v
        }
        if(v>max) {
            max_c+=1
            max = v
        }
    }
    return []int32{max_c, min_c}
    //return make([]int32,(max, min))
    
  
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer outputFile.Close()

    writer := bufio.NewWriterSize(outputFile, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    scoreTemp := strings.Split(readLine(reader), " ")

    var score []int32

    for scoreItr := 0; scoreItr < int(n); scoreItr++ {
        scoreItemTemp, err := strconv.ParseInt(scoreTemp[scoreItr], 10, 64)
        checkError(err)
        scoreItem := int32(scoreItemTemp)
        score = append(score, scoreItem)
    }

    result := breakingRecords(score)

    for resultItr, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if resultItr != len(result) - 1 {
            fmt.Fprintf(writer, " ")
        }
    }

    fmt.Fprintf(writer, "\n")

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

