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

	var nums []float64
	for i := 0; i <= len(operands)-1; i++ {
		var num float64
		_, err := fmt.Sscanf(operands[i], "%f", &num)
		if err != nil {
			fmt.Printf("Неверный формат %v-го операнда.\n", i+1)
			return
		}
		if num < 1 || num > 10 {
			fmt.Println("Калькулятор принимает только значения от 1 до 10")
			return
		}
		nums = append(nums, num)
	}

	var result float64 // Пока не понял, как выставить приоритет первостепенности умножения и деления перед сложением и вычитанием
	switch operator {
	case "*":
		result = 1
		for _, num := range nums {
			result *= num
		}
	case "/":
		for i, num := range nums {
			if i == 0 {
				result = num
			} else if num == 0 {
				fmt.Println("Мы не в техническом ВУЗе")
				return
			} else {
				result /= num
			}
		}
	case "+":
		for _, num := range nums {
			result += num
		}
	case "-":
		for i, num := range nums {
			if i == 0 {
				result = num
			} else {
				result -= num
			}
		}
	}

	inputWithResult := strings.Join(operands, operator) + " = " + fmt.Sprintf("%g", result)

	fmt.Println(inputWithResult)
}
