package main

import(
	"os"
	"testing"
)

func BenchmarkCPU(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Simulate CPU-intensive work
		add := 0
		for j := 0; j < 10000; j++ {
			add += j 
		}
	}
}


func BenchmarkIO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Simulate I/O-intensive work
		f,_ := os.Create("test.txt")
		f.WriteString(".")
		f.Close()
	}
	os.Remove("test.txt")
}