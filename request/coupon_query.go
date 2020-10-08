/*
@Time : 2020/10/7 20:35
@File:  coupon_query
@Author: yandongit
@Description: jd.union.open.coupon.query 优惠券领取情况查询接口【申请】
通过领券链接查询优惠券的平台、面额、期限、可用状态、剩余数量等详细信息，通常用于和商品信息一起展示优惠券券信息。需向cps-qxsq@jd.com申请权限。
*/
package request

import (
	"encoding/json"
	"fmt"
	"github.com/jarvis4901/jd-sdk-go/entity"
	"strconv"
)

type CouponQueryRequest struct {
	CouponUrls []string `json:couponUrls`
}

// 优惠券
type Coupon struct {
	TakeEndTime   int64   `json:"takeEndTime"`   // 券领取结束时间(时间戳，毫秒)
	TakeBeginTime int64   `json:"takeBeginTime"` // 券领取开始时间(时间戳，毫秒)
	RemainNum     int64   `json:"remainNum"`     //券剩余张数
	Yn            string  `json:"yn"`            // 券有效状态
	Num           int64   `json:"num"`           // 券总张数
	Quota         float32 `json:"quota"`         // 券消费限额
	Link          string  `json:"link"`          // 券链接
	Discount      float32 `json:"discount"`      //券面额
	BeginTime     int64   `json:"beginTime"`     // 券有效使用开始时间(时间戳，毫秒)
	EndTime       int64   `json:"endTime"`       //券有效使用结束时间(时间戳，毫秒)
	Platform      string  `json:"platform"`      // 券使用平台
}

//type CouponQueryReq struct {
//	SkuIds []int64  `json:"skuIds"` //skuId集合
//	Fields []string `json:"fields"` //查询域集合，不填写则查询全部，目目前支持：categoryInfo（类目信息）,imageInfo（图片信息）,baseBigFieldInfo（基础大字段信息）,bookBigFieldInfo（图书大字段信息）,videoBigFieldInfo（影音大字段信息）,detailImages（商详图）
//}

type CouponQueryResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    []*Coupon `json:"data"`
}

func (c *JdClient) CouponQuery(req CouponQueryRequest) (queryResult *CouponQueryResponse, e error) {
	methodName := "jd.union.open.coupon.query"
	responseName := "jd_union_open_coupon_query_response"

	goodsReq := map[string]interface{}{
		"couponUrls": &req.CouponUrls,
	}
	respBytes, err := c.Execute(methodName, goodsReq)
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
		e := &CouponQueryResponse{
			Code:    code,
			Message: errResponseBody.Zh_desc,
			Data:    []*Coupon{},
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
