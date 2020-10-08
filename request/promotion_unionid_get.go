/*
@Time : 2020/10/8 10:01
@File:  unionid_promotion_get
@Author: yandongit
@Description: jd.union.open.promotion.byunionid.get 工具商获取推广链接接口【申请】
工具商媒体帮助子站长获取普通推广链接和优惠券二合一推广链接，可传入PID参数以区分子站长的推广位，该参数可在订单查询接口返回。需向cps-qxsq@jd.com申请权限。
*/
package request

import (
	"encoding/json"
	"fmt"
	"github.com/jarvis4901/jd-sdk-go/entity"
	"strconv"
)

type UnionIdPromotionRequest struct {
	PromotionCodeReq *UnionIdPromotionReq `json:promotionCodeReq`
}

//入参
type UnionIdPromotionReq struct {
	MaterialId    string `json:"materialId"`              // 推广物料
	PositionId    int64  `json:"positionId,omitempty"`    // 推广位id
	UnionId       int64  `json:"unionId,omitempty"`       // 目标推客的联盟ID
	SubUnionId    string `json:"subUnionId,omitempty"`    // 子联盟ID（需申请，申请方法请见https://union.jd.com/helpcenter/13246-13247-46301），该字段为自定义参数，建议传入字母数字和下划线的格式
	Pid           string `json:"pid,omitempty"`           // 联盟子站长身份标识，格式：子站长ID_子站长网站ID_子站长推广位ID
	CouponUrl     string `json:"couponUrl,omitempty"`     //优惠券领取链接，在使用优惠券、商品二合一功能时入参，且materialId须为商品详情页链接
	ChainType     int    `json:"chainType,omitempty"`     //转链类型，1：长链， 2 ：短链 ，3： 长链+短链，默认短链
	GiftCouponKey string `json:"giftCouponKey,omitempty"` // 礼金批次号
}

type UnionIdPromotionResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    UnionIdPromotionResp `json:"data"`
}

type UnionIdPromotionResp struct {
	ShortURL string `json:"shortURL"` //生成的推广目标链接，以短链接形式，有效期60天
	ClickURL string `json:"clickURL"` //生成推广目标的长链，长期有效
}

func (c *JdClient) UnionIdPromotionGet(req UnionIdPromotionRequest) (queryResult *UnionIdPromotionResponse, e error) {
	methodName := "jd.union.open.promotion.byunionid.get"
	responseName := "jd_union_open_promotion_byunionid_get_response"

	promotionReq := map[string]interface{}{
		"promotionCodeReq": &req.PromotionCodeReq,
	}
	respBytes, err := c.Execute(methodName, promotionReq)
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
		e := &UnionIdPromotionResponse{
			Code:    code,
			Message: errResponseBody.Zh_desc,
			Data:    UnionIdPromotionResp{},
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
