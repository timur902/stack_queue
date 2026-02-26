package stack

import "testing"

func TestQueue_Order(t *testing.T) {
	var q QueueInt

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	
	v, ok := q.Dequeue()
	if v != 1 {
		t.Fatalf("expected 1 from Dequeue, got %d", v)
	}
	if !ok {
		t.Fatalf("expected ok=true")
	}
	v1, ok1 := q.Dequeue()
	if v1 != 2 {
		t.Fatalf("expected 2 from Dequeue, got %d", v1)
	}
	if !ok1 {
		t.Fatalf("expected ok=true")
	}
	v2, ok2 := q.Dequeue()
	if v2 != 3 {
		t.Fatalf("expected 3 from Dequeue, got %d", v2)
	}
	if !ok2 {
		t.Fatalf("expected ok=true")
	}
	_, ok3 := q.Dequeue()
	if ok3 {
		t.Fatalf("expected ok=false")
	}
}

func TestQueue_PeekDoesNotRemove(t *testing.T) {
	var q QueueInt
	q.Enqueue(42)
	v, ok := q.Peek()
	if q.Len() != 1 {
		t.Fatalf("expected len=1 after Peek")
	}
	if !ok {
		t.Fatalf("expected ok=true from Peek")
	}
	if v != 42 {
		t.Fatalf("expected 42 from Peek, got %d", v)
	}
	if q.IsEmpty() {
		t.Fatalf("expected IsEmpty=false")
	}
	v2, ok2 := q.Dequeue()
	if !ok2 {
		t.Fatalf("expected ok=true from Dequeue")
	}
	if v2 != 42 {
		t.Fatalf("expected 42 from Dequeue, got %d", v2)
	}
	if !q.IsEmpty() {
		t.Fatalf("expected IsEmpty=true")
	}
}