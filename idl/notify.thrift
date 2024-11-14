namespace go notify
include "base.thrift"

struct SendNotifyReq{
    1: required string platform
    2: required string title
    3: required string message
}
struct SendNotifyResp{
    1: required base.BaseResp baseResp
}

service NotifyService {
    SendNotifyResp SendNotify(SendNotifyReq req)
}
