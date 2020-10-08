/*
@Time : 2020/10/8 15:14
@File:  position_query
@Author: yandongit
@Description: jd.union.open.position.query 查询推广位【申请】
工具商媒体帮助推客查询推广位，推广位分为两种类型：3：私域推广位，上限5000个，不在联盟后台展示，无对应 PID；4：联盟后台推广位，上限500个，
在联盟后台展示，可用于内容平台推广。业务参数key需要由推客进入联盟官网-我的工具-我的API中查询，有效期365天。此接口需向cps-qxsq@jd.com申请权限。
*/
package request

import (
	"encoding/json"
	"fmt"
	"github/bepicolombo/jd-sdk-go/entity"
	"strconv"
)

type PositionQueryRequest struct {
	PositionReq *PositionReq `json:positionReq`
}
type PositionReq struct {
	UnionId   int64  `json:unionId,omitempty`   //需要查询的目标联盟id
	Key       string `json:key,omitempty`       // 推客unionid对应的“授权Key”，在联盟官网-我的工具-我的API中查询，授权Key具有唯一性，有效期365天，删除或创建新的授权Key后原有的授权Key自动失效
	UnionType int64  `json:unionType,omitempty` // 3：私域推广位，上限5000个，不在联盟后台展示，无对应 PID；4：联盟后台推广位，上限500个，会在推客联盟后台展示，可用于内容平台推广
	PageIndex int64  `json:pageIndex,omitempty` //页码，上限100
	PageSize  int64  `json:pageSize,omitempty`  //每页条数，上限100
}
type PositionQueryResponse struct {
	Code    int                  `json:code`
	Message string               `json:message`
	Data    []*PositionQueryResp `json:data,omitempty`
	Total   int64                `json:total`
}
type PositionQueryResp struct {
	Id        int64  `json:id`        //推广位ID
	SiteId    int64  `json:siteId`    // 站点ID，如网站ID/appID/snsID
	SpaceName string `json:spaceName` //推广位名称
	Type      int64  `json:type`      //站点类型 1.网站推广位2.APP推广位3.导购媒体推广位4.聊天工具推广位
	Pid       int64  `json:pid`       //pid：仅uniontype传4时，展示pid；可用于内容平台推广

}

func (c *JdClient) PositionQuery(req PositionQueryRequest) (queryResult *PositionQueryResponse, e error) {
	methodName := "jd.union.open.position.query"
	responseName := "jd_union_open_position_query_response"

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
		e := &PositionQueryResponse{
			Code:    code,
			Message: errResponseBody.Zh_desc,
			Data:    []*PositionQueryResp{},
			Total:   0,
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
