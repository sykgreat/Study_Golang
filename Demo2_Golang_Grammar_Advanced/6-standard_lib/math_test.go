package __standard_lib

import (
	"fmt"
	"math"
	"testing"
)

func Test_Math(t *testing.T) {
	fmt.Println("2的10次方：", 2<<9, math.Pow(2, 10))
	fmt.Println("返回二为底，1024的对数", math.Log2(1024))
	fmt.Println("返回两个数中较大的数", math.Max(1, 2))
	fmt.Println("向上取值", math.Ceil(2.1))
	fmt.Println("向下取值", math.Floor(2.9))
	fmt.Println("四舍五入", math.Round(2.5), math.Round(2.4))
	fmt.Println("绝对值", math.Abs(-2.5))
	fmt.Println("45的正弦值", math.Sin(45))
	fmt.Println("45的余弦值", math.Cos(45))
	fmt.Println("45的正切值", math.Tan(45))
	fmt.Println("45的反正切值", math.Atan(45))
}
