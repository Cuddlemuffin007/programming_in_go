package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

var bigDigits = [][]string{
    {"  000  ", " 0   0 ", "0     0", "0     0", "0     0", " 0   0 ", "  000  "},
    {" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
    {"  222  ", " 2   2 ", "    2  ", "   2   ", " 2     ", "2      ", "2222222"},
}

func main() {
    if len(os.Args) == 1 {
        fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
        os.Exit(1)
    }

    stringOfDigits := os.Args[1]
    for row := range bigDigits[0] {
        line := ""
        for col := range stringOfDigits {
            digit := stringOfDigits[col] - '0'
            if 0 <= digit && digit <= 2 {
                line += bigDigits[digit][row] + "  "
            } else {
                log.Fatal("Invalid whole number")
            }
        }
        fmt.Println(line)
    }
}
