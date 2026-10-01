package main

import "fmt"

func main() {
	m := make(map[int]int)
	for i := 0; i < 50; i++ {
		go func(i int) {
			m[i] = i
			fmt.Println(i)
		}(i)
	}

	select {}
}
