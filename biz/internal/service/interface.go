package service

import (
	"context"
	"github.com/li1553770945/personal-notify-service/biz/infra/config"
	"github.com/li1553770945/personal-notify-service/kitex_gen/notify"
)

type NotifyService struct {
	Cfg *config.Config
}

type INotifyService interface {
	SendNotify(ctx context.Context, req *notify.SendNotifyReq) (*notify.SendNotifyResp, error)
	sendWeChat(ctx context.Context, req *notify.SendNotifyReq) (*notify.SendNotifyResp, error)
}

func NewNotifyService(cfg *config.Config) INotifyService {
	return &NotifyService{
		Cfg: cfg,
	}
}
