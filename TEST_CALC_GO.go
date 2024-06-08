package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isArabic(s string) bool { // проверка на арбскость
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}
func isRomanNumeral(s string) bool { // проверка числа на римскость
	romanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, numeral := range romanNumerals {
		if strings.EqualFold(numeral, s) {
			return true
		}
	}
	return false
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

func main() {

	fmt.Print("Введите операцию : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")

	if len(parts) != 3 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	operator := parts[1]

	var operand1, operand2 int

	if isArabic(parts[0]) && isRomanNumeral(parts[2]) || isRomanNumeral(parts[0]) && isArabic(parts[2]) {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}
	if isArabic(parts[0]) && isArabic(parts[2]) {
		operand1, _ = strconv.Atoi(parts[0])
		operand2, _ = strconv.Atoi(parts[2])
		var result = calcInt(operator, operand1, operand2)
		fmt.Println(result)
	}
	if isRomanNumeral(parts[0]) && isRomanNumeral(parts[2]) { // если оба числа римские приводим к арабским
		romanToArabic := map[string]int{
			"I":    1,
			"II":   2,
			"III":  3,
			"IV":   4,
			"V":    5,
			"VI":   6,
			"VII":  7,
			"VIII": 8,
			"IX":   9,
			"X":    10,
		}
		operand1 = romanToArabic[parts[0]]
		operand2 = romanToArabic[parts[2]]
		var result = calcInt(operator, operand1, operand2)
		if result <= 0 {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
		}
		fmt.Println(intToRoman(result))
	}

}
