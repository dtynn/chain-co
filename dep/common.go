package dep

import (
	"context"

	"github.com/dtynn/dix"
	metricsi "github.com/ipfs/go-metrics-interface"

	"github.com/filecoin-project/lotus/node/modules/helpers"
)

func MetricsCtxOption(ctx context.Context, scope string) dix.Option {
	return dix.Override(new(helpers.MetricsCtx), func() context.Context {
		return metricsi.CtxScope(ctx, scope)
	})
}
