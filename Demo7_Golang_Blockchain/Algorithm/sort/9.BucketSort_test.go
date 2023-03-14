package sort

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	d := sort.Search(len(a), func(i int) bool { return a[i] >= 3 })
	fmt.Println(d)
}
