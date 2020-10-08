package test

import (
	"fmt"
	"github/bepicolombo/jd-sdk-go/request"
	"testing"
)

const (
	APPKEY    = ""
	APPSECRET = ""
)

/**
京粉商品
*/
func TestJfGoodsMethod(t *testing.T) {
	client := request.NewClient(APPKEY, APPSECRET, "")
	jfGoodsQueryReq := request.JingfenGoodsQueryRequest{}
	jfGoodsReq := request.JingfenGoodsReq{}
	jfGoodsReq.PageIndex = 1
	jfGoodsReq.PageSize = 1
	jfGoodsReq.EliteId = 1
	jfGoodsQueryReq.GoodsReq = &jfGoodsReq
	res, err := client.JingfenGoodsQuery(jfGoodsQueryReq)
	if err == nil && res.Code == 200 {
		fmt.Println(res.Data)
	}
}

/**
关键词查询
*/
func TestGoodsQueryMethod(t *testing.T) {
	client := request.NewClient(APPKEY, APPSECRET, "")
	goodsQueryReq := request.GoodsQueryRequest{}
	goodsReq := request.GoodsQueryReq{}
	goodsReq.PageSize = 1
	goodsReq.PageIndex = 1
	goodsQueryReq.GoodsReqDTO = &goodsReq
	res, err := client.GoodsQuery(goodsQueryReq)
	if err == nil && res.Code == 200 {
		fmt.Println(res.Data)

	}
}

/**
根据sku查询商品
*/
func TestSkuGoodsQueryMethod(t *testing.T) {
	client := request.NewClient(APPKEY, APPSECRET, "")
	skuGoodsQueryReq := request.SkuGoodsQueryRequest{}
	skuids := "10020991470547,26160761615"
	skuGoodsQueryReq.SkuIds = skuids
	res, err := client.SkuGoodsQuery(skuGoodsQueryReq)
	if err == nil && res.Code == 200 {
		fmt.Println(res.Data)
	}
}

/**
秒杀商品
*/
func TestSeckillGoodsQueryMethod(t *testing.T) {
	client := request.NewClient(APPKEY, APPSECRET, "")
	seckillGoodsQueryReq := request.SeckillGoodsQueryRequest{}
	goodsReq := request.SeckillGoodsQueryReq{}
	goodsReq.PageIndex = 1
	goodsReq.PageSize = 2
	seckillGoodsQueryReq.GoodsReq = &goodsReq
	res, err := client.SeckillGoodsQuery(seckillGoodsQueryReq)
	if err == nil && res.Code == 200 {
		fmt.Println(res.Data)
	} else {
		fmt.Println(res)
	}
}

/**
学生价商品查询
*/
func TestStupriceGoodsQueryMethod(t *testing.T) {
	client := request.NewClient(APPKEY, APPSECRET, "")
	stupriceGoodsQueryReq := request.StupriceGoodsQueryRequest{}
	goodsReq := request.StupriceGoodsQueryReq{}
	goodsReq.PageIndex = 1
	goodsReq.PageSize = 2
	stupriceGoodsQueryReq.GoodsReq = &goodsReq
	res, err := client.StupriceGoodsQuery(stupriceGoodsQueryReq)
	if err == nil && res.Code == 200 {
		fmt.Println(res.Data)
	} else {
		fmt.Println(res)
	}
}

/**
商品类目查询
*/
func TestGoodsCategoriesQueryMethod(t *testing.T) {
	client := request.NewClient(APPKEY, APPSECRET, "")
	goodsCategoriesQueryReq := request.GoodsCategoriesQueryRequest{}
	req := request.GoodsCategoriesQueryReq{}
	req.ParentId = 0
	req.Grade = 1
	goodsCategoriesQueryReq.Req = &req
	res, err := client.GoodsCategoriesQuery(goodsCategoriesQueryReq)
	if err == nil && res.Code == 200 {
		fmt.Println(res.Data)
	} else {
		fmt.Println(res)
	}
}

