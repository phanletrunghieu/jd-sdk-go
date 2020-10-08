/*
@Time : 2020/9/23 15:17
@File:  sku_good_query
@Author: yandongit
@Description:通过SKUID查询推广商品的名称、主图、类目、价格、物流、是否自营、30天引单数量等详细信息，支持批量获取。通常用于在媒体侧展示商品详情。
*/
package request

import (
	"encoding/json"
	"fmt"
	"github.com/BepiColombo/jd-sdk-go/entity"
	"strconv"
)

type SkuGoodsQueryRequest struct {
	//京东skuID串，逗号分割，最多100个，开发示例如param_json={'skuIds':'5225346,7275691'}
	//（非常重要 请大家关注：如果输入的sk串中某个skuID的商品不在推广中[就是没有佣金]，返回结果中不会包含这个商品的信息）
	SkuIds string `json:skuIds`
}
type SkuGoodsQueryResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    []*SkuGoodQueryResp `json:"data"`
}
type SkuGoodQueryResp struct {
	SkuId             int64   `json:skuId`             //商品ID
	UnitPrice         float64 `json:unitPrice`         //商品单价即京东价
	MaterialUrl       string  `json:materialUrl`       //商品落地页
	EndDate           int64   `json:endDate`           //推广结束日期(时间戳，毫秒)
	IsFreeFreightRisk uint8   `json:isFreeFreightRisk` //是否支持运费险(1:是,0:否)
	IsFreeShipping    uint8   `json:isFreeShipping`    //是否包邮(1:是,0:否,2:自营商品遵从主站包邮规则)
	CommisionRatioWl  float64 `json:commisionRatioWl`  //无线佣金比例
	CommisionRatioPc  float64 `json:commisionRatioPc`  //PC佣金比例
	ImgUrl            string  `json:imgUrl`            //图片地址
	Vid               int64   `json:vid`               //商家ID
	CidName           string  `json:cidName`           //一级类目名称
	WlUnitPrice       float64 `json:wlUnitPrice`       //商品无线京东价（单价为-1表示未查询到该商品单价）
	Cid2Name          string  `json:cid2Name`          //二级类目名称
	IsSeckill         uint8   `json:isSeckill`         //是否秒杀(1:是,0:否)
	Cid2              int64   `json:cid2`              //二级类目ID
	Cid3Name          string  `json:cid3Name`          //三级类目名称
	InOrderCount      int64   `json:inOrderCount`      //30天引单数量
	Cid3              int64   `json:cid3`              //三级类目ID
	ShopId            int64   `json:shopId`            //商品ID
	IsJdSale          uint8   `json:isJdSale`          //是否自营(1:是,0:否)
	GoodsName         string  `json:goodsName`         //商品名称
	StartDate         int64   `json:startDate`         //推广开始日期（时间戳，毫秒）
	Cid               int64   `json:cid`               //一级类目ID
}

func (c *JdClient) SkuGoodsQuery(req SkuGoodsQueryRequest) (queryResult *SkuGoodsQueryResponse, e error) {
	methodName := "jd.union.open.goods.promotiongoodsinfo.query"
	responseName := "jd_union_open_goods_promotiongoodsinfo_query_response"

	goodsReq := map[string]interface{}{
		"skuIds": &req.SkuIds,
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
		e := &SkuGoodsQueryResponse{
			Code:    code,
			Message: errResponseBody.Zh_desc,
			Data:    []*SkuGoodQueryResp{},
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
