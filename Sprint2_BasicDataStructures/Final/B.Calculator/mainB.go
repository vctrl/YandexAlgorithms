/*

ID посылки 46138818

-- ПРИНЦИП РАБОТЫ --

Считываем входную строку, сплитим на массив, проходим по массиву, попутно совершая операции со стеком.
Структура данных стек реализована на массиве.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

Алгоритм реализован в соответствии со спецификацией:

Для вычисления значения выражения, записанного в обратной польской нотации, нужно считывать выражение слева направо и придерживаться следующих шагов:

Обработка входного символа:
Если на вход подан операнд, он помещается на вершину стека.
Если на вход подан знак операции, то эта операция выполняется над требуемым количеством значений из стека, взятых в порядке добавления. Результат выполненной операции помещается на вершину стека.
Если входной набор символов обработан не полностью, перейти к шагу 1.
После полной обработки входного набора символов результат вычисления выражения находится в вершине стека. Если в стеке осталось несколько чисел, то надо вывести только верхний элемент.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --

O(n), т.к. стек используется для сохранения промежуточных результатов

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	parts := strings.Split(s.Text(), " ")

	var stack Stack

	for _, p := range parts {
		if isSign(p) {
			op1str, ok := stack.Pop()
			if !ok {
				log.Fatal("Stack is empty")
			}

			op1, _ := strconv.Atoi(op1str)
			op2str, ok := stack.Pop()
			if !ok {
				log.Fatal("Stack is empty")
			}

			op2, _ := strconv.Atoi(op2str)
			res := doOperation(p, op2, op1)
			stack.Push(strconv.Itoa(res))
		} else {
			stack.Push(p)
		}

		fmt.Println(stack)
	}

	result, ok := stack.Pop()
	if !ok {
		log.Fatal("Stack is empty")
	}

	fmt.Println(result)
}

func isSign(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}

func doOperation(op string, op1, op2 int) int {
	switch op {
	case "+":
		return op1 + op2
	case "-":
		return op1 - op2
	case "*":
		return op1 * op2
	case "/":
		// математическое целочисленное деление
		// Это значит, что округление всегда происходит вниз.
		// А именно: если a / b = c, то b ⋅ c – это наибольшее число, которое не превосходит a и одновременно делится без остатка на b.
		if op1 < 0 && op2 > 0 && op1%op2 != 0 {
			return op1/op2 - 1
		}

		return op1 / op2
	}

	return 0
}

// Stack структура данных стек на массиве
type Stack []string

// Push добавляет элемент в стек
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// Pop возвращает элемент из стека, удаляя его, при успехе возвращает значение элемента и true
// Если стек пустой, возвращает false вторым значением
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	index := len(*s) - 1
	res := (*s)[index]
	*s = (*s)[:index]
	return res, true
}

// IsEmpty возвращает true, если в стеке нет элементов, и true в обратном случае
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
