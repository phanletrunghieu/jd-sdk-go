/*
@Time : 2020/10/8 15:36
@File:  position_create
@Author: yandongit
@Description: jd.union.open.position.create 创建推广位【申请】
工具商媒体帮助推客批量创建推广位，推广位分为两种类型：3：私域推广位，上限5000个，不在联盟后台展示，无对应 PID；4：联盟后台推广位，上限500个，
在联盟后台展示，自动生成PID，可用于内容平台推广。业务参数key需要由推客进入联盟官网-我的工具-我的API中查询，有效期365天。此接口需向cps-qxsq@jd.com申请权限。
*/
package request

import (
	"encoding/json"
	"fmt"
	"github/bepicolombo/jd-sdk-go/entity"
	"strconv"
)

type PositionCreateRequest struct {
	PositionReq *PositionCreateReq `json:positionReq`
}
type PositionCreateReq struct {
	UnionId       int64    `json:unionId,omitempty`       //需要创建的目标联盟id
	Key           string   `json:key,omitempty`           // 推客unionid对应的“授权Key”，在联盟官网-我的工具-我的API中查询，授权Key具有唯一性，有效期365天，删除或创建新的授权Key后原有的授权Key自动失效
	UnionType     int      `json:unionType,omitempty`     // 3：私域推广位，上限5000个，不在联盟后台展示，无对应 PID；4：联盟后台推广位，上限500个，会在推客联盟后台展示，自动生成对应PID，可用于内容平台推广
	Type          int      `json:type,omitempty`          // 站点类型 1.网站推广位2.APP推广位3.导购媒体推广位4.聊天工具推广位
	SpaceNameList []string `json:spaceNameList,omitempty` // 推广位名称集合，英文,分割；上限50
	SiteId        int64    `json:siteId,omitempty`        // 站点ID：网站的ID/app ID/snsID 。当type非4(聊天工具)时，siteId必填
}
type PositionCreateResponse struct {
	Code    int                 `json:code`
	Message string              `json:message`
	Data    *PositionCreateResp `json:data`
}
type PositionCreateResp struct {
	ResultList *PositionCreateResultList    `json:resultList` //推广位结果集合
	SiteId     int64                        `json:siteId`     // 站点ID，如网站ID/appID/snsID
	Type       int64                        `json:type`       //站点类型 1.网站推广位2.APP推广位3.导购媒体推广位4.聊天工具推广位
	Pid        *PositionCreateResultPidList `json:pid`        //pid：仅uniontype传4时，展示pid；可用于内容平台推广
	UnionId    int64                        `json:unionId`    // 联盟ID
}

type PositionCreateResultList struct {
	Result string `json:result` //推广位结果，但是对应的pid不能作为母子分佣使用。
}
type PositionCreateResultPidList struct {
	Pid string `json:pid` //pid结果，仅uniontype传4时，展示pid；可用于内容平台推广
}

func (c *JdClient) PositionCreate(req PositionCreateRequest) (queryResult *PositionCreateResponse, e error) {
	methodName := "jd.union.open.position.create"
	responseName := "jd_union_open_position_create_response"

	positionReq := map[string]interface{}{
		"positionReq": &req.PositionReq,
	}
	respBytes, err := c.Execute(methodName, positionReq)
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
		e := &PositionCreateResponse{
			Code:    code,
			Message: errResponseBody.Zh_desc,
			Data:    &PositionCreateResp{},
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
