package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid argument")
		os.Exit(0)
	}
	for _, v := range os.Args[1:] {
		if m, e := regexp.MatchString(`[\w\-\._]+@[\w\-\._]+\.[A-Za-z]+`, v); len(v) < 256 && m && e == nil {
			fmt.Println(v, "is a valid email address.")
		} else {
			fmt.Println(v, "is NOT a valid email address.")
		}
	}
}
