package nats_client

import "time"

type Option func(*Client)

func Timeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.timeout = timeout
	}
}
