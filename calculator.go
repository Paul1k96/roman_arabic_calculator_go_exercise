package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//var numInput1, numInput2 string
	var isRoman1, isRoman2, isNum1, isNum2 bool

	reader := bufio.NewReader(os.Stdin)
	roman := map[string]string{
		"I": "1", "II": "2",
		"III": "3", "IV": "4",
		"V": "5", "VI": "6",
		"VII": "7", "VIII": "8",
		"IX": "9", "X": "10",
		"L": "50", "C": "100",
	}

	test := "1"
	fmt.Println(romanCheck(test, roman))
	for {
		fmt.Println("Введите значение")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		list := strings.Split(input, " ")

		isRoman1 = romanCheck(list[0], roman)
		isRoman2 = romanCheck(list[2], roman)
		isNum1 = numCheck(list[0], roman)
		isNum2 = numCheck(list[2], roman)

		fmt.Println(isRoman1, isRoman2, isNum1, isNum2, list)
		switch {
		case isRoman1 == true && isRoman2 == true:
			num1, _ := strconv.Atoi(roman[list[0]])
			expression := list[1]
			num2, _ := strconv.Atoi(roman[list[2]])

			// -----------------------------------------
			result := counting(num1, num2, expression)
			if result == 99999 {
				fmt.Println("Используйсте в примере арифметичесий оператор (+, -, *, /)")
			} else {
				fmt.Println(result, "roman")
			}

		case isNum1 == true && isNum2 == true:
			num1, _ := strconv.Atoi(list[0])
			expression := list[1]
			num2, _ := strconv.Atoi(list[2])

			// -----------------------------------------
			result := counting(num1, num2, expression)
			if result == 99999 {
				fmt.Println("Используйсте в примере арифметичесий оператор (+, -, *, /)")
			} else {
				fmt.Println(result, "numeric")
			}
		}

		//if isRoman1 == false || isRoman2 == false {
		//	if isNum1 == true && isNum2 == true {
		//
		//	} else {
		//		fmt.Println("Используйте пожалуйста римские или арабские числа от 1 до 10(от I до X)")
		//	}
		//}

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
