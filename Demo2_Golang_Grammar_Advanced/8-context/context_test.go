package __context

import (
	"context"
	"log"
	"testing"
	"time"
)

func Test_Context_Foundation(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "syk", "xh")
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	//ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	defer cancel()
}

func Test_Context_1(t *testing.T) {
	data := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	ch := make(chan []int)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "bzd", "wsm")
	ctx, cancelFunc := context.WithTimeout(ctx, time.Second*5)
	defer cancelFunc()

	go calculate(ctx, ch)
	for _, v := range data {
		ch <- v
	}

	time.Sleep(time.Second * 10)
}

func calculate(ctx context.Context, ch chan []int) {
	for {
		select {
		case i := <-ch:
			ctx = context.WithValue(ctx, "syk", "xhfyy")

			chSum := make(chan []int)
			go sumContext(ctx, chSum)
			chSum <- i

			chMulti := make(chan []int)
			go multiContext(ctx, chMulti)
			chMulti <- i

		case <-ctx.Done():
			log.Printf("calculate done, ctx syk: %s, ctx err: %s", ctx.Value("bzd").(string), ctx.Err())
			return
		}
	}
}

func sumContext(ctx context.Context, ch chan []int) {
	for {
		select {
		case i := <-ch:
			a, b := i[0], i[1]
			log.Printf("%d + %d = %d", a, b, sum(a, b))
		case <-ctx.Done():
			log.Printf("sumContext done, ctx syk: %s, ctx err: %s", ctx.Value("syk").(string), ctx.Err())
			return
		}
	}
}

func multiContext(ctx context.Context, ch chan []int) {
	for {
		select {
		case i := <-ch:
			a, b := i[0], i[1]
			log.Printf("%d * %d = %d", a, b, multi(a, b))
		case <-ctx.Done():
			log.Printf("multiContext done, ctx syk: %s, ctx, err: %s", ctx.Value("syk").(string), ctx.Err())
			return
		}
	}
}

func sum(a, b int) int {
	return a + b
}

func multi(a, b int) int {
	return a * b
}
