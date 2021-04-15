/*

ID посылки 46136228

-- ПРИНЦИП РАБОТЫ --
Дек на массиве, см. структуру Deck
Сканируем файл, сначала считываем числовые параметры, потом команды, преобразуем их в соответствущие методы дека и делаем нужные операции.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

Держим два указателя, соответственно перемещаем их при операциях с деком - PROFIT!
При добавлении элемента спереди мы его добавляем с правого края, предварительно сдвигая соответствующий указатель.
При добавлении элемента сзади элемент добавляется с левого края.
В обоих случаях размер дека увеличивается.
При удалении элемента спереди мы берём значение, двигаем соответствующий указатель, значение возвращается.
В таком случае размер дека уменьшается.

Стоит выделить важный краевой случай: когда в деке 0 элементов, то оба указателя должны двигаться вместе.
Чтобы не выходить за пределы массива, мы вычисляем индекс по модулю capacity, таким образом дек занимает фиксированное количество памяти

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --

O(1), размер константный вне зависимости от количества операций. Определяется при инициализации структуры.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var cmdCount, cap int
	fmt.Scanf("%d\n%d", &cmdCount, &cap)

	s := bufio.NewScanner(os.Stdin)
	d := NewDeck(cap)
	for s.Scan() {
		parseCommand(d, s.Text())
	}
}

// parseCommand парсит команду и вызывает соответствующий метод на деке
func parseCommand(d *Deck, cmd string) {
	cmdParts := strings.Split(cmd, " ")
	switch cmdParts[0] {
	case "push_back":
		val, _ := strconv.Atoi(cmdParts[1])
		ok := d.PushBack(val)
		if !ok {
			fmt.Println("error")
		}
	case "push_front":
		val, _ := strconv.Atoi(cmdParts[1])
		ok := d.PushFront(val)
		if !ok {
			fmt.Println("error")
		}
	case "pop_front":
		val, ok := d.PopFront()
		if !ok {
			fmt.Println("error")
		} else {
			fmt.Println(val)
		}
	case "pop_back":
		val, ok := d.PopBack()
		if !ok {
			fmt.Println("error")
		} else {
			fmt.Println(val)
		}
	}
	fmt.Println(d.slice)
}

// Deck двунаправленная очередь на массиве
type Deck struct {
	slice []int // внутренний массив

	front int // указатель на крайний левый элемент
	rear  int // указатель на крайний правый элемент

	cap  int // общая вместимость дека
	size int // количество элементов в деке
}

// NewDeck инициализирует пустой дек
func NewDeck(cap int) *Deck {
	return &Deck{slice: make([]int, cap), cap: cap, size: 0, rear: cap - 1, front: 1}
}

// PushBack добавляет элемент в конец дека, возвращает true
// Если в деке уже находится максимальное число элементов, возвращает false
func (d *Deck) PushBack(val int) bool {
	if d.size == d.cap {
		return false
	}

	d.rear = (d.rear + 1) % d.cap
	d.slice[d.rear] = val

	if d.size == 0 {
		d.front = d.rear
	}

	d.size++
	return true
}

// PushFront добавляет элемент в начало дека, возвращает true
// Если в деке уже находится максимальное число элементов, возвращает false
func (d *Deck) PushFront(val int) bool {
	if d.size == d.cap {
		return false
	}

	d.front = (d.front - 1 + d.cap) % d.cap
	d.slice[d.front] = val
	if d.size == 0 {
		d.rear = d.front
	}

	d.size++
	return true
}

// PopBack возвращает последний элемент дека и удаляет его, возвращает true
// Если дек был пуст, то возвращает false
func (d *Deck) PopBack() (int, bool) {
	if d.size == 0 {
		return 0, false
	}

	val := d.slice[d.rear]
	d.rear = (d.rear - 1 + d.cap) % d.cap
	d.size--
	return val, true
}

// PopFront возвращает первый элемент дека и удаляет его, возвращает true
// Если дек был пуст, то возвращает false
func (d *Deck) PopFront() (int, bool) {
	if d.size == 0 {
		return 0, false
	}

	val := d.slice[d.front]
	d.front = (d.front + 1) % d.cap
	d.size--
	return val, true
}
