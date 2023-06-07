package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/dtynn/dix"
	"github.com/etherlabsio/healthcheck/v2"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/ipfs-force-community/metrics"
	"github.com/ipfs-force-community/metrics/ratelimit"
	"github.com/ipfs-force-community/sophon-auth/core"
	"github.com/ipfs-force-community/sophon-auth/jwtclient"
	local_api "github.com/ipfs-force-community/sophon-co/cli/api"
	apiInfo "github.com/ipfs-force-community/venus-common-utils/apiinfo"
	logging "github.com/ipfs/go-log/v2"
	"go.opencensus.io/plugin/ochttp"
)

func serveRPC(ctx context.Context, authApi apiInfo.APIInfo, rateLimitRedis, listen string, mCnf *metrics.TraceConfig, jwt jwtclient.IJwtAuthClient, full api.FullNode, localApi local_api.LocalAPI, stop dix.StopFunc, maxRequestSize int64) error {
	serverOptions := []jsonrpc.ServerOption{}
	if maxRequestSize > 0 {
		serverOptions = append(serverOptions, jsonrpc.WithMaxRequestSize(maxRequestSize))
	}

	var remoteJwtCli *jwtclient.AuthClient
	if len(authApi.Addr) > 0 {
		if len(authApi.Token) == 0 {
			return fmt.Errorf("auth token is need when auth api is set")
		}
		remoteJwtCli, _ = jwtclient.NewAuthClient(authApi.Addr, string(authApi.Token))
	}

	pma := api.PermissionedFullAPI(full)
	if len(rateLimitRedis) > 0 && remoteJwtCli != nil {
		log.Infof("use rate limit %s", rateLimitRedis)
		limiter, err := ratelimit.NewRateLimitHandler(
			rateLimitRedis,
			nil, &core.ValueFromCtx{},
			jwtclient.WarpLimitFinder(remoteJwtCli),
			logging.Logger("rate-limit"))
		_ = logging.SetLogLevel("rate-limit", "debug")
		if err != nil {
			return err
		}

		var rateLimitAPI api.FullNodeStruct
		limiter.WrapFunctions(full, &rateLimitAPI.Internal)
		limiter.WrapFunctions(full, &rateLimitAPI.NetStruct.Internal)
		limiter.WrapFunctions(full, &rateLimitAPI.VenusAPIStruct.Internal)
		limiter.WrapFunctions(full, &rateLimitAPI.CommonStruct.Internal)
		pma = &rateLimitAPI
	}

	mux := http.NewServeMux()

	serveRpc := func(path string, hnd interface{}, rpcSer *jsonrpc.RPCServer) {
		rpcSer.Register("Filecoin", hnd)

		var handler http.Handler
		if remoteJwtCli != nil {
			handler = (http.Handler)(jwtclient.NewAuthMux(jwt, jwtclient.WarpIJwtAuthClient(remoteJwtCli), rpcSer))
		} else {
			handler = (http.Handler)(jwtclient.NewAuthMux(jwt, nil, rpcSer))
		}
		mux.Handle(path, handler)
	}

	serveRpc("/rpc/v0", &v0api.WrapperV1Full{FullNode: pma}, jsonrpc.NewServer(serverOptions...))
	serveRpc("/rpc/v1", pma, jsonrpc.NewServer(serverOptions...))
	serveRpc("/rpc/admin/v0", localApi, jsonrpc.NewServer(serverOptions...))
	mux.Handle("/healthcheck", healthcheck.Handler())

	allHandler := (http.Handler)(mux)
	if reporter, err := metrics.RegisterJaeger(mCnf.ServerName, mCnf); err != nil {
		log.Fatalf("register %s JaegerRepoter to %s failed:%s", mCnf.ServerName, mCnf.JaegerEndpoint, err)
	} else if reporter != nil {
		log.Infof("register jaeger-tracing exporter to %s, with node-name:%s", mCnf.JaegerEndpoint, mCnf.ServerName)
		defer metrics.UnregisterJaeger(reporter)
		allHandler = &ochttp.Handler{Handler: allHandler}
	}

	server := http.Server{
		Addr:    listen,
		Handler: allHandler,
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}

	sigCh := make(chan os.Signal, 2)

	go func() {
		select {
		case <-ctx.Done():

		case sig := <-sigCh:
			log.Infof("signal %s captured", sig)
		}

		if err := server.Shutdown(context.Background()); err != nil {
			log.Warnf("shutdown http server: %s", err)
		}

		if err := stop(context.Background()); err != nil {
			log.Warnf("call app stop func: %s", err)
		}

		log.Sync() // nolint:errcheck
	}()

	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	log.Infow("start http server", "addr", listen)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	log.Info("gracefull down")
	return nil
}
