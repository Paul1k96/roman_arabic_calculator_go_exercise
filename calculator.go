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

	//test := "1"
	//fmt.Println(romanCheck(test, roman))
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
				fmt.Println(numToRoman(result), "roman")
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
