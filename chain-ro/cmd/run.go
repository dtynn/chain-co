package main

import (
	"context"

	"github.com/filecoin-project/lotus/api"
	"github.com/urfave/cli/v2"

	"github.com/dtynn/chain-co/chain-ro/service"
	"github.com/dtynn/chain-co/dep"
)

var runCmd = &cli.Command{
	Name:  "run",
	Usage: "start the chain-ro daemon",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "listen",
			Usage: "listen address for the service",
			Value: ":1234",
		},
		&cli.Int64Flag{
			Name:  "max-req-size",
			Usage: "max request size",
			Value: 10 << 20,
		},
		&cli.StringSliceFlag{
			Name:  "node",
			Usage: "node info",
		},
	},
	Action: func(cctx *cli.Context) error {
		appCtx, appCancel := context.WithCancel(cctx.Context)
		defer appCancel()

		var full api.FullNode

		stop, err := service.Build(
			appCtx,

			dep.MetricsCtxOption(appCtx, cliName),

			service.ParseNodeInfoList(cctx.StringSlice("node")),
			service.FullNode(&full),
		)

		if err != nil {
			return nil
		}

		defer stop(context.Background())

		return serveRPC(
			appCtx,
			cctx.String("listen"),
			full,
			func(ctx context.Context) error {
				appCancel()
				stop(ctx)
				return nil
			},
			cctx.Int64("max-req-size"),
		)
	},
}
