// Code generated by hertz generator.

package hertz_demo

import (
	"context"

	hertz_demo "code.byted.org/motor/hertz-project/biz/model/hertz_demo"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetUserInfo .
// @router /hertz/user/get [POST]
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hertz_demo.GetUserInfoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(hertz_demo.GetUserInfoResp)
	resp.UUID = req.UUID

	c.JSON(consts.StatusOK, resp)
}
