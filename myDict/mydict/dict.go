package mydict

import (
	"errors"
	"fmt"
)

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("not found")
var errWordExists = errors.New("already exist word")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		fmt.Println("value :", value, "exists :", exists)
		return value, nil
	}
	fmt.Println("value :", value, "exists :", exists)
	return "", errNotFound

}

// Add a word to the Dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil

}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		return errNotFound
	} else if err == nil {
		delete(d, word)
	}
	return nil
}
