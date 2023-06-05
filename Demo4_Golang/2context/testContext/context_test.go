package testContext

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContextEmptyCtx(t *testing.T) {
	ctx := context.Background()
	fmt.Println(ctx)
}

func TestContextCancelCtx(t *testing.T) {
	background := context.Background()
	ctx, cancel := context.WithCancel(background)

	go goRun(ctx, "goT1")
	go goRun(ctx, "goT2")
	go goRun(ctx, "goT3")

	time.Sleep(time.Second * 5)
	cancel()
	fmt.Println("goT停止")
	time.Sleep(time.Second * 1)
}

func TestContextTimerCtxWT(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)

	go goRun(ctx, "goT1")
	go goRun(ctx, "goT2")
	go goRun(ctx, "goT3")

	time.Sleep(time.Second * 5)
	fmt.Println("goT停止")
}

func TestContextTimerCtxWD(t *testing.T) {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))

	go goRun(ctx, "goT1")
	go goRun(ctx, "goT2")
	go goRun(ctx, "goT3")

	time.Sleep(time.Second * 5)
	fmt.Println("goT停止")
}

func goRun(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s,关闭了\n", name)
			return
		default:
			fmt.Printf("%s,等待关闭\n", name)
			time.Sleep(time.Second * 1)
		}
	}
}

func TestContestValueCtx(t *testing.T) {
	background := context.Background()
	value1 := "value1"
	ctx1 := context.WithValue(background, "key1", value1)

	value2 := "value2"
	ctx2 := context.WithValue(ctx1, "key2", value2)

	fmt.Println("key1 ==>", ctx1.Value("key1"))
	fmt.Println("key1 ==>", ctx1.Value("key2"))
	fmt.Println("key1 ==>", ctx2.Value("key1"))
	fmt.Println("key2 ==>", ctx2.Value("key2"))
}

func TestContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
