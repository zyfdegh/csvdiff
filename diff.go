package main

import (
	"fmt"
	"runtime"
	"sync"
)

// diff finds strings in b but not in a
func diff(a, b []string) []string {
	res := make([]string, 0, len(b))
	for _, p := range b {
		var found bool
		for _, q := range a {
			if p == q {
				found = true
				break
			}
		}
		if !found {
			res = append(res, p)
		}
	}
	return res
}

func diffmulti(a, b []string) []string {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	res := make([]string, 0, len(b))
	chunks := Split(b, runtime.NumCPU())
	for _, ck := range chunks {
		wg.Add(1)
		fmt.Println("========= ck", ck)
		go func(arr []string) {
			defer wg.Done()
			adds := diff(a, arr)
			mu.Lock()
			res = append(res, adds...)
			mu.Unlock()
		}(ck)
	}
	wg.Wait()
	return res
}
