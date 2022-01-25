package tp_test

import (
	"fmt"
	tp "github.com/ansjsun/tptool"
	"go.uber.org/atomic"
	"strconv"
	"sync"
	"testing"
	"time"
)

var testMap = sync.Map{}
var enable = atomic.NewBool(true)

func TestTp(t *testing.T) {
	for i := 0; i < 1000; i++ {
		testMap.Store(strconv.Itoa(i), strconv.Itoa(i))
	}

	tt := tp.NewTpTest(tp.NewSpace("test"), 2000000, 1, tp.PrintCallback)

	go func() {
		time.Sleep(3000 * time.Millisecond)
		fmt.Println("to open")
		tt.SetEnable(true)
		time.Sleep(3000 * time.Millisecond)
		fmt.Println("to close")
		tt.SetEnable(false)
		time.Sleep(3000 * time.Millisecond)
		fmt.Println("to open")
		tt.SetEnable(true)
	}()

	wg := &sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				tt.Exe(func() {
					for i := 0; i < 2000; i++ {
						testMap.Load(strconv.Itoa(i))
						enable.Load()
					}
				})
			}
		}()
	}

	wg.Wait()
	time.Sleep(1000000)
}

func TestTp2(t *testing.T) {
	for i := 0; i < 1000; i++ {
		testMap.Store(strconv.Itoa(i), strconv.Itoa(i))
	}

	name := "abc"

	tp.NewTp(name)

	err := tp.TpEnable(name, true)
	if err != nil {
		panic(err)
	}

	wg := &sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				tp.Exe(name, func() {
					for i := 0; i < 2000; i++ {
						testMap.Load(strconv.Itoa(i))
						enable.Load()
					}
				})
			}
		}()
	}

	wg.Wait()
	time.Sleep(1000000)
}

func TestTp3(t *testing.T) {
	for i := 0; i < 1000; i++ {
		testMap.Store(strconv.Itoa(i), strconv.Itoa(i))
	}

	tt := tp.NewTpTest(tp.NewSpace("test"), 2000000, 1, tp.PrintCallback)
	tt.SetEnable(true)

	tt1 := tp.NewTpTest(tp.NewSpace("test1"), 2000000, 1, tp.PrintCallback)
	tt1.SetEnable(true)

	wg := &sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				tt.Exe(func() {
					for i := 0; i < 2000; i++ {
						testMap.Load(strconv.Itoa(i))
						enable.Load()
					}
				})
			}
		}()
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				tt1.Exe(func() {
					for i := 0; i < 2000; i++ {
						testMap.Load(strconv.Itoa(i))
						enable.Load()
					}
				})
			}
		}()
	}
	wg.Wait()
	time.Sleep(1000000)
}
