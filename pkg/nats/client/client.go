package nats_client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	nats_rpc "github.com/ImitationOfCoder/music_platform/pkg/nats"

	"github.com/nats-io/nats.go"
)

const (
	_defaultWaitTime = 5 * time.Second
	_defaultAttempts = 10
	_defaultTimeout  = 2 * time.Second
)

type Client struct {
	subject    string
	connection *nats.Conn

	timeout time.Duration
}

func New(
	url string,
	serverSubject string,
	opts ...Option,
) (*Client, error) {
	connection, err := nats.Connect(
		url,
		nats.ReconnectWait(_defaultWaitTime),
		nats.MaxReconnects(_defaultAttempts),
		nats.Timeout(_defaultWaitTime),
	)
	if err != nil {
		return nil, fmt.Errorf("nats_client - New - nats.Connect: %w", err)
	}

	c := &Client{
		subject:    serverSubject,
		connection: connection,
		timeout:    _defaultTimeout,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.connection = connection

	return c, nil
}

func (c *Client) Shutdown() error {
	c.connection.Close()

	return nil
}

func (c *Client) RemoteCall(handler string, request, response any) error {
	var (
		requestBody []byte
		err         error
	)

	if request != nil {
		requestBody, err = json.Marshal(request)

		if err != nil {
			return err
		}
	}

	header := nats.Header{
		"Handler": []string{handler},
	}

	requestMessage := nats.Msg{
		Subject: c.subject,
		Header:  header,
		Data:    requestBody,
	}

	message, err := c.connection.RequestMsg(&requestMessage, c.timeout)
	if errors.Is(err, context.DeadlineExceeded) {
		return nats_rpc.ErrTimeout
	}

	if err != nil {
		return fmt.Errorf("nats_client - RemoteCall - c.connection.RequestMsg: %w", err)
	}

	switch message.Header.Get("Status") {
	case nats_rpc.Success:
		err = json.Unmarshal(message.Data, &response)
		if err != nil {
			return fmt.Errorf("nats_client - RemoteCall - json.Unmarshal: %w", err)
		}
	case nats_rpc.ErrBadHandler.Error():
		return nats_rpc.ErrBadHandler
	case nats_rpc.ErrInternalServer.Error():
		return nats_rpc.ErrInternalServer
	}

	return nil
}
