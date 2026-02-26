package stack

import "testing"

func TestStack_PopOrder(t *testing.T) {
    var s StackInt

    s.Push(1)
    s.Push(2)
    s.Push(3)

    value, ok := s.Pop()

    if !ok {
        t.Fatalf("expected ok=true, got false")
    }

    if value != 3 {
        t.Fatalf("expected 3, got %d", value)
    }

    value, ok = s.Pop()

    if !ok {
        t.Fatalf("expected ok=true, got false")
    }

    if value != 2 {
        t.Fatalf("expected 2, got %d", value)
    }

    value, ok = s.Pop()

    if !ok {
        t.Fatalf("expected ok=true, got false")
    }

    if value != 1 {
        t.Fatalf("expected 1, got %d", value)
    }

    if !s.IsEmpty() {
    t.Fatalf("expected stack to be empty")
    }

    _, ok = s.Pop()
    if ok {
        t.Fatalf("expected ok=false on empty stack")
    }
}

func TestStack_PopEmpty(t *testing.T) {
    var s StackInt
    _, ok := s.Pop()
    if ok {
        t.Fatalf("expected ok=false on empty stack")
    }
}

func TestStack_PeekDoesNotRemove(t *testing.T) {
    var s StackInt
    s.Push(42)
    v, ok := s.Peek()
    if s.Len() != 1 {
        t.Fatalf("expected len=1 after Peek")
    }
    if !ok {
        t.Fatalf("expexted ok=true from Peek")
    }
    if v != 42 {
        t.Fatalf("expected 42 from Peek, got %d", v)
    }
    v2, ok2 := s.Pop()
    if !ok2 {
        t.Fatalf("expected ok=true from Pop")
    }
    if v2 != 42 {
        t.Fatalf("expected 42 from Pop, got %d", v2)
    }
    if !s.IsEmpty() {
        t.Fatalf("expected stack to be empty")
    }
}

