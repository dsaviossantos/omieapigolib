package omie_api

// OmieApiRateLimits - Rate limits for Omie API
// 960 req/min per IP
// 240 req/min per IP+AppKey+Method
// 4 req at same time per IP+AppKey+Method - Some are special cases
// Redundant request - wait 60 seconds
// 10 requests with error - Block for 30 minutes
type OmieApiRateLimits struct {
	IpRateLimit                    int
	IpRateLimitRemaining           int
	AppKeyMethodRateLimit          int
	AppKeyMethodRateLimitRemaining int
	AppKeyMethodRateLimitReset     int
	SimultaneousRequests           int
	RedundantRequest               int
	Errors                         int
	Block                          int
	ExecutingMethods               []string
}

func OmieApiRateLimitsFactory() *OmieApiRateLimits {
	return &OmieApiRateLimits{
		IpRateLimit:                    960,
		IpRateLimitRemaining:           960,
		AppKeyMethodRateLimit:          240,
		AppKeyMethodRateLimitRemaining: 240,
		AppKeyMethodRateLimitReset:     60,
		SimultaneousRequests:           4,
		RedundantRequest:               60,
		Errors:                         10,
		Block:                          30,
	}
}
