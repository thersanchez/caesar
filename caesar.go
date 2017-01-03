package main

import (
	"log"
	"math"
)

func main() {
	test()
}


func caesar(r rune, n int) rune {
	i := int(r) + n
	if i < 0 {
		log.Fatalf("cannot convert negative number to rune, was %d",
			i)
	}
	if i > math.MaxInt32 {
		log.Fatalf("cannot convert to rune: number too big, was %d",
			i)
	}
	return rune(i)
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

