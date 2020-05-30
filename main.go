package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		log.Println("need at lease 2 args")
		return
	}
	f1, f2 := os.Args[1], os.Args[2]

	cores := runtime.NumCPU()
	if len(os.Args) > 3 {
		n, _ := strconv.Atoi(os.Args[3])
		if n > 0 && n < cores {
			cores = n
		}
	}

	filediff(f1, f2, cores)
}

func filediff(f1, f2 string, cores int) {
	log.Printf("reading: %s, %s\n", f1, f2)
	arr1, err := readLines(f1)
	if err != nil {
		log.Printf("read file %s error: %v\n", f1, err)
		return
	}
	arr2, err := readLines(f2)
	if err != nil {
		log.Printf("read file %s error: %v\n", f2, err)
		return
	}
	log.Printf("%s: %d lines\n", f1, len(arr1))
	log.Printf("%s: %d lines\n", f1, len(arr2))

	log.Printf("using %d core(s)...\n", cores)
	var adds []string
	if cores <= 1 {
		adds = diff(arr1, arr2)
	} else {
		adds = diffmulti(arr1, arr2, cores)
	}
	log.Printf("diff on %d core(s): %d adds\n", cores, len(adds))
	for _, v := range adds {
		fmt.Println(v)
	}
	log.Printf("all done\n")
}
