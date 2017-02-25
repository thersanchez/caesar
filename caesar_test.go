package caesar

import "testing"

func TestCaesar(t *testing.T) {
	input := "esther#.79: Alberto"
	n := 1
	expected := "ftuifs#.79: Bmcfsup"
	obtained := Caesar(input, n)
	if obtained != expected {
		t.Errorf("expected=%q, obtained=%q", expected,
			obtained)
	}
}

func TestCaesarEmptyInput(t *testing.T) {
	input := ""
	n := 1
	expected := ""
	obtained := Caesar(input, n)
	if obtained != expected {
		t.Errorf("expected=%q, obtained=%q", expected,
			obtained)
	}
}

func TestCaesarRuneZero(t *testing.T) {
	inputs := [...]rune{'a', 'b', 'y', 'z', 'A', 'B', 'Y', 'Z'}
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		expected := input
		obtained := CaesarRune(input, 0)
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
		obtained := CaesarRune(input, 3)
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
		obtained := CaesarRune(input, -3)
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
		obtained := CaesarRune(input, 5525) // 5525 % 26 = 13
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
		obtained := CaesarRune(input, -5525)
		if obtained != expected {
			t.Errorf("input=%q, expected=%q, obtained=%q",
				input, expected, obtained)
		}
	}
}

func TestCaesarRuneIgnoresWeirdRunes(t *testing.T) {
	weirds := [...]rune{'3', ' ', '@', '[', '`', '{', '¡'}
	for i := 0; i < len(weirds); i++ {
		if CaesarRune(weirds[i], 3) != weirds[i] {
			t.Errorf("%q\n", weirds[i])
		}
	}
}
