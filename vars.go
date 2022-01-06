package tp

import (
	"fmt"
	"go.uber.org/atomic"
	"sync"
)

type Space struct {
	name   string
	enable *atomic.Bool
}

func NewSpace(name string) *Space {
	return &Space{name: name, enable: atomic.NewBool(false)}
}

var tpMap = &sync.Map{}

var count = atomic.NewInt32(0)

func NewTp(name string) {
	tpMap.Store(name, NewTpTest(NewSpace(name), 100000, 1, PrintCallback))
}

func Exe(name string, fn func()) {
	if count.Load() == 0 {
		fn()
		return
	}
	store, loaded := tpMap.Load(name)
	if !loaded {
		fn()
		return
	}
	store.(*TpTest).Exe(fn)
}

func TpEnable(name string, enable bool) error {
	if store, loaded := tpMap.Load(name); loaded {
		store.(*TpTest).SetEnable(enable)
		if enable {
			count.Inc()
		} else {
			count.Dec()
		}
		return nil
	}
	return fmt.Errorf("not found:%s", name)
}
