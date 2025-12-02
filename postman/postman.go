package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var ptm = map[int]string{
	1: "Letter from a friend",
	2: "Invitation to a bd",
	3: "Tax letter",
}

func Postman(wg *sync.WaitGroup, transferPoint chan<- string, ctx context.Context, n int, mail string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("I'm a postman number:", n, "my shift has finished")
			return
		default:
			fmt.Println("I'm a postman number:", n, "started delivering mail")
			time.Sleep(1 * time.Second)
			fmt.Println("I'm a postamn number:", n, "has delivered mail")
			transferPoint <- mail
			fmt.Println("I'm a postman number:", n, "has transfered mail", mail)
		}
	}
}

func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	mailTransferPoint := make(chan string)
	wg := &sync.WaitGroup{}

	for i := range postmanCount {
		wg.Add(1)
		go Postman(wg, mailTransferPoint, ctx, i+1, postmanToMail(i+1))
	}
	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()
	return mailTransferPoint
}

func postmanToMail(postmanNumber int) string {

	mail, ok := ptm[postmanNumber]
	if !ok {
		return "Lottery"
	}

	return mail
}
