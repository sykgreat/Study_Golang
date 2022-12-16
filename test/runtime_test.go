package test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestC(t *testing.T) {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOROOT())
}
