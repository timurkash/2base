package main

import (
	"testing"
)

func Test2Base(t *testing.T) {
	tests := []struct {
		I      uint64
		Is2    bool
		Bas    uint8
		Result bool
	}{
		{
			I:      1,
			Is2:    true,
			Bas:    0,
			Result: true,
		},
		{
			I:      10,
			Is2:    false,
			Bas:    0,
			Result: true,
		},
		{
			I:      16,
			Is2:    true,
			Bas:    4,
			Result: true,
		},
		{
			I:      1024,
			Is2:    true,
			Bas:    10,
			Result: true,
		},
		{
			I:      1025,
			Is2:    false,
			Bas:    10,
			Result: false,
		},
		{
			I:      1025,
			Is2:    false,
			Bas:    0,
			Result: true,
		},
		{
			I:      1024,
			Is2:    true,
			Bas:    11,
			Result: false,
		},
	}
	for _, test := range tests {
		is2, bas := Is2Base(test.I)
		if (is2 != test.Is2 || bas != test.Bas) && test.Result {
			t.Errorf("bad result")
		} else if is2 == test.Is2 && bas == test.Bas && !test.Result {
			t.Errorf("expected wrong result")
		}
	}
}

func Benchmark2Base_1(b *testing.B) {
	var i uint64 = 1024
	for i = 0; i < uint64(b.N); i++ {
		_, _ = Is2Base(i)
	}
}

func Benchmark2Base_2(b *testing.B) {
	var i uint64 = 1024
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = Is2Base(i)
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
		_, _ = Is2Base32(uint64(i))
	}
}

func Benchmark2Base4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Is2Base4(uint64(i))
	}
}
