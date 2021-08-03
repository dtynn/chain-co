package main

import (
	"context"
	"io/ioutil"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/api/v1api"

	"github.com/ipfs-force-community/chain-co/chain-ro/service"
	"github.com/ipfs-force-community/chain-co/dep"
	"github.com/ipfs-force-community/chain-co/localwt"
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
		&cli.StringFlag{
			Name:  "auth-url",
			Usage: "specify url for connect to venus-auth",
			Value: "",
		},
		&cli.StringFlag{
			Name:  "rate_limit_redis",
			Usage: "config redis to request api limit",
		},
		&cli.StringFlag{
			Name:        "version",
			Usage:       "rpc api version",
			Value:       "v1",
			DefaultText: "v1",
		},
	},
	Action: func(cctx *cli.Context) error {
		appCtx, appCancel := context.WithCancel(cctx.Context)
		defer appCancel()

		var full v1api.FullNode
		localJwt, err := localwt.NewLocalJwt()
		if err != nil {
			return err
		}
		token, err := localJwt.Token()
		if err != nil {
			return err
		}
		err = ioutil.WriteFile("./token", token, 0666)
		if err != nil {
			return err
		}
		stop, err := service.Build(
			appCtx,

			dep.MetricsCtxOption(appCtx, cliName),

			dep.APIVersionOption(cctx.String("version")),
			service.ParseNodeInfoList(cctx.StringSlice("node"), cctx.String("version")),
			service.FullNode(&full),
		)
		if err != nil {
			return nil
		}

		defer stop(context.Background())

		return serveRPC(
			appCtx,
			cctx.String("auth-url"),
			cctx.String("rate_limit_redis"),
			cctx.String("listen"),
			localJwt,
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
