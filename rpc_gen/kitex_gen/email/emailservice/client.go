// Code generated by Kitex v0.9.1. DO NOT EDIT.

package emailservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	email "gomall_demo/rpc_gen/kitex_gen/email"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Send(ctx context.Context, Req *email.EmailReq, callOptions ...callopt.Option) (r *email.EmailResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kEmailServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kEmailServiceClient struct {
	*kClient
}

func (p *kEmailServiceClient) Send(ctx context.Context, Req *email.EmailReq, callOptions ...callopt.Option) (r *email.EmailResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Send(ctx, Req)
}
