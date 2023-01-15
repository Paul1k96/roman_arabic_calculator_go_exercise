package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var b = "\u001B[34m" // blue
	var g = "\u001B[32m" // green
	var r = "\u001B[31m" // red
	reader := bufio.NewReader(os.Stdin)
	roman := map[string]string{
		"I": "1", "II": "2",
		"III": "3", "IV": "4",
		"V": "5", "VI": "6",
		"VII": "7", "VIII": "8",
		"IX": "9", "X": "10",
		"L": "50", "C": "100",
	}

	// Output for users
	fmt.Println("Добро пожаловать в калькулятор для арабских и римских чисел!")
	fmt.Println("Калькулятор может обрабатывать только 2 целых числа от 1 до 10(от I до X) включительно. ")
	fmt.Println("Оба числа должны быть " + clr(r, "ТОЛЬКО") + " арабскими, либо " + clr(r, "ТОЛЬКО") + " римскими.")
	fmt.Println("Ввод данных осуществляется так:" + clr(g, "'число1'") +
		clr(b, "'выражение'") + clr(g, "'число2'") + ", разделённые пробелами между собой.")
	fmt.Println("Примеры: '" + clr(g, "10") + clr(b, " + ") + clr(g, "9") + "," +
		clr(g, " V") + clr(b, " + ") + clr(g, "II"))
	fmt.Println("Для выхода введите" + clr(r, " exit") + ".\n")
	fmt.Println("Начнём работу!")
out:

	for {

		fmt.Println("Введите пример:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Завершаю работу.")
			break out
		}

		list := strings.Split(input, " ")
		if len(list) != 3 {
			fmt.Println("Ошибка: Числа и математическая операция должны быть разделены одиночным пробелом.")
			continue
		}

		isRoman1 := romanCheck(list[0], roman)
		isRoman2 := romanCheck(list[2], roman)
		isNum1 := numCheck(list[0], roman)
		isNum2 := numCheck(list[2], roman)

		fmt.Println(isRoman1, isRoman2, isNum1, isNum2, list)
		switch {
		case isRoman1 == true && isRoman2 == true:
			num1, _ := strconv.Atoi(roman[list[0]])
			expression := list[1]
			num2, _ := strconv.Atoi(roman[list[2]])

			// -----------------------------------------
			result := counting(num1, num2, expression)
			if result == 99999 {
				fmt.Println("Ошибка Используйте в примере арифметичесий оператор (+, -, *, /)")
			} else if result <= 0 {
				fmt.Println("Ошибка: При вычитании первая римская цифра не может быть меньше второй.")
			} else {
				fmt.Println(numToRoman(result))
			}

		case isNum1 == true && isNum2 == true:
			num1, _ := strconv.Atoi(list[0])
			expression := list[1]
			num2, _ := strconv.Atoi(list[2])

			// -----------------------------------------
			result := counting(num1, num2, expression)
			if result == 99999 {
				fmt.Println("Используйте в примере арифметичесий оператор (+, -, *, /)")
			} else {
				fmt.Println(result)
			}
		default:
			fmt.Println("Ошибка: Оба числа должны быть целые в диапазоне I-X/1-10 и быть либо только римскими, либо только арабскими.")
			fmt.Println("Завершаю работу.")
			break out
		}
	}
}

func counting(num1, num2 int, expression string) int {
	switch expression {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	}
	return 99999
}

func numToRoman(num int) string {
	numeric := map[int]string{
		1: "I", 2: "II",
		3: "III", 4: "IV",
		5: "V", 6: "VI",
		7: "VII", 8: "VIII",
		9: "IX", 10: "X",
		50: "L", 100: "C",
	}
	switch {
	case num <= 10 || num == 100:
		return numeric[num]
	case num <= 39:
		int1, int2 := separIntToArray(num)
		output1 := strings.Repeat(numeric[10], int1)
		return output1 + numeric[int2]
	case num <= 49:
		_, int2 := separIntToArray(num)
		return numeric[10] + numeric[50] + numeric[int2]
	case num <= 89:
		int1, int2 := separIntToArray(num)
		output1 := strings.Repeat(numeric[10], int1-5)
		return numeric[50] + output1 + numeric[int2]
	case num <= 99:
		_, int2 := separIntToArray(num)
		return numeric[10] + numeric[100] + numeric[int2]
	}
	return "0"
}

func separIntToArray(num int) (int, int) {
	symb := strconv.Itoa(num)
	list := strings.Split(symb, "")
	int1, _ := strconv.Atoi(list[0])
	int2, _ := strconv.Atoi(list[1])
	return int1, int2
}

func romanCheck(num string, roman map[string]string) bool {
	for k, _ := range roman {
		if num == k {
			return true
		}
	}
	return false
}

func numCheck(num string, roman map[string]string) bool {
	for _, v := range roman {
		if num == v {
			return true
		}
	}
	return false
}

func clr(col, txt string) string {
	var rst = "\u001B[0m" // reset
	return col + txt + rst
}
