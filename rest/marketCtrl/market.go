package marketCtrl

import (
	"apiRequest/consts"
	"apiRequest/signature"
	"apiRequest/structs"
	"fmt"
)

type MarketCtrl struct {
	env structs.Env
}

func NewMarketCtrl(env *structs.Env) *MarketCtrl {
	return &MarketCtrl{
		env: *env,
	}
}

func (m *MarketCtrl) GetMarketCandles() {
	requestURL := fmt.Sprintf(m.env.Url+consts.MARKET_CANDLES+"?instId=%s", "BTC-USDT")
	requestPath := fmt.Sprintf(consts.MARKET_CANDLES+"?instId=%s", "BTC-USDT")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &m.env)
}

func (m *MarketCtrl) GetMarketTickers() {
	requestURL := fmt.Sprintf(m.env.Url+consts.MARKET_TICKERS+"?instType=%s&uly=%s", consts.SPOT, "SAGE-USDT")
	requestPath := fmt.Sprintf(consts.MARKET_TICKERS+"?instType=%s&uly=%s", consts.SPOT, "SAGE-USDT")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &m.env)
}

func (m *MarketCtrl) GetMarketInstruments() {
	requestURL := fmt.Sprintf(m.env.Url+consts.MARKET_INSTRUMENTS+"?instType=%s&uly=%s", consts.SWAP, "BTC-USDT")
	requestPath := fmt.Sprintf(consts.MARKET_INSTRUMENTS+"?instType=%s&uly=%s", consts.SWAP, "BTC-USDT")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &m.env)
}
