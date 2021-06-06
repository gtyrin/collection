package collection

import (
	"reflect"
	"testing"
)

type testData struct {
	Name string
	Age  int
}

func TestContains(t *testing.T) {
	objs := []*testData{{Name: "John Doe", Age: 33}}
	if !Contains(&testData{Name: "John Doe", Age: 33}, objs) {
		t.Fail()
	}
	if Contains(&testData{Name: "John Doe", Age: 10}, objs) {
		t.FailNow()
	}
}

func TestContainsStr(t *testing.T) {
	objs := []string{"John Doe"}
	if !ContainsStr("John Doe", objs) {
		t.Fail()
	}
	if ContainsStr("", objs) {
		t.FailNow()
	}
}

func TestIndex(t *testing.T) {
	objs := []string{"John Doe"}
	if Index("John Doe", objs) != 0 {
		t.FailNow()
	}
	if Index("John Travolta", objs) == 0 {
		t.FailNow()
	}
}

func TestRemoveStr(t *testing.T) {
	objs := []string{"A", "B"}
	if !reflect.DeepEqual(RemoveStr("A", objs), []string{"B"}) {
		t.Fail()
	}
	if !reflect.DeepEqual(RemoveStr("C", objs), objs) {
		t.FailNow()
	}
}

func TestSplitWithTrim(t *testing.T) {
	if !reflect.DeepEqual(SplitWithTrim("1, 2, 3", ","), []string{"1", "2", "3"}) {
		t.FailNow()
	}
}

func TestSliceDiff(t *testing.T) {
	sl1 := []int{1, 2, 3}
	upd, del, ins := SliceDiff(sl1, []int8{1, 2, 3})
	if len(ins) != 0 || len(upd) != 0 || len(del) != 0 {
		t.Fail()
	}
	sl2 := []int{1, 2, 3}
	upd, del, ins = SliceDiff(sl1, sl2)
	if len(ins) != 0 || len(upd) != 0 || len(del) != 0 {
		t.Fail()
	}
	sl2[1] = 4
	upd, del, ins = SliceDiff(sl1, sl2)
	if len(ins) != 0 || len(upd) != 1 || upd[0].elem != 4 || len(del) != 0 {
		t.Fail()
	}
	sl2 = append(sl2[:1], sl2[2])
	upd, del, ins = SliceDiff(sl1, sl2) // {1,2,3} vs {1,3}
	if len(ins) != 0 || len(upd) != 0 || len(del) != 1 || del[0].elem != 2 {
		t.Fail()
	}
	upd, del, ins = SliceDiff(sl2, sl1) // {1,3} vs {1,2,3}
	if len(ins) != 1 || ins[0].elem != 2 || len(upd) != 0 || len(del) != 0 {
		t.Fail()
	}
	sl1[2] = 5
	upd, del, ins = SliceDiff(sl2, sl1) // {1,3} vs {1,2,5}
	if len(ins) != 2 || len(upd) != 0 || len(del) != 1 {
		t.Fail()
	}
}
