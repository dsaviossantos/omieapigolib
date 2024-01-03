package omie_api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// Params - Params for Omie API
type params map[string]interface{}

// Get Method for Omie API
var Get = Payload{
	BaseUrl: "https://app.omie.com.br/api",
	Body:    []byte{},
}

// ListCompanies - List all companies registered in Omie Account
func (o Payload) ListCompanies() []byte {
	p := make(params)
	p["pagina"] = 1
	p["registros_por_pagina"] = 50
	par := [1]params{p}
	o.Body = Body{}.buildBody("ListarEmpresas", par[:])
	b := bytes.NewReader(o.Body)
	response, err := http.Post(o.BaseUrl+"/v1/geral/empresas/", "application/json", b)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	r, errRead := io.ReadAll(response.Body)
	if errRead != nil {
		fmt.Println(errRead)
	}
	return r
}
