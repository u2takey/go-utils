package buffer

import (
	"sync"
)

type LockQueue struct {
	data []interface{}
	sync.Mutex
	pos int // next position = length
}

// NewLockQueue with init capacity
func NewLockQueue(capacity uint64) *LockQueue {
	return &LockQueue{data: make([]interface{}, capacity)}
}

func (l *LockQueue) Put(a interface{}) {
	l.Lock()
	l.data[l.pos] = a
	l.pos += 1
	if l.pos >= len(l.data) {
		tmp := make([]interface{}, len(l.data)*2)
		copy(tmp, l.data)
		l.data = tmp
	}
	l.Unlock()
}

func (l *LockQueue) Get() (interface{}, bool) {
	l.Lock()
	if l.pos == 0 {
		l.Unlock()
		return nil, false
	}
	ret := l.data[l.pos-1]
	l.pos--
	l.Unlock()
	return ret, true
}

func (l *LockQueue) GetAll() []interface{} {
	l.Lock()
	ret := make([]interface{}, l.pos)
	copy(ret, l.data[:l.pos])
	l.pos = 0
	l.Unlock()
	return ret
}
