package main

import (
	"testing"
)

func Benchmark2Base(b *testing.B) {
	var i uint64
	for i = 0; i < uint64(b.N); i++ {
		_, _ = Is2Base(uint64(i))
	}
}

func Benchmark2Base_2(b *testing.B) {
	var i uint64 = 1024
	b.RunParallel(func (pb *testing.PB) {
		for pb.Next() {
			_, _ = Is2Base(uint64(i))
		}
	})
}

func Benchmark2Base2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Is2Base2(uint64(i))
	}
}

func Benchmark2Base3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Is2Base3(uint64(i))
	}
}

func Benchmark2Base3_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Is2Base3_2(uint64(i))
	}
}

func Benchmark2Base4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Is2Base4(uint64(i))
	}
}