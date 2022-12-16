package Concurrency

import (
	"fmt"
	"testing"
)

func Benchmark_primeNumber(b *testing.B) {
	primeNumber(100)
}

func primeNumber(num int) {
	for i := 0; i < num; i++ {
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i)
		}
	}
}
