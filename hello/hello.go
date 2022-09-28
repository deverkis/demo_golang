package main

import (
    "fmt"
    "log"
    "example.com/greetings"
)

func main() {
    log.SetPrefix("greetings:")
    log.SetFlags(0)

    //fmt.Println("Hello, World!")
    //message := greetings.Hello("Gladys")
    //message, err := greetings.Hello("")
    message, err := greetings.Hello("Gladys")

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
}