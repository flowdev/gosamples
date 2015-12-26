package main

import "testing"

func BenchmarkSample1(b *testing.B) {
    flow := NewSayHello()
    data := []string{"Rani", "Sina", "Lina"}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
	flow.InPort(data)
    }
}
