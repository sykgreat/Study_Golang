package ArrayList

import (
	"testing"
)

func TestArrayListIterator(t *testing.T) {
	for it := list.Iterator(); it.HasNext(); {
		value := it.Next()
		if value == 2 {
			it.Remove()
		}
		t.Log(value)
	}
	t.Log(list)
}
