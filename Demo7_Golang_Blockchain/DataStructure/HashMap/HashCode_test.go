package HashMap

import (
	"fmt"
	"testing"
	"unsafe"
)

type User struct {
	Name   string
	Age    int
	Sex    bool
	Friend []User
}

func init() {
	var user = User{
		Name: "张三",
		Age:  18,
		Sex:  true,
	}
	fmt.Println(Memhash(unsafe.Pointer(&user), 1, 36))

	fmt.Println(Strhash(user.Name))
}

func Test_HashCode(t *testing.T) {
	var user = &User{
		Name: "张三",
		Age:  18,
		Sex:  true,
	}
	fmt.Println(Memhash(unsafe.Pointer(&user), 1, 36))

	fmt.Println(Strhash(user.Name))
}
