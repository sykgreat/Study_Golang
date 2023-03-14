package test

import (
	"testing"
	"time"
)

func proc() {
	panic("panic")
}

func Test_recover_panic(t *testing.T) {
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							t.Log(err)
						}
					}()
					t.Log(time.Now())
					proc()
				}()
			}
		}
	}()
	select {}
}
