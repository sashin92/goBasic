package main

import (
	"dictionary/myDict"
	"fmt"
)

func main() {
	dictionary := myDict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)

	dictionary.Delete(baseWord)
	word, err = dictionary.Search(baseWord)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(word)
	}

}

// add 에러처리 확인
// func main() {
// 	dictionary := myDict.Dictionary{"first": "First Word"}

// 	word := "hello"
// 	definition := "Greeting"
// 	err := dictionary.Add(word, definition)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	hello, _ := dictionary.Search(word)
// 	fmt.Println(hello)
// 	err2 := dictionary.Add(word, definition)
// 	if err2 != nil {
// 		fmt.Println(err2)
// 	}

// }
