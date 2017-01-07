package main

import (
	"log"
)

func main() {
	test()
}

// Function caesar apply a Caesar cypher with the key n to the rune r.
// E.g.: caesar('a', 3) → 'd'.
func caesar(r rune, n int) rune {
	// Turns a negative n into its equivalent modulo 26 positive n.
	// eg: -1 → 25
	if n<0 {
		n = 26+(n % 26)
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
		return int(r-'A')
	}
	if isLowercase(r) {
		return int(r-'a')
	}
	log.Fatal("offsetFromA needs a letter")
	return 0 // unreachable
}

func isUppercase(r rune) bool {
	return r>='A' && r<='Z'
}

func isLowercase(r rune) bool {
	return r>='a' && r<='z'
}

// Function mustTranslate returns true if r is a letter.
func mustTranslate(r rune) bool {
	for i:=0; i<len(validRunes); i++ {
		if r==validRunes[i] {
			return true
		}
	}
	return false
}

var validRunes = [...]rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

func test() {
	testMustTrasnslateOrdinaryLetters()
	testMustTranslateWeirdLetters()
	testCaesarZero()
	testCaesarNormal()
	testCaesarNegative()
	testCaesarBig()
	testCaesarNegativeBig()
}

func testCaesarZero() {
	inputs := [...]rune{'a', 'b', 'y', 'z', 'A', 'B', 'Y', 'Z'}
	for i:=0; i<len(inputs); i++ {
		input := inputs[i]
		expected := input
		obtained := caesar(input, 0)
		if obtained != expected {
			log.Printf("testCaesarZero failed: input=%q, expected=%q, obtained=%q",
				input, expected, obtained)
		}
	}
}

func testCaesarNormal() {
	inputs := [...]rune{'a', 'b', 'y', 'z', 'A', 'B', 'Y', 'Z'}
	expecteds := [...]rune{'d', 'e', 'b', 'c', 'D', 'E', 'B', 'C'}
	for i:=0; i<len(inputs); i++ {
		input := inputs[i]
		expected := expecteds[i]
		obtained := caesar(input, 3)
		if obtained != expected {
			log.Printf("testCaesarNormal failed: input=%q, expected=%q, obtained=%q",
				input, expected, obtained)
		}
	}
}

func testCaesarNegative() {
	inputs := [...]rune{'a', 'b', 'y', 'z', 'A', 'B', 'Y', 'Z'}
	expecteds := [...]rune{'x', 'y', 'v', 'w', 'X', 'Y', 'V', 'W'}
	for i:=0; i<len(inputs); i++ {
		input := inputs[i]
		expected := expecteds[i]
		obtained := caesar(input,-3)
		if obtained != expected {
			log.Printf("testCaesarNegative failed: input=%q, expected=%q, obtained=%q",
				input, expected, obtained)
		}
	}
}

func testCaesarBig() {
	inputs := [...]rune{'a', 'b', 'y', 'z', 'A', 'B', 'Y', 'Z'}
	expecteds := [...]rune{'n', 'o', 'l', 'm', 'N', 'O', 'L', 'M'} // 13
	for i:=0; i<len(inputs); i++ {
		input := inputs[i]
		expected := expecteds[i]
		obtained := caesar(input,5525) // 5525 % 26 = 13
		if obtained != expected {
			log.Printf("testCaesarBig failed: input=%q, expected=%q, obtained=%q",
				input, expected, obtained)
		}
	}
}

func testCaesarNegativeBig() {
	inputs := [...]rune{'a', 'b', 'y', 'z', 'A', 'B', 'Y', 'Z'}
	expecteds := [...]rune{'n', 'o', 'l', 'm', 'N', 'O', 'L', 'M'}
	for i:=0; i<len(inputs); i++ {
		input := inputs[i]
		expected := expecteds[i]
		obtained := caesar(input,-5525)
		if obtained != expected {
			log.Printf("testCaesarNegativeBig failed: input=%q, expected=%q, obtained=%q",
				input, expected, obtained)
		}
	}
}

func testMustTrasnslateOrdinaryLetters() {
	ok := [...]rune{'a', 'A', 'z', 'Z'}
	for i:=0; i<len(ok); i++ {
		if !mustTranslate(ok[i]) {
			log.Fatalf("testMustTranslateOrdinaryLetters failed: %q\n", ok[i])
		}
	}
}

func testMustTranslateWeirdLetters() {
	ko := [...]rune{'.', ',', ' ', '4', 'á'}
	for i:=0; i<len(ko); i++ {
		if mustTranslate(ko[i]) {
			log.Fatalf("testMustTranslatedWeirdLetters failed: %q\n", ko[i])
		}
	}
}

