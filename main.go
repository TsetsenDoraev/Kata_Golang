/*
Input:
1 + 2
Output:
3

Input:
VI / III
Output:
II

Input:
I - II
Output:
Вывод ошибки, так как в римской системе нет отрицательных чисел.

Input:
I + 1
Output:
Вывод ошибки, так как используются одновременно разные системы счисления.

Input:
1
Output:
Вывод ошибки, так как строка не является математической операцией.

Input:
1 + 2 + 3
Output:
Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).
*/

package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
	//"strconv"
)

func main() {
	Greeting := "\n\nДобро пожаловать в калькулятор" // Приветственная часть
	fmt.Println(Greeting)
	Delimiter := strings.Repeat("—", utf8.RuneCountInString(Greeting)-2)
	fmt.Println(Delimiter)

	var expression string
	fmt.Print("Введите выражение")
	fmt.Scan(&expression)

	expression = strings.ReplaceAll(expression, " ", "")
	action := []string{"+", "-", "*", "/"}

	switch action { //Валидация
	case "+":
	case "-":
	case "*":
	case "/":
	default:
		fmt.Println("Введён нерелевантный символ")
		os.Exit(0) // Или лучше использовать «panic()»?
	}

}
