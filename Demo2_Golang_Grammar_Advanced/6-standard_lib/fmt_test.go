package __standard_lib

import (
	"fmt"
	"os"
	"testing"
)

func Test_fmt(t *testing.T) {
	fmt.Println("hello world")

	fmt.Printf("hello world, %s\n", "nick1")

	sprintf := fmt.Sprintf("hello world, %s\n", "nick2")
	fmt.Println(sprintf)

	n, err := fmt.Fprint(os.Stderr, sprintf)
	fmt.Println(n, err)
}

func Test_Placeholder(t *testing.T) {
	type simple struct {
		key   int
		value string
	}

	s := simple{key: 1, value: "nick"}

	// %v	值的默认格式表示
	fmt.Printf("%v\n", s)

	// %+v	类似 %v，但输出结构体时会添加字段名
	fmt.Printf("%+v\n", s)

	// %#v	值的Go语法表示
	fmt.Printf("%#v\n", s)

	// %T	值的类型的Go语法表示
	fmt.Printf("%T\n", s)

	// %%	字面上的百分号，并非值的占位符
	fmt.Printf("%%\n")

	// %b	表示为二进制
	fmt.Printf("%b\n", 10)

	// %c	该值对应的unicode码值
	fmt.Printf("%c\n", 20170)

	// %d	表示为十进制
	fmt.Printf("%d\n", 10)

	// %o	表示为八进制
	fmt.Printf("%o\n", 10)

	// %q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	fmt.Printf("%q\n", "nick")

	// %x	表示为十六进制，使用a-f
	fmt.Printf("%x\n", 11)

	// %X	表示为十六进制，使用A-F
	fmt.Printf("%X\n", 11)

	// %U	表示为Unicode格式：U+1234，等价于 "U+%04X"
	fmt.Printf("%U\n", 10)

	// %f	表示为浮点数
	fmt.Printf("%f\n", 10.1)

	// %e	表示为科学计数法，e.g. -1.234456e+78
	fmt.Printf("%e\n", 10.1)

	// %E	表示为科学计数法，e.g. -1.234456E+78
	fmt.Printf("%E\n", 10.1)

	// %s	表示为基本的字符串
	fmt.Printf("%s\n", "nick")

	// %q	表示为双引号括起来的字符串
	fmt.Printf("%q\n", "nick")

	// %p	表示为十六进制，并加上前缀 0x
	fmt.Printf("%p\n", &s)
}
