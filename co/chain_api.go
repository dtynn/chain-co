package co

import (
	"context"
	"time"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/store"
)

func (c *Coordinator) ChainNotify(ctx context.Context) (<-chan []*api.HeadChange, error) {
	subch := c.tspub.Sub(tipsetChangeTopic)

	c.headMu.RLock()
	head := c.head
	c.headMu.RUnlock()

	out := make(chan []*api.HeadChange, 32)
	out <- []*api.HeadChange{{
		Type: store.HCCurrent,
		Val:  head,
	}}

	done := make(chan struct{}, 0)
	go func() {
		select {
		case <-ctx.Done():

		case <-c.ctx.lc.Done():

		}

		close(done)
	}()

	go func() {
		defer func() {
			close(out)
			c.tspub.Unsub(subch)
			for range subch {
			}
		}()

		for {
			select {
			case val, ok := <-subch:
				if !ok {
					log.Info("ChainNotify: request done")
					return
				}

				if len(out) > 0 {
					log.Warnf("ChainNotify: head change sub is slow, has %d buffered entries", len(out))
				}

				select {
				case out <- val.([]*api.HeadChange):

				case <-done:
					return

				case <-time.After(time.Minute):
					log.Warn("ChainNotify: stucked for 1min")
					return
				}

			case <-done:
				return
			}
		}
	}()

	return out, nil
}
