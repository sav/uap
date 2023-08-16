package lib

import (
	"container/list"
	"fmt"
)

type ListEventType int8

const (
	ElementAdded ListEventType = iota
	ElementRemoved
)

type ListEvent[T any] struct {
	Kind    ListEventType
	Element T
}

type ListResult struct {
	Element *list.Element
	Found   bool
}

func ListRepr(l *list.List) string {
	s := "["
	for e := l.Front(); e != nil; e = e.Next() {
		s += fmt.Sprintf("%#v", e.Value)
		if e.Next() != nil {
			s += " "
		}
	}
	s += "]"
	return s
}

func ListIn[T any](
	l *list.List,
	m *list.List,
	compare func(a T, b T) bool,
	callback func(*list.Element, bool),
) {
	var res = []ListResult{}
	for a := l.Front(); a != nil; a = a.Next() {
		found := false
		for b := m.Front(); b != nil; b = b.Next() {
			if compare(a.Value.(T), b.Value.(T)) {
				found = true
				break
			}
		}
		res = append(res, ListResult{a, found})
	}
	for _, r := range res {
		callback(r.Element, r.Found)
	}
}

func ListMap[T any](l *list.List, callback func(T)) {
	for e := l.Front(); e != nil; e = e.Next() {
		callback(e.Value.(T))
	}
}

// ListMirror the elements of `m' in `l', i.e., remove elements from `l' that do
// not belong to `m' and add elements from `m' that do not belong to `l'.
func ListMirror[T any](
	l *list.List,
	m *list.List,
	compare func(T, T) bool,
	callback func(ListEvent[T]),
) {
	ListIn(m, l, compare, func(e *list.Element, found bool) {
		if !found {
			l.PushBack(e.Value.(T))
			callback(ListEvent[T]{ElementAdded, e.Value.(T)})
		}
	})
	ListIn(l, m, compare, func(e *list.Element, found bool) {
		if !found {
			l.Remove(e)
			callback(ListEvent[T]{ElementRemoved, e.Value.(T)})
		}
	})
}
