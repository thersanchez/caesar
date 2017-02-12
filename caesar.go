package main

import (
	"fmt"
	"os"
	"strconv"
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

	fmt.Println(caesar(text, n))
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println("\tcaesar [n] text")
}

func caesar(s string, n int) string {
	ret := make([]rune, len(s))
	for i, r := range s {
		ret[i] = caesarRune(r, n)
	}
	return string(ret)
}

// Function caesarRune apply a Caesar cypher with the key n to the letter r.
// E.g.: caesarRune('a', 3) → 'd'.
// If r is not a letter, the rune is returned unchanged.
func caesarRune(r rune, n int) rune {
	if !mustCypher(r) {
		return r
	}

	// Turns a negative n into its equivalent modulo 26 positive n.
	// eg: -1 → 25
	if n < 0 {
		n = 26 + (n % 26)
	}

	// Cyphering is independent of the input letter case,
	// therefore we cypher the offset of the letter from its base ('A' or 'a')
	// and apply the cyphered offset back to the base at the end.
	index := alphabeticIndex(r)
	cypheredIndex := (index + n) % 26

	base := 'a'
	if isUppercase(r) {
		base = 'A'
	}
	return base + rune(cypheredIndex)
}

// Function alphabeticIndex returns the Unicode distance between r and 'A' (or
// 'a') or an unspecified number if r is not a letter.
func alphabeticIndex(r rune) int {
	if isUppercase(r) {
		return int(r - 'A')
	}
	return int(r - 'a')
}

func isUppercase(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func mustCypher(r rune) bool {
	return isLetter(r)
}

func isLetter(r rune) bool {
	switch {
	case r >= 'A' && r <= 'Z':
		return true
	case r >= 'a' && r <= 'z':
		return true
	default:
		return false
	}
}
