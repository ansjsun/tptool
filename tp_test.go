package tp

import (
	"sync"
	"testing"
	"time"
)

func TestTp(t *testing.T) {
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
	time.Sleep(100000)
	tp.End(now)
}
