package main

import (
	"log"
)

func main() {
	test()
}


func caesar(r rune, n int) rune {
	return 'A'
}

/*

func testCaesar() {
	if 
}
*/

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

