package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func WHY_NEED_SYNC() {
	var data int32 = 0

	for i := 0; i < 2000; i++ {
		go func() {
			data += 1
		}()
	}

	for i := 0; i < 1000; i++ {
		go func() {
			data -= 1
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(data)
}

func SOLVE_SYNC() {
	var data int32 = 0
	var mu sync.Mutex = sync.Mutex{}

	for i := 0; i < 2000; i++ {
		go func() {
			mu.Lock()
			data += 1
			mu.Unlock()
		}()
	}

	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			data -= 1
			mu.Unlock()
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(data)
}

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

	sigForNext := make(chan bool)
	tries := 100
	sum := 0
	for i := 0; i < tries; i++ {
		myOrder := i
		go func(ndb *NumDataBox, syd *Synchronized, order int, sig chan bool) {
			fmt.Println("order: ", order)
			syd.Locking()
			ndb.Num++
			sum += ndb.Num
			syd.Unlocking()

			sig <- true

		}(&numData, &syncController, myOrder, sigForNext)
	}

	for i := 0; i < tries; i++ {
		<-sigForNext
	}

	fmt.Println(sum)
	fmt.Println(numData.Num)
}

func USAGE_COND() {
	mutex := new(sync.Mutex)
	cond := sync.NewCond(mutex)

	ch := make(chan bool)
	for i := 0; i < 3; i++ {
		go func(n int) {
			mutex.Lock()
			ch <- true
			fmt.Println("start waiting: ", n)
			cond.Wait()
			fmt.Println("end waiting: ", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 3; i++ {
		<-ch
	}

	for i := 0; i < 3; i++ {
		mutex.Lock()
		fmt.Println("signal : ", i)
		cond.Signal()
		mutex.Unlock()
		time.Sleep(time.Second)
	}

	time.Sleep(time.Second * 3)
}

func USAGE_COND_BROAD() {
	mutex := new(sync.Mutex)
	cond := sync.NewCond(mutex)

	ch := make(chan bool)
	for i := 0; i < 3; i++ {
		go func(n int) {
			mutex.Lock()
			ch <- true
			fmt.Println("start waiting: ", n)
			cond.Wait()
			fmt.Println("end waiting: ", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 3; i++ {
		<-ch
	}

	cond.Broadcast()
	time.Sleep(time.Second * 3)
}

type Semaphore struct {
	mutex     sync.Mutex
	condition *sync.Cond
	count     int
}

func NewSemaphore(initialCount int) *Semaphore {
	semaphore := &Semaphore{
		count: initialCount,
	}
	semaphore.condition = sync.NewCond(&semaphore.mutex)
	return semaphore
}

func (s *Semaphore) Acquire() {
	s.mutex.Lock()
	for s.count <= 0 {
		s.condition.Wait()
	}
	s.count--
	s.mutex.Unlock()
}

func (s *Semaphore) Release() {
	s.mutex.Lock()
	s.count++
	s.condition.Signal()
	s.mutex.Unlock()
}

func USAGE_WAITGROUP() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println(n)
			wg.Done()
		}(i)
	}
	wg.Wait() // Add 된 숫자 만큼 Done이 될 떄까지 기달림
	fmt.Println("the end")
}

type MsgData struct {
	isUsed bool
	msg    int
}

func USAGE_POOL() {
	wg := sync.WaitGroup{}
	pool := sync.Pool{
		New: func() interface{} {
			data := new(MsgData)
			data.msg = 0
			return data
		},
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data, _ := pool.Get().(*MsgData)
			fmt.Println(data)
			data.msg = rand.Intn(99)
			data.isUsed = true
			pool.Put(data)
			fmt.Println(" => ", data.msg)
		}()
	}
	fmt.Println("waiting ...")
	wg.Wait()
	fmt.Println("wait done")
	fmt.Println(pool.Get())
}

func USAGE_ATOMIC() {
	var data int32 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&data, 1)
			wg.Done()
		}()
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&data, -1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(data)
}

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("CPU NUM: ", runtime.NumCPU())
	fmt.Println("PROCS: ", runtime.GOMAXPROCS(0))
	// WHY_NEED_SYNC()
	// SOLVE_SYNC()
	// USAGE_MUTEX()
	// USAGE_COND()
	// USAGE_COND_BROAD()
	// USAGE_WAITGROUP()
	// USAGE_POOL()
	USAGE_ATOMIC()
}
