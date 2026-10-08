package main
import "fmt"

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val T
}

func (Lst *List[T]) Push(v T) {
	if Lst.tail == nil {
		Lst.head = &element[T]{val: v}
		Lst.tail = Lst.head
	} else {
		Lst.tail.next = &element[T]{val: v}
		Lst.tail = Lst.tail.next
	}
}

func (Lst *List[T]) AllElements() []T {
	var elems []T
	for e := Lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))
	_ = SlicesIndex[[]string, string](s, "zoo")

	Lst := List[int]{}
	Lst.Push(10)
	Lst.Push(13)
	Lst.Push(23)
	fmt.Println("list:", Lst.AllElements())
}
