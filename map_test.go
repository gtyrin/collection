package collection

import (
	"fmt"
	"testing"
)

func TestMapDiff(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	upd, del, ins := MapDiff(m1, map[string]int8{"a": 3, "b": 2, "c": 1})
	if len(ins) != 0 || len(upd) != 0 || len(del) != 0 {
		t.Fail()
	}
	m2 := map[string]int{"a": 1, "b": 2, "c": 3}
	upd, del, ins = MapDiff(m1, m2)
	if len(ins) != 0 || len(upd) != 0 || len(del) != 0 {
		t.Fail()
	}
	m2["a"] = 11
	upd, del, ins = MapDiff(m1, m2)
	fmt.Println(ins, upd, del)
	if len(ins) != 0 || len(upd) != 1 || len(del) != 0 {
		t.Fail()
	}
	m2["a"] = 1
	delete(m1, "b")
	upd, del, ins = MapDiff(m1, m2)
	if len(ins) != 1 || len(upd) != 0 || len(del) != 0 {
		t.Fail()
	}
	upd, del, ins = MapDiff(m2, m1)
	if len(ins) != 0 || len(upd) != 0 || len(del) != 1 {
		t.Fail()
	}
}
