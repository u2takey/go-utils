package workqueue

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/panjf2000/ants/v2"
	"github.com/u2takey/go-utils/sets"
)

type TestPool struct {
	pool *ants.Pool
}

func NewTestPool(size int) *TestPool {
	p := &TestPool{}
	p.pool, _ = ants.NewPool(size)
	return p
}

func (p *TestPool) Submit(f func()) {
	_ = p.pool.Submit(f)
}

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func TestParallelizeWithPoolUntil(t *testing.T) {
	var l = 100
	for _, size := range []int{1, 2, 4} {
		c := make([]int, l)
		pool := NewTestPool(size)
		ParallelizeWithPoolUntil(context.Background(), l, l, pool, func(piece int) {
			c[piece] = GoID()
		})
		if sets.NewInt(c...).Len() != size {
			t.Errorf("expect: %d, got: %d", size, sets.NewInt(c...).Len())
		}
	}
}
