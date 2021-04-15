/*

ID посылки 50367573

-- ПРИНЦИП РАБОТЫ --

Хэш целого числа считается по модулю большого простого числа 81131
Коллизии разрешаются с помощью метода цепочек

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

В каждой ячейке массива(корзине) хранится ссылка на звено односвязного списка.
Звено хранит ключ и значение.
При вычислении номера корзины по модулю размера массива мы не выходим за границы массива.
Коллизии разрешаются следующим образом: если попадаем в ту же корзину, идём по списку и значения ключей сравниваются.
Если ключи равны, записываем в текущей ячейке.
Если не находим ячейки с тем же ключом, то дописываем в конец связного списка новый узел.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

O(1) в среднем для каждой из операций

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --

т.к. масштабирование не требуется, беру за размер 81131 например, получается тоже в данной реализации О(1)

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const M = 81131

type hashtable struct {
	arr []*node
}

type node struct {
	next *node
	key  int
	val  int
}

func NewHashtable() *hashtable {
	return &hashtable{arr: make([]*node, M)}
}

func (h *hashtable) put(key, value int) {
	i := key % M // todo move to hash function
	if h.arr[i] == nil {
		h.arr[i] = &node{key: key, val: value}
		return
	}

	// helper node
	curr := &node{next: h.arr[i]}
	for curr.next != nil {
		if key == curr.next.key {
			curr.next.val = value
			return
		}
		curr = curr.next
	}

	curr.next = &node{key: key, val: value}
}

func (h *hashtable) get(key int) (int, bool) {
	i := key % M // todo move to hash function
	if h.arr[i] == nil {
		return 0, false
	}

	// search in linked list
	curr := h.arr[i]
	for curr != nil {
		if key == curr.key {
			return curr.val, true
		}
		curr = curr.next
	}

	return 0, false
}

func (h *hashtable) delete(key int) (int, bool) {
	i := key % M
	if h.arr[i] == nil {
		return 0, false
	}

	curr := h.arr[i]
	var prev *node
	for curr != nil {
		if key == curr.key {
			val := curr.val
			// remove from linked list
			if prev == nil {
				h.arr[i] = nil
				return val, true
			}
			prev.next = curr.next
			return val, true
		}
		prev = curr
		curr = curr.next
	}

	return 0, false
}

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	ht := NewHashtable()
	s := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 4096)
	s.Buffer(buf, n*(10e9+10e9+10e9))

	var res strings.Builder
	for n > 0 && s.Scan() {
		cmdres, _ := parseCmd(s.Text(), ht)
		// fmt.Println(ht)
		if cmdres == "" {
			n--
			continue
		}
		res.WriteString(cmdres)

		if n > 1 {
			res.WriteString("\n")
		}
		n--
	}

	fmt.Print(res.String())
}

func parseCmd(cmd string, ht *hashtable) (string, error) {
	parts := strings.Split(cmd, " ")
	switch parts[0] {
	case "get":
		key, _ := strconv.Atoi(parts[1])
		v, ok := ht.get(key)
		if !ok {
			return "None", nil
		}
		return strconv.Itoa(v), nil
	case "put":
		key, _ := strconv.Atoi(parts[1])
		val, _ := strconv.Atoi(parts[2])

		ht.put(key, val)
		// return strconv.Itoa(val), nil
		return "", nil
	case "delete":
		key, _ := strconv.Atoi(parts[1])
		v, ok := ht.delete(key)
		if !ok {
			return "None", nil
		}
		return strconv.Itoa(v), nil
	}

	return "", fmt.Errorf("unknown command")
}
