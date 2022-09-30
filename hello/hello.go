package main

import (
    "fmt"
    "log"
    "example.com/greetings"
    "github.com/google/go-cmp/cmp"
)

func main() {
    log.SetPrefix("greetings:")
    log.SetFlags(0)

    //fmt.Println("Hello, World!")
    //message := greetings.Hello("Gladys")
    //message, err := greetings.Hello("")
    //message, err := greetings.Hello("Gladys")
    names := []string{"Gladys","Samantha","Darrin"}
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }

    //fmt.Println(message)
    fmt.Println(messages)

    fmt.Println(cmp.Diff("Hello World", "Hello Go"))

}