package main

type user struct {
	name  string
	email string
}

func main() {
	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", &u2)
}

//go:noinline	//阻止编译器内联
func createUserV1() user {
	u1 := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V1", &u1)
	return u1
}

//go:noinline	//阻止编译器内联
func createUserV2() *user {
	u2 := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V2", &u2)
	return &u2
}
