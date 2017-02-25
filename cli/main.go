package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/thersanchez/caesar"
)

func main() {
	var text string
	var n int
	var err error

	switch len(os.Args) {
	case 2:
		text = os.Args[1]
		n = 3
	case 3:
		text = os.Args[2]
		n, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Cannot translate n: %s.\n", err)
			usage()
			os.Exit(1)
		}
	default:
		fmt.Println("Wrong number of arguments.")
		usage()
		os.Exit(1)
	}

	fmt.Println(caesar.String(text, n))
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println("\tcaesar [n] text")
}
