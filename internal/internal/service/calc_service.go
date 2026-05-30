package service

func Soma(a, b int) int {
	return a + b
}

func Calc(x int) string {

	if x > 10 {
		return "Grande"
	}

	if x > 5 {
		return "Medio"
	}

	return "Pequeno"
}
