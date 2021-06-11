package main

import "fmt"

type FortyTwo struct{}

func main() {
	vars := []interface{}{"42", uint(42), 42, byte(42), int16(42), uint32(42), int64(42), 42 != 42, float32(42), 42.0, complex64(42), 42 + 21i, FortyTwo{}, [1]int{42}, map[string]int{"42": 42}, (*int)(nil), []int{}, chan int(nil), nil}
	for _, v := range vars {
		if t := fmt.Sprintf("%T", v); t[0:2] == "*i" {
			fmt.Printf("%p is an %s.\n", v, t)
		} else if t[0:1] == "i" {
			fmt.Printf("%v is an %s.\n", v, t)
		} else {
			fmt.Printf("%v is a %s.\n", v, t)
		}
	}
}
