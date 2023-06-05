package main

import (
	"sync"
)

type NumDataBox struct {
	Num int
}

type Synchronized struct {
	m sync.Mutex
}

func (controller *Synchronized) Locking() {
	controller.m.Lock()
}

func (controller *Synchronized) Unlocking() {
	controller.m.Unlock()
}

func USAGE_MUTEX() {
	numData := NumDataBox{}
	syncController := Synchronized{}

	var sig
	tries := 10000
	for i := 0; i < tries; i++ {
		go func(ndb *NumDataBox, syd *Synchronized, order int) {
			ndb.Num++
		}(&numData, &syncController, i)
	}

}

// type Semaphore struct {
// 	m     sync.Mutex
// 	cond  *sync.Cond
// 	count int
// }

func main() {
	USAGE_MUTEX()
}
