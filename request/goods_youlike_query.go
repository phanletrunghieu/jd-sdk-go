/*
@Time : 2020/10/8 16:09
@File:  goods_youlike_query
@Author: yandongit
@Description: jd.union.open.goods.material.query 猜你喜欢商品推荐
输入频道id、userid即可获取个性化推荐的商品信息，目前联盟推荐的精选频道包含猜你喜欢、实时热销、大额券、9.9包邮等各种实时数据，
适用于toc搭建频道页，千人千面商品推荐模块场景。建议使用clickURL转链长链接，千人千面推荐效果会更好。注意：请勿传入排序参数，以免影响推荐效果。
*/
package request

import (
	"encoding/json"
	"fmt"
	"github.com/BepiColombo/jd-sdk-go/entity"
	"strconv"
)

type YouLikeGoodsQueryRequest struct {
	GoodsReq *YouLikeGoodsReq `json:goodsReq`
}

type YouLikeGoodsReq struct {
	EliteId    int    `json:"eliteId,omitempty"`    // 频道ID
	PageIndex  int    `json:"pageIndex,omitempty"`  // 页码 默认1
	PageSize   int    `json:"PageSize,omitempty"`   // 每页数量，默认20，上限50
	Pid        string `json:"pid,omitempty"`        // 联盟id_应用id_推广位id，三段式
	SubUnionId string `json:"subUnionId,omitempty"` // 子联盟ID（需申请，申请方法请见https://union.jd.com/helpcenter/13246-13247-46301），该字段为自定义参数，建议传入字母数字和下划线的格式
	SiteId     string `json:"siteId,omitempty"`     // 站点ID是指在联盟后台的推广管理中的网站Id、APPID（1、通用转链接口禁止使用社交媒体id入参；2、订单来源，即投放链接的网址或应用必须与传入的网站ID/AppID备案一致，否则订单会判“无效-来源与备案网址不符”）
	PositionId string `json:"positionId,omitempty"` // 推广位id
	Ext1       string `json:"ext1,omitempty"`       // 系统扩展参数，无需传入
	//SkuId    int64 `json:"skuId,omitempty"`    // 预留字段，请勿传入
	HasCoupon  int `json:"hasCoupon,omitempty"`  // 1：只查询有最优券商品，不传值不做限制
	UserIdType int `json:"userIdType,omitempty"` // 支用户ID类型，传入此参数可获得个性化推荐结果。
	// 当前userIdType支持的枚举值包括：8、16、32、64、128、32768。userIdType和userId需同时传入，且一一对应。
	//userIdType各枚举值对应的userId含义如下：8(安卓移动设备Imei); 16(苹果移动设备Openudid)；32(苹果移动设备idfa); 64(安卓移动设备imei的md5编码，32位，大写，匹配率略低);
	//128(苹果移动设备idfa的md5编码，32位，大写，匹配率略低); 32768(安卓移动设备oaid)
	UserId string `json:"userId,omitempty"` // userIdType对应的用户设备ID，传入此参数可获得个性化推荐结果，userIdType和userId需同时传入
	//示例1： userIdType设置为8时，此时userId需要设置为安卓移动设备Imei，如861794042953717
	//示例2： userIdType设置为16时，此时userId需要设置为苹果移动设备Openudid，如f99dbd2ba8de45a65cd7f08b7737bc919d6c87f7
	//示例3： userIdType设置为32时，此时userId需要设置为苹果移动设备idfa，如DCC77BDA-C2CA-4729-87D6-B7F65C8014D6
	//示例4： userIdType设置为64时，此时userId需要设置为安卓移动设备imei的32位大写的MD5编码，如1097787632DB8876D325C356285648D0（原始imei：861794042953717）
	//示例5： userIdType设置为128时，此时userId需要设置为苹果移动设备idfa的32位大写的MD5编码，如01D0C2D675F700BA3716C05F39BDA0EB（原始idfa：DCC77BDA-C2CA-4729-87D6-B7F65C8014D6）
	//示例6： userIdType设置为32768时，此时userId需要设置为安卓移动设备oaid，如7dafe7ff-bffe-a28b-fdf5-7fefdf7f7e85
	Fields      string `json:"fields,omitempty"`      // 支持出参数据筛选，逗号','分隔，目前可用：videoInfo(视频信息)
	ForbidTypes string `json:"forbidTypes,omitempty"` //10微信京东购物小程序禁售，11微信京喜小程序禁售

}

type YouLikeGoodsQueryResponse struct {
	Code       int                      `json:"code"`
	Message    string                   `json:"message"`
	RequestId  string                   `json:"requestId"`
	TotalCount int64                    `json:"totalCount"`
	Data       []*YouLikeGoodsQueryResp `json:"data"`
}

type YouLikeGoodsQueryResp struct {
	CategoryInfo          *entity.CategoryInfo   `json:"categoryInfo"`          //类目信息
	Comments              int64                  `json:"comments"`              //评论数
	CommissionInfo        *entity.CommissionInfo `json:"commissionInfo"`        //佣金信息
	CouponInfo            *entity.CouponInfo     `json:"couponInfo"`            //优惠券信息，返回内容为空说明该SKU无可用优惠券
	PinGouInfo            *entity.PinGouInfo     `json:"pinGouInfo"`            //拼购信息
	ResourceInfo          *entity.ResourceInfo   `json:"resourceInfo"`          //资源信息
	InOrderCount30DaysSku int64                  `json:"inOrderCount30DaysSku"` //30天引单数量(sku维度)
	SeckillInfo           *entity.SeckillInfo    `json:"seckillInfo"`           //秒杀信息
	JxFlgs                []uint                 `json:jxFlags`                 //京喜商品类型，1京喜、2京喜工厂直供、3京喜优选（包含3时可在京东APP购买）
	VideoInfo             *entity.VideoInfo      `json:videoInfo`               //视频信息
	PromotionInfo         *entity.PromotionInfo  `json:promotionInfo`           //推广信息
	ForbidTypes           []uint8                `json:forbidTypes`             //0普通商品，10微信京东购物小程序禁售，11微信京喜小程序禁售
	DeliveryType          uint8                  `json:deliveryType`            //京东配送 1：是，0：不是
}

func (c *JdClient) YouLikeGoodsQuery(req YouLikeGoodsQueryRequest) (queryResult *YouLikeGoodsQueryResponse, e error) {
	methodName := "jd.union.open.goods.material.query"
	responseName := "jd_union_open_goods_material_query_response"

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
		e := &YouLikeGoodsQueryResponse{
			Code:       code,
			Message:    errResponseBody.Zh_desc,
			Data:       []*YouLikeGoodsQueryResp{},
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
