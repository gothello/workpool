package pool

import (
	"context"
	"sync"
)

type Task interface {
	Run()
}

type Worker chan Task

type Pool struct {
	Workers chan Worker
	wg      *sync.WaitGroup
	cancel  context.CancelFunc
}

func (p *Pool) Start(ctx context.Context) {
	worker := make(Worker)
	for {
		select {
		case p.Workers <- worker:
			task := <-worker
			task.Run()
		case <-ctx.Done():
			p.wg.Done()
			return
		}

	}
}

func (p *Pool) Stop() {
	p.cancel()
	p.wg.Wait()
}

func New(workers int) *Pool {

	ctx, cancel := context.WithCancel(
		context.Background(),
	)

	pool := &Pool{
		Workers: make(chan Worker),
		wg:      &sync.WaitGroup{},
		cancel:  cancel,
	}

	for i := 0; i < workers; i++ {
		pool.wg.Add(1)
		go pool.Start(ctx)
	}

	return pool
}
