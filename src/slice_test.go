package collection

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.Zero(t, Index("John Doe", objs))
	assert.NotZero(t, Index("John Travolta", objs))
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
	assert.Empty(t, ins)
	assert.Empty(t, upd)
	assert.Empty(t, del)
	sl2 := []int{1, 2, 3}
	upd, del, ins = SliceDiff(sl1, sl2)
	assert.Empty(t, ins)
	assert.Empty(t, upd)
	assert.Empty(t, del)
	sl2[1] = 4
	upd, del, ins = SliceDiff(sl1, sl2)
	assert.Empty(t, ins)
	assert.Len(t, upd, 1)
	assert.Equal(t, upd[0].elem, 4)
	assert.Empty(t, del)
	sl2 = append(sl2[:1], sl2[2])
	upd, del, ins = SliceDiff(sl1, sl2) // {1,2,3} vs {1,3}
	assert.Empty(t, ins)
	assert.Empty(t, upd)
	assert.Len(t, del, 1)
	assert.Equal(t, del[0].elem, 2)
	upd, del, ins = SliceDiff(sl2, sl1) // {1,3} vs {1,2,3}
	assert.Len(t, ins, 1)
	assert.Equal(t, ins[0].elem, 2)
	assert.Empty(t, upd)
	assert.Empty(t, del)
	sl1[2] = 5
	upd, del, ins = SliceDiff(sl2, sl1) // {1,3} vs {1,2,5}
	assert.Len(t, ins, 2)
	assert.Empty(t, upd)
	assert.Len(t, del, 1)
}
