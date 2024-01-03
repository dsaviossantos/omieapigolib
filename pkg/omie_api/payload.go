package omie_api

import (
	"encoding/json"
	"fmt"
)

// RequestPayload - Request Payload for Omie API
type Payload struct {
	BaseUrl string
	Body    []byte
}

// RequestBody - Request Body for Omie API
type Body struct {
	Call      string   `json:"call"`
	AppKey    string   `json:"app_key"`
	AppSecret string   `json:"app_secret"`
	Param     []params `json:"param"`
}

func (b Body) buildBody(call string, params []params) []byte {
	b.Call = call
	b.AppKey = Config.Key
	b.AppSecret = Config.Secret
	b.Param = params[:]
	bodyJson, errMarshalling := json.Marshal(b)
	fmt.Println(string(bodyJson))
	if errMarshalling != nil {
		panic(errMarshalling)
	}
	return bodyJson
}
