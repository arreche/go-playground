package domain

import (
	"testing"
)

func TestSave(t *testing.T) {
	r := NewRepository()
	want := &Entity{Name: "Test"}
	got, err := r.Save(want)
	if err != nil {
		t.Error("unexpected error")
		return
	}
	if got.ID != want.ID {
		t.Errorf("Save(%v), want %v, got %v", want, got, want)
		return
	}
}
