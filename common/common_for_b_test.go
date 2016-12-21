package common

import (
	// "fmt"
	"testing"
)

func Benchmark_For(b *testing.B) {
	//	i := 0
	//	for ; i < b.N; i++ { //use b.N for looping
	//		b.Log(i)
	//	}

	for i := 0; i < b.N; i++ { //use b.N for looping
		b.Log(i)
	}
}
