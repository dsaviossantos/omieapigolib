package omie_api

// OmieApiRateLimits - Rate limits for Omie API
type omieApiRateLimits struct {
	// 960 req/min per IP
	ipRateLimitRemaining int
	// 240 req/min per IP+AppKey+Method
	appKeyMethodRateLimitRemaining int
	// 10 requests with error - Block for 30 minutes
	errors int
	// Block Time
	block int
	// Methods being executed
	executingMethods []string
}

func (o *omieApiRateLimits) DecreaseRemaining() {
	o.ipRateLimitRemaining -= 1
}

func omieApiRateLimitsFactory() *omieApiRateLimits {
	return &omieApiRateLimits{
		ipRateLimitRemaining:           960,
		appKeyMethodRateLimitRemaining: 240,
		errors:                         10,
		block:                          0,
		executingMethods:               []string{},
	}
}

var RateLimits = omieApiRateLimitsFactory()
