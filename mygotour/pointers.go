package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // read i through the pointer

	p = &j         // point to j
	*p /= 37       // (*p = *p / 37) divide j through the pointer
	fmt.Println(j) // read j through the pointer
}
