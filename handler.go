package main

import (
	"context"
	"github.com/li1553770945/personal-notify-service/biz/infra/container"
	notify "github.com/li1553770945/personal-notify-service/kitex_gen/notify"
)

// NotifyServiceImpl implements the last service interface defined in the IDL.
type NotifyServiceImpl struct{}

// SendNotify implements the NotifyServiceImpl interface.
func (s *NotifyServiceImpl) SendNotify(ctx context.Context, req *notify.SendNotifyReq) (resp *notify.SendNotifyResp, err error) {
	App := container.GetGlobalContainer()
	resp, err = App.NotifyService.SendNotify(ctx, req)
	return
}
