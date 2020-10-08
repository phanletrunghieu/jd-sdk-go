/*
@Time : 2020/10/8 9:54
@File:  subunionid_promotion_get
@Author: yandongit
@Description:jd.union.open.promotion.bysubunionid.get 社交媒体获取推广链接接口【申请】
通过商品链接、领券链接、活动链接获取普通推广链接或优惠券二合一推广链接，支持传入subunionid参数，可用于区分媒体自身的用户ID。需向cps-qxsq@jd.com申请权限。功能同宙斯接口的优惠券,商品二合一转接API-通过subUnionId获取推广链接、联盟微信手q通过subUnionId获取推广链接。
*/
package request

import (
	"encoding/json"
	"fmt"
	"github/bepicolombo/jd-sdk-go/entity"
	"strconv"
)

type SubUnionIdPromotionRequest struct {
	PromotionCodeReq *SubUnionIdPromotionReq `json:promotionCodeReq`
}

//入参
type SubUnionIdPromotionReq struct {
	MaterialId    string `json:"materialId"`              // 推广物料
	PositionId    int64  `json:"positionId,omitempty"`    // 推广位id
	SubUnionId    string `json:"subUnionId,omitempty"`    // 子联盟ID（需申请，申请方法请见https://union.jd.com/helpcenter/13246-13247-46301），该字段为自定义参数，建议传入字母数字和下划线的格式
	Ext1          string `json:"ext1,omitempty"`          // 系统扩展参数，无需传入
	Pid           string `json:"pid,omitempty"`           // 联盟子站长身份标识，格式：子站长ID_子站长网站ID_子站长推广位ID
	CouponUrl     string `json:"couponUrl,omitempty"`     //优惠券领取链接，在使用优惠券、商品二合一功能时入参，且materialId须为商品详情页链接
	ChainType     int    `json:"chainType,omitempty"`     //转链类型，1：长链， 2 ：短链 ，3： 长链+短链，默认短链
	GiftCouponKey string `json:"giftCouponKey,omitempty"` // 礼金批次号
}

type SubUnionIdPromotionResponse struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Data    *SubUnionIdPromotionResp `json:"data"`
}

type SubUnionIdPromotionResp struct {
	ShortURL string `json:"shortURL"` //生成的推广目标链接，以短链接形式，有效期60天
	ClickURL string `json:"clickURL"` //生成推广目标的长链，长期有效
}

func (c *JdClient) SubUnionIdPromotionGet(req SubUnionIdPromotionRequest) (queryResult *SubUnionIdPromotionResponse, e error) {
	methodName := "jd.union.open.promotion.bysubunionid.get"
	responseName := "jd_union_open_promotion_bysubunionid_get_response"

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
		e := &SubUnionIdPromotionResponse{
			Code:    code,
			Message: errResponseBody.Zh_desc,
			Data:    &SubUnionIdPromotionResp{},
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
