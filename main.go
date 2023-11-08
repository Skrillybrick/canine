package canine

import (
    "fmt"

    "github.com/GoesToEleven/puppy"
    "canine/config"
)

func Bark() {
    fmt.Println(puppy.Bark())
}

func BetterBark() {
    s := puppy.Bark()
    fmt.Println(s, s, s)
}

func Version() {
    fmt.Println(config.Version)
}

