package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

func fetchAll(ctx context.Context) error {
	errs, ctx := errgroup.WithContext(ctx)
	rand.Seed(time.Now().Unix())
	// run all the http requests in parallel
	for i := 0; i < 4; i++ {
		errs.Go(func() error {
			rnd := rand.Int() % 10
			var err error
			if rnd == 0 {
				err = fmt.Errorf("error in go routine, bailing")
			}
			select {
			case <-time.Tick(time.Duration(rand.Intn(100)) * time.Millisecond):
				if err == nil {
					fmt.Println("ok")
				}
				return err
			case <-ctx.Done():
				return nil
			}
		})
	}

	// Wait for completion and return the first error (if any)
	return errs.Wait()
}

func main() {
	fmt.Println(fetchAll(context.Background()))
}
