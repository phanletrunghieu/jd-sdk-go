/*
@Time : 2020/9/24 13:53
@File:  category_goods_query
@Author: yandongit
@Description: jd.union.open.category.goods.get 商品类目查询接口
   根据商品的父类目id查询子类目id信息，通常用获取各级类目对应关系，以便将推广商品归类。业务参数parentId、grade都输入0可查询所有一级类目ID，之后再用其作为parentId查询其子类目。
*/
package request

import (
	"encoding/json"
	"fmt"
	"github.com/BepiColombo/jd-sdk-go/entity"
	"strconv"
)

type GoodsCategoriesQueryRequest struct {
	Req *GoodsCategoriesQueryReq `json:req`
}

type GoodsCategoriesQueryReq struct {
	ParentId int `json:"parentId"` //父类目id(一级父类目为0)
	Grade    int `json:"grade"`    //类目级别(类目级别 0，1，2 代表一、二、三级类目)
}

type GoodsCategoriesQueryResponse struct {
	Code    int                         `json:"code"`
	Message string                      `json:"message"`
	Data    []*GoodsCategoriesQueryResp `json:"data"`
}

type GoodsCategoriesQueryResp struct {
	Id       int    `json:"id"`       //类目Id
	Name     string `json:"name"`     //类目名称
	Grade    int    `json:grade`      //类目级别(类目级别 0，1，2 代表一、二、三级类目)
	ParentId int    `json:"parentId"` //父类目Id

}

func (c *JdClient) GoodsCategoriesQuery(req GoodsCategoriesQueryRequest) (queryResult *GoodsCategoriesQueryResponse, e error) {
	methodName := "jd.union.open.category.goods.get"
	responseName := "jd_union_open_category_goods_get_response"

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
		e := &GoodsCategoriesQueryResponse{
			Code:    code,
			Message: errResponseBody.Zh_desc,
			Data:    []*GoodsCategoriesQueryResp{},
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
