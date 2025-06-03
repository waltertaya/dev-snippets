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
 * Complete the 'fizzBuzz' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func fizzBuzz(n int32) {
    var i int32
    for i = 1; i <= n; i++ {
        // modulus of i to both 3 and 5
        result3 := i % 3
        result5 := i % 5
        // if i is mulitiple of both 3 and 5, print FizzBuzz
        if result3 == 0 && result5 == 0 {
            fmt.Println("FizzBuzz")
        }else if result3 == 0 && result5 != 0 {
            fmt.Println("Fizz")
        } else if result3 != 0 && result5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    fizzBuzz(n)
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
