package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	if n == 1 {
		fmt.Println(0)
		return
	}

	if n == 2 {
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &a[i])
		}
		if (a[0] == 1 && a[1] == 2) || (a[0] == 2 && a[1] == 1) { // 1,2 or 2,1
			fmt.Println(1)
			return
		}
		fmt.Println(n)
		return
	}

	fmt.Println(n)

	return
}
