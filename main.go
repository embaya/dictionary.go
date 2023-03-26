package main

import (
	dic "dictionary/dictionary"
	"fmt"
	"os"
)

func main() {
	d, err := dic.New("./badger")
	handleErr(err)
	defer d.Close()

	d.Add("php", "is a good langage")
	words, entries, _ := d.List()
	for _, word := range words {
		fmt.Println(entries[word])
	}

}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error: %v\n", err)
		os.Exit(1)
	}
}
