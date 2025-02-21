package ws

import (
	"apiRequest/consts"
	"apiRequest/signature"
	"apiRequest/structs"
	"fmt"
)

type PrivateWsCtrl struct {
	env structs.Env
}

func NewPrivateWsCtrl(env *structs.Env) *PrivateWsCtrl {
	return &PrivateWsCtrl{
		env: *env,
	}
}

func (t *PrivateWsCtrl) GetListenKey() {
	requestURL := fmt.Sprintf(t.env.Url + consts.ListenKey)
	requestPath := fmt.Sprintf(consts.ListenKey)
	
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

func (t *PrivateWsCtrl) Extend() {
	requestURL := fmt.Sprintf(t.env.Url + consts.ExtendListenKey+"?listenKey=%s", "3f8021a44a262e69344d5c522b613006")
	requestPath := fmt.Sprintf(consts.ExtendListenKey+"?listenKey=%s", "3f8021a44a262e69344d5c522b613006")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

