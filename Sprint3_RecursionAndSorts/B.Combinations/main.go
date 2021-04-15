package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)

	h := map[string]string{"2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz"}

	chars := make([]string, len(s))
	for i, n := range s {
		chars[i] = h[string(n)]
	}

	// fmt.Println(chars)

	getCombinations("", chars)
}

func getCombinations(prefix string, chars []string) {
	if len(chars) == 0 {
		return
	}

	for _, c := range chars[0] {
		if len(chars) == 1 {
			fmt.Printf("%s%s ", prefix, string(c))
			continue
		}

		getCombinations(prefix+string(c), chars[1:])
	}
}
