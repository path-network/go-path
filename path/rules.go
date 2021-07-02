package path

// Rules holds a list of Rule objects
type Rules struct {
	Rules []Rule `json:"rules"`
}

// Rule represents a rule entry in the firewall
type Rule struct {
	Protocol string `json:"protocol,omitempty"`
	// omitempty is required in the event that these values are not provided, as Go will default to 0, which will be
	// recognized as an invalid port number
	DstPort  int    `json:"dst_port,omitempty"`
	SrcPort  int    `json:"src_port,omitempty"`
	// If we do not want to rate limit the rule, then we must send a null value for rate_limter_id. An empty string does
	// not suffice, as it will be recognized as an invalid UUID. If we do not specify a value for the RateLimiterID member,
	// the uninitialized pointer iw
	RateLimiterID *string `json:"rate_limiter_id"`
	Whitelist     bool    `json:"whitelist"`
	Destination   string  `json:"destination"`
	Source        string  `json:"source"`
	Priority      bool    `json:"priority"`
	Comment       string  `json:"comment"`
	ID            string  `json:"id,omitempty"`
}
