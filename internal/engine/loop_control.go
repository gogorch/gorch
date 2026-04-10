package engine

import "sync/atomic"

type loopControl struct {
	breakLoop atomic.Bool
	until     atomic.Bool
}

func newLoopControl() *loopControl {
	return &loopControl{}
}

func (lc *loopControl) resetRound() {
	lc.breakLoop.Store(false)
	lc.until.Store(false)
}

func (lc *loopControl) setBreak() {
	lc.breakLoop.Store(true)
}

func (lc *loopControl) setUntil(done bool) {
	lc.until.Store(done)
}

func (lc *loopControl) shouldBreak() bool {
	return lc.breakLoop.Load()
}

func (lc *loopControl) isUntil() bool {
	return lc.until.Load()
}
