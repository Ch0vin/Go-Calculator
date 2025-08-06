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

	for {
		fmt.Print("Enter a first number: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		num1, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid number ,please try again")
			continue
		}
		fmt.Println("Enter an operator (+, -, *, /):")
		operator, _ := reader.ReadString('\n')
		operator = strings.TrimSpace(operator)
		operator = strings.TrimSpace(operator)

		fmt.Print("Enter a second number: ")
		input1, _ := reader.ReadString('\n')
		input1 = strings.TrimSpace(input1)
		num2, err := strconv.ParseFloat(input1, 64)
		if err != nil {
			fmt.Println("Invalid number , please try again")
			continue
		}

		var result float64

		switch operator {

		case "+":
			result = num1 + num2

		case "-":
			result = num1 - num2

		case "/":
			if num2 == 0 {
				fmt.Println("Error: Division by zero.")
				continue

			}
			result = num1 / num2

		case "*":
			result = num1 * num2

		default:
			fmt.Println("Invalid operator. Please use one of +, -, *, /.")
			continue

		}
		fmt.Printf("Result: %f\n", result)

	}

}
