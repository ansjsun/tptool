package tp

import (
	"strconv"
	"sync"
	"testing"
	"time"
)

var testMap = sync.Map{}

func TestTp(t *testing.T) {
	for i := 0; i < 1000; i++ {
		testMap.Store(strconv.Itoa(i), strconv.Itoa(i))
	}

	tp := NewTpTest(newSpace("test"), 2000000, 1, PrintCallback)


	wg := &sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				run(tp)
			}
		}()
	}

	wg.Wait()
	time.Sleep(1000000)
}

func run(tp *TpTest) {
	now := NowMicro()
	for i := 0; i < 2000; i++ {
		testMap.Store(strconv.Itoa(i), strconv.Itoa(i))
	}
	tp.End(now)
}
