package tp

import (
	"fmt"
	"go.uber.org/atomic"
	"sort"
	"sync"
	"time"
	"unsafe"
)

func PrintCallback(name string, start int64, sortUses []int) {
	l := len(sortUses)
	if l < 10 {
		fmt.Println("data so less", sortUses)
		return
	}
	fmt.Println(
		"name", name,
		"start", start,
		"all", len(sortUses),
		"min", sortUses[0],
		"max", sortUses[l-1],
		"tp90", sortUses[int(float32(l)*0.9)],
		"tp99", sortUses[int(float32(l)*0.99)],
		"tp999", sortUses[int(float32(l)*0.999)],
	)
}

type TpTest struct {
	sync.Mutex
	space           *Space
	uses            *atomic.UnsafePointer
	periodSecond    int64
	startTimeSecond *atomic.Int64
	callBack        func(name string, start int64, sortUses []int)
	chanLen         int
}

func NewTpTest(space *Space, chanLen int, periodSecond int64, callBack func(name string, start int64, sortUses []int)) *TpTest {
	l := chanLen
	if space.enable.Load() {
		l = chanLen
	}
	c := make(chan int64, l)

	return &TpTest{
		Mutex:           sync.Mutex{},
		space:           space,
		uses:            atomic.NewUnsafePointer(unsafe.Pointer(&c)),
		periodSecond:    periodSecond,
		startTimeSecond: atomic.NewInt64(time.Now().Unix()),
		callBack:        callBack,
		chanLen:         chanLen,
	}
}

func (t *TpTest) Exe(f func()) {
	if t.space.enable.Load() {
		now := NowMicro()
		f()
		t.End(now)
	} else {
		f()
	}
}

func (t *TpTest) SetEnable(flag bool) bool {
	if flag {
		c := make(chan int64, t.chanLen)
		t.uses.Swap(unsafe.Pointer(&c))
	} else {
		c := make(chan int64, 0)
		t.uses.Swap(unsafe.Pointer(&c))
	}

	return t.space.enable.Swap(flag)
}

func (t *TpTest) End(startUnixMicro int64) {

	go func() {
		now := time.Now()
		use := now.UnixNano()/1000 - startUnixMicro
		if use == 0 {
			panic("aaaaaa")
		}
		c := t.uses.Load()

		select {
		case *(*chan int64)(c) <- use:
		default:
			fmt.Println("channel is full")
		}
		if now.Unix()-t.startTimeSecond.Load() >= t.periodSecond {
			t.Lock()
			if now.Unix()-t.startTimeSecond.Load() < t.periodSecond {
				t.Unlock()
				return
			}

			c := make(chan int64, t.chanLen)
			uses := t.uses.Swap(unsafe.Pointer(&c))
			start := t.startTimeSecond.Swap(now.Unix())
			t.Unlock()
			go explain(t.space.name, start, *((*chan int64)(uses)), t.callBack)

		}
	}()
}

func explain(name string, start int64, uses chan int64, callBack func(name string, start int64, sortUses []int)) {
	l := len(uses)
	s := make([]int, 0, l)
	for i := 0; i < l; i++ {
		s = append(s, int(<-uses))

	}
	sort.Ints(s)
	callBack(name, start, s)
}

func NowMicro() int64 {
	return time.Now().UnixNano() / 1000
}
