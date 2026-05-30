
package service

import "testing"

func TestSoma(t *testing.T) {

	result := Soma(2, 3)

	expected := 5

	if result != expected {

		t.Errorf(
			"expected %d got %d",
			expected,
			result,
		)
	}
}

func TestSomaNegative(t *testing.T) {

	result := Soma(-1, 1)

	expected := 0

	if result != expected {

		t.Errorf(
			"expected %d got %d",
			expected,
			result,
		)
	}
}

func TestSomaZero(t *testing.T) {

	result := Soma(0, 0)

	expected := 0

	if result != expected {

		t.Errorf(
			"expected %d got %d",
			expected,
			result,
		)
	}
}

func TestCalcGrande(t *testing.T) {

	result := Calc(11)

	expected := "Grande"

	if result != expected {

		t.Errorf(
			"expected %s got %s",
			expected,
			result,
		)
	}
}

func TestCalcMedio(t *testing.T) {

	result := Calc(7)

	expected := "Medio"

	if result != expected {

		t.Errorf(
			"expected %s got %s",
			expected,
			result,
		)
	}
}

func TestCalcPequeno(t *testing.T) {

	result := Calc(2)

	expected := "Pequeno"

	if result != expected {

		t.Errorf(
			"expected %s got %s",
			expected,
			result,
		)
	}
}

func TestCalcBoundary10(t *testing.T) {

	result := Calc(10)

	expected := "Medio"

	if result != expected {

		t.Errorf(
			"expected %s got %s",
			expected,
			result,
		)
	}
}

func TestCalcBoundary5(t *testing.T) {

	result := Calc(5)

	expected := "Pequeno"

	if result != expected {

		t.Errorf(
			"expected %s got %s",
			expected,
			result,
		)
	}
}