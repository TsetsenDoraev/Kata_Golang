package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	Greeting := "\n\nДобро пожаловать в калькулятор" // Приветственная часть
	fmt.Println(Greeting)
	Delimiter := strings.Repeat("—", utf8.RuneCountInString(Greeting)-2)
	fmt.Println(Delimiter)

	var input string //Ввод значений и удаление пробелов
	fmt.Print("Введите выражение: ")
	fmt.Scanln(&input)
	input = strings.ReplaceAll(input, " ", "") //input = strings.Replace(input, " ", "", -1) есть ли разнциа помимо более широкго диапозона возможностей у Replace перед ReplaceAll?
	// не получается убирать пробелы до и после операнда, strings.TrimSpace()» тоже не помогает

	operators := [...]string{"+", "-", "*", "/"}

	var operator string // Валидация оператора
	for _, op := range operators {
		if strings.Contains(input, op) {
			operator = op
			break
		} else if op == operators[len(operators)-1] {
			fmt.Println("Оператор не найден.")
			return
		}
	}

	operands := strings.Split(input, operator) // Сепарация ввода на операнды
	if len(operands) != 2 {
		fmt.Println("Неверный формат выражения.")
		return
	}

	var num1, num2 float64 //Смена типа данных и присваивание
	_, err := fmt.Sscanf(operands[0], "%f", &num1)
	if err != nil {
		fmt.Println("Неверный формат первого операнда.")
		return
	}

	_, err = fmt.Sscanf(operands[1], "%f", &num2)
	if err != nil {
		fmt.Println("Неверный формат второго операнда.")
		return
	}

	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 { // Валидация операндов
		fmt.Println("Калькулятор принимает только значения от 1 до 10")
		return
	}

	var result float64 // Вариация математических операций
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Мы не в техническом ВУЗе")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Неверный оператор.")
		return
	}

	fmt.Printf("%.0f %s %.0f = %.0f\n", num1, operator, num2, result)
}
