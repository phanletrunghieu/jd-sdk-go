/*
@Time : 2020/10/8 10:11
@File:  order_query
@Author: yandongit
@Description: jd.union.open.order.query 订单查询接口
查询推广订单及佣金信息，会随着订单状态变化更新数据，支持按下单时间、完成时间或状态更新时间查询，通常可按更新时间每分钟调用一次来获取订单的最新状态。
支持subunionid、推广位、PID、工具商角色订单查询。功能相当于原宙斯接口的订单查询、 查询引入订单、查询业绩订单、工具商订单查询、工具商引入数据查询接口、工具商业绩数据查询接口、PID订单查询、PID引入订单查询、PID业绩订单查询。
*/
package request

import (
	"encoding/json"
	"fmt"
	"github/bepicolombo/jd-sdk-go/entity"
	"strconv"
)

type OrderQueryRequest struct {
	Req *OrderQueryReq `json:orderReq`
}

type OrderQueryReq struct {
	PageNo       int    `json:"pageNo,omitempty"`       // 页码，返回第几页结果
	PageSize     int    `json:"pageSize,omitempty"`     //每页包含条数，上限为500
	Type         int    `json:"type,omitempty"`         // 订单时间查询类型(1：下单时间，2：完成时间，3：更新时间)
	Time         string `json:"time,omitempty"`         // 查询时间，建议使用分钟级查询，格式：yyyyMMddHH、yyyyMMddHHmm或yyyyMMddHHmmss，如201811031212 的查询范围从12:12:00--12:12:59
	Key          string `json:"key,omitempty"`          // 其他推客的授权key，查询工具商订单需要填写此项，childUnionid和key不能同时传入
	ChildUnionId int64  `json:"childUnionId,omitempty"` //子站长ID（需要联系运营开通PID账户权限才能拿到数据），childUnionId和key不能同时传入
}

type OrderQueryResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    []*OrderQueryResp `json:"data"`
	HasMore bool              `json:"hasMore"`
}

