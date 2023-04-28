package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func toInt(s string) int {
    number, _ := strconv.Atoi(s)
    return number
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    val := toInt(scanner.Text())
    fmt.Println(abs(val))
}
