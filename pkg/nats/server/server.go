package nats_server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	nats_rpc "profile_microservice/pkg/nats"
	"time"

	"github.com/ImitationOfCoder/music_platform/pkg/logger"
	"github.com/nats-io/nats.go"
	"golang.org/x/sync/errgroup"
)

const (
	_defaultWaitTime = 5 * time.Second
	_defaultAttempts = 10
	_defaultTimeout  = 2 * time.Second
)

type CallHandler func(context.Context, *nats.Msg) (any, error)

type Server struct {
	ctx context.Context
	eg  *errgroup.Group

	subject      string
	connection   *nats.Conn
	subscription *nats.Subscription
	router       map[string]CallHandler
	stop         chan struct{}
	notify       chan error

	timeout time.Duration

	log *logger.Logger
}

func New(
	log *logger.Logger,
	url string,
	subject string,
	router map[string]CallHandler,
	opts ...Option,
) (*Server, error) {
	group, ctx := errgroup.WithContext(context.Background())
	group.SetLimit(1)

	connection, err := nats.Connect(
		url,
		nats.ReconnectWait(_defaultWaitTime),
		nats.MaxReconnects(_defaultAttempts),
		nats.Timeout(_defaultWaitTime),
	)
	if err != nil {
		return nil, fmt.Errorf("nats_server - New - nats.Connect: %w", err)
	}

	s := &Server{
		ctx:        ctx,
		eg:         group,
		subject:    subject,
		connection: connection,
		router:     router,
		stop:       make(chan struct{}),
		notify:     make(chan error, 1),
		timeout:    _defaultTimeout,
		log:        log,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s, nil
}

func (s *Server) Start() {
	s.eg.Go(func() error {
		err := s.subscribe()
		if err != nil {
			s.notify <- err

			close(s.notify)

			return err
		}

		<-s.stop

		return nil
	})

	s.log.Info("nats_server - Start - Started")
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	var shutdownErrors []error

	close(s.stop)

	// Wait for all goroutines to finish and get any error
	err := s.eg.Wait()
	if err != nil && !errors.Is(err, context.Canceled) {
		s.log.Error(fmt.Sprintf("nats_server - Shutdown - s.eg.Wait: %s", err.Error()))

		shutdownErrors = append(shutdownErrors, err)
	}

	// Unsubscribe
	if s.subscription != nil {
		err := s.subscription.Unsubscribe()
		if err != nil {

			s.log.Error(fmt.Sprintf("nats_server - Shutdown - s.conn.Subscription.Unsubscribe: %s", err.Error()))

			shutdownErrors = append(shutdownErrors, err)
		}
	}

	// Close connection
	s.connection.Close()

	s.log.Info("nats_server - Shutdown")

	return errors.Join(shutdownErrors...)
}

func (s *Server) subscribe() error {
	subscription, err := s.connection.Subscribe(s.subject, s.handleMessage)
	if err != nil {
		return fmt.Errorf("nats_rpc server - subscribe - s.conn.AttemptConnect: %w", err)
	}

	s.subscription = subscription

	return nil
}

func (s *Server) handleMessage(msg *nats.Msg) {
	handler := msg.Header.Get("Handler")

	callHandler, ok := s.router[handler]
	if !ok {
		s.publish(msg, nil, nats_rpc.ErrBadHandler.Error())
		return
	}

	ctx := context.Background()

	response, err := callHandler(ctx, msg)
	if err != nil {
		s.publish(msg, nil, nats_rpc.ErrInternalServer.Error())
		s.log.Error(fmt.Sprintf("nats_server - handleMessage - callHandler: %s", err.Error()))
		return
	}

	body, err := json.Marshal(response)
	if err != nil {
		s.publish(msg, nil, nats_rpc.ErrInternalServer.Error())
		s.log.Error(fmt.Sprintf("nats_server - handleMessage - json.Marshal: %s", err.Error()))
		return
	}

	s.publish(msg, body, nats_rpc.Success)
}

func (s *Server) publish(msg *nats.Msg, body []byte, status string) {
	respondMsg := nats.NewMsg(msg.Reply)
	respondMsg.Header.Set("Status", status)
	respondMsg.Data = body

	err := s.connection.PublishMsg(respondMsg)
	if err != nil {
		s.log.Error(fmt.Sprintf("nats_server - publish - s.connection.PublishMsg: %s", err.Error()))
	}
}
