package main

import (
	"context"
	"fmt"
	"hw13/miner"
	"hw13/postman"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var coal atomic.Int64
	var mails []string
	var mtx sync.Mutex
	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	coalTransferPoint := miner.MinerPool(minerContext, 300)
	mailTransferPoint := postman.PostmanPool(postmanContext, 300)

	tnow := time.Now()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("-----MINER SHIFT FINISHED-----")
		minerCancel()
	}()

	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println("-----POSTMAN SHIFT FINISHED-----")
		postmanCancel()
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for v := range coalTransferPoint {
			coal.Add(int64(v))
		}
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()
		for m := range mailTransferPoint {
			mtx.Lock()
			mails = append(mails, m)
			mtx.Unlock()
		}
	}()

	wg.Wait()

	fmt.Println("all coal", coal.Load())

	mtx.Lock()
	fmt.Println("mails:", len(mails))
	mtx.Unlock()

	fmt.Println("total time:", time.Since(tnow))
}
