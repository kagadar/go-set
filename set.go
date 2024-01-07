package set

import "github.com/kagadar/go-pipeline"

type Empty struct{}

type Set[K comparable] map[K]Empty

func New[K comparable](k ...K) Set[K] {
	s := Set[K]{}
	s.Put(k...)
	return s
}

func (s Set[K]) Has(k K) bool {
	_, ok := s[k]
	return ok
}

func (s Set[K]) Put(k ...K) {
	for _, e := range k {
		s[e] = Empty{}
	}
}

func (s Set[K]) Elements() []K {
	return pipeline.MapToSlice(s, func(k K, _ Empty) K { return k })
}
