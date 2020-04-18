package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type calc struct {}

func (calc) operate (input1 string, input2 string, operator string) int {
	inputValue1 := intParser(input1)
	inputValue2 := intParser(input2)

	switch operator {
	case "+":
		return inputValue1 + inputValue2
	case "-":
		return inputValue1 - inputValue2
	case "*":
		return inputValue1 * inputValue2
	case "/":
		if (inputValue2) == 0 {
			fmt.Println("Any number could be divided by zero")
			return 0
		} else {
			return inputValue1 / inputValue2
		}
	default:
		fmt.Println("Operator not supported")
		return 0
	}
}

func intParser (input string) int {
	// String.Atoi returns two values (value, err), so we need to use "_" (underscore) as a wildcard to avoid «Not used variable»
	// If you manage the error you need to assign the second value in a variable
	intValue, _ := strconv.Atoi(input)
	return intValue
}

func main() {
	fmt.Println("Qué operacion deseas hacer? (+, -, *, /)")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// Read values from console
	operator := scanner.Text()
	fmt.Println("Ingresa el primer número")
	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Scan()
	fmt.Println("Ingresa el segundo número")
	inputValue1 := scanner2.Text()
	scanner3 := bufio.NewScanner(os.Stdin)
	scanner3.Scan()
	inputValue2 := scanner3.Text()
	c := calc{}
	fmt.Println("Resultado:" , c.operate(inputValue1, inputValue2,operator))
}
