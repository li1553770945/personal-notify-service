package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/li1553770945/personal-notify-service/biz/constant"
	"github.com/li1553770945/personal-notify-service/kitex_gen/base"
	"github.com/li1553770945/personal-notify-service/kitex_gen/notify"
	"net/http"
)

func (s NotifyService) SendNotify(ctx context.Context, req *notify.SendNotifyReq) (*notify.SendNotifyResp, error) {
	switch req.Platform {
	case "wechat":
		return s.sendWeChat(ctx, req)
	default:
		return &notify.SendNotifyResp{
			BaseResp: &base.BaseResp{
				Code:    constant.NotFound,
				Message: "未知的消息发送平台",
			},
		}, nil
	}
}

func (s NotifyService) sendWeChat(ctx context.Context, req *notify.SendNotifyReq) (*notify.SendNotifyResp, error) {
	url := fmt.Sprintf("https://sctapi.ftqq.com/%s.send", s.Cfg.KeysConfig.SctKey)
	data := map[string]string{"title": req.Title, "desp": req.Message}
	payload, _ := json.Marshal(data)

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if httpReq == nil {
		klog.CtxErrorf(ctx, "发送消息失败，构建req为nil")
		return &notify.SendNotifyResp{
			BaseResp: &base.BaseResp{
				Code:    constant.SystemError,
				Message: "发送消息失败，构建req为nil",
			},
		}, nil
	}
	if err != nil {
		klog.CtxErrorf(ctx, fmt.Sprintf("发送消息失败:%v", err))
		return &notify.SendNotifyResp{
			BaseResp: &base.BaseResp{
				Code:    constant.SystemError,
				Message: fmt.Sprintf("发送消息失败:%v", err),
			},
		}, nil
	}

	httpReq.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if httpResp == nil {
		klog.CtxErrorf(ctx, "发送消息失败，返回值为nil")
		return &notify.SendNotifyResp{
			BaseResp: &base.BaseResp{
				Code:    constant.SystemError,
				Message: "发送消息失败，返回值为nil",
			},
		}, nil
	}
	if err != nil {
		klog.CtxErrorf(ctx, fmt.Sprintf("发送消息失败:%v", err))
		return &notify.SendNotifyResp{
			BaseResp: &base.BaseResp{
				Code:    constant.SystemError,
				Message: fmt.Sprintf("发送消息失败:%v", err),
			},
		}, nil
	}
	if httpResp.StatusCode != 200 {
		klog.CtxErrorf(ctx, fmt.Sprintf("发送消息失败,状态码:%v", httpResp.StatusCode))
		return &notify.SendNotifyResp{
			BaseResp: &base.BaseResp{
				Code:    constant.SystemError,
				Message: fmt.Sprintf("发送消息失败,状态码:%v", httpResp.StatusCode),
			},
		}, nil
	}
	return &notify.SendNotifyResp{
		BaseResp: &base.BaseResp{
			Code: constant.Success,
		},
	}, nil
}
