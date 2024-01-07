package set

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kagadar/go-pipeline"
)

func TestHas(t *testing.T) {
	if !New(1, 2, 3).Has(1) {
		t.Error("Has(1) returned false, expected true")
	}
	if New(1, 2, 3).Has(4) {
		t.Error("Has(4) returned true, expected false")
	}
}

func TestPut(t *testing.T) {
	s := New(1)
	if s.Has(2) {
		t.Error("Has(2 before put) returned true, expected false")
	}
	s.Put(2)
	if !s.Has(2) {
		t.Error("Has(2 after put) returned false, expected true")
	}
}

func TestElements(t *testing.T) {
	if diff := cmp.Diff(
		New(1, 2, 3).Elements(),
		[]int{1, 2, 3},
		cmpopts.SortSlices(pipeline.Less[int]),
	); diff != "" {
		t.Errorf("Elements() unexpected diff (-got +want):\n%s", diff)
	}
}
