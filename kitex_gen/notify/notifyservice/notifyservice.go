// Code generated by Kitex v0.7.2. DO NOT EDIT.

package notifyservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	notify "github.com/li1553770945/personal-notify-service/kitex_gen/notify"
)

func serviceInfo() *kitex.ServiceInfo {
	return notifyServiceServiceInfo
}

var notifyServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "NotifyService"
	handlerType := (*notify.NotifyService)(nil)
	methods := map[string]kitex.MethodInfo{
		"SendNotify": kitex.NewMethodInfo(sendNotifyHandler, newNotifyServiceSendNotifyArgs, newNotifyServiceSendNotifyResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "notify",
		"ServiceFilePath": `idl/notify.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.2",
		Extra:           extra,
	}
	return svcInfo
}

func sendNotifyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*notify.NotifyServiceSendNotifyArgs)
	realResult := result.(*notify.NotifyServiceSendNotifyResult)
	success, err := handler.(notify.NotifyService).SendNotify(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNotifyServiceSendNotifyArgs() interface{} {
	return notify.NewNotifyServiceSendNotifyArgs()
}

func newNotifyServiceSendNotifyResult() interface{} {
	return notify.NewNotifyServiceSendNotifyResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) SendNotify(ctx context.Context, req *notify.SendNotifyReq) (r *notify.SendNotifyResp, err error) {
	var _args notify.NotifyServiceSendNotifyArgs
	_args.Req = req
	var _result notify.NotifyServiceSendNotifyResult
	if err = p.c.Call(ctx, "SendNotify", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
