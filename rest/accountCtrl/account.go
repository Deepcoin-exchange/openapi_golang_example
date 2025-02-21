package accountCtrl

import (
	"apiRequest/consts"
	"apiRequest/signature"
	"apiRequest/structs"
	"encoding/json"
	"fmt"
)

type AccountCtrl struct {
	env structs.Env
}

func NewAccountCtrl(env *structs.Env) *AccountCtrl {
	return &AccountCtrl{
		env: *env,
	}
}

func (a *AccountCtrl) GetAccountBalance() {
	requestURL := fmt.Sprintf(a.env.Url+consts.ACCOUNT_BALANCE+"?instType=%s&ccy=", consts.SWAP)
	requestPath := fmt.Sprintf(consts.ACCOUNT_BALANCE+"?instType=%s&ccy=", consts.SWAP)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &a.env)
}

func (a *AccountCtrl) GetAccountBills() {
	requestURL := fmt.Sprintf(a.env.Url+consts.ACCOUNT_BILLS+"?instType=%s", consts.SPOT)
	requestPath := fmt.Sprintf(consts.ACCOUNT_BILLS+"?instType=%s", consts.SPOT)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &a.env)
}

func (a *AccountCtrl) SetLeverage() {
	type setLeverageRequest struct {
		InstId      string `json:"instId,omitempty"`
		Lever       string `json:"lever,omitempty"`
		MgnMode     string `json:"mgnMode,omitempty"`
		MrgPosition string `json:"mrgPosition,omitempty"`
	}

	setLeverage := &setLeverageRequest{}
	setLeverage.InstId = "BTC-USDT-SWAP"
	setLeverage.Lever = "17"
	setLeverage.MgnMode = consts.CROSS
	setLeverage.MrgPosition = consts.MERGE

	requestBody, err := json.Marshal(setLeverage)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := a.env.Url + consts.SET_LEVERAGE
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.SET_LEVERAGE, string(requestBody), &a.env)
}

func (a *AccountCtrl) GetPositions() {
	requestURL := fmt.Sprintf(a.env.Url+consts.POSITIONS+"?instType=%s", consts.SWAP)
	requestPath := fmt.Sprintf(consts.POSITIONS+"?instType=%s", consts.SWAP)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &a.env)
}
