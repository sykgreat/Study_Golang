package _2_opterator

import "testing"

func Test_Arithmetic_Operator(t *testing.T) {
	a := 10
	b := 3
	var c int
	c = a + b
	t.Log("a + b = ", c)
	c = a - b
	t.Log("a - b = ", c)
	c = a * b
	t.Log("a * b = ", c)
	c = a / b
	t.Log("a / b = ", c)
	c = a % b
	t.Log("a % b = ", c)
	c++
	t.Log("c++ = ", c)
	c--
	t.Log("c-- = ", c)
}

func Test_Relational_Operator(t *testing.T) {
	a := 10
	b := 3
	t.Log("a = ", a, "b = ", b)
	t.Log("a == b = ", a == b)
	t.Log("a != b = ", a != b)
	t.Log("a > b = ", a > b)
	t.Log("a < b = ", a < b)
	t.Log("a >= b = ", a >= b)
	t.Log("a <= b = ", a <= b)
}

func Test_Logical_Operators(t *testing.T) {
	a := true
	b := false
	t.Log("a = ", a, "b = ", b)
	t.Log("a && b = ", a && b)
	t.Log("a || b = ", a || b)
	t.Log("!a = ", !a)
	t.Log("!b = ", !b)
}

func Test_Bitwise_Operator(t *testing.T) {
	var a uint8 = 60
	var b uint8 = 13
	var c uint8
	t.Logf("	  a = %08b\n", a)
	t.Logf("	  b = %08b\n", b)
	t.Logf("	  c = %08b\n", c)
	t.Log("----------------------------")
	c = a & b
	t.Logf("	  a = %08b\n", a)
	t.Logf("	  b = %08b\n", b)
	t.Logf("a & b = %08b", c)
	t.Log("----------------------------")
	c = a | b
	t.Logf("	  a = %08b\n", a)
	t.Logf("	  b = %08b\n", b)
	t.Logf("a | b = %08b", c)
	t.Log("----------------------------")
	c = a ^ b
	t.Logf("	  a = %08b\n", a)
	t.Logf("	  b = %08b\n", b)
	t.Logf("a ^ b = %08b", c)
	t.Log("----------------------------")
	c = a << 2
	t.Logf("	   a = %08b\n", a)
	t.Logf("a << 2 = %08b", c)
	t.Log("----------------------------")
	c = a >> 2
	t.Logf("	   a = %08b\n", a)
	t.Logf("a >> 2 = %08b", c)
	t.Log("----------------------------")
	c = ^a
	t.Logf("	  a = %08b\n", a)
	t.Logf("	 ^a = %08b", c)
	t.Log("----------------------------")
	c = a &^ b
	t.Logf("	   a = %08b\n", a)
	t.Logf("	   b = %08b\n", b)
	t.Logf("a &^ b = %08b", c)
	t.Log("----------------------------")
	c = a & ^b
	t.Logf("	   a = %08b\n", a)
	t.Logf("	   b = %08b\n", b)
	t.Logf("a & ^b = %08b", c)
	t.Log("----------------------------")
}

func Test_Assignment_Operator(t *testing.T) {
	a := 10
	c := a
	t.Log("a = c, c的值为：", c)
	c += a
	t.Log("c += a, c的值为：", c)
	c -= a
	t.Log("c -= a, c的值为：", c)
	c *= a
	t.Log("c *= a, c的值为：", c)
	c /= a
	t.Log("c /= a, c的值为：", c)
	c %= a
	t.Log("c %= a, c的值为：", c)

	var b uint8 = 60
	t.Logf("             b = %08b\n", b)
	b <<= 2
	t.Logf("b <<= 2, b的值为: %08b\n", b)
	b >>= 2
	t.Logf("b >>= 2, b的值为: %08b\n", b)
	b &= 2
	t.Logf("b &= 2,  b的值为: %08b\n", b)
	b ^= 2
	t.Logf("b ^= 2,  b的值为: %08b\n", b)
	b |= 2
	t.Logf("b |= 2,  b的值为: %08b\n", b)
}
