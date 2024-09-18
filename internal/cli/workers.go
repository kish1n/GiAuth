package cli

import (
	"context"
	"github.com/kish1n/GiAuth/internal/config"
	"github.com/kish1n/GiAuth/internal/service"
	"sync"
)

func runServices(ctx context.Context, cfg config.Config, wg *sync.WaitGroup) {
	run := func(f func()) {
		wg.Add(1)
		go func() {
			f()
			wg.Done()
		}()
	}

	run(func() { service.Run(ctx, cfg) })
}
