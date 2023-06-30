package main

import "StandardProject/sync/ebag/questions"

func main() {
	go questions.GetZsQsAgain(COROUTINE_NUM)
	go questions.GetEbagQsAgain(COROUTINE_NUM)
}
