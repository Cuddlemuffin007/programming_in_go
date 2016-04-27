package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    who := "World"
    if len(os.Args) > 1 {/* os.Args[0] is 'hello' */
        who = strings.Join(os.Args[1:], " ")
    }
    fmt.Println("Hello", who)
    var age int = 28
    var favNum float64 = 7.777
    var 名 string = "ブレノン"
    fmt.Println(age, favNum, 名)
}
