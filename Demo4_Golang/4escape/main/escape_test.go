package main

import (
	"testing"
)

func Test_Main1(t *testing.T) {
	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", &u2)
}
