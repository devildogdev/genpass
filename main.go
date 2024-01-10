package main

import (
    "fmt"
    "strings"
    // "os"
    "flag"
    "math/rand"
)

const asciiOffset int = 33
var asciiLower = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var asciiUpper = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var asciiNum = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
var asciiSym = []string{"!", "\"", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",", "-", ".", "/", ":", ";", "<", ">", "?", "@", "[", "\\", "]", "_", "`"}
var length int
var upper, num, sym bool

func init() {
    flag.Usage = func() {
        fmt.Println("Genpass is a simple CLI for generating passwords.")
        fmt.Println("The default behavior only uses lowercase letters")
        flag.PrintDefaults()
    }
    flag.IntVar(&length, "l", 8, "Length of password")
    flag.BoolVar(&upper, "u", false, "Include uppercase letters")
    flag.BoolVar(&num, "d", false, "Include numbers")
    flag.BoolVar(&sym, "s", false, "Include symbols")
}

func createPassword(length int) string {
    buildChars := asciiLower
    passwordChars := make([]string, length)
    if upper {
        buildChars = append(buildChars, asciiUpper...)
    }
    if num {
        buildChars = append(buildChars, asciiNum...)
    }
    if sym {
        buildChars = append(buildChars, asciiSym...)
    }
    for i := 0; i < length; i++ {
        choice := rand.Intn(len(buildChars))
        char := buildChars[choice]
        passwordChars[i] = char
    }
    password := strings.Join(passwordChars, "")
    return password
}

func main() {
    flag.Parse()
    password := createPassword(length)
    fmt.Println(password)
}
