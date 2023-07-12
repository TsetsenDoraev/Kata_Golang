package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	// Приветственная часть
	Greeting := "\n\nДобро пожаловать в калькулятор"
	fmt.Println(Greeting)
	Delimiter := strings.Repeat("—", utf8.RuneCountInString(Greeting)-2)
	fmt.Println(Delimiter)

	// Запрос выражения у пользователя
	fmt.Println("Введите выражение: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Проверка, содержит ли ввод римские цифры
	var isRoman bool
	if strings.ContainsAny(input, "IVX") {
		isRoman = true
	}

	// Определение оператора в выражении для Split и Switch
	var operator string
	operators := [...]string{"+", "-", "*", "/"}
	for _, op := range operators {
		if strings.Contains(input, op) {
			operator = op
			break
		} else if op == operators[len(operators)-1] {
			fmt.Println("Оператор не найден.")
			return
		}
	}

	// Сепарация ввода на операнды по оператору
	operands := strings.Split(input, operator)
	if len(operands) > 2 {
		fmt.Println("Неверный формат выражения, не более двух операндов в выражении")
		return
	}
	if len(operands) < 2 {
		fmt.Println("Cтрока не является математической операцией")
		return
	}

	// Смена типа данных и присваивание
	var num1, num2 float64
	if isRoman {
		num1 = float64(Roman2Arabic(operands[0]))
		num2 = float64(Roman2Arabic(operands[1]))
	} else {
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
	}
	// Валидация операндов
	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		fmt.Println("Калькулятор принимает только значения от 1 до 10")
		return
	}
	// Определение математической операции и валидация оператора — Switch
	var result float64
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
	// Отображение результата
	if isRoman {
		if result < 0 {
			fmt.Println("В римской системе нет отрицательных чисел")
		} else {
			fmt.Printf("%s\n", Arabic2Roman(float64(result)))
		}
	} else {
		fmt.Printf("%.0f %s %.0f = %.0f\n", num1, operator, num2, result)
	}
}

// Вычисляем арабское значение римской цифры

func Roman2Arabic(num string) int {
	result := 0
	DicR2A := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
	}

	for i := 0; i < len(num)-1; i++ {
		if DicR2A[rune(num[i])] < DicR2A[rune(num[i+1])] {
			result -= DicR2A[rune(num[i])]
		} else {
			result += DicR2A[rune(num[i])]
		}
	}
	result += DicR2A[rune(num[len(num)-1])]
	return result
}

// Вычисляем римское значение арабской цифры
func Arabic2Roman(num float64) string {
	DicA2R := map[float64]string{
		100: "C",
		90:  "XC",
		50:  "L",
		40:  "XL",
		10:  "X",
		9:   "IX",
		5:   "V",
		4:   "IV",
		1:   "I",
	}
	var result string

	for value, numeral := range DicA2R {
		for num >= value {
			result += numeral
			num -= value
		}
	}
	return result
}

/*Input:
I + 1

Output:
Вывод ошибки, так как используются одновременно разные системы счисления.
*/
