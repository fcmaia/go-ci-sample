package main

import (
	"fmt"
)

func main() {
	/*
		if len(os.Args) < 3 {
			fmt.Println("expected 2 numbers")
			os.Exit(1)
		}

		s1, err1 := strconv.ParseInt(os.Args[1], 10, 32)
		s2, err2 := strconv.ParseInt(os.Args[2], 10, 32)

		if err1 != nil && err2 != nil {
			fmt.Print(Sum(int(s1), int(s2)))
		} else {
			fmt.Print(err1, err2)
			os.Exit(1)
		}
	*/
	fmt.Print("5 + 5 = ", Sum(5, 5))
}

func Sum(x int, y int) int {
	return x + y
}
