package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	s "strings"
)

func Calculator() {
	scanner := bufio.NewScanner(os.Stdin)
	var firstNumberHolder float64
	var secondNumberHolder float64

	//First Number
	fmt.Println("Enter your first decimal number: ")
	for scanner.Scan() {
		nStr := scanner.Text()
		firstNumber, err := strconv.ParseFloat(nStr, 64)

		if err != nil {
			fmt.Printf("%s is not a number.\n", nStr)
			fmt.Println("Enter your first decimal number: ")
			continue
		} else {
			firstNumberHolder = firstNumber
			break
		}
	}

	//Second Number
	fmt.Println("Enter your second decimal number: ")
	for scanner.Scan() {
		nStr := scanner.Text()
		secondNumber, err := strconv.ParseFloat(nStr, 64)

		if err != nil {
			fmt.Printf("%s is not a number.\n", nStr)
			fmt.Println("Enter your second decimal number: ")
			continue
		} else {
			secondNumberHolder = secondNumber
			break
		}
	}

	//Operation
	fmt.Println("Would you like to add, subtract, multiply, or divide these numbers? ")
out:
	for scanner.Scan() {
		operationType := scanner.Text()

		switch s.ToLower(operationType) {
		case "add":
			fmt.Printf("Addition of %.2f and %.2f is: %.2f", firstNumberHolder, secondNumberHolder, firstNumberHolder+secondNumberHolder)
			break out
		case "subtract":
			fmt.Printf("Subtraction of %.2f and %.2f is: %.2f", firstNumberHolder, secondNumberHolder, firstNumberHolder-secondNumberHolder)
			break out
		case "multiply":
			fmt.Printf("Multiplication of %.2f and %.2f is: %.2f", firstNumberHolder, secondNumberHolder, firstNumberHolder*secondNumberHolder)
			break out
		case "divide":
			if secondNumberHolder == 0 {
				fmt.Print("Cannot divide by 0, thank you and goodbye")
				break out
			}
			fmt.Printf("Division of %.2f and %.2f is: %.2f", firstNumberHolder, secondNumberHolder, firstNumberHolder/secondNumberHolder)
			break out
		default:
			fmt.Println("Oops that was not a valid operation. Please enter add, subtract, multiply, or divide")
			continue out
		}
	}

}
