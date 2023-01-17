package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var b = "\u001B[34m" // blue
	var g = "\u001B[32m" // green
	var r = "\u001B[31m" // red

	roman := map[string]string{
		"I": "1", "II": "2",
		"III": "3", "IV": "4",
		"V": "5", "VI": "6",
		"VII": "7", "VIII": "8",
		"IX": "9", "X": "10",
	}

	// Output for users
	fmt.Println("----------------------------------------------------------------------------------------------")
	fmt.Println("Добро пожаловать в калькулятор для арабских и римских чисел!")
	fmt.Println("Калькулятор может обрабатывать только 2" + clr(r, " ЦЕЛЫХ ") + "числа от 1 до 10(от I до X) включительно. ")
	fmt.Println("Оба числа должны быть " + clr(r, "ТОЛЬКО") + " арабскими, либо " + clr(r, "ТОЛЬКО") + " римскими.")
	fmt.Println("Ввод данных осуществляется так:" + clr(g, "'число1'") +
		clr(b, "'выражение'") + clr(g, "'число2'") + ", разделённые пробелами между собой.")
	fmt.Println("Примеры: '" + clr(g, "10") + clr(b, " + ") + clr(g, "9") + "', '" +
		clr(g, "V") + clr(b, " + ") + clr(g, "II") + "'")
	fmt.Println("Для выхода введите" + clr(r, " exit") + ".")
	fmt.Println("----------------------------------------------------------------------------------------------\n")
	fmt.Println("Начнём работу!")

out:
	for {

		fmt.Println("Введите пример:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // stripping spaces and tabs from both ends of lines
		input = strings.ToUpper(input)

		// terminates the program
		if input == "EXIT" {
			fmt.Println("Завершаю работу.")
			break out
		}

		list := strings.Split(input, " ")

		// check
		if len(list) != 3 {
			fmt.Println("Ошибка: Введите пожалуйста два числа и математическую операцию, разделённые одиночным пробелом.")
			continue
		}

		// flags to check
		isRoman1 := romanCheck(list[0], roman)
		isRoman2 := romanCheck(list[2], roman)
		isNum1 := numCheck(list[0], roman)
		isNum2 := numCheck(list[2], roman)

		switch {
		case isRoman1 == true && isRoman2 == true: // if both numbers are roman
			num1, _ := strconv.Atoi(roman[list[0]])
			expression := list[1]
			num2, _ := strconv.Atoi(roman[list[2]])

			// -----------------------------------------
			result := counting(num1, num2, expression)
			if result == 99999 {
				fmt.Println("Ошибка: Используйте в примере арифметичесий оператор (+, -, *, /)")
			} else if result <= 0 {
				fmt.Println("Ошибка: При вычитании или делении первая римская цифра должна быть больше второй.")
			} else {
				fmt.Println("Ответ:", numToRoman(result))
			}

		case isNum1 == true && isNum2 == true: // if both numbers are arabic
			num1, _ := strconv.Atoi(list[0])
			expression := list[1]
			num2, _ := strconv.Atoi(list[2])

			// -----------------------------------------
			result := counting(num1, num2, expression)
			if result == 99999 {
				fmt.Println("Ошибка: Используйте в выражении арифметичесий оператор (+, -, *, /).")
			} else {
				fmt.Println("Ответ:", result)
			}

		case (isNum1 || isNum2) && (isRoman1 || isRoman2): //  if one is Arabic and the other is Roman
			fmt.Println("Ошибка: Операции между римскими и арабскими числами запрещены.")
			break out

		case isNum1 || isNum2: // if only one arabic
			fmt.Println("Ошибка: Было введено только одно арабское целое число от 1 до 10.")
			break out

		case isRoman1 || isRoman2: // if only one roman
			fmt.Println("Ошибка: Было введено только одно римское число в диапазоне от I до X.")
			break out

		default:
			fmt.Println("Ошибка: Должны быть два целых в диапазоне I-X/1-10.")
			fmt.Println("Завершаю работу.")
			break out
		}
	}
}

func counting(num1, num2 int, expression string) int {
	// performs mathematical operations and returns the result

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
	// convert arabic number to roman number

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
		int1, int2 := sepIntToArray(num)
		output1 := strings.Repeat(numeric[10], int1)
		return output1 + numeric[int2]
	case num <= 49:
		_, int2 := sepIntToArray(num)
		return numeric[10] + numeric[50] + numeric[int2]
	case num <= 89:
		int1, int2 := sepIntToArray(num)
		output1 := strings.Repeat(numeric[10], int1-5)
		return numeric[50] + output1 + numeric[int2]
	case num <= 99:
		_, int2 := sepIntToArray(num)
		return numeric[10] + numeric[100] + numeric[int2]
	}
	return "0"
}

func sepIntToArray(num int) (int, int) {
	// divides a two-digit int into two parts and returns them as an int

	str := strconv.Itoa(num)
	list := strings.Split(str, "")
	int1, _ := strconv.Atoi(list[0])
	int2, _ := strconv.Atoi(list[1])
	return int1, int2
}

func romanCheck(num string, roman map[string]string) bool {
	// checks if there is a variable 'num' in the keys in the 'roman' map and returns the flag

	for k, _ := range roman {
		if num == k {
			return true
		}
	}
	return false
}

func numCheck(num string, roman map[string]string) bool {
	// checks if there is a variable 'num' in the 'roman' map and returns the flag

	for _, v := range roman {
		if num == v {
			return true
		}
	}
	return false
}

func clr(col, txt string) string {
	// returns the received string in the received color

	var rst = "\u001B[0m" // reset
	return col + txt + rst
}
