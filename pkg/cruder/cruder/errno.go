package cruder

import (
	"errors"
)

var (
	ErrCreateNothing = errors.New("create nothing")
	ErrUpdateNothing = errors.New("update nothing")
)
