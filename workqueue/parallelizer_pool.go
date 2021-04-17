package workqueue

import (
	"context"
	"sync"

	utilruntime "github.com/u2takey/go-utils/runtime"
)

type GoroutinePool interface {
	Submit(func())
}

func ParallelizeWithPoolUntil(ctx context.Context, workers, pieces int, pool GoroutinePool, doWorkPiece DoWorkPieceFunc) {
	var stop <-chan struct{}
	if ctx != nil {
		stop = ctx.Done()
	}

	toProcess := make(chan int, pieces)
	for i := 0; i < pieces; i++ {
		toProcess <- i
	}
	close(toProcess)

	if pieces < workers {
		workers = pieces
	}

	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		pool.Submit(
			func() {
				defer utilruntime.HandleCrash()
				defer wg.Done()
				for piece := range toProcess {
					select {
					case <-stop:
						return
					default:
						doWorkPiece(piece)
					}
				}
			})
	}
	wg.Wait()
}
