package main

import (
    "fmt"
    "strings"
    "errors"
    "flag"
    "math/rand"
)

const asciiOffset int = 33
var asciiLower = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var asciiUpper = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var asciiNum = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
var asciiSym = []string{"!", "\"", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",", "-", ".", "/", ":", ";", "<", ">", "?", "@", "[", "\\", "]", "_", "`"}
var length int
var upper, num, sym, all bool

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
    flag.BoolVar(&all, "a", false, "Include any ascii characters")
}

func createPassword(length int) (error, string) {
    if length <= 0 {
        return errors.New("Error: Length must be greater than 0"), ""
    }
    buildChars := asciiLower
    passwordChars := make([]string, length)
    if all {
        buildChars = append(buildChars, asciiUpper...)
        buildChars = append(buildChars, asciiNum...)
        buildChars = append(buildChars, asciiSym...)
    }
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
    return nil, password
}

func main() {
    flag.Parse()
    err, password := createPassword(length)
    if err != nil {
        fmt.Println(err)
        flag.PrintDefaults()
    } else {
        fmt.Println(password)
    }
}
