package main

import (
	"context"
	"github.com/ipfs-force-community/metrics"
	"go.opencensus.io/plugin/ochttp"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	logging "github.com/ipfs/go-log/v2"

	"github.com/dtynn/dix"

	"github.com/filecoin-project/go-jsonrpc"

	"github.com/filecoin-project/lotus/api"

	"github.com/filecoin-project/venus-auth/cmd/jwtclient"

	"github.com/ipfs-force-community/chain-co/localwt"
	"github.com/ipfs-force-community/metrics/ratelimit"
)

func serveRPC(ctx context.Context, authEndpoint, rate_limit_redis, listen string, mCnf *metrics.TraceConfig, jwt *localwt.LocalJwt, full api.FullNode, stop dix.StopFunc, maxRequestSize int64) error {
	serverOptions := []jsonrpc.ServerOption{}
	if maxRequestSize > 0 {
		serverOptions = append(serverOptions, jsonrpc.WithMaxRequestSize(maxRequestSize))
	}

	rpcServer := jsonrpc.NewServer(serverOptions...)
	rpcServer.Register("Filecoin", full)

	var remoteJwtCli *jwtclient.JWTClient
	if len(authEndpoint) > 0 {
		remoteJwtCli = jwtclient.NewJWTClient(authEndpoint)
	}

	//register hander to verify token in venus-auth
	var handler http.Handler
	if remoteJwtCli != nil {
		handler = (http.Handler)(jwtclient.NewAuthMux(jwt, jwtclient.WarpIJwtAuthClient(remoteJwtCli), rpcServer, logging.Logger("Auth")))
	} else {
		handler = (http.Handler)(jwtclient.NewAuthMux(jwt, nil, rpcServer, logging.Logger("Auth")))
	}

	if repoter, err := metrics.RegisterJaeger(mCnf.ServerName, mCnf); err != nil {
		log.Fatalf("register %s JaegerRepoter to %s failed:%s", mCnf.ServerName, mCnf.JaegerEndpoint)
	} else if repoter != nil {
		log.Infof("register jaeger-tracing exporter to %s, with node-name:%s", mCnf.JaegerEndpoint, mCnf.ServerName)
		defer metrics.UnregisterJaeger(repoter)
		handler = &ochttp.Handler{Handler: handler}
	}

	serveRpc := func(path string, hnd interface{}) {
		rpcServer := jsonrpc.NewServer(serverOptions...)
		rpcServer.Register("Filecoin", hnd)
		http.Handle(path, handler)
	}

	limitWrapper := full
	if len(rate_limit_redis) > 0 && remoteJwtCli != nil {
		limiter, err := ratelimit.NewRateLimitHandler(
			rate_limit_redis,
			nil, &jwtclient.ValueFromCtx{},
			jwtclient.WarpLimitFinder(remoteJwtCli),
			logging.Logger("rate-limit"))
		_ = logging.SetLogLevel("rate-limit", "info")
		if err != nil {
			return err
		}

		var rateLimitAPI api.FullNodeStruct
		limiter.WarpFunctions(full, &rateLimitAPI.Internal)
		limiter.WarpFunctions(full, &rateLimitAPI.VenusAPIStruct.Internal)
		limiter.WarpFunctions(full, &rateLimitAPI.CommonStruct.Internal)
		limitWrapper = &rateLimitAPI
	}

	pma := api.PermissionedFullAPI(limitWrapper)

	serveRpc("/rpc/v0", pma)
	serveRpc("/rpc/v1", pma)

	server := http.Server{
		Addr:    listen,
		Handler: http.DefaultServeMux,
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

		log.Sync()
	}()

	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	log.Infow("start http server", "addr", listen)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	log.Info("gracefull down")
	return nil
}
