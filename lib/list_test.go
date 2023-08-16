package lib

import (
	"testing"
	"container/list"
)

func TestList(t *testing.T) {
    var l1 = list.New()
    if ListRepr(l1) != "[]" {
        t.Errorf("expected: '[]', found: '%s'", ListRepr(l1))
    }
    l1.PushFront(1)
    if ListRepr(l1) != "[1]" {
        t.Errorf("expected: '[1]', found: '%s'", ListRepr(l1))
    }
    l1.PushBack(2)
    if ListRepr(l1) != "[1 2]" {
        t.Errorf("expected: '[1 2]', found: '%s'", ListRepr(l1))
    }
    l1.PushFront(0)
    if ListRepr(l1) != "[0 1 2]" {
        t.Errorf("expected: '[0 1 2]', found: '%s'", ListRepr(l1))
    }

    var l2 = list.New()
    var sum int = 0
    ListMirror(l2, l1, func(x, y int) bool {
        return x == y
    }, func (ev ListEvent[int]) {
        if ev.Kind != ElementAdded {
            t.Errorf("expected: %v, found: %v", ElementAdded, ev.Kind)
        } else {
            sum += ev.Element
        }
    })
    if sum != 3 {
        t.Errorf("expected: 3, found: %v", sum)
    }

    ListMap(l2, func (x int) {
        sum -= x
    })
    if sum != 0 {
        t.Errorf("expected: 0, found: %v", sum)
    }
}
