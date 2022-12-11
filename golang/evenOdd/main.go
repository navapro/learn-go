package main

import "fmt"

func main () {
	s := []int {}
	for i := 0 ; i <=10; i ++ {
		s = append(s, i)
	}

	for _, num := range s {
		if num%2 == 0 {
			fmt.Println(num,"is", "Even")
		} else {
			fmt.Println(num, "is", "Odd")
		}
	}
}