/**
商品详情查询
*/
func TestGoodsDetailsQueryMethod(t *testing.T) {
	client := request.NewClient(APPKEY, APPSECRET, "")
	goodsDetailQueryReq := request.GoodsDetailQueryRequest{}
	req := request.GoodsDetailQueryReq{}
	req.SkuIds = []int64{57298641577}
	goodsDetailQueryReq.Req = &req
	res, err := client.GoodsDetailQuery(goodsDetailQueryReq)
	if err == nil && res.Code == 200 {
		fmt.Println(res.Data)
	} else {
		fmt.Println(res)
	}
}

/**
优惠券领取情况查询
*/
func TestCouponQueryMethod(t *testing.T) {
	client := request.NewClient(APPKEY, APPSECRET, "")
	couponQueryRequest := request.CouponQueryRequest{}
	couponQueryRequest.CouponUrls = []string{"https://coupon.m.jd.com/coupons/show.action?linkKey=AAROH_xIpeffAs_-naABEFoeNk3q1Xt8kYalOOTEZZJs8ph5HzyiV-OpUbXy-xek_lYu6yOR1Ng0HbWltUlGOyGX0bgrFg"}
	res, err := client.CouponQuery(couponQueryRequest)
	if err == nil && res.Code == 200 {
		fmt.Println(res.Data)
	} else {
		fmt.Println(res)
	}
}

/**
网站/APP获取推广链接接口
*/
func TestCommonPromotionMethod(t *testing.T) {
	client := request.NewClient(APPKEY, APPSECRET, "")
	commonPromotionRequest := request.CommonPromotionRequest{}
	promotionCodeReq := request.PromotionCodeReq{}
	promotionCodeReq.MaterialId = "https://item.jd.com/71095044412.html"
	promotionCodeReq.SiteId = "435676"
	commonPromotionRequest.PromotionCodeReq = &promotionCodeReq
	res, err := client.CommonPromotionGet(commonPromotionRequest)
	if err == nil && res.Code == 200 {
		fmt.Println(res.Data)
	} else {
		fmt.Println(res)
	}
}

/**
社交媒体获取推广链接
*/
func TestSubUnionIdPromotionnMethod(t *testing.T) {
	client := request.NewClient(APPKEY, APPSECRET, "")
	subUnionIdPromotionRequest := request.SubUnionIdPromotionRequest{}
	promotionCodeReq := request.SubUnionIdPromotionReq{}
	promotionCodeReq.MaterialId = "https://item.jd.com/71095044412.html"
	subUnionIdPromotionRequest.PromotionCodeReq = &promotionCodeReq
	res, err := client.SubUnionIdPromotionGet(subUnionIdPromotionRequest)
	if err == nil && res.Code == 200 {
		fmt.Println(res.Data)
	} else {
		fmt.Println(res)
	}
}

/**
工具商获取推广链接
*/
func TestUnionIdPromotionnMethod(t *testing.T) {
	client := request.NewClient(APPKEY, APPSECRET, "")
	unionIdPromotionRequest := request.UnionIdPromotionRequest{}
	promotionCodeReq := request.UnionIdPromotionReq{}
	promotionCodeReq.MaterialId = "https://item.jd.com/71095044412.html"
	promotionCodeReq.UnionId = 1003085935
	unionIdPromotionRequest.PromotionCodeReq = &promotionCodeReq
	res, err := client.UnionIdPromotionGet(unionIdPromotionRequest)
	if err == nil && res.Code == 200 {
		fmt.Println(res.Data.ShortURL)
	} else {
		fmt.Println(res)
	}
}

