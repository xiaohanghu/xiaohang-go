package test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

// 协程可以理解为纯用户态的线程。协程的思想本质上就是控制流的主动让出（yield）和恢复（resume）机制
// 协程 Coroutine 的开销远远小于线程的开销
// 一般协程占用栈大小远小于线程（协程 KB 级别，线程 MB 级别），所以可以开更多的协程
// 在处理大规模并发连接（IO密集型任务）时，协程要优于线程。
func Test_GOMAXPROCS(t *testing.T) {
	// 逻辑CPU个数
	numCPU := runtime.NumCPU()
	fmt.Println("numCPU:", numCPU)
	// As of Go 1.5, the default value of GOMAXPROCS is the number of CPUs
	runtime.GOMAXPROCS(numCPU - 1)
}

func factorial(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

func setFactorial(n int, factorialMap *map[int]int, lock *sync.RWMutex) {
	fac := factorial(n)
	lock.Lock()
	(*factorialMap)[n] = fac
	lock.Unlock()
	fmt.Printf("finished: %v \n", n)
}

// set race for unit test:
// https://dev.to/s0xzwasd/run-tests-with-race-flag-in-goland-512j
// Go tool arguments:-race
func Test_factorial(t *testing.T) {
	//fmt.Printf("%v", factorial(3))

	// WARNING: DATA RACE
	// fatal error: concurrent map writes
	var factorialMap *map[int]int = &map[int]int{}
	//var factorialMap *sync.Map = &sync.Map[int]int{}
	var lock *sync.RWMutex = &sync.RWMutex{}
	fmt.Printf("lock: %T \n", lock)

	var wg sync.WaitGroup
	for n := 1; n <= 15; n++ {
		wg.Add(1)
		n1 := n // n会有并发问题，需要拷贝一次
		go func() {
			setFactorial(n1, factorialMap, lock)
			defer wg.Done()
		}()
	}

	//time.Sleep(3 * time.Second)
	wg.Wait()

	lock.RLock()
	fmt.Printf("factorialMap: %v \n", *factorialMap)
	lock.RUnlock()
}
