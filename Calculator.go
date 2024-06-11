package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"errors"
)

func calculate(num1, num2 int, operation string ) int {
	switch operation {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		panic("Invalid operation!")
	}
}

func romanNumCheck (num string) int {
	switch num {
		case "I":
			return 1
		case "II":
			return 2
		case "III":
			return 3
		case "IV":
			return 4
		case "V":
			return 5
		case "VI":
			return 6
		case "VII":
			return 7
		case "VIII":
			return 8
		case "IX":
			return 9
		case "X":
			return 10
		default:
			return 0
	}
}

func romanNumGet (num int) string {
	switch num {
		case 1:
			return "I"
		case 2:
			return "II"
		case 3:
			return "III"
		case 4:
			return "IV"
		case 5:
			return "V"
		case 6:
			return "VI"
		case 7:
			return "VII"
		case 8:
			return "VIII"
		case 9:
			return "IX"
		case 10:
			return "X"
		default:
			return "0"
	}
}

func main() {
	var 
		num1, num2 int
	var
		operation, inpt string

	err := errors.New("error!")
	
	fmt.Println("Input an expression in the format: number1 operation number2")
	if inpt, err = bufio.NewReader(os.Stdin).ReadString('\n'); err != nil {
		panic(err)
	}
	slices := strings.Split(inpt, " ")
	if len(slices) > 3 {
		panic("Too many arguments!")
	}

	// discard last two elements "\n\r"
	num2Tmp := slices[2]
	trimmedNum2 := strings.Trim(num2Tmp, "\n\r")

	if num1 = romanNumCheck(slices[0]); num1 != 0 {
		if num2 = romanNumCheck(trimmedNum2); num2 != 0 {
			if result := calculate(num1, num2, slices[1]); result > 0 {
				fmt.Println("Result: ", romanNumGet(result))
			} else {
				panic("The result is out of range for roman numerals!")
			}
		} else {
			panic("Invalid arguments")
		}
	} else {
		if num1, err = strconv.Atoi(slices[0]); (err != nil) || ((num1 < 1) || (num1 > 10)) {
			panic("Invalid arguments!")
		}
		operation = string(slices[1])
		if num2, err = strconv.Atoi(trimmedNum2); (err != nil) || ((num2 < 1) || (num2 > 10)) {
			panic("Invalid arguments!")
		}
		fmt.Println("Result: ", calculate(num1, num2, operation))
	}
	
}