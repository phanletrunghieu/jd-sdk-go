/*
@Time : 2020/10/8 9:36
@File:  common_promotion_get
@Author: yandongit
@Description:jd.union.open.promotion.common.get 网站/APP获取推广链接接口
网站/APP来获取的推广链接，功能同宙斯接口的自定义链接转换、 APP领取代码接口通过商品链接、活动链接获取普通推广链接，支持传入subunionid参数，可用于区分媒体自身的用户ID。
*/
package request

import (
	"encoding/json"
	"fmt"
	"github/bepicolombo/jd-sdk-go/entity"
	"strconv"
)

type CommonPromotionRequest struct {
	PromotionCodeReq *PromotionCodeReq `json:promotionCodeReq`
}

//入参
type PromotionCodeReq struct {
	MaterialId string `json:"materialId"`           // 推广物料
	SiteId     string `json:"siteId"`               // 站点ID是指在联盟后台的推广管理中的网站Id、APPID（1、通用转链接口禁止使用社交媒体id入参；2、订单来源，即投放链接的网址或应用必须与传入的网站ID/AppID备案一致，否则订单会判“无效-来源与备案网址不符”）
	PositionId int64  `json:"positionId,omitempty"` // 推广位id
	SubUnionId string `json:"subUnionId,omitempty"` // 子联盟ID（需申请，申请方法请见https://union.jd.com/helpcenter/13246-13247-46301），该字段为自定义参数，建议传入字母数字和下划线的格式
	Ext1       string `json:"ext1,omitempty"`       // 系统扩展参数，无需传入
	//Protocol      int    `json:"protocol"`      // 请勿再使用，后续会移除此参数，请求成功一律返回https协议链接
	Pid           string `json:"pid,omitempty"`           // 联盟子站长身份标识，格式：子站长ID_子站长网站ID_子站长推广位ID
	CouponUrl     string `json:"couponUrl,omitempty"`     //优惠券领取链接，在使用优惠券、商品二合一功能时入参，且materialId须为商品详情页链接
	GiftCouponKey string `json:"giftCouponKey,omitempty"` // 礼金批次号

}

//type CouponQueryReq struct {
//	SkuIds []int64  `json:"skuIds"` //skuId集合
//	Fields []string `json:"fields"` //查询域集合，不填写则查询全部，目目前支持：categoryInfo（类目信息）,imageInfo（图片信息）,baseBigFieldInfo（基础大字段信息）,bookBigFieldInfo（图书大字段信息）,videoBigFieldInfo（影音大字段信息）,detailImages（商详图）
//}

type CommonPromotionResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (c *JdClient) CommonPromotionGet(req CommonPromotionRequest) (queryResult *CommonPromotionResponse, e error) {
	methodName := "jd.union.open.promotion.common.get"
	responseName := "jd_union_open_promotion_common_get_response"

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
		e := &CommonPromotionResponse{
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
