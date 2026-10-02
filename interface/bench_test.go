package main

import "testing"

func BenchmarkDirect(b *testing.B) {
	s := Square{Side: 5}
	for i := 0; i < b.N; i++ {
		s.Area()
		s.Perimeter()
	}
}

func BenchmarkInterface(b *testing.B) {
	var shape Shape = Square{Side: 5}
	for i := 0; i < b.N; i++ {
		shape.Area()
		shape.Perimeter()
	}
}
