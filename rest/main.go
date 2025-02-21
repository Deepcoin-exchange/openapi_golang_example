package main

import (
	"apiRequest/accountCtrl"
	assetCtrl "apiRequest/asset"
	"apiRequest/consts"
	"apiRequest/copytrading"
	"apiRequest/marketCtrl"
	"apiRequest/structs"
	"apiRequest/tradeCtrl"
	"apiRequest/ws"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func getConfigs() *structs.Env {
	configName := consts.MASTER

	fmt.Println(configName)

	viper.SetConfigName(configName)
	viper.AddConfigPath("./config")

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	url := viper.GetString("api.url")
	key := viper.GetString("api.key")
	secretKey := viper.GetString("api.secret_key")
	passphrase := viper.GetString("api.passphrase")

	return &structs.Env{
		Url:        url,
		Key:        key,
		SecretKey:  secretKey,
		Passphrase: passphrase,
	}
}

type CommandHandler func()

type CommandRegistry struct {
	handlers map[string]CommandHandler
}

func NewCommandRegistry(env *structs.Env) *CommandRegistry {
	account := accountCtrl.NewAccountCtrl(env)
	market := marketCtrl.NewMarketCtrl(env)
	trade := tradeCtrl.NewTradeCtrl(env)
	ws := ws.NewPrivateWsCtrl(env)
	copytrading := copytrading.NewCopyTradingCtrl(env)
	asset := assetCtrl.NewAssetCtrl(env)

	return &CommandRegistry{
		handlers: map[string]CommandHandler{
			// account
			"getAccountBalance": account.GetAccountBalance,
			"getAccountBills":   account.GetAccountBills,
			"setLeverage":       account.SetLeverage,
			"getPositions":      account.GetPositions,

			// market
			"getMarketCandles":     market.GetMarketCandles,
			"getMarketTickers":     market.GetMarketTickers,
			"getMarketInstruments": market.GetMarketInstruments,

			// trade
			"order":                            trade.Order,
			"cancelOrder":                      trade.CancelOrder,
			"tradeFills":                       trade.TradeFills,
			"historyOrder":                     trade.HistoryOrder,
			"spotHistoryOrder":                 trade.SpotHistoryOrder,
			"pendingOrder":                     trade.PendingOrder,
			"spotPendingOrder":                 trade.SpotPendingOrder,
			"swapQueryOrderByOrderSysID":       trade.SwapQueryOrderByOrderSysID,
			"spotQueryOrderByOrderSysID":       trade.SpotQueryOrderByOrderSysID,
			"swapFinishQueryOrderByOrderSysID": trade.SwapFinishQueryOrderByOrderSysID,
			"spotFinishQueryOrderByOrderSysID": trade.SpotFinishQueryOrderByOrderSysID,
			"getPosition":                      trade.GetPosition,
			"getFundingRate":                   trade.GetFundingRate,
			"replace-order":                    trade.ReplaceOrder,
			"batch-cancel-order":               trade.BatchCancelOrder,
			"swapQueryPendingOrders":           trade.SwapQueryPendingOrders,
			"swapCalcelAllOrders":              trade.SwapCalcelAllOrders,
			"swapReplaceOrderSlTp":             trade.SwapReplaceOrderSlTp,
			"swapReplacePositionSlTp":          trade.SwapReplacePositionSlTp,

			// copytrading
			"leader-settings":   copytrading.LeaderSettings,
			"support-contracts": copytrading.SupportContracts,
			"set-contracts":     copytrading.SetContracts,
			"leader-position":   copytrading.LeaderPosition,
			"estimate-profit":   copytrading.EstimateProfit,
			"history-profit":    copytrading.HistoryProfit,
			"follower-rank":     copytrading.FollowerRank,
			"getAccountIDs":     copytrading.GetAccountIDs,

			// WebSocket
			"getListenKey":    ws.GetListenKey,
			"extendListenKey": ws.Extend,

			// asset
			"depositList":                asset.GetDepositList,
			"withdrawList":               asset.GetWithdrawList,
			"getInternalTransferSupport": asset.GetInternalTransferSupport,
			"postInternalTransfer":       asset.PostInternalTransfer,
			"getInternalTransferHistory": asset.GetInternalTransferHistory,
		},
	}
}

func (r *CommandRegistry) Execute(command string) error {
	handler, exists := r.handlers[command]
	if !exists {
		return fmt.Errorf("unknown command: %s", command)
	}
	handler()
	return nil
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("command is required")
		return
	}

	env := getConfigs()
	registry := NewCommandRegistry(env)

	if err := registry.Execute(args[1]); err != nil {
		fmt.Println(err)
	}
}
