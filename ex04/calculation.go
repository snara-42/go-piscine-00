package main

import (
	"fmt"
	"os"
	"strconv"
)

func calculationStr(a []string) (s string, ok bool) {
	if len(a) < 3 {
		return
	}
	n, e1 := strconv.ParseFloat(a[1], 64)
	m, e2 := strconv.ParseFloat(a[2], 64)
	if e1 != nil || e2 != nil || n-n != m-m || n*m/m != n {
		return
	}
	return fmt.Sprintf("sum: %v\ndifference: %v\nproduct: %v\nquotient: %v\n", n+m, n-m, n*m, n/m), true
}

const ERROR_MSG string = "Arguments is invalid."

func main() {
	s, ok := calculationStr(os.Args)
	if !ok {
		fmt.Println(ERROR_MSG)
		os.Exit(1)
	}
	fmt.Print(s)
}
