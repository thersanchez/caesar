package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	test()

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
	// therefore we cypher the offset of the letter from it base ('A' or 'a')
	// and apply the cyphered offset back to the base at the end.
	offset := offsetFromA(r)
	cypheredOffset := (offset + n) % 26

	base := 'a'
	if isUppercase(r) {
		base = 'A'
	}
	return base + rune(cypheredOffset)
}

// Function offsetFromA of a rune r returns the distance between r and 'A' (or 'a').
// If r is not a letter, the program exits in error.
func offsetFromA(r rune) int {
	if isUppercase(r) {
		return int(r - 'A')
	}
	if isLowercase(r) {
		return int(r - 'a')
	}
	log.Fatal("offsetFromA needs a letter")
	return 0 // unreachable
}

func isUppercase(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isLowercase(r rune) bool {
	return r >= 'a' && r <= 'z'
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

func test() {
	testMustCypherOrdinaryLetters()
	testMustCypherWeirdLetters()
	testCaesarRuneZero()
	testCaesarRuneThree()
	testCaesarRuneNegative()
	testCaesarRuneBig()
	testCaesarRuneNegativeBig()
}

func testCaesarRuneZero() {
	inputs := [...]rune{'a', 'b', 'y', 'z', 'A', 'B', 'Y', 'Z'}
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		expected := input
		obtained := caesarRune(input, 0)
		if obtained != expected {
			log.Printf("testCaesarZero failed: input=%q, expected=%q, obtained=%q",
				input, expected, obtained)
		}
	}
}

func testCaesarRuneThree() {
	inputs := [...]rune{'a', 'b', 'y', 'z', 'A', 'B', 'Y', 'Z'}
	expecteds := [...]rune{'d', 'e', 'b', 'c', 'D', 'E', 'B', 'C'}
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		expected := expecteds[i]
		obtained := caesarRune(input, 3)
		if obtained != expected {
			log.Printf("testCaesarNormal failed: input=%q, expected=%q, obtained=%q",
				input, expected, obtained)
		}
	}
}

func testCaesarRuneNegative() {
	inputs := [...]rune{'a', 'b', 'y', 'z', 'A', 'B', 'Y', 'Z'}
	expecteds := [...]rune{'x', 'y', 'v', 'w', 'X', 'Y', 'V', 'W'}
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		expected := expecteds[i]
		obtained := caesarRune(input, -3)
		if obtained != expected {
			log.Printf("testCaesarNegative failed: input=%q, expected=%q, obtained=%q",
				input, expected, obtained)
		}
	}
}

func testCaesarRuneBig() {
	inputs := [...]rune{'a', 'b', 'y', 'z', 'A', 'B', 'Y', 'Z'}
	expecteds := [...]rune{'n', 'o', 'l', 'm', 'N', 'O', 'L', 'M'} // 13
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		expected := expecteds[i]
		obtained := caesarRune(input, 5525) // 5525 % 26 = 13
		if obtained != expected {
			log.Printf("testCaesarBig failed: input=%q, expected=%q, obtained=%q",
				input, expected, obtained)
		}
	}
}

func testCaesarRuneNegativeBig() {
	inputs := [...]rune{'a', 'b', 'y', 'z', 'A', 'B', 'Y', 'Z'}
	expecteds := [...]rune{'n', 'o', 'l', 'm', 'N', 'O', 'L', 'M'}
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		expected := expecteds[i]
		obtained := caesarRune(input, -5525)
		if obtained != expected {
			log.Printf("testCaesarNegativeBig failed: input=%q, expected=%q, obtained=%q",
				input, expected, obtained)
		}
	}
}

func testMustCypherOrdinaryLetters() {
	ok := [...]rune{'A', 'P', 'Z', 'a', 'p', 'z'}
	for i := 0; i < len(ok); i++ {
		if !mustCypher(ok[i]) {
			log.Fatalf("testMustTranslateOrdinaryLetters failed: %q\n", ok[i])
		}
	}
}

func testMustCypherWeirdLetters() {
	ko := [...]rune{'3', '@', '[', '`', '{', '¡'}
	for i := 0; i < len(ko); i++ {
		if mustCypher(ko[i]) {
			log.Fatalf("testMustTranslatedWeirdLetters failed: %q\n", ko[i])
		}
	}
}
