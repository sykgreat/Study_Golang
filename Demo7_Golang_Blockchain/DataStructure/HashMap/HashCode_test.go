package HashMap

import (
	"testing"
	"unsafe"
)

type User struct {
	Name   string
	Age    int
	Sex    bool
	Friend []User
}

func Test_HashCode(t *testing.T) {
	user := &User{
		Name: "张三",
		Age:  18,
		Sex:  true,
	}
	u := memhash(unsafe.Pointer(&user), 1, 36)
	t.Log(u % 1000)
}
