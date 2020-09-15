package path

// RateLimiters holds a list of RateLimiter objects
type RateLimiters struct {
	RateLimiters []RateLimiter `json:"rate_limiter"`
}

// RateLimiter stores the data that a single rate limiter entry contains
type RateLimiter struct {
	// PacketsPerSecond is the amount of packets that can be stored in the token bucket each second before subsequent
	// packets get rejected
	PacketsPerSecond int `json:"packets_per_second"`
	// Comment allows you to indicate what use the rate limiter has
	Comment string `json:"comment"`
	// ID is a unique identifier for each rate limiter. This allows it to be attached to rules
	ID string `json:"id,omitempty"`
}
