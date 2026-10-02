package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			mu.Lock()
			m[i] = i
			mu.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("有效的map长度:", len(m))

}
