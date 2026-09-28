package stats

import (
	"sync"
	"testing"
)

func TestStore_MostFrequent(t *testing.T) {
	s := NewStore()
	s.Record(Request{Int1: 3, Int2: 5, Limit: 15, Str1: "fizz", Str2: "buzz"})
	s.Record(Request{Int1: 2, Int2: 3, Limit: 6, Str1: "tic", Str2: "tac"})
	s.Record(Request{Int1: 2, Int2: 3, Limit: 6, Str1: "tic", Str2: "tac"})

	got, hits, ok := s.MostFrequent()
	want := Request{Int1: 2, Int2: 3, Limit: 6, Str1: "tic", Str2: "tac"}

	if !ok || got != want || hits != 2 {
		t.Errorf("got (%v, %d, %v), want (%v, 2, true)", got, hits, ok, want)
	}
}

func TestStore_ConcurrentRecord(t *testing.T) {
	const goroutines = 100
	const perGoroutine = 100

	s := NewStore()
	req := Request{Int1: 3, Int2: 5, Limit: 15, Str1: "fizz", Str2: "buzz"}

	var wg sync.WaitGroup
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < perGoroutine; j++ {
				s.Record(req)
			}
		}()
	}
	wg.Wait()

	_, hits, _ := s.MostFrequent()
	want := goroutines * perGoroutine
	if hits != want {
		t.Errorf("hits = %d, want %d", hits, want)
	}
}