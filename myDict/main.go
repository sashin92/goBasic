package main

import (
	"fmt"

	"github.com/sashin92/goBasic/myDict/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search("word")
	fmt.Println("Found :", word, "definition :", hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}

	err3 := dictionary.Delete(word)
	if err3 != nil {
		fmt.Println(err3)
	}
	hellow, _ := dictionary.Search(word)
	fmt.Println("Found :", word, "definition :", hellow)

}
