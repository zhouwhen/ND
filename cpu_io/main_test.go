package main

import(
	"os"
	"testing"
)

func BenchmarkCPU(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add := 0
		for j := 0; j < 10000; j++ {
			add += j 
		}
	}
}


func BenchmarkIO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f,_ := os.Create("test.txt")
		f.WriteString(".")
		f.Close()
	}
	os.Remove("test.txt")
}
/*
测试结果：
goos: windows
goarch: amd64
pkg: prime/cpu_io
cpu: Intel(R) Core(TM) Ultra 7 251HX
BenchmarkCPU-18          1092860              1018 ns/op
BenchmarkIO-18              5630            203958 ns/op
PASS
ok      prime/cpu_io    4.167s
*/