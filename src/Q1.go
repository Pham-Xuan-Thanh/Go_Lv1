package lv1

import (
	"fmt"
)

func Q1() {
	var n int
	var sum int
	var i int = 1

	fmt.Print("Nhap vao n: ")
	fmt.Scan(&n)

	for ; i < n; i++ {
		sum += i
		fmt.Print(i, "+")
	}

	fmt.Print(n, "=", sum+n)

}
