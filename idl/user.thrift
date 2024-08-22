namespace go hertz_demo
//引入其他 Thrift 文件
//include "user.thrift"
struct GetUserInfoReq{
    1: required string uuid,
    2: optional string openid,
}

struct GetUserInfoResp{
    1: required string uuid
    2: required string openid
    3: required string name
    4: required string projectName
}

service User{
    GetUserInfoResp GetUserInfo(1: GetUserInfoReq req)(api.post="/hertz/user/get")
}