package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 2 {
		fmt.Println("need 2 args")
		return
	}
	f1, f2 := os.Args[1], os.Args[2]
	// fmt.Printf("reading: %s, %s\n", f1, f2)
	arr1, err := readLines(f1)
	if err != nil {
		fmt.Printf("read file %s error: %v\n", f1, err)
		return
	}
	arr2, err := readLines(f2)
	if err != nil {
		fmt.Printf("read file %s error: %v\n", f2, err)
		return
	}
	adds := diff(arr1, arr2)
	for _, v := range adds {
		fmt.Println(v)
	}
}
