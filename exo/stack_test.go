package exo

import (
    "errors"
    "testing"
)

func TestStackInt(t *testing.T) {
    s := NewStack[int]()

    if !s.IsEmpty() {
        t.Error("nouveau stack devrait être vide")
    }

    _, err := s.Pop()
    if !errors.Is(err, ErrEmptyStack) {
        t.Errorf("Pop sur stack vide: attendait ErrEmptyStack, got %v", err)
    }

    s.Push(1)
    s.Push(2)
    s.Push(3)

    if s.Size() != 3 {
        t.Errorf("Size = %d, want 3", s.Size())
    }

    top, err := s.Peek()
    if err != nil || top != 3 {
        t.Errorf("Peek = %d, %v, want 3, nil", top, err)
    }

    v, _ := s.Pop()
    if v != 3 {
        t.Errorf("Pop = %d, want 3", v)
    }
    v, _ = s.Pop()
    if v != 2 {
        t.Errorf("Pop = %d, want 2", v)
    }

    if s.Size() != 1 {
        t.Errorf("après 2 pops, Size = %d, want 1", s.Size())
    }
}

func TestStackString(t *testing.T) {
    s := NewStack[string]()
    s.Push("a")
    s.Push("b")
    v, _ := s.Pop()
    if v != "b" {
        t.Errorf("got %q, want %q", v, "b")
    }
}