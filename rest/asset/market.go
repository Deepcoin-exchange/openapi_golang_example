package assetCtrl

import (
	"apiRequest/consts"
	"apiRequest/signature"
	"apiRequest/structs"
	"encoding/json"
	"fmt"
)

type AssetCtrl struct {
	env structs.Env
}

func NewAssetCtrl(env *structs.Env) *AssetCtrl {
	return &AssetCtrl{
		env: *env,
	}
}

func (a *AssetCtrl) GetDepositList() {
	requestURL := fmt.Sprintf(a.env.Url+consts.ASSET_DEPOSIT_LIST+"?size=%s", "50")
	requestPath := fmt.Sprintf(consts.ASSET_DEPOSIT_LIST+"?size=%s", "50")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &a.env)
}

func (a *AssetCtrl) GetWithdrawList() {
	requestURL := fmt.Sprintf(a.env.Url+consts.ASSET_WITHDARW_LIST+"?size=%s", "1")
	requestPath := fmt.Sprintf(consts.ASSET_WITHDARW_LIST+"?size=%s", "1")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &a.env)
}

func (a *AssetCtrl) GetInternalTransferSupport() {
	requestURL := fmt.Sprintf(a.env.Url + consts.INTERNAL_TRANSFER_SUPPORT)
	requestPath := fmt.Sprintf(consts.INTERNAL_TRANSFER_SUPPORT)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &a.env)
}

func (a *AssetCtrl) GetInternalTransferHistory() {
	requestURL := fmt.Sprintf(a.env.Url + consts.INTERNAL_TRANSFER_HISTORY)
	requestPath := fmt.Sprintf(consts.INTERNAL_TRANSFER_HISTORY)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &a.env)
}

func (a *AssetCtrl) PostInternalTransfer() {
	type internalTransferRequest struct {
		Amount          string `json:"amount"`
		Coin            string `json:"coin"`
		ReceiverAccount string `json:"receiverAccount,omitempty"`
		AccountType     string `json:"accountType,omitempty"`
		ReceiverUID     string `json:"receiverUID"`
	}

	req := &internalTransferRequest{}
	req.Amount = "10"
	req.Coin = "USDT"
	req.ReceiverUID = "36007196"

	requestBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := a.env.Url + consts.INTERNAL_TRANSFER
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.INTERNAL_TRANSFER, string(requestBody), &a.env)
}
