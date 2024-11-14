// Code generated by Kitex v0.7.2. DO NOT EDIT.

package notifyservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	notify "github.com/li1553770945/personal-notify-service/kitex_gen/notify"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SendNotify(ctx context.Context, req *notify.SendNotifyReq, callOptions ...callopt.Option) (r *notify.SendNotifyResp, err error)
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
	return &kNotifyServiceClient{
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

type kNotifyServiceClient struct {
	*kClient
}

func (p *kNotifyServiceClient) SendNotify(ctx context.Context, req *notify.SendNotifyReq, callOptions ...callopt.Option) (r *notify.SendNotifyResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SendNotify(ctx, req)
}