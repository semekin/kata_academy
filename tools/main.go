package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите значение")
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		s := strings.ReplaceAll(text, " ", "")
		s = strings.ToUpper(s)
		s = strings.TrimSpace(s)
		obrabotka(s)

	}
}
func obrabotka(s string) {
	var index int
	var roman1, roman2 string
	razdel := strings.Split(s, "")
	for s1, s2 := range razdel {
		if s2 == "/" || s2 == "*" || s2 == "+" || s2 == "-" {
			index = s1
		}
	}
	operatorIndex := razdel[index]                  //Индекс оператора
	result := (strings.Join(razdel[0:index], ""))   //разделяем срезы по индексу оператора слева
	result2 := (strings.Join(razdel[index+1:], "")) //разделяем срезы по индексу оператора справо

	resultInt, err := strconv.Atoi(result) //Преобразуем строковые значение среза в инт
	if err != nil {
		roman1 = result //Если значения римского формата оставляем оставляем строковое значение
	}
	resultInt2, err := strconv.Atoi(result2)
	if err != nil {
		roman2 = result2
	}

	if (resultInt > 0 && resultInt < 11) && (resultInt2 > 0 && resultInt2 < 11) { // условия задачи

		fmt.Printf("%d %s %d = %d \n", resultInt, operatorIndex, resultInt2, calc(resultInt, resultInt2, operatorIndex))
	} else {
		rome(roman1, roman2, operatorIndex)
	}

}
