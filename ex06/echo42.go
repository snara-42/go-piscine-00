package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	s := flag.String("s", " ", "separator")
	n := flag.Bool("n", false, "omit trailing newline")
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *s), map[bool]string{true: "", false: "\n"}[*n])
}
