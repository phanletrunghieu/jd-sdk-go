/*
@Time : 2020/10/8 15:55
@File:  pid_create
@Author: yandongit
@Description: jd.union.open.user.pid.get 获取PID【申请】
工具商媒体帮助子站长创建PID，该参数可在媒体和子站长之间建立关联，并通过获取推广链接、订单查询来跟踪。需向cps-qxsq@jd.com申请权限。此接口创建的pid,同一联盟用户不能超过10个，和官网手动创建的推广位相加不能超过500个。
*/
package request

import (
	"encoding/json"
	"fmt"
	"github.com/jarvis4901/jd-sdk-go/entity"
	"strconv"
)

type PidCreateRequest struct {
	PidReq *PidCreateReq `json:pidReq`
}
type PidCreateReq struct {
	UnionId       int64  `json:unionId,omitempty`       //需要创建的目标联盟id
	ChildUnionId  int64  `json:childUnionId,omitempty`  // 子站长ID
	PromotionType int    `json:promotionType,omitempty` //推广类型,1APP推广 2聊天工具推广
	PositionName  string `json:positionName,omitempty`  // 子站长的推广位名称，如不存在则创建，不填则由联盟根据母账号信息创建
	MediaName     string `json:mediaName,omitempty`     // 媒体名称，即子站长的app应用名称，推广方式为app推广时必填，且app名称必须为已存在的app名称
}
type PidCreateResponse struct {
	Code    int    `json:code`
	Message string `json:message`
	Data    string `json:data`
}

func (c *JdClient) PidCreate(req PidCreateRequest) (queryResult *PidCreateResponse, e error) {
	methodName := "jd.union.open.user.pid.get"
	responseName := "jd_union_open_user_pid_get_response"

	pidReq := map[string]interface{}{
		"pidReq": &req.PidReq,
	}
	respBytes, err := c.Execute(methodName, pidReq)
	if err != nil {
		return nil, err
	}
	var respObj map[string]interface{}
	if err := json.Unmarshal(respBytes, &respObj); err != nil {
		fmt.Println("JSON Unmarshal failed:", err)
		return nil, err
	}
	//错误信息处理
	if respObj["error_response"] != nil {
		var errResponseBody *entity.JdError
		errJson, err := json.Marshal(respObj["error_response"])
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(errJson, &errResponseBody); err != nil {
			fmt.Println("error response json unmarshal failed")
			return nil, err
		}
		code, err := strconv.Atoi(errResponseBody.Code)
		e := &PidCreateResponse{
			Code:    code,
			Message: errResponseBody.Zh_desc,
			Data:    "",
		}
		return e, nil
	}
	responseMessage := respObj[responseName].(map[string]interface{})
	respResult := responseMessage["result"].(string)
	if err := json.Unmarshal([]byte(respResult), &queryResult); err != nil {
		return nil, err
	}
	return queryResult, nil
}
