package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrMap(t *testing.T) {
	m := NewStrMap()
	assert.Empty(t, m)
	m.Add("a", "1")
	assert.True(t, m.Exists("a"))
	assert.Equal(t, m.Value("a"), "1")
	m.Delete("a")
	assert.Empty(t, m)
	assert.True(t, m.IsEmpty())
	m.Clean()
	assert.Empty(t, m)
}

func TestMapDiff(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	upd, del, ins := MapDiff(m1, map[string]int8{"a": 3, "b": 2, "c": 1})
	assert.Empty(t, ins)
	assert.Empty(t, upd)
	assert.Empty(t, del)
	m2 := map[string]int{"a": 1, "b": 2, "c": 3}
	upd, del, ins = MapDiff(m1, m2)
	assert.Empty(t, ins)
	assert.Empty(t, upd)
	assert.Empty(t, del)
	m2["a"] = 11
	upd, del, ins = MapDiff(m1, m2)
	assert.Empty(t, ins)
	assert.Len(t, upd, 1)
	assert.Empty(t, del)
	m2["a"] = 1
	delete(m1, "b")
	upd, del, ins = MapDiff(m1, m2)
	assert.Len(t, ins, 1)
	assert.Empty(t, upd)
	assert.Empty(t, del)
	upd, del, ins = MapDiff(m2, m1)
	assert.Empty(t, ins)
	assert.Empty(t, upd)
	assert.Len(t, del, 1)
}
