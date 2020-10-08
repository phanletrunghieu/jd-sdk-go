/*
@Time : 2020/9/23 14:13
@File:  good_query
@Author: yandongit
@Description: jd.union.open.goods.query 关键词商品查询接口
	查询商品及优惠券信息，返回的结果可调用转链接口生成单品或二合一推广链接。支持按SKUID、关键词、优惠券基本属性、是否拼购、是否爆款等条件查询，
	建议不要同时传入SKUID和其他字段，以获得较多的结果。支持按价格、佣金比例、佣金、引单量等维度排序。用优惠券链接调用转链接口时，需传入搜索接口link字段返回的原始优惠券链接，
	切勿对链接进行任何encode、decode操作，否则将导致转链二合一推广链接时校验失败。
*/
package request

import (
	"encoding/json"
	"fmt"
	"github.com/BepiColombo/jd-sdk-go/entity"
	"strconv"
)

type GoodsQueryRequest struct {
	GoodsReqDTO *GoodsQueryReq `json:goodsReqDTO`
}

type GoodsQueryReq struct {
	Cid1                 int64   `json:"cid1,omitempty"`             //一级类目id
	Cid2                 int64   `json:"cid2,omitempty"`             //二级类目id
	Cid3                 int64   `json:"cid3,omitempty"`             //三级类目id
	PageIndex            int     `json:"pageIndex,omitempty"`        //页码
	PageSize             int     `json:"pageSize,omitempty"`         // 每页数量，单页数最大30，默认20
	SkuIds               []int64 `json:"skuIds,omitempty"`           // skuid集合(一次最多支持查询100个sku)，数组类型开发时记得加[]
	Keyword              string  `json:"keyword,omitempty"`          // 关键词，字数同京东商品名称一致，目前未限制
	Pricefrom            float64 `json:"pricefrom,omitempty"`        // 商品价格下限
	Priceto              float64 `json:"priceto,omitempty"`          // 商品价格上限
	CommissionShareStart int     `json:"commissionShareStart"`       // 佣金比例区间开始
	CommissionShareEnd   int     `json:"commissionShareEnd"`         // 佣金比例区间结束
	Owner                string  `json:"owner,omitempty"`            // 商品类型：自营[g]，POP[p]
	SortName             string  `json:"sortName,omitempty"`         // 排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30Days：30天引单量， inOrderComm30Days：30天支出佣金)
	Sort                 string  `json:"sort,omitempty"`             // asc,desc升降序,默认降序
	IsCoupon             uint8   `json:"isCoupon"`                   // 是否是优惠券商品，1：有优惠券，0：无优惠券
	IsPG                 uint8   `json:"isPG"`                       // 是否是拼购商品，1：拼购商品，0：非拼购商品
	PingouPriceStart     float64 `json:"pingouPriceStart,omitempty"` // 拼购价格区间开始
	PingouPriceEnd       float64 `json:"pingouPriceEnd,omitempty"`   // 拼购价格区间结束
	IsHot                int     `json:"isHot,omitempty"`            // 已废弃，请勿使用
	BrandCode            string  `json:"brandCode,omitempty"`        // 品牌code
	ShopId               int     `json:"shopId,omitempty"`           // 店铺Id
	HasContent           uint8   `json:"hasContent,omitempty"`       // 1：查询内容商品；其他值过滤掉此入参条件。
	HasBestCoupon        uint8   `json:"hasBestCoupon,omitempty"`    // 1：查询有最优惠券商品；其他值过滤掉此入参条件。
	Pid                  string  `json:"pid,omitempty"`              // 联盟id_应用iD_推广位id
	Fields               string  `json:"fields,omitempty"`           // 支持出参数据筛选，逗号','分隔，目前可用：videoInfo(视频信息),hotWords(热词),similar(相似推荐商品),documentInfo(段子信息)
	ForbidTypes          string  `json:"forbidTypes,omitempty"`      // 微信京东购物小程序禁售商品过滤规则，入参表示不展示该规则数据，支持多个逗号','分隔，2:OTC商品;3:加油卡;4:游戏充值卡;5:合约机;6:京保养;7:虚拟组套;8:定制商品
	JxFlags              []uint  `json:"jxFlags,omitempty"`          // 京喜商品类型，1京喜、2京喜工厂直供、3京喜优选（包含3时可在京东APP购买），入参多个值表示或条件查询
	ShopLevelFrom        float64 `json:"shopLevelFrom"`              // 支持传入0.0、2.5、3.0、3.5、4.0、4.5、4.9，默认为空表示不筛选评分
}

type GoodsQueryResponse struct {
	Code           int               `json:"code"`
	Message        string            `json:"message"`
	SimilarSkuList string            `json:"similarSkuList"`
	TotalCount     int64             `json:"totalCount"`
	Data           []*GoodsQueryResp `json:"data"`
	HotWords       string            `json:"hotWords"`
}

type GoodsQueryResp struct {
	entity.CategoryInfo   `json:"categoryInfo"`   //类目信息
	Comments              int                     `json:"comments"` //评论数
	entity.CommissionInfo `json:"commissionInfo"` //佣金信息
	entity.CouponInfo     `json:"couponInfo"`     //优惠券信息，返回内容为空说明该SKU无可用优惠券
	GoodCommentsShare     float64                 `json:"goodCommentsShare"` //商品好评率
	entity.ImageInfo      `json:"imageinfo"`      //图片信息
	InOrderCount30DaysSku int64                   `json:"inOrderCount30DaysSku"` //30天引单数量(sku维度)
	MaterialUrl           string                  `json:"materialUrl"`           //商品落地页
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
	entity.VideoInfo      `json:videoInfo`        //视频信息
	entity.CommentInfo    `json:commentInfo`      //评价信息
	JxFlgs                []uint                  `json:jxFlags` //京喜商品类型，1京喜、2京喜工厂直供、3京喜优选（包含3时可在京东APP购买）
	entity.DocumentInfo   `json:documentInfo`     //段子信息
}

func (c *JdClient) GoodsQuery(req GoodsQueryRequest) (queryResult *GoodsQueryResponse, e error) {
	methodName := "jd.union.open.goods.query"
	responseName := "jd_union_open_goods_query_response"

	goodsReq := map[string]interface{}{
		"goodsReqDTO": &req.GoodsReqDTO,
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
		e := &GoodsQueryResponse{
			Code:           code,
			Message:        errResponseBody.Zh_desc,
			Data:           []*GoodsQueryResp{},
			SimilarSkuList: "",
			TotalCount:     0,
			HotWords:       "",
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
