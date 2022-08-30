package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

func (calculator calc) operate(entrada, operador string) int {
	valores := strings.Split(entrada, operador)
	operador1 := parsear(valores[0])
	operador2 := parsear(valores[1])
	switch operador {
	case "+":
		return operador1 + operador2
	case "-":
		return operador1 - operador2

	case "*":
		return operador1 * operador2

	case "/":
		return operador1 / operador2

	default:
		fmt.Println(operador, "No soportado")
		return 0
	}
}

func parsear(entrada string) int {
	operador, err := strconv.Atoi(entrada)
	pushError(err)
	return operador
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	return scanner.Text()
}

func pushError(err error) {
	if err != nil {
		panic(err)
	}
}
