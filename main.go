package main

import (
	dic "dictionary/dictionary"
	"flag"
	"fmt"
	"os"
)

func main() {
	action := flag.String("action", "list", "action to perform n dectionary")

	d, err := dic.New("./badger")
	handleErr(err)
	defer d.Close()

	flag.Parse()

	switch *action {
	case "list":
		actionList(d)
	case "add":
		actionAdd(d, flag.Args())
	case "define":
		actionDefine(d, flag.Args())
	case "remove":
		actionRemove(d, flag.Args())
	default:
		fmt.Printf("unknown action : %v\n", *action)
	}

	d.Add("php", "is a good langage")

}

func actionRemove(d *dic.Dictionary, args []string) {
	word := args[0]
	err := d.Remove(word)
	handleErr(err)
	fmt.Printf("'%v' was removed from the dictionary\n", word)
}

func actionDefine(d *dic.Dictionary, args []string) {
	word := args[0]
	entry, err := d.Get(word)
	handleErr(err)
	fmt.Println(entry)
}

func actionAdd(d *dic.Dictionary, args []string) {
	word := args[0]
	defintion := args[1]
	err := d.Add(word, defintion)
	handleErr(err)
	fmt.Printf("%v added to dictionary\n", word)
}

func actionList(d *dic.Dictionary) {
	words, entries, err := d.List()
	handleErr(err)
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
