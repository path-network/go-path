package path

import "time"

// AttackHistory stores the history of all attacks detected for hosts under your account
type AttackHistory struct {
	AttackHistory []AttackDetails `json:"attack_history"`
}

// AttackDetails gives us the details of an attack that occurred
type AttackDetails struct {
	Host    string     `json:"host"`
	Reason  string     `json:"reason"`
	Start   time.Time  `json:"start"`
	End     time.Time  `json:"end"`
	PeakBPS AttackPeak `json:"peak_bps"`
	PeakPPS AttackPeak `json:"peak_pps"`
}

// AttackPeak represents a peak value that happened during a particular attack, such as the bits per second or packets
// per second
type AttackPeak struct {
	Value     int       `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
