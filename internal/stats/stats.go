package stats

import "sync"

//Request = combination of fizzbuzz params
type Request struct {
	Int1  int 
	Int2  int
	Limit int
	Str1  string
	Str2  string
}

//store = nb of times each Request has been received
type Store struct {
	mu       sync.Mutex 
	counts   map[Request]int
	top      Request
	nbHits   int
}

func NewStore() *Store {
	return &Store{counts: make(map[Request]int)}
}

func (s *Store) Record(r Request) {
	s.mu.Lock()
	defer s.mu.Unlock()


	s.counts[r]++
	if s.counts[r] > s.nbHits {
		s.top = r
		s.nbHits = s.counts[r]
	}
}

func (s *Store) MostFrequent() (req Request, hits int, ok bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.top, s.nbHits, s.nbHits > 0
}