func (c *JdClient) OrderQuery(req OrderQueryRequest) (queryResult *OrderQueryResponse, e error) {
	methodName := "jd.union.open.order.query"
	responseName := "jd_union_open_order_query_response"

	orderReq := map[string]interface{}{
		"orderReq": &req.Req,
	}
	respBytes, err := c.Execute(methodName, orderReq)
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
		e := &OrderQueryResponse{
			Code:    code,
			Message: errResponseBody.Zh_desc,
			Data:    []*OrderQueryResp{},
			HasMore: false,
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

type OrderQueryResp struct {
	FinishTime int64      `json:"finishTime"` //订单完成时间(时间戳，毫秒)
	OrderEmt   int        `json:"orderEmt"`   //下单设备(1:PC,2:无线)
	OrderId    int64      `json:orderId`      //订单ID
	OrderTime  int64      `json:"orderTime"`  //下单时间(时间戳，毫秒)
	ParentId   int64      `json:"parentId"`   //父单的订单ID，仅当发生订单拆分时返回， 0：未拆分，有值则表示此订单为子订单
	PayMonth   int        `json:"payMonth"`   //订单维度预估结算时间,不建议使用，可以用订单行sku维度paymonth字段参考，（格式：yyyyMMdd），0：未结算，订单'预估结算时间'仅供参考。账号未通过资质审核或订单发生售后，会影响订单实际结算时间。
	Plus       int        `json:plus`         //下单用户是否为PLUS会员 0：否，1：是
	PopId      int64      `json:"popId"`      //订单维度商家ID，不建议使用，可以用订单行sku维度popId参考
	UnionId    int64      `json:"unionId"`    //推客的联盟ID
	SkuList    []*SkuInfo `json:"skuList"`    //订单包含的商品信息列表
	Ext1       string     `json:ext1`         //订单维度的推客生成推广链接时传入的扩展字段，不建议使用，可以用订单行sku维度ext1参考,（需要联系运营开放白名单才能拿到数据）
	ValidCode  int        `json:"validCode"`  //sku维度的有效码（-1：未知,2.无效-拆单,3.无效-取消,4.无效-京东帮帮主订单,5.无效-账号异常,6.无效-赠品类目不返佣,7.无效-校园订单,8.无效-企业订单,9.无效-团购订单,10.无效-开增值税专用发票订单,11.无效-乡村推广员下单,12.无效-自己推广自己下单,13.无效-违规订单,14.无效-来源与备案网址不符,15.待付款,16.已付款,17.已完成,18.已结算（5.9号不再支持结算状态回写展示））,20.无效-复购订单对应的首购订单无效,21.无效-云店订单

}

type SkuInfo struct {
	ActualCosPrice   float32 `json:"actualCosPrice"` //实际计算佣金的金额。订单完成后，会将误扣除的运费券金额更正。如订单完成后发生退款，此金额会更新。
	ActualFee        float32 `json:"actualFee"`      //推客获得的实际佣金（实际计佣金额*佣金比例*最终比例）。如订单完成后发生退款，此金额会更新。
	CommissionRate   float32 `json:"commissionRate"` //佣金比例
	EstimateCosPrice float32 `json:estimateCosPrice` //预估计佣金额，即用户下单的金额(已扣除优惠券、白条、支付优惠、进口税，未扣除红包和京豆)，有时会误扣除运费券金额，完成结算时会在实际计佣金额中更正。如订单完成前发生退款，此金额也会更新。
	EstimateFee      float32 `json:"estimateFee"`    // 推客的预估佣金（预估计佣金额*佣金比例*最终比例），如订单完成前发生退款，此金额也会更新。
	FinalRate        float32 `json:"finalRate"`      // 最终比例（分成比例+补贴比例）
	Cid1             int64   `json:"cid1"`           //一级类目ID
	FrozenSkuNum     int64   `json:frozenSkuNum`     //商品售后中数量
	Pid              string  `json:"pid"`            // 联盟子站长身份标识，格式：子站长ID_子站长网站ID_子站长推广位ID
	PositionId       int64   `json:"positionId"`     //推广位ID,0代表无推广位
	Price            float32 `json:"price"`          //商品单价
	Cid2             int64   `json:cid2`             //二级类目ID
	SiteId           int64   `json:"siteId"`         // 网站ID，0：无网站
	SkuId            int64   `json:"skuId"`          //商品ID
	SkuName          string  `json:"skuName"`        //商品名称
	SkuNum           int64   `json:skuNum`           //商品数量
	SkuReturnNum     int64   `json:"skuReturnNum"`   //商品已退货数量
	SubSideRate      float32 `json:"subSideRate"`    //分成比例
	SubsidyRate      float32 `json:subsidyRate`      //补贴比例
	Cid3             int64   `json:"cid3"`           //三级类目ID
	UnionAlias       string  `json:"unionAlias"`     //PID所属母账号平台名称（原第三方服务商来源）
	UnionTag         string  `json:"unionTag"`
	//联盟标签数据（整型的二进制字符串，返回32位：00000000000000000000000000000001。数据从右向左进行，每一位为1表示符合联盟的标签特征，
	//第1位：红包，第2位：组合推广，第3位：拼购，第5位：有效首次购（0000000000011XXX表示有效首购，最终奖励活动结算金额会结合订单状态判断，以联盟后台对应活动效果数据报表https://union.jd.com/active为准）,
	//第8位：复购订单，第9位：礼金，第10位：联盟礼金，第11位：推客礼金，第12位：京喜APP首购，第13位：京喜首购，第14位：京喜复购，第15位：京喜订单，
	//第16位：京东极速版APP首购，第17位白条首购，
	//第18位校园订单，例如：00000000000000000000000000000001:红包订单，00000000000000000000000000000010:组合推广订单，00000000000000000000000000000100:
	//拼购订单，00000000000000000000000000011000:有效首购，00000000000000000000000000000111：红包+组合推广+拼购等）
	UnionTrafficGroup int `json:unionTrafficGroup` //渠道组 1：1号店，其他：京东
	ValidCode         int `json:"validCode"`
	//sku维度的有效码（-1：未知,2.无效-拆单,3.无效-取消,4.无效-京东帮帮主订单,5.无效-账号异常,6.无效-赠品类目不返佣,7.无效-校园订单,
	//8.无效-企业订单,9.无效-团购订单,10.无效-开增值税专用发票订单,11.无效-乡村推广员下单,12.无效-自己推广自己下单,
	//13.无效-违规订单,14.无效-来源与备案网址不符,15.待付款,16.已付款,17.已完成,18.已结算（5.9号不再支持结算状态回写展示））
	SubUnionId          string  `json:"subUnionId"`        //子联盟ID(需要联系运营开放白名单才能拿到数据)
	TraceType           int     `json:"traceType"`         // 2：同店；3：跨店
	PayMonth            int     `json:payMonth`            // 订单行维度预估结算时间（格式：yyyyMMdd） ，0：未结算。订单'预估结算时间'仅供参考。账号未通过资质审核或订单发生售后，会影响订单实际结算时间。
	PopId               int64   `json:"popId"`             //商家ID。'订单行维度'
	Ext1                string  `json:"ext1"`              //推客生成推广链接时传入的扩展字段（需要联系运营开放白名单才能拿到数据）。'订单行维度'
	CpActId             int64   `json:"cpActId"`           //招商团活动id，正整数，为0时表示无活动
	UnionRole           int     `json:unionRole`           //站长角色，1： 推客、 2： 团长
	GiftCouponKey       string  `json:"giftCouponKey"`     //礼金批次ID
	GiftCouponOcsAmount float32 `json:giftCouponOcsAmount` //礼金分摊金额
	ProPriceAmount      float32 `json:"proPriceAmount"`    //价保金额
}
