/*

ID посылки 50366656

-- ПРИНЦИП РАБОТЫ --

Строим индекс, представляет собой мапку, где для каждого слова учитываем сколько оно в каком документе встречалось(тоже через мапку)
Потом идём по каждому запросу, ищем слово в индексе, смотрим количество вхождений в документах. Суммируем и получаем релевантность каждого документа
Остаётся только корректно отсортировать

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

Если мы имеем для каждого слова количество вхождений в каждый из документов,
то суммируя эти значения с помощью вспомогательной мапы,
получим как раз искомое значение релевантности для каждого документа в случае какой-либо конкретной строки.
А благодаря хэш-таблицам сложность каждого поиска будет константной

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

O(m+n) + O(dlogd), где m - кол-во слов в документах, столько занимает индекс
n - кол-во слов в запросах, по ним нужно пройти
d - кол-во документов, их нужно отсортировать в конце

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --

O(m+n+d), m для индекса, n для агрегирующей мапы с суммой вхождений каждого слова из запроса, d для отсортированного слайса документов

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const limit = 5

type docr struct {
	n int // порядковый номер
	r int // релевантность
}

type adsf struct {
}

func main() {
	fmt.Print(&adsf{})
	s := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 4096)
	s.Buffer(buf, 6+1001*10000+6+1001*10000)
	s.Scan()
	nd, _ := strconv.Atoi(s.Text())

	index := make(map[string]map[int]int)
	n := 1
	// строим индекс
	for n <= nd && s.Scan() {
		doc := s.Text()
		for _, w := range strings.Split(doc, " ") {
			_, ok := index[w]
			if !ok {
				index[w] = map[int]int{n: 0}
			}
			index[w][n]++
		}

		n++
	}

	s.Scan()
	nr, _ := strconv.Atoi(s.Text())
	var res strings.Builder

	for nr > 0 && s.Scan() {
		sorted := make([]*docr, 0, nd)
		relevance := make(map[int]int) // релевантность каждого документа
		alreadySeen := make(map[string]struct{})
		req := s.Text()
		for _, w := range strings.Split(req, " ") {
			if _, ok := alreadySeen[w]; ok {
				continue
			}
			if docCnts, ok := index[w]; ok {
				for k, v := range docCnts {
					relevance[k] += v
				}
				alreadySeen[w] = struct{}{}
			}
		}

		// сортируем
		for n, r := range relevance {
			sorted = append(sorted, &docr{n: n, r: r})
		}

		sort.Slice(sorted, func(i, j int) bool {
			if sorted[i].r > sorted[j].r {
				return true
			}
			if sorted[i].r < sorted[j].r {
				return false
			}

			return sorted[i].n < sorted[j].n
		})

		for i, d := range sorted {
			if i >= limit {
				break
			}
			res.WriteString(strconv.Itoa(d.n))
			if i < limit-1 && i != len(sorted)-1 {
				res.WriteString(" ")
			}
		}
		if nr >= 2 {
			res.WriteString("\n")
		}
		nr--
	}

	fmt.Print(res.String())
}
