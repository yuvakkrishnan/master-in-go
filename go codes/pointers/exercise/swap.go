package main

import "fmt"

func main() {

	v, n := 3.55, 8.99
	swapVar(&v, &n)
	fmt.Printf("v :%g - n :%g\n", v, n)
	pa, pb := &v, &n
	pa, pb = swapAddress(pa, pb)
	fmt.Printf("pa : %p - pb : %p\n", pa, pb)
	fmt.Printf("pa : %p - pb : %p\n", *pa, *pb)

}

func swapVar(v *float64, n *float64) {
	*v, *n = *n, *v
	return
}

func swapAddress(v, n *float64) (*float64, *float64) {
	return n, v

}
