package path

import "time"

// AnnouncementDetails provides the details for a particular announcement
type AnnouncementDetails struct {
	// The network that the BGP announcement was made for
	Net    string
	Reason string
	Start  time.Time
	End    time.Time
}

// AnnouncementHistory possesses the history of BGP announcements
type AnnouncementHistory struct {
	AnnouncementHistory []AnnouncementDetails `json:"announcement_history"`
}
