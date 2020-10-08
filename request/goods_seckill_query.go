/*
@Time : 2020/9/24 10:44
@File:  seckill_goods_query 秒杀商品查询接口【申请】
@Author: yandongit
@Description: jd.union.open.goods.seckill.query 秒杀商品查询接口【申请】
	根据SKUID、类目等信息查询秒杀商品信息，秒杀商品的价格通常为近期低价，有利于促成购买。需向cps-qxsq@jd.com申请权限。
*/
package request

import (
	"encoding/json"
	"fmt"
	"github/bepicolombo/jd-sdk-go/entity"
	"strconv"
)

type SeckillGoodsQueryRequest struct {
	GoodsReq *SeckillGoodsQueryReq `json:goodsReq`
}

type SeckillGoodsQueryReq struct {
	SkuIds               []int64 `json:"skuIds,omitempty"`               // sku id集合，长度最大30
	PageIndex            int     `json:"pageIndex,omitempty"`            //页码
	PageSize             int     `json:"pageSize,omitempty"`             // 每页数量，单页数最大30，默认20
	IsBeginSecKill       int     `json:"isBeginSecKill"`                 // 是否返回未开始秒杀商品。1=返回，0=不返回
	SecKillPriceFrom     float64 `json:"secKillPriceFrom,omitempty"`     // 秒杀价区间开始（单位：元）
	SecKillPriceTo       float64 `json:"secKillPriceTo,omitempty"`       // 秒杀价区间结束
	Cid1                 int64   `json:"cid1,omitempty"`                 //一级类目id
	Cid2                 int64   `json:"cid2,omitempty"`                 //二级类目id
	Cid3                 int64   `json:"cid3,omitempty"`                 //三级类目id
	Owner                string  `json:"owner,omitempty"`                // g=自营，p=pop
	CommissionShareStart int     `json:"commissionShareStart,omitempty"` // 佣金比例区间开始
	CommissionShareEnd   int     `json:"commissionShareEnd,omitempty"`   // 佣金比例区间结束
	SortName             string  `json:"sortName,omitempty"`             // 排序字段，可为空。 （默认搜索综合排序。允许的排序字段：seckillPrice、commissionShare、inOrderCount30Days、inOrderComm30Days）
	Sort                 string  `json:"sort,omitempty"`                 // desc=降序，asc=升序，可为空（默认降序）
}

type SeckillGoodsQueryResponse struct {
	Code       int                       `json:"code"`
	Message    string                    `json:"message"`
	TotalCount int64                     `json:"totalCount"`
	Data       []*SecklillGoodsQueryResp `json:"data"`
}

type SecklillGoodsQueryResp struct {
	SkuName               string  `json:"skuName"`               //商品名称
	SkuId                 int64   `json:"skuId"`                 //商品ID
	ImageUrl              string  `json:imageUrl`                //图片url
	IsSecKill             uint8   `json:"isSecKill"`             //是秒杀。1：是商品 0：非秒杀商品
	OriPrice              float64 `json:"oriPrice"`              //原价
	SecKillPrice          float64 `json:"secKillPrice"`          //秒杀价
	SecKillStartTime      int64   `json:"secKillStartTime"`      //秒杀开始展示时间（时间戳：毫秒）
	SecKillEndTime        int64   `json:"secKillEndTime"`        //秒杀结束时间（时间戳：毫秒）
	Cid1Id                int64   `json:"cid1Id"`                //一级类目id
	Cid2Id                int64   `json:"cid2Id"`                //二级类目id
	Cid3Id                int64   `json:"cid3Id"`                //三级类目id
	Cid1Name              string  `json:"cid1Name"`              //一级类目名称
	Cid2Name              string  `json:"cid2Name"`              //二级类目名称
	Cid3Name              string  `json:"cid3Name"`              //三级类目名称
	CommissionShare       float64 `json:"commissionShare"`       //通用佣金比例，百分比
	Commission            float64 `json:commission`              //通用佣金
	Owner                 string  `json:"owner"`                 //g=自营，p=pop
	InOrderCount30DaysSku int64   `json:"inOrderCount30DaysSku"` //30天引单数量(sku维度)
	InOrderComm30Days     float64 `json:inOrderComm30Days`       //30天支出佣金（spu
	JdPrice               float64 `json:jdPrice`                 //京东价
}

func (c *JdClient) SeckillGoodsQuery(req SeckillGoodsQueryRequest) (queryResult *SeckillGoodsQueryResponse, e error) {
	methodName := "jd.union.open.goods.seckill.query"
	responseName := "jd_union_open_goods_seckill_query_response"

	goodsReq := map[string]interface{}{
		"goodsReq": &req.GoodsReq,
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
		e := &SeckillGoodsQueryResponse{
			Code:       code,
			Message:    errResponseBody.Zh_desc,
			Data:       []*SecklillGoodsQueryResp{},
			TotalCount: 0,
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
