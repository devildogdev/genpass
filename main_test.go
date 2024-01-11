package main

import (
    "testing"
)

func TestCreatePassword(t *testing.T) {
    err, password := CreatePassword(8)
    if len(password) != 8 && err == nil {
        t.Error("Password is not 8 characters!")
    }
}