/**
订单查询接口
*/
func TestOrderQueryMethod(t *testing.T) {
	client := request.NewClient(APPKEY, APPSECRET, "")
	orderQueryRequest := request.OrderQueryRequest{}
	orderQueryReq := request.OrderQueryReq{}
	orderQueryReq.PageSize = 10
	orderQueryReq.PageNo = 1
	orderQueryReq.Type = 1
	orderQueryReq.Time = "202009181212"
	orderQueryRequest.Req = &orderQueryReq
	res, err := client.OrderQuery(orderQueryRequest)
	if err == nil && res.Code == 200 {
		fmt.Println(res.Data)
	} else {
		fmt.Println(res)
	}
}

/**
奖励订单查询接口
*/
func TestBonusOrderQueryMethod(t *testing.T) {
	client := request.NewClient(APPKEY, APPSECRET, "")
	orderQueryRequest := request.BonusOrderQueryRequest{}
	orderQueryReq := request.BonusOrderQueryReq{}
	orderQueryReq.PageSize = 10
	orderQueryReq.PageNo = 1
	orderQueryRequest.Req = &orderQueryReq
	res, err := client.BonusOrderQuery(orderQueryRequest)
	if err == nil && res.Code == 200 {
		fmt.Println(res.Data)
	} else {
		fmt.Println(res)
	}
}

/**
订单行查询接口
*/
func TestRowOrderQueryMethod(t *testing.T) {
	client := request.NewClient(APPKEY, APPSECRET, "")
	orderQueryRequest := request.RowOrderQueryRequest{}
	orderQueryReq := request.RowOrderQueryReq{}
	orderQueryReq.PageSize = 10
	orderQueryReq.PageIndex = 1
	orderQueryReq.Type = 1
	orderQueryReq.StartTime = "2020-09-14 17:23:00"
	orderQueryReq.EndTime = "2020-09-14 18:23:00"
	orderQueryReq.Fields = "goodsInfo,categoryInfo"

	orderQueryRequest.Req = &orderQueryReq
	res, err := client.RowOrderQuery(orderQueryRequest)
	if err == nil && res.Code == 200 {
		fmt.Println(res.Data[0].SiteId)
	} else {
		fmt.Println(err)
	}
}

/**
查询推广位
*/
func TestPositionQueryMethod(t *testing.T) {
	client := request.NewClient(APPKEY, APPSECRET, "")
	positionQueryRequest := request.PositionQueryRequest{}
	positionQueryReq := request.PositionReq{}
	positionQueryReq.PageSize = 10
	positionQueryReq.PageIndex = 1

	positionQueryRequest.PositionReq = &positionQueryReq
	res, err := client.PositionQuery(positionQueryRequest)
	if err == nil && res.Code == 200 {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}
}

/**
创建推广位
*/
func TestPositionCreateMethod(t *testing.T) {
	client := request.NewClient(APPKEY, APPSECRET, "")
	positionCreateRequest := request.PositionCreateRequest{}
	positionCreateReq := request.PositionCreateReq{}
	positionCreateReq.UnionId = 10000618
	positionCreateReq.Key = "sdasdswqeqweq"
	positionCreateReq.UnionId = 1
	positionCreateReq.Type = 1
	positionCreateReq.SpaceNameList = []string{"sdd", "dsad"}

	positionCreateRequest.PositionReq = &positionCreateReq
	res, err := client.PositionCreate(positionCreateRequest)
	if err == nil && res.Code == 200 {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}
}

/**
猜你喜欢
*/
func TestYouLikeGoodsQueryMethod(t *testing.T) {
	client := request.NewClient(APPKEY, APPSECRET, "")
	goodsQueryRequest := request.YouLikeGoodsQueryRequest{}
	goodQueryReq := request.YouLikeGoodsReq{}
	goodQueryReq.EliteId = 1
	goodQueryReq.UserIdType = 8
	goodQueryReq.UserId = "868377031634655"
	goodsQueryRequest.GoodsReq = &goodQueryReq
	res, err := client.YouLikeGoodsQuery(goodsQueryRequest)
	if err == nil && res.Code == 200 {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}
}
