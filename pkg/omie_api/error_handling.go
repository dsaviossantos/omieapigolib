package omie_api

type OmieError struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	Referer     string `json:"referer"`
	Fatal       bool   `json:"fatal"`
}
