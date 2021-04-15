package main

// func main() {
// 	arr := []int{}
// 	merge_sort(arr, 0, len(arr))
// 	fmt.Println("after", arr)
// }

func merge(arr []int, lf int, mid int, rg int) []int {
	// fmt.Printf("merging %d %d %d\n", lf, mid, rg)
	i, j, k := lf, mid, 0
	res := make([]int, rg-lf)
	for i < mid && j < rg {
		if arr[i] <= arr[j] {
			res[k] = arr[i]
			i++
		} else {
			res[k] = arr[j]
			j++
		}

		k++
	}

	for i < mid {
		res[k] = arr[i]
		i++
		k++
	}

	for j < rg {
		res[k] = arr[j]
		j++
		k++
	}

	// fmt.Println("merged")
	return res
}

func merge_sort(arr []int, lf int, rg int) {
	// fmt.Println(lf, rg)
	if rg-lf <= 1 {
		return
	}

	mid := (lf + rg) / 2
	merge_sort(arr, lf, mid)
	merge_sort(arr, mid, rg)

	res := merge(arr, lf, mid, rg)

	k := 0
	for i := lf; i < rg; i++ {
		arr[i] = res[k]
		k++
	}
}
