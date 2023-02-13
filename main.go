package main

import (
	"fmt"
	"math"
	"math/bits"
)

func Is2Base(i uint64) (bool, uint8) {
	if bits.OnesCount64(i) != 1 {
		return false, 0
	}
	return true, uint8(bits.TrailingZeros64(i))
}

func Is2Base2(i uint64) (bool, uint8) {
	var n uint64 = 1
	var d uint8 = 0
	for {
		if n > i {
			break
		}
		if n == i {
			return true, d
		}
		n = n << 1
		d++
	}
	return false, 0
}

func Is2Base3(i uint64) (bool, uint8) {
	var n uint64 = 1
	var d uint8 = 0
	for {
		if n > i {
			break
		}
		if n == i {
			return true, d
		}
		n *= 2
		d++
	}
	return false, 0
}

func Is2Base3_2(i uint64) (bool, uint8) {
	var n uint64 = 1
	var d uint8 = 0
	for {
		if n > i {
			break
		}
		if n == i {
			return true, d
		}
		n = n * 2
		d++
	}
	return false, 0
}

func Is2Base4(i uint64) (bool, uint8) {
	base := math.Log2(float64(i))
	if base == math.Trunc(base) {
		return true, uint8(base)
	}
	return false, 0
}

func main() {
	var i uint64 = 4
	fmt.Println(Is2Base(i))
	fmt.Println(Is2Base2(i))
	fmt.Println(Is2Base3(i))
	fmt.Println(Is2Base4(i))
}