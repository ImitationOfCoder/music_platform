package grpc_server

import (
	"context"
	"errors"
	"net"

	"github.com/ImitationOfCoder/music_platform/pkg/logger"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

const (
	_defaultAddress = ":80"
)

type Server struct {
	ctx    context.Context
	log    *logger.Logger
	eg     *errgroup.Group
	notify chan error

	address string

	App  *grpc.Server
	opts []grpc.ServerOption
}

func New(log *logger.Logger, opts ...Option) *Server {
	group, ctx := errgroup.WithContext(context.Background())
	group.SetLimit(1)

	s := &Server{
		ctx:     ctx,
		log:     log,
		eg:      group,
		notify:  make(chan error, 1),
		address: _defaultAddress,
	}

	for _, opt := range opts {
		opt(s)
	}

	s.App = grpc.NewServer(s.opts...)

	return s
}

func (s *Server) Start() {
	s.eg.Go(func() error {
		var lc net.ListenConfig

		ln, err := lc.Listen(s.ctx, "tcp", s.address)
		if err != nil {
			s.notify <- err

			close(s.notify)

			return err
		}

		err = s.App.Serve(ln)
		if err != nil {
			s.notify <- err

			close(s.notify)

			return err
		}

		return nil
	})

	s.log.Info("gRPC server - Server - Started")
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	var shutdownErrors []error

	s.App.GracefulStop()

	err := s.eg.Wait()
	if err != nil && !errors.Is(err, context.Canceled) {
		s.log.Error("gRPC server - Server - Shutdown - s.eg.Wait")
		shutdownErrors = append(shutdownErrors, err)
	}

	s.log.Info("gRPC server - Server - Shutdown")
	return errors.Join(shutdownErrors...)
}
