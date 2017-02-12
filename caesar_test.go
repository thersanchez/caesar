package main

import "testing"

func TestCaesarRuneZero(t *testing.T) {
	inputs := [...]rune{'a', 'b', 'y', 'z', 'A', 'B', 'Y', 'Z'}
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		expected := input
		obtained := caesarRune(input, 0)
		if obtained != expected {
			t.Errorf("input=%q, expected=%q, obtained=%q\n",
				input, expected, obtained)

		}
	}
}

func TestCaesarRuneThree(t *testing.T) {
	inputs := [...]rune{'a', 'b', 'y', 'z', 'A', 'B', 'Y', 'Z'}
	expecteds := [...]rune{'d', 'e', 'b', 'c', 'D', 'E', 'B', 'C'}
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		expected := expecteds[i]
		obtained := caesarRune(input, 3)
		if obtained != expected {
			t.Errorf("input=%q, expected=%q, obtained=%q",
				input, expected, obtained)
		}
	}
}

func TestCaesarRuneNegative(t *testing.T) {
	inputs := [...]rune{'a', 'b', 'y', 'z', 'A', 'B', 'Y', 'Z'}
	expecteds := [...]rune{'x', 'y', 'v', 'w', 'X', 'Y', 'V', 'W'}
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		expected := expecteds[i]
		obtained := caesarRune(input, -3)
		if obtained != expected {
			t.Errorf("input=%q, expected=%q, obtained=%q",
				input, expected, obtained)
		}
	}
}

func TestCaesarRuneBig(t *testing.T) {
	inputs := [...]rune{'a', 'b', 'y', 'z', 'A', 'B', 'Y', 'Z'}
	expecteds := [...]rune{'n', 'o', 'l', 'm', 'N', 'O', 'L', 'M'} // 13
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		expected := expecteds[i]
		obtained := caesarRune(input, 5525) // 5525 % 26 = 13
		if obtained != expected {
			t.Errorf("input=%q, expected=%q, obtained=%q",
				input, expected, obtained)
		}
	}
}

func TestCaesarRuneNegativeBig(t *testing.T) {
	inputs := [...]rune{'a', 'b', 'y', 'z', 'A', 'B', 'Y', 'Z'}
	expecteds := [...]rune{'n', 'o', 'l', 'm', 'N', 'O', 'L', 'M'}
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		expected := expecteds[i]
		obtained := caesarRune(input, -5525)
		if obtained != expected {
			t.Errorf("input=%q, expected=%q, obtained=%q",
				input, expected, obtained)
		}
	}
}

func TestMustCypherOrdinaryLetters(t *testing.T) {
	ok := [...]rune{'A', 'P', 'Z', 'a', 'p', 'z'}
	for i := 0; i < len(ok); i++ {
		if !mustCypher(ok[i]) {
			t.Errorf("%q\n", ok[i])
		}
	}
}

func TestMustCypherWeirdLetters(t *testing.T) {
	ko := [...]rune{'3', '@', '[', '`', '{', '¡'}
	for i := 0; i < len(ko); i++ {
		if mustCypher(ko[i]) {
			t.Errorf("%q\n", ko[i])
		}
	}
}
