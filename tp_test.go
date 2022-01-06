package tp_test

import (
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

	tt.SetEnable(true)

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
