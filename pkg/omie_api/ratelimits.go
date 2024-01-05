package omie_api

// OmieApiRateLimits - Rate limits for Omie API
type omieApiRateLimits struct {
	// 4 req/sec per IP
	LimitBySecond int
	// 960 req/min per IP
	ipRateLimitRemaining int
	// 240 req/min per IP+AppKey+Method
	appKeyMethodRateLimitRemaining int
	// 10 requests with error - Block for 30 minutes
	errors int
	// Blocking Time
	block int
	// Methods being executed
	executingMethods [4]string
}

// AddMethod - Add method to executingMethods
// Return wait time 1 if method already exists
// Return wait time 0 if method added
func (o *omieApiRateLimits) AddMethod(method string) int {
	if len(o.executingMethods) <= 4 {
		for i := range o.executingMethods {
			var emptySpot int
			if o.executingMethods[i] == method {
				return 1
			} else if o.executingMethods[i] == "" {
				emptySpot = i
			} else {
				o.executingMethods[emptySpot] = method
				return 0
			}
		}
	}
	return 1
}

// RemoveMethod - Remove method from executingMethods
func (o *omieApiRateLimits) RemoveMethod(method string) {
	for i := range o.executingMethods {
		if o.executingMethods[i] == method {
		}
	}
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
		executingMethods:               [4]string{},
	}
}

var RateLimits = omieApiRateLimitsFactory()
