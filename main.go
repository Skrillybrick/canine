package canine

import (
    "fmt"

    "github.com/GoesToEleven/puppy"
)

func Bark() {
    fmt.Println(puppy.Bark())
}

func BetterBark() {
    s := puppy.Bark()
    fmt.Println(s, s, s)
}

func Version() {
    fmt.Println("v1.2.0")
}

