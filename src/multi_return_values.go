package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func allotOfVals() (int, int, int) {
	return 1, 2, 3
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	d, e, f := allot_of_vals()
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	_, c := vals()
	fmt.Println(c)
}
