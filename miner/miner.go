package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Miner(ctx context.Context, wg *sync.WaitGroup, transferPoint chan<- int, n int, power int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("i'm a miner number:", n, "my shift finished")
			return
		default:
			fmt.Println("i'm a miner number:", n, "started mining coal")
			time.Sleep(1 * time.Second)
			fmt.Println("i'm a miner number:", n, "finished mining coal")
			transferPoint <- power
			fmt.Println("i'm a miner number:", n, "transfered this amount of coal", power)
		}
	}
}

func MinerPool(ctx context.Context, minerCount int) <-chan int {
	coalTransferPoint := make(chan int)

	wg := &sync.WaitGroup{}

	for i := range minerCount {
		wg.Add(1)
		go Miner(ctx, wg, coalTransferPoint, i+1, (i+1)*10)
	}

	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}
