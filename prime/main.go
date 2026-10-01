package main

import "fmt"

func numGenerator(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}

}
 int, ch1 chan<- int, prime int) {
	for {
		val := <-ch
		if val%prime != 0 {
			ch1 <- val
		}
	}
}

func main() {

	ch := make(chan int)

	go numGenerator(ch)
	for i := 0; i < 50; i++ {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go primeChecker(ch, ch1, prime)
		ch = ch1
	}
}
