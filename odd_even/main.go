package main

import "fmt"

func main() {
	done1 := make(chan bool, 1) //偶数信号
	done2 := make(chan bool)    //奇数信号
	finish := make(chan bool)   //结束
	//奇数
	go func() {
		for i := 1; i <= 10; i += 2 {
			done1 <- true
			fmt.Println("奇数：", i)
			done2 <- true
		}

	}()
	//偶数
	go func() {
		for i := 2; i <= 10; i += 2 {
			<-done2
			fmt.Println("偶数：", i)
			<-done1
		}
		finish <- true
	}()
	<-finish
}
