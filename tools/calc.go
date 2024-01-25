package main

func calc(firstNumb, secNumb int, operator string) int {
	switch operator {
	case "+":
		return firstNumb + secNumb
	case "-":
		return firstNumb - secNumb
	case "*":
		return firstNumb * secNumb
	case "/":
		if secNumb == 0 {
			panic("Деление на ноль невозможно")
		}
		return firstNumb / secNumb
	default:
		panic("Недопустимая операция")

	}

}
