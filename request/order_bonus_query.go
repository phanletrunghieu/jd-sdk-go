/*
@Time : 2020/10/8 13:57
@File:  bonus_order_query
@Author: yandongit
@Description: jd.union.open.order.bonus.query 奖励订单查询接口【申请】
奖励订单查询接口，奖励活动效果的奖励订单明细查询接口，每日奖励订单大于五千单可申请该接口权限。未达到该标准的站长可在联盟官网—效果报表—导出对应订单明细。
*/
package request

import (
	"encoding/json"
	"fmt"
	"github.com/jarvis4901/jd-sdk-go/entity"
	"strconv"
)

type BonusOrderQueryRequest struct {
	Req *BonusOrderQueryReq `json:orderReq`
}

type BonusOrderQueryReq struct {
	PageNo    int    `json:"pageNo,omitempty"`    // 页码，返回第几页结果
	PageSize  int    `json:"pageSize,omitempty"`  //每页包含条数，上限为500
	OptType   int    `json:"optType,omitempty"`   // 时间类型 (1：下单时间，sortValue和pageSize组合使用； 2：更新时间，pageNo和pageSize组合使用)
	StartTime int64  `json:"startTime,omitempty"` //订单开始时间，时间戳（毫秒），起止时间限制10min内
	EndTime   int64  `json:"endTime,omitempty"`   //订单开始时间，时间戳（毫秒），起止时间限制10min内
	SortValue string `json:"sortValue,omitempty"` // 时间类型按'下单'查询时，和pageSize组合使用。获取最后一条记录的sortValue，指定拉取条数（pageSize），以此方式查询数据。
}

type BonusOrderQueryResponse struct {
	Code    int                               `json:"code"`
	Message string                            `json:"message"`
	Data    []BonusOrderQueryResponseDataItem `json:"data"`
}

func (c *JdClient) BonusOrderQuery(req BonusOrderQueryRequest) (queryResult *BonusOrderQueryResponse, e error) {
	methodName := "jd.union.open.order.bonus.query"
	responseName := "jd_union_open_order_bonus_query_response"

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
		e := &BonusOrderQueryResponse{
			Code:    code,
			Message: errResponseBody.Zh_desc,
			Data:    []BonusOrderQueryResponseDataItem{},
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

type BonusOrderQueryResponseDataItem struct {
	UnionId          int64   `json:"unionId"`          //推客的联盟ID
	BonusInvalidCode string  `json:"bonusInvalidCode"` // 无效状态码，-1:无效、2:无效-拆单、3:无效-取消、4:无效-京东帮帮主订单、5:无效-账号异常、6:无效-赠品类目不返佣 等
	BonusInvalidText string  `json:"bonusInvalidText"` // 无效状态码对应的无效状态文案
	PayPrice         string  `json:"payPrice"`         // 实际支付金额
	EstimateCosPrice float32 `json:estimateCosPrice`   //预估计佣金额，即用户下单的金额(已扣除优惠券、白条、支付优惠、进口税，未扣除红包和京豆)，有时会误扣除运费券金额，完成结算时会在实际计佣金额中更正。如订单完成前发生退款，此金额也会更新。
	EstimateFee      float32 `json:"estimateFee"`      // 推客的预估佣金（预估计佣金额*佣金比例*最终比例），如订单完成前发生退款，此金额也会更新。
	ActualCosPrice   float32 `json:"actualCosPrice"`   //实际计算佣金的金额。订单完成后，会将误扣除的运费券金额更正。如订单完成后发生退款，此金额会更新。
	ActualFee        float32 `json:"actualFee"`        //推客获得的实际佣金（实际计佣金额*佣金比例*最终比例）。如订单完成后发生退款，此金额会更新。
	OrderTime        int64   `json:"orderTime"`        //下单时间(时间戳，毫秒)
	PositionId       int64   `json:"positionId"`       //推广位ID,0代表无推广位
	OrderId          int64   `json:orderId`            //订单ID
	BonusState       int     `json:bonusState`         //奖励状态，0:无效、1:有效
	BonusText        string  `json:bonusText`          //奖励状态文案
	SkuName          string  `json:"skuName"`          //商品名称
	CommissionRate   float32 `json:"commissionRate"`   //佣金比例
	SubUnionId       string  `json:"subUnionId"`       //子联盟ID(需要联系运营开放白名单才能拿到数据)
	Pid              string  `json:"pid"`              // 联盟子站长身份标识，格式：子站长ID_子站长网站ID_子站长推广位ID
	Ext1             string  `json:"ext1"`             //推客生成推广链接时传入的扩展字段（需要联系运营开放白名单才能拿到数据）。'订单行维度'
	UnionAlias       string  `json:"unionAlias"`       //PID所属母账号平台名称（原第三方服务商来源）
	SubSideRate      float32 `json:"subSideRate"`      //分成比例
	SubsidyRate      float32 `json:subsidyRate`        //补贴比例
	FinalRate        float32 `json:"finalRate"`        // 最终比例（分成比例+补贴比例）
	ActivityName     string  `json:"activityName"`     // 活动名称
	ParentId         int64   `json:"parentId"`         //父单的订单ID，仅当发生订单拆分时返回， 0：未拆分，有值则表示此订单为子订单
	SkuId            int64   `json:"skuId"`            //商品ID
	EstimateBonusFee float32 `json:"estimateBonusFee"` // 预估奖励金额
	ActualBonusFee   float32 `json:"actualBonusFee"`   // 实际奖励金额
	OrderState       float32 `json:"orderState"`       // 订单奖励状态，1:已完成、2:已付款、3:待付款
	OrderText        string  `json:"orderText"`        // 订单奖励状态
	SortValue        string  `json:"sortValue"`        //排序值，按'下单时间'分页查询时使用

}
