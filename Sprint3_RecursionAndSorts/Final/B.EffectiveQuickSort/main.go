/*

ID посылки 49776569

-- ПРИНЦИП РАБОТЫ --

Алгоритм работы 1 в 1 как в описании задания. По коду понять легко, читаться должно без проблем

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

Индексы идут навстречу друг другу, если не пересеклись меняем местами.
Если пересеклись - продолжаем рекурсивно сортировать.
Если длина массива 1, то возврат.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

O(NlogN), то есть нужно пройти/поменять местами все элементы массива ровно столько раз, сколько можно поделить его пополам :)

*/

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Student struct {
	Name    string
	Solved  int
	Penalty int
}

func (s *Student) String() string {
	return s.Name
}

func (s1 *Student) Compare(s2 *Student) int8 {
	if s1.Solved < s2.Solved {
		return 1
	}

	if s1.Solved > s2.Solved {
		return -1
	}

	if s1.Penalty < s2.Penalty {
		return -1
	}

	if s1.Penalty > s2.Penalty {
		return 1
	}

	return int8(strings.Compare(s1.Name, s2.Name))
}

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	s := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 4096)

	s.Buffer(buf, (20+10e9+10e9+3)*n)
	students := make([]*Student, 0, n)

	var nn int
	for s.Scan() && nn < n {
		str := strings.Split(s.Text(), " ")
		solved, _ := strconv.Atoi(str[1])
		penalty, _ := strconv.Atoi(str[2])
		students = append(students, &Student{Name: str[0], Solved: solved, Penalty: penalty})
		nn++
	}

	rand.Seed(time.Now().UnixNano())
	inplaceqs(students, 0, n-1)

	var resBytes bytes.Buffer

	for i := 0; i < n; i++ {
		resBytes.WriteString(students[i].String() + "\n")
	}

	fmt.Printf(resBytes.String())
}

func inplaceqs(ss []*Student, l, r int) {
	if r-l <= 0 || l < 0 || r >= len(ss) {
		return
	}

	pi := rand.Intn(r-l) + l
	p := ss[pi]

	lt := l
	rt := r

	for rt >= lt {
		for rt >= lt && ss[lt].Compare(p) == -1 {
			lt++
		}

		for rt >= lt && ss[rt].Compare(p) == 1 {
			rt--
		}

		if rt < lt {
			break
		}

		swap(ss, lt, rt)
		lt++
		rt--
	}

	inplaceqs(ss, l, rt)
	inplaceqs(ss, lt, r)
}

func swap(ss []*Student, i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}
