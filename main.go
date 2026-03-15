// Programa de exemplo com funções, múltiplos retornos, closure e higher-order function.
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	sayHello()
	fmt.Println(sum(1, 2))
	fmt.Println(swap(10, 20))
	fmt.Println(divide(10, 3))

	// Função anônima atribuída a uma variável (closure).
	foo := func(a int) int {
		return a + 1
	}
	fmt.Println(foo(1))
	fmt.Println(sumVariadic(1, 2, 3, 4, 5))
}

func sayHello() {
	fmt.Println("Hello")
}

// sum retorna a soma de dois inteiros.
func sum(a, b int) int {
	return a + b
}

// swap retorna os dois argumentos com a ordem invertida (múltiplos retornos nomeados implícitos).
func swap(a, b int) (int, int) {
	return b, a
}

// divide faz divisão inteira e retorna quociente e resto.
func divide(a, b int) (int, int) {
	result := a / b
	remainder := a % b
	return result, remainder
}

// sumHigherOrder é uma higher-order function: retorna uma função que soma 'a' (capturado) com o argumento recebido.
func sumHigherOrder(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

// function with variadic arguments
func sumVariadic(args ...int) int {
	total := 0
	for _, arg := range args {
		total += arg
	}
	return total
}
