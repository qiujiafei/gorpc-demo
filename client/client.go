package client

import (
	"context"
	"gorpc-demo/utils"
)

type Client interface {
	Invoke(ctx context.Context, req, rsp interface{}, path string, opts ...Option) error
}

type defaultClient struct {
	opts *Options
}

var New = func() *defaultClient {
	return &defaultClient{
		opts: &Options{
			protocol: "proto",
		},
	}
}

var DefaultClient = New()

func (c *defaultClient) Call(ctx context.Context, servicePath string, req, rsp interface{}, opts ...Option) error {
	callOpts := make([]Option, 0 ,len(opts) + 1)
	callOpts = append(callOpts, opts...)
	//callOpts = append(callOpts, WithSerializationType())
	err := c.Invoke(ctx, req, rsp, servicePath, callOpts...)
	if err != nil {
		return err
	}

	return nil
}

func (c *defaultClient) Invoke(ctx context.Context, req, rsp interface{}, path string, opts ...Option) error {
	for _, o := range opts {
		o(c.opts)
	}

	if c.opts.timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, c.opts.timeout)
		defer cancel()
	}

	service, method, err := utils.ParseServicePath(path)
	if err != nil {
		return err
	}

	c.opts.serviceName = service
	c.opts.method = method
	return nil
}