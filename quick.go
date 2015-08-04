package main

import "fmt"

func main() {

	array := []int{12, 5, 9, 99, 16, 3}

	fmt.Println(array)
	quickSort(array, 0, len(array)-1)
	fmt.Println(array)
}

func quickSort(a []int, l int, r int) {

	i := l
	j := r

	p := a[(l+r)/2]

	for i <= j {
		for a[i] < p {
			i++
		}

		for p < a[j] {
			j--
		}
		if i <= j {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
	}

	if l < j {
		quickSort(a, l, j)
	}
	if i < r {
		quickSort(a, i, r)
	}

}
