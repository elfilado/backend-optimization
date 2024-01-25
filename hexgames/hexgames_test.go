package main

import "testing"

func BenchmarkTt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Tt()
	}
}

func BenchmarkTt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Tt2()
	}
}

func BenchmarkT3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		T3()
	}
}

func BenchmarkT4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		T4()
	}
}
