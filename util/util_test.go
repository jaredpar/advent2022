package util

import (
	"reflect"
	"testing"

	"github.com/jaredpar/advent2022/testUtil"
)

func TestSplitOnBlankSimple(t *testing.T) {
	items := SplitOnWhiteSpace(" 42 13")
	if !reflect.DeepEqual(items, []string{"42", "13"}) {
		t.Error("wrong items")
	}
}

type fakeInt int

func TestInsertSortedSimple(t *testing.T) {
	assert := testUtil.NewAssert(t)
	data := make([]fakeInt, 0, 10)
	for i := 0; i < 10; i++ {
		data = InsertSorted(data, fakeInt(i))
	}

	assert.EqualInt(10, len(data))
	for i := 0; i < len(data); i++ {
		assert.EqualInt(i, int(data[i]))
	}
}

type fakeString struct {
	value string
}

func TestInsertSortedFSimple(t *testing.T) {
	assert := testUtil.NewAssert(t)
	input := "adbc"
	data := make([]fakeString, 0)
	for _, r := range input {
		f := fakeString{value: string(r)}
		data = InsertSortedF(data, f, func(l, r fakeString) bool {
			return l.value < r.value
		})
	}

	assert.EqualString("a", data[0].value)
	assert.EqualString("b", data[1].value)
	assert.EqualString("c", data[2].value)
	assert.EqualString("d", data[3].value)
}

func TestStartsWithString(t *testing.T) {
	assert := testUtil.NewAssert(t)
	assert.True(StartsWithString("hello world", "hello"))
	assert.True(StartsWithString("hello world", ""))
	assert.True(StartsWithString("hello world", "h"))
	assert.False(StartsWithString("hello world", "dog"))
	assert.False(StartsWithString("", "dog"))
}
