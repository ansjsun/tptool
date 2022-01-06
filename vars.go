package tp

import "go.uber.org/atomic"

type Space struct {
	name   string
	enable *atomic.Bool
}

func NewSpace(name string) *Space {
	return &Space{name: name, enable: atomic.NewBool(false)}
}

