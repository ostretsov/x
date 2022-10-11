package x

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	t.Run("nilSlice", func(t *testing.T) {
		var s []int
		got := Filter(s, func(_ int) bool { return true })
		if got != nil {
			t.Errorf("want nil slice, got %#v", got)
		}
	})

	t.Run("emptySlice", func(t *testing.T) {
		s := []int{}
		got := Filter(s, func(_ int) bool { return true })
		if len(got) != 0 {
			t.Error("want 0 length; got ", len(got))
		}
	})

	t.Run("filterNegativeInts", func(t *testing.T) {
		s := []int{-5, -3, -1, 0, 1, 3, 5}
		want := []int{0, 1, 3, 5}
		got := Filter(s, func(v int) bool { return v >= 0 })
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %#v; got %#v", want, got)
		}
	})
}

func TestContains(t *testing.T) {
	if Contains(nil, 0) {
		t.Error("nil slice shouldn't contain 0")
	}
	if Contains([]int{}, 0) {
		t.Error("empty slice shouldn't contain 0")
	}
	if !Contains([]int{0, 1, 2, 3, 1}, 1) {
		t.Errorf("should contain 1")
	}
}
