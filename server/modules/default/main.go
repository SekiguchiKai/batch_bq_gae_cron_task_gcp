package main

import "fmt"

func init() {
	ans := sum(2, 3)
	fmt.Println(ans)

}

func sum(a, b int)int {
	return a + b
}