package main

import (
	"StandardProject/common/util"
	"errors"
	"fmt"
	"time"
)

type KL struct {
	Id int
}

func AA(i int) {
	time.Sleep(3000)
	fmt.Println(i)
}
func main() {
	//list := make([]KL, 10)
	//mMap := make(map[int]int)
	//for i := 0; i < 10; i++ {
	//	list[i] = KL{
	//		i,
	//	}
	//	mMap[i] = rand.Intn(100)
	//}
	//
	//var mu sync.Mutex
	//
	//err := util.ExecuteCoroutineList(len(mMap), func(x int) error {
	//	if x != rand.Intn(50) {
	//		mu.Lock()
	//		mMap[rand.Intn(50)] = x
	//		mu.Unlock()
	//	}
	//
	//	return nil
	//})
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(mMap)

	coroutines := util.NewCoroutines()
	for i := 0; i < 5; i++ {
		coroutines.Add(func() error {

			if i == 3 {
				fmt.Println(i)
				return errors.New("scscasc")
			}
			time.Sleep(3 * time.Second)
			return nil
		})
	}
	time.Sleep(3 * time.Second)
	err := coroutines.Wait()

	fmt.Println(err)

}
