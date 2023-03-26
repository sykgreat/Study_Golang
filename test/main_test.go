package test

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func Test_reflect(t *testing.T) {
	//t.Log(GetValueSize(true))
	t.Log(GetValueSize(map[string]interface{}{
		"1": 123,
		"2": "456",
		//"3": "789",
		//"4": []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
	}))
	//t.Log(GetValueSize([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
	//t.Log(GetValueSize([]string{"1", "2"}))
	//t.Log(GetValueSize(123))
	//t.Log(GetValueSize("123"))
	t.Log(GetValueSize(
		users{
			Name: "123",
			Age:  123,
			Sex:  true,
			Bzd: map[string]interface{}{
				"1": 123,
				"2": "456",
			},
		},
	))
}

func GetValueSize(value interface{}) int64 {
	var size int64 = 0
	vo := reflect.ValueOf(value)
	switch vo.Kind() {
	case reflect.Map:
		for _, key := range vo.MapKeys() {
			size += GetValueSize(key.Interface())
			size += GetValueSize(vo.MapIndex(key).Interface())
		}
	case reflect.Slice:
		for i := 0; i < vo.Len(); i++ {
			size += GetValueSize(vo.Index(i).Interface())
		}
	case reflect.Struct:
		for i := 0; i < vo.NumField(); i++ {
			size += GetValueSize(vo.Field(i).Interface())
		}
	}
	size += int64(vo.Type().Size())
	return size
}

type users struct {
	Name string
	Sex  bool
	Age  int
	Bzd  map[string]interface{}
}

func bibao() (val int) {
	val = 10
	defer func() {
		val++
	}()
	return 100
}

func Test_bibao(t *testing.T) {
	t.Log(bibao())
}

func defer1() func() {
	defer fmt.Println("before return")
	return func() {
		defer fmt.Println("in return")
	}
}
func Test_defer(t *testing.T) {
	m := 10
	defer fmt.Printf("first defer %d\n", m)
	m = 100
	defer func() {
		fmt.Printf("second defer %d\n", m)
	}()

	m *= 10
	defer fmt.Printf("third defer %d\n", m)

	funcVal := defer1()
	funcVal()
}

func TestName1(t *testing.T) {
	t.Log(4 << (^uintptr(0) >> 63))
}

func Test_SizeOf(t *testing.T) {
	var boolSize bool
	t.Log("bool size of: ", unsafe.Sizeof(boolSize))
	t.Log()

	var uint8Size uint8
	t.Log("uint8 size of: ", unsafe.Sizeof(uint8Size))
	var uint16Size uint16
	t.Log("uint16 size of: ", unsafe.Sizeof(uint16Size))
	var uint32Size uint32
	t.Log("uint32 size of: ", unsafe.Sizeof(uint32Size))
	var uint64Size uint64
	t.Log("uint64 size of: ", unsafe.Sizeof(uint64Size))
	t.Log()

	var int8Size int8
	t.Log("int8 size of: ", unsafe.Sizeof(int8Size))
	var int16Size int16
	t.Log("int16 size of: ", unsafe.Sizeof(int16Size))
	var int32Size int32
	t.Log("int32 size of: ", unsafe.Sizeof(int32Size))
	var int64Size int64
	t.Log("int64 size of: ", unsafe.Sizeof(int64Size))
	t.Log()

	var float32Size float32
	t.Log("float32 of: ", unsafe.Sizeof(float32Size))
	var float64Size float64
	t.Log("float64 size of: ", unsafe.Sizeof(float64Size))
	t.Log()

	var byteSize byte
	t.Log("byte size of: ", unsafe.Sizeof(byteSize))
	var runeSize rune
	t.Log("rune size of: ", unsafe.Sizeof(runeSize))
	t.Log()

	var stringSize string
	t.Log("string size of: ", unsafe.Sizeof(stringSize))
	t.Log()

	var preSize uintptr
	t.Log("uintptr size of: ", unsafe.Sizeof(preSize))
	t.Log()
}

func Test_(t *testing.T) {

}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	for l1 != nil || l2 != nil {

	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
