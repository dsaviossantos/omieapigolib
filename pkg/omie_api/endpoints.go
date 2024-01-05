package omie_api

import (
	"encoding/json"
	"io"
	"log/slog"
)

// Params - Params for Omie API
type params map[string]interface{}

func (p params) addParam(key string, value interface{}) {
	p[key] = value
}

func (p params) toList() []params {
	return []params{p}
}

// Get Method for Omie API
var Get = Payload{
	baseUrl: "https://app.omie.com.br/api",
	body:    []byte{},
}

// Módulo Geral - Empresas - Listar Empresas
//
// https://app.omie.com.br/api/v1/geral/empresas/#ListarEmpresas
//
// page: Número da Página a ser exibida
//
// nRegistriesPage: Número de registros por página
func (o Payload) ListCompanies(page, nRegistriesPage int) CompaniesList {
	var cl CompaniesList
	var method string = "ListarEmpresas"

	p := make(params)
	p.addParam("pagina", page)
	p.addParam("registros_por_pagina", nRegistriesPage)

	o.body = NewBody(method, p.toList())

	slog.Info("Requesting:", "URL", o.baseUrl+"/v1/geral/empresas/", "METHOD", method, "PARAMS", p)
	response := RequestWithRetry(o.baseUrl+"/v1/geral/empresas/", o.body, 3)
	r, errRead := io.ReadAll(response.Body)
	if errRead != nil {
		slog.Error("IO Read ResponseBody", "LOCATION", "endpoints.go:53", "MSG", errRead, "CAUSE", response.Body)
	}
	errUnmarshal := json.Unmarshal(r, &cl)
	if errUnmarshal != nil {
		slog.Error("JSON Unmarshal", "LOCATION", "endpoints.go:58", "MSG", errUnmarshal, "CAUSE", r)
	}
	slog.Info("Response:", "Company", cl.EmpresasCadastro[0].NomeFantasia)
	return cl
}
