package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isArabic(s string) bool { // проверка на арбскость
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func calcInt(operator string, operand1, operand2 int) int {
	var result int
	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		result = operand1 / operand2
	default:
		panic("Неизвестный оператор!")
	}
	return result
}
func intToRoman(num int) string {
	romanMap := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	result := ""

	for i := 0; i < len(romanMap); i++ {
		for num >= romanMap[i].value {
			result += romanMap[i].symbol
			num -= romanMap[i].value
		}
	}

	return result
}
func romanToInt(s string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	result := 0
	for i := 0; i < len(s); i++ {
		current := romanMap[string(s[i])]
		next := 0
		if i+1 < len(s) {
			next = romanMap[string(s[i+1])]
		}

		if current < next {
			result -= current
		} else {
			result += current
		}
	}

	return result
}

// isValidRoman проверяет, является ли строка корректным римским числом.
func isValidRoman(num string) bool {
	pattern := regexp.MustCompile(`^M{0,3}(CM|CD|D?C{0,3})?(XC|XL|L?X{0,3})?(IX|IV|V?I{0,3})?$`)
	return pattern.MatchString(num)
}

func main() {

	//fmt.Print("Введите операцию : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")

	//fmt.Println(parts[0])
	//fmt.Println(parts[1])
	//fmt.Println(parts[2])
	//fmt.Println("=")

	if len(parts) != 3 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	if !isValidRoman(parts[0]) || !isValidRoman(parts[2]) {
		panic("Некорректное Римское число")
	}

	operator := parts[1]

	var operand1, operand2 int

	if isArabic(parts[0]) && isValidRoman(parts[2]) || isValidRoman(parts[0]) && isArabic(parts[2]) {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}
	if isArabic(parts[0]) && isArabic(parts[2]) {

		operand1, _ = strconv.Atoi(parts[0])
		operand2, _ = strconv.Atoi(parts[2])
		if operand1 < 1 || operand1 > 10 || operand2 < 1 || operand2 > 10 {
			panic("значение вводимые должны быть от 1 до 10 вкл")
		}
		var result = calcInt(operator, operand1, operand2)
		fmt.Println(result)
	}
	if isValidRoman(parts[0]) && isValidRoman(parts[2]) {

		operand1 = romanToInt(parts[0])
		operand2 = romanToInt(parts[2])
		if operand1 < 1 || operand1 > 10 || operand2 < 1 || operand2 > 10 {
			panic("значение вводимые должны быть от 1 до 10 вкл")
		}

		var result = calcInt(operator, operand1, operand2)
		if result <= 0 {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
		}
		fmt.Println(intToRoman(result))
	}

}
