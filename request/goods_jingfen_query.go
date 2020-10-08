/*
@Time : 2020/9/23 9:13
@File:  jingfen_goods_query
@Author: yandongit
@Description: jd.union.open.goods.jingfen.query 京粉精选商品查询接口
	京东联盟精选优质商品，每日更新，可通过频道ID查询各个频道下的精选商品。
	用获取的优惠券链接调用转链接口时，需传入搜索接口link字段返回的原始优惠券链接，
	切勿对链接进行任何encode、decode操作，否则将导致转链二合一推广链接时校验失败。
*/
package request

import (
	"encoding/json"
	"fmt"
	"github.com/jarvis4901/jd-sdk-go/entity"
	"strconv"
)

type JingfenGoodsQueryRequest struct {
	GoodsReq *JingfenGoodsReq `json:goodsReq`
}

type JingfenGoodsReq struct {
	EliteId   int    `json:"eliteId,omitempty"`   // 频道ID
	PageIndex int    `json:"pageIndex,omitempty"` // 页码 默认1
	PageSize  int    `json:"PageSize,omitempty"`  // 每页数量，默认20，上限50
	SortName  string `json:"sortName,omitempty"`  // 排序字段
	Sort      string `json:"sort,omitempty"`      // asc,desc升降序,默认降序
	Pid       string `json:"pid,omitempty"`       // 联盟id_应用id_推广位id，三段式
	Fields    string `json:"fields,omitempty"`    // 支持出参数据筛选，逗号','分隔，目前可用：videoInfo,documentInfo
}

type JingfenGoodsQueryResponse struct {
	Code       int                      `json:"code"`
	Message    string                   `json:"message"`
	RequestId  string                   `json:"requestId"`
	TotalCount int64                    `json:"totalCount"`
	Data       []*JingfenGoodsQueryResp `json:"data"`
}

type JingfenGoodsQueryResp struct {
	entity.CategoryInfo   `json:"categoryInfo"`   //类目信息
	Comments              int                     `json:"comments"` //评论数
	entity.CommissionInfo `json:"commissionInfo"` //佣金信息
	entity.CouponInfo     `json:"couponInfo"`     //优惠券信息，返回内容为空说明该SKU无可用优惠券
	GoodCommentsShare     float64                 `json:"goodCommentsShare"` //商品好评率
	entity.ImageInfo      `json:"imageinfo"`      //图片信息
	InOrderCount30Days    int64                   `json:"inOrderCount30Days"` //30天引单数量
	MaterialUrl           string                  `json:"materialUrl"`        //商品落地页
	entity.PriceInfo      `json:"priceInfo"`      //价格信息
	entity.ShopInfo       `json:"shopinfo"`       //店铺信息
	SkuId                 int64                   `json:"skuId"`     //商品ID
	SkuName               string                  `json:"skuName"`   //商品名称
	IsHot                 int                     `json:"isHot"`     //已废弃，请勿使用
	Spuid                 int64                   `json:"spuid"`     //spuid，其值为同款商品的主skuid
	BrandCode             string                  `json:"brandCode"` //品牌code
	BrandName             string                  `json:"brandName"` //品牌名
	Owner                 string                  `json:"owner"`     //g=自营，p=pop
	entity.PinGouInfo     `json:"pinGouInfo"`     //拼购信息
	entity.ResourceInfo   `json:"resourceInfo"`   //资源信息
	InOrderCount30DaysSku int64                   `json:"inOrderCount30DaysSku"` //30天引单数量(sku维度)
	entity.SeckillInfo    `json:"seckillInfo"`    //秒杀信息
	JxFlgs                []uint                  `json:jxFlags` //京喜商品类型，1京喜、2京喜工厂直供、3京喜优选（包含3时可在京东APP购买）
	entity.VideoInfo      `json:videoInfo`        //视频信息
	entity.DocumentInfo   `json:documentInfo`     //段子信息
}

func (c *JdClient) JingfenGoodsQuery(req JingfenGoodsQueryRequest) (queryResult *JingfenGoodsQueryResponse, e error) {
	methodName := "jd.union.open.goods.jingfen.query"
	responseName := "jd_union_open_goods_jingfen_query_response"

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
		e := &JingfenGoodsQueryResponse{
			Code:       code,
			Message:    errResponseBody.Zh_desc,
			Data:       []*JingfenGoodsQueryResp{},
			RequestId:  "",
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
