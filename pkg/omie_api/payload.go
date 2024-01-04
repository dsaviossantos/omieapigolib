package omie_api

import (
	"encoding/json"
)

// RequestPayload - Request Payload for Omie API
type Payload struct {
	baseUrl string
	body    []byte
}

// RequestBody - Request Body for Omie API
type body struct {
	Call      string   `json:"call"`
	AppKey    string   `json:"app_key"`
	AppSecret string   `json:"app_secret"`
	Param     []params `json:"param"`
}

func (b *body) buildBody(call string, params []params) []byte {
	b.Call = call
	b.AppKey = Config.Key
	b.AppSecret = Config.Secret
	b.Param = params[:]
	bodyJson, errMarshalling := json.Marshal(b)
	if errMarshalling != nil {
		panic(errMarshalling)
	}
	return bodyJson
}

func NewBody(call string, params []params) []byte {
	b := body{}
	return b.buildBody(call, params)
}
