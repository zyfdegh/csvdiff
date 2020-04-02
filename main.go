package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func main() {
	if len(os.Args) <= 2 {
		fmt.Println("need 2 args")
		return
	}
	f1, f2 := os.Args[1], os.Args[2]
	log.Printf("reading: %s, %s\n", f1, f2)

	runtime.GOMAXPROCS(runtime.NumCPU())
	diffseq(f1, f2)
}

func diffseq(f1, f2 string) {
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
	log.Printf("%s: %d lines\n", f1, len(arr1))
	log.Printf("%s: %d lines\n", f1, len(arr2))

	// adds := diff(arr1, arr2)
	adds := diffmulti(arr1, arr2)
	log.Printf("diff: %d adds\n", len(adds))
	for _, v := range adds {
		fmt.Println(v)
	}
	log.Printf("all done\n")
}
