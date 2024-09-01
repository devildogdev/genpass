package main

import (
    "fmt"
    "strings"
    "errors"
    "flag"
    "math/rand"
    "github.com/devildogdev/genpass/internal/chars"
)

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

func CreatePassword(length int) (string, error) {
    if length <= 0 {
        return "", errors.New("Error: Length must be greater than 0")
    }
    buildChars := chars.AsciiLower
    passwordChars := make([]string, length)
    if all {
        buildChars = append(buildChars, chars.AsciiUpper...)
        buildChars = append(buildChars, chars.AsciiNum...)
        buildChars = append(buildChars, chars.AsciiSym...)
    }
    if upper {
        buildChars = append(buildChars, chars.AsciiUpper...)
    }
    if num {
        buildChars = append(buildChars, chars.AsciiNum...)
    }
    if sym {
        buildChars = append(buildChars, chars.AsciiSym...)
    }
    for i := 0; i < length; i++ {
        choice := rand.Intn(len(buildChars))
        char := buildChars[choice]
        passwordChars[i] = char
    }
    password := strings.Join(passwordChars, "")
    return password, nil
}

func main() {
    flag.Parse()
    password, err := CreatePassword(length)
    if err != nil {
        fmt.Println(err)
        flag.PrintDefaults()
    } else {
        fmt.Println(password)
    }
}
