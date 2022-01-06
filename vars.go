package tp

import "go.uber.org/atomic"

type space struct {
	name   string
	enable *atomic.Bool
}

func newSpace(name string) *space {
	return &space{name: name, enable: atomic.NewBool(false)}
}

