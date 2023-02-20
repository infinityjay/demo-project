package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	stopCh := make(chan bool)
	defer func() {
		close(stopCh)
		log.Info().Msg("websocket connection closed")
	}()
	count := 0
	g.Go(func() error {
		for {
			select {
			case <-stopCh:
				fmt.Println("first go routine stoped by ch!")
				return nil
			case <-ctx.Done():
				fmt.Println("first go routine Canceled!")
				return ctx.Err()
			default:
				for _, r := range `-\|/` {
					fmt.Printf("\r%c\n", r)
					time.Sleep(1 * time.Second)
				}
				count++
				if count > 20 {
					return nil
				}
			}
		}
	})

	for i := 0; i < 20; i++ {
		time.Sleep(1 * time.Second)
		if i > 10 {
			return
		}
		g.Go(func() error {
			select {
			case <-stopCh:
				fmt.Println("stoped by ch!")
				return nil
			case <-ctx.Done():
				fmt.Println("goroutine Canceled:", i)
				return nil
			default:
				if i > 15 {
					//fmt.Println("go routine Error:", i)
					return fmt.Errorf("go routine Error: %d", i)
				}
				fmt.Println("go routine End:", i)
			}
			return nil
		})

	}
	if err := g.Wait(); err != nil {
		log.Info().Msg(err.Error())
	}
}
