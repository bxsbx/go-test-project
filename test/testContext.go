package main

import (
	"StandardProject/common/util"
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	waitGroup := util.NewWaitGroupWithContext(ctx, 2)

	for i := 0; i < 10; i++ {
		x := i
		waitGroup.GoCancel(func() error {
			time.Sleep(3 * time.Second)
			fmt.Println(x)
			if x == 2 {
				return errors.New("err")
			}
			return nil
		})
	}

	waitGroup.Wait()

	//cancel, cancelFunc := context.WithCancel(ctx)
	//go func() {
	//	time.Sleep(3 * time.Second)
	//	fmt.Println("1")
	//	select {
	//	case <-cancel.Done():
	//		fmt.Println("任务取消")
	//	default:
	//		time.Sleep(10 * time.Second)
	//		fmt.Println("1")
	//		cancelFunc()
	//	}
	//}()
	//
	//go func() {
	//	select {
	//	case <-cancel.Done():
	//		fmt.Println("任务取消")
	//	default:
	//		time.Sleep(2 * time.Second)
	//		fmt.Println("2")
	//		cancelFunc()
	//	}
	//}()
	//time.Sleep(5 * time.Second)
	//<-cancel.Done()
}
