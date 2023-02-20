package cache

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetValueSize(t *testing.T) {
	type args struct {
		value interface{}
	}
	var tests []struct {
		name string
		args args
		want int64
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValueSize(tt.args.value); got != tt.want {
				t.Errorf("GetValueSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSize(t *testing.T) {
	type args struct {
		size string
	}
	var tests []struct {
		name  string
		args  args
		want  int64
		want1 string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ParseSize(tt.args.size)
			if got != tt.want {
				t.Errorf("ParseSize() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ParseSize() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_reflect(t *testing.T) {
	//t.Log(GetValueSizeA(true))
	t.Log(GetValueSize(map[string]interface{}{
		"1":         "1",
		"2":         "2",
		"232432":    "2",
		"2fds32432": []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
	}))
	//t.Log(GetValueSize([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 15215158, 12, 0, 30, 03, 26, 0}))
	//t.Log(GetValueSize([]string{"1", "2", "232432", "2fds32432", "1", "2", "232432", "2fds32432", "1", "2", "232432", "2fds32432"}))

	//t.Log(GetValueSizeA(NewMemCache()))
}

func GetValueSizeA(value interface{}) int64 {
	vo := reflect.ValueOf(value)
	fmt.Println(vo.Type())
	return 0
}
