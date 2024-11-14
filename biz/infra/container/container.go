package container

import (
	"github.com/li1553770945/personal-notify-service/biz/infra/config"
	"github.com/li1553770945/personal-notify-service/biz/infra/log"
	"github.com/li1553770945/personal-notify-service/biz/infra/trace"
	"github.com/li1553770945/personal-notify-service/biz/internal/service"
	"sync"
)

type Container struct {
	Logger        *log.TraceLogger
	Trace         *trace.TraceStruct
	Config        *config.Config
	NotifyService service.INotifyService
}

var APP *Container
var once sync.Once

func GetGlobalContainer() *Container {
	if APP == nil {
		panic("APP在使用前未初始化")
	}
	return APP
}

func InitGlobalContainer(env string) {
	once.Do(func() {
		APP = GetContainer(env)
	})
}

func NewContainer(config *config.Config,

	logger *log.TraceLogger,
	traceStruct *trace.TraceStruct,

	notifyService service.INotifyService,
) *Container {
	return &Container{
		Config:        config,
		Logger:        logger,
		NotifyService: notifyService,
		Trace:         traceStruct,
	}

}
