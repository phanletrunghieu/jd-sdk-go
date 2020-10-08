/*
@Time : 2020/10/7 20:07
@File:  good_detail_query
@Author: yandongit
@Description: jd.union.open.goods.bigfield.query 商品详情查询接口【申请】
  商品详情查询接口,大字段信息
*/
package request

import (
	"encoding/json"
	"fmt"
	"github.com/BepiColombo/jd-sdk-go/entity"
	"strconv"
)

type GoodsDetailQueryRequest struct {
	Req *GoodsDetailQueryReq `json:goodsReq`
}

type GoodsDetailQueryReq struct {
	SkuIds []int64  `json:"skuIds"`           //skuId集合
	Fields []string `json:"fields,omitempty"` //查询域集合，不填写则查询全部，目目前支持：categoryInfo（类目信息）,imageInfo（图片信息）,baseBigFieldInfo（基础大字段信息）,bookBigFieldInfo（图书大字段信息）,videoBigFieldInfo（影音大字段信息）,detailImages（商详图）
}

type GoodsDetailQueryResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    *BigFieldGoodsResp `json:"data"`
}

func (c *JdClient) GoodsDetailQuery(req GoodsDetailQueryRequest) (queryResult *GoodsDetailQueryResponse, e error) {
	methodName := "jd.union.open.goods.bigfield.query"
	responseName := "jd_union_open_goods_bigfield_query_response"

	goodsReq := map[string]interface{}{
		"req": &req.Req,
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
		e := &GoodsDetailQueryResponse{
			Code:    code,
			Message: errResponseBody.Zh_desc,
			Data:    &BigFieldGoodsResp{},
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

// 图书大字段信息
type BookBigFieldInfo struct {
	Comments        string `json:"comments"`        //媒体评论
	Image           string `json:"image"`           // 精彩文摘与插图(插图
	VontentDesc     string `json:"contentDesc"`     // 内容摘要(内容简介)
	TelatedProducts string `json:"relatedProducts"` // 产品描述(相关商品)
	EditerDesc      string `json:"editerDesc"`      // 编辑推荐
	Catalogue       string `json:"catalogue"`       // 目录
	BookAbstract    string `json:"bookAbstract"`    // 精彩摘要(精彩书摘)
	AuthorDesc      string `json:"authorDesc"`      // 作者简介
	Introduction    string `json:"introduction"`    //前言(前言/序言)
	ProductFeatures string `json:"productFeatures"` //产品特色
}

//影音大字段信息
type VideoBigFieldInfo struct {
	Comments             string `json:"comments"`             //媒体评论
	Image                string `json:"image"`                // 商品描述(精彩剧照)
	ContentDesc          string `json:"contentDesc"`          // 内容摘要(内容简介)
	Box_Contents         string `json:"box_Contents"`         // 包装清单
	EditerDesc           string `json:"editerDesc"`           // 编辑推荐
	Catalogue            string `json:"catalogue"`            // 目录
	Material_Description string `json:"material_Description"` // 特殊说明
	Manual               string `json:"manual"`               //说明书
	ProductFeatures      string `json:"productFeatures"`      //产品特色
}

//基础大字段信息
type BaseBigFieldInfo struct {
	Wdis     string `json:"wdis"`     //商品介绍
	PropCode string `json:"propCode"` //规格参数
	WareQD   string `json:"wareQD"`   // 包装清单(仅自营商品)

}

// 类目信息
type CategoryInfo struct {
	Cid1     int64  `json:"cid1"`     // 一级类目ID
	Cid1Name string `json:"cid1Name"` // 一级类目名称
	Cid2     int64  `json:"cid2"`     // 二级类目ID
	Cid2Name string `json:"cid2Name"` // 二级类目名称
	Cid3     int64  `json:"cid3"`     // 三级类目ID
	Cid4Name string `json:"cid3Name"` // 三级类目名称
}

// 图片信息
type ImageInfo struct {
	ImageList []*UrlInfo `json:"imageList"` // 图片合集
}

// 图片合集
type UrlInfo struct {
	Url string `json:"url"` // 图片链接地址，第一个图片链接为主图链接
}

//商品详情
type BigFieldGoodsResp struct {
	SkuId             int64              `json:"skuId"`             // skuId
	SkuName           string             `json:"high"`              // 商品名称
	CategoryInfo      *CategoryInfo      `json:"categoryInfo"`      //类目信息
	ImageInfo         *ImageInfo         `json:"imageInfo"`         //图片信息
	BaseBigFieldInfo  *BaseBigFieldInfo  `json:"baseBigFieldInfo"`  //基础大字段信息
	BookBigFieldInfo  *BookBigFieldInfo  `json:"bookBigFieldInfo"`  //图书大字段信息
	VideoBigFieldInfo *VideoBigFieldInfo `json:"videoBigFieldInfo"` //影音大字段信息
	MainSkuId         int64              `json:"mainSkuId"`         //自营主skuId
	ProductId         int64              `json:"productId"`         //非自营商品Id
	SkuStatus         int                `json:"skuStatus"`         // 宽
	Owner             string             `json:"owner"`             // 高
	DetailImages      string             `json:"detailImages"`      // 商详图
}
