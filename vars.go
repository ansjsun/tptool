package tp

import "go.uber.org/atomic"

type space struct {
	name   string
	enable *atomic.Bool
}

func newSpace(name string) *space {
	return &space{name: name, enable: atomic.NewBool(false)}
}

const (
	BIN  = "BIN"
	BOUT = "BOUT"
	CIN  = "CIN"
	COUT = "COUT"
)

var (
	BIN_TP  = NewTpTest(newSpace(BIN), 100000, 1, PrintCallback)
	BOUT_TP = NewTpTest(newSpace(BOUT), 100000, 1, PrintCallback)
	CIN_TP  = NewTpTest(newSpace(CIN), 100000, 1, PrintCallback)
	COUT_TP = NewTpTest(newSpace(COUT), 100000, 1, PrintCallback)
)
