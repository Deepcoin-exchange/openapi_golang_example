package copytrading

import (
	"apiRequest/consts"
	"apiRequest/signature"
	"apiRequest/structs"
	"encoding/json"
	"fmt"
)

type CopyTradingCtrl struct {
	env structs.Env
}

func NewCopyTradingCtrl(env *structs.Env) *CopyTradingCtrl {
	return &CopyTradingCtrl{
		env: *env,
	}
}

func (c *CopyTradingCtrl) LeaderSettings() {
	type leaderSettingsRequest struct {
		Status           int    `json:"status"`
		Profile          string `json:"profile,omitempty"`
		HomeMode         int    `json:"homeMode,omitempty"`
		IsClosedCopyCode bool   `json:"isClosedCopyCode,omitempty"`
		CopyCode         string `json:"copyCode,omitempty"`
	}

	l := &leaderSettingsRequest{}
	l.Status = 0
	l.HomeMode = 1
	l.IsClosedCopyCode = true
	l.CopyCode = ""

	requestBody, err := json.Marshal(l)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := c.env.Url + consts.COPYTRADING_LEADER_SETTINGS
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.COPYTRADING_LEADER_SETTINGS, string(requestBody), &c.env)
}

func (c *CopyTradingCtrl) SupportContracts() {
	requestURL := fmt.Sprintf(c.env.Url + consts.COPYTRADING_SUPPORT_CONTRACT)
	requestPath := fmt.Sprintf(consts.COPYTRADING_SUPPORT_CONTRACT)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &c.env)
}

func (c *CopyTradingCtrl) SetContracts() {
	type setContractsRequest struct {
		Constracts []string `json:"constracts"`
	}

	l := &setContractsRequest{}
	l.Constracts = []string{"BTCUSDT", "ETHUSDT"}

	requestBody, err := json.Marshal(l)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := c.env.Url + consts.COPYTRADING_SET_CONTRACT
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.COPYTRADING_SET_CONTRACT, string(requestBody), &c.env)
}

func (c *CopyTradingCtrl) LeaderPosition() {
	requestURL := fmt.Sprintf(c.env.Url+consts.COPYTRADING_LEADER_POSITION+"?pageNum=%d&pageSize=%d", 1, 1)
	requestPath := fmt.Sprintf(consts.COPYTRADING_LEADER_POSITION+"?pageNum=%d&pageSize=%d", 1, 1)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &c.env)
}

func (c *CopyTradingCtrl) EstimateProfit() {
	requestURL := fmt.Sprintf(c.env.Url+consts.COPYTRADING_ESTIMATE_PROFIT+"?pageNum=%d&pageSize=%d", 1, 20)
	requestPath := fmt.Sprintf(consts.COPYTRADING_ESTIMATE_PROFIT+"?pageNum=%d&pageSize=%d", 1, 20)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &c.env)
}

func (c *CopyTradingCtrl) HistoryProfit() {
	requestURL := fmt.Sprintf(c.env.Url + consts.COPYTRADING_HISTORY_PROFIT)
	requestPath := fmt.Sprintf(consts.COPYTRADING_HISTORY_PROFIT)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &c.env)
}

func (c *CopyTradingCtrl) FollowerRank() {
	requestURL := fmt.Sprintf(c.env.Url+consts.COPYTRADING_FOLLOWER_RANK+"?status=%d", 2)
	requestPath := fmt.Sprintf(consts.COPYTRADING_FOLLOWER_RANK+"?status=%d", 2)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &c.env)
}

func (c *CopyTradingCtrl) GetAccountIDs() {
	requestURL := fmt.Sprintf(c.env.Url + consts.COPYTRADING_GET_ACCOUNTID)
	requestPath := fmt.Sprintf(consts.COPYTRADING_GET_ACCOUNTID)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &c.env)
}
