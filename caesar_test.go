package caesar

import "testing"

func TestString(t *testing.T) {
	input := "esther#.79: Alberto"
	n := 1
	expected := "ftuifs#.79: Bmcfsup"
	obtained := String(input, n)
	if obtained != expected {
		t.Errorf("expected=%q, obtained=%q", expected,
			obtained)
	}
}

func TestEmptyString(t *testing.T) {
	input := ""
	n := 1
	expected := ""
	obtained := String(input, n)
	if obtained != expected {
		t.Errorf("expected=%q, obtained=%q", expected,
			obtained)
	}
}

func TestRuneZero(t *testing.T) {
	inputs := [...]rune{'a', 'b', 'y', 'z', 'A', 'B', 'Y', 'Z'}
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		expected := input
		obtained := Rune(input, 0)
		if obtained != expected {
			t.Errorf("input=%q, expected=%q, obtained=%q\n",
				input, expected, obtained)

		}
	}
}

func TestRuneThree(t *testing.T) {
	inputs := [...]rune{'a', 'b', 'y', 'z', 'A', 'B', 'Y', 'Z'}
	expecteds := [...]rune{'d', 'e', 'b', 'c', 'D', 'E', 'B', 'C'}
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		expected := expecteds[i]
		obtained := Rune(input, 3)
		if obtained != expected {
			t.Errorf("input=%q, expected=%q, obtained=%q",
				input, expected, obtained)
		}
	}
}

func TestRuneNegative(t *testing.T) {
	inputs := [...]rune{'a', 'b', 'y', 'z', 'A', 'B', 'Y', 'Z'}
	expecteds := [...]rune{'x', 'y', 'v', 'w', 'X', 'Y', 'V', 'W'}
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		expected := expecteds[i]
		obtained := Rune(input, -3)
		if obtained != expected {
			t.Errorf("input=%q, expected=%q, obtained=%q",
				input, expected, obtained)
		}
	}
}

func TestRuneBig(t *testing.T) {
	inputs := [...]rune{'a', 'b', 'y', 'z', 'A', 'B', 'Y', 'Z'}
	expecteds := [...]rune{'n', 'o', 'l', 'm', 'N', 'O', 'L', 'M'} // 13
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		expected := expecteds[i]
		obtained := Rune(input, 5525) // 5525 % 26 = 13
		if obtained != expected {
			t.Errorf("input=%q, expected=%q, obtained=%q",
				input, expected, obtained)
		}
	}
}

func TestRuneNegativeBig(t *testing.T) {
	inputs := [...]rune{'a', 'b', 'y', 'z', 'A', 'B', 'Y', 'Z'}
	expecteds := [...]rune{'n', 'o', 'l', 'm', 'N', 'O', 'L', 'M'}
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		expected := expecteds[i]
		obtained := Rune(input, -5525)
		if obtained != expected {
			t.Errorf("input=%q, expected=%q, obtained=%q",
				input, expected, obtained)
		}
	}
}

func TestRuneIgnoresWeirdRunes(t *testing.T) {
	weirds := [...]rune{'3', ' ', '@', '[', '`', '{', '¡'}
	for i := 0; i < len(weirds); i++ {
		if Rune(weirds[i], 3) != weirds[i] {
			t.Errorf("%q\n", weirds[i])
		}
	}
}
