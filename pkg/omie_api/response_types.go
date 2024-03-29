package omie_api

type CompaniesList struct {
	Pagina           int `json:"pagina"`
	TotalDePaginas   int `json:"total_de_paginas"`
	Registros        int `json:"registros"`
	TotalDeRegistros int `json:"total_de_registros"`
	EmpresasCadastro []struct {
		Bairro                                string `json:"bairro"`
		Cep                                   string `json:"cep"`
		Cidade                                string `json:"cidade"`
		Cnae                                  string `json:"cnae"`
		CnaeMunicipal                         string `json:"cnae_municipal"`
		Cnpj                                  string `json:"cnpj"`
		CodigoEmpresa                         string `json:"codigo_empresa"`
		CodigoPais                            string `json:"codigo_pais"`
		Complemento                           string `json:"complemento"`
		CscHomologacao                        string `json:"csc_homologacao"`
		CscIDHomologacao                      string `json:"csc_id_homologacao"`
		CscIDProducao                         string `json:"csc_id_producao"`
		CscProducao                           string `json:"csc_producao"`
		EcdCodigoCadastral                    string `json:"ecd_codigo_cadastral"`
		EcdCodigoInstituicaoResponsavel       string `json:"ecd_codigo_instituicao_responsavel"`
		EfdAtividadeIndustrial                string `json:"efd_atividade_industrial"`
		EfdPerfilArquivoFiscal                string `json:"efd_perfil_arquivo_fiscal"`
		Email                                 string `json:"email"`
		Endereco                              string `json:"endereco"`
		EnderecoNumero                        string `json:"endereco_numero"`
		Estado                                string `json:"estado"`
		FaxDdd                                string `json:"fax_ddd"`
		FaxNumero                             string `json:"fax_numero"`
		GeraNfse                              string `json:"gera_nfse"`
		Inativa                               string `json:"inativa"`
		InscricaoEstadual                     string `json:"inscricao_estadual"`
		InscricaoMunicipal                    string `json:"inscricao_municipal"`
		InscricaoSuframa                      string `json:"inscricao_suframa"`
		Logradouro                            string `json:"logradouro"`
		NomeFantasia                          string `json:"nome_fantasia"`
		OptanteSimplesNacional                string `json:"optante_simples_nacional"`
		PdvSincrAnalitica                     string `json:"pdv_sincr_analitica"`
		RazaoSocial                           string `json:"razao_social"`
		RegimeTributario                      string `json:"regime_tributario"`
		SpedCodigoCriterioEscrituracao        string `json:"sped_codigo_criterio_escrituracao"`
		SpedCodigoIncidenciaTributaria        string `json:"sped_codigo_incidencia_tributaria"`
		SpedCodigoIndicadorApropriacaoCredito string `json:"sped_codigo_indicador_apropriacao_credito"`
		SpedCodigoTipoAtividade               string `json:"sped_codigo_tipo_atividade"`
		SpedCodigoTipoContribuicao            string `json:"sped_codigo_tipo_contribuicao"`
		SpedCpfContador                       string `json:"sped_cpf_contador"`
		SpedCrcContador                       string `json:"sped_crc_contador"`
		SpedEmailContador                     string `json:"sped_email_contador"`
		SpedJuntaComercial                    string `json:"sped_junta_comercial"`
		SpedMatriz                            string `json:"sped_matriz"`
		SpedNaturezaPessoaJuridica            string `json:"sped_natureza_pessoa_juridica"`
		SpedNomeContador                      string `json:"sped_nome_contador"`
		SpedRegistroJuntaComercial            string `json:"sped_registro_junta_comercial"`
		SpedUsaContabilidadeTerceirizada      string `json:"sped_usa_contabilidade_terceirizada"`
		Telefone1Ddd                          string `json:"telefone1_ddd"`
		Telefone1Numero                       string `json:"telefone1_numero"`
		Telefone2Ddd                          string `json:"telefone2_ddd"`
		Telefone2Numero                       string `json:"telefone2_numero"`
		Website                               string `json:"website"`
	} `json:"empresas_cadastro"`
	ProdutoServicoResumido []any `json:"produto_servico_resumido"`
}
