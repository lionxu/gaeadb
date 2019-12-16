package prefix

import (
	"gaeadb/cache"
	"gaeadb/locker"
	"gaeadb/stack"
	"gaeadb/suffix"
)

const (
	RootPage = int64(0)
)

const (
	R = iota // root node
	E        // entry of branch
	P        // prefix node
	S        // suffix node
	C        // character
)

type Tree interface {
	Close() error

	Get([]byte) (uint64, error)
	Del([]byte, suffix.Writer) error
	Set([]byte, uint64, suffix.Writer) error

	NewForwardIterator([]byte) (Iterator, error)
	NewBackwardIterator([]byte) (Iterator, error)
}

type Iterator interface {
	Close() error

	Next() error
	Valid() bool
	Key() []byte
	Value() uint64
}

type resource struct {
	pg cache.Page
	le locker.Locker
}

type forwardElement struct {
	typ  int
	cnt  int
	val  uint64
	pref []byte
	rsrc *resource
	itr  suffix.Iterator
}

type forwardIterator struct {
	t *tree
	s stack.Stack
}

type backwardElement struct {
	typ  int
	cnt  int
	val  uint64
	pref []byte
	rsrc *resource
	itr  suffix.Iterator
}

type backwardIterator struct {
	t *tree
	s stack.Stack
}

type tree struct {
	c cache.Cache
	t locker.Table
}