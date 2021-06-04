package main

import (
	"context"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/dtynn/dix"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
)

func serveRPC(ctx context.Context, listen string, full api.FullNode, stop dix.StopFunc, maxRequestSize int64) error {
	rpcOpts := []jsonrpc.ServerOption{}
	if maxRequestSize > 0 {
		rpcOpts = append(rpcOpts, jsonrpc.WithMaxRequestSize(maxRequestSize))
	}

	rpcServer := jsonrpc.NewServer(rpcOpts...)
	rpcServer.Register("Filecoin", full)

	http.Handle("/rpc/v0", rpcServer)

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
