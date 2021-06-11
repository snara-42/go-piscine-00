package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, e := strconv.Atoi(os.Args[1])
	for s, i := "*", 1; e == nil && i*(i+1)/2 <= n; s, i = s+"*", i+1 {
		fmt.Println(s)
	}
}
