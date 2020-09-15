package path

// Holds a list of diversion
type Diversions struct {
	Diversions []Diversion `json:"diversion"`
}

// Represents a Path diversion
type Diversion struct {
	Subnet      string        `json:"subnet"`
	// Whether or not the diversion was performed manually
	Manual      bool          `json:"manual"`
	UnderAttack []UnderAttack `json:"under_attack"`
}

// Provides details for a diversion
type UnderAttack struct {
	Host   string `json:"host"`
	Reason string `json:"reason"`
	Since  string `json:"since"`
}
