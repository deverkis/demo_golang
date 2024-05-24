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
	//test1()
	test2()
}

func test1() {
	messages, err := greetings.Hello("kis")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}

func test2() {
	names := []string{"Gladys", "Samantha", "Darrin"}

	//php $ss = ["a","b"];
	//go ss := []string{"a","b"}

	// i, value := range ss { }
